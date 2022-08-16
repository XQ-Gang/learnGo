# Q&A

## 什么是协程（Goroutine）

协程是**用户态轻量级线程**，它是**线程调度的基本单位**。通常在函数前加上go关键字就能实现并发。一个Goroutine会以一个很小的栈启动2KB或4KB，当遇到栈空间不足时，栈会**自动伸缩**， 因此可以轻易实现成千上万个goroutine同时启动。

## 如何高效地拼接字符串

拼接字符串的方式有："+", fmt.Sprintf, strings.Builder, bytes.Buffer, strings.Join

1. **"+"**：使用`+`操作符进行拼接时，会对字符串进行遍历，计算并开辟一个新的空间来存储原来的两个字符串。

2. **fmt.Sprintf**：由于采用了接口参数，必须要用**反射**获取值，因此有性能损耗。

3. **strings.Builder**：用 WriteString() 进行拼接，内部实现是**指针+切片**，同时 String() 返回拼接后的字符串，它是直接把 []byte 转换为 string，从而避免变量拷贝。

`strings.builder`的实现原理很简单，结构如下：

```go
 type Builder struct {
     addr *Builder // of receiver, to detect copies by value
     buf  []byte // 1
 }
```

`addr`字段主要是做`copycheck`，`buf`字段是一个`byte`类型的切片，这个就是用来存放字符串内容的，提供的`WriteString()`方法就是向切片`buf`中追加数据：

```go
 func (b *Builder) WriteString(s string) (int, error) {
     b.copyCheck()
     b.buf = append(b.buf, s...)
     return len(s), nil
 }
```

提供的`String`方法就是将`[]byte`转换为`string`类型，这里为了避免内存拷贝的问题，使用了强制转换来避免内存拷贝：

```go
 func (b *Builder) String() string {
     return *(*string)(unsafe.Pointer(&b.buf))
 }
```

4. **bytes.Buffer**：`bytes.Buffer`是一个一个缓冲`byte`类型的缓冲器，这个缓冲器里存放着都是`byte`。

`bytes.buffer`底层也是一个`[]byte`切片，结构体如下：

```Go
type Buffer struct {
    buf      []byte // contents are the bytes buf[off ：len(buf)]
    off      int    // read at &buf[off], write at &buf[len(buf)]
    lastRead readOp // last read operation, so that Unread* can work correctly.
}
```

因为`bytes.Buffer`可以持续向`Buffer`尾部写入数据，从`Buffer`头部读取数据，所以`off`字段用来记录读取位置，再利用切片的`cap`特性来知道写入位置，重点看一下`WriteString`方法是如何拼接字符串的：

```go
func (b *Buffer) WriteString(s string) (n int, err error) {
    b.lastRead = opInvalid
    m, ok := b.tryGrowByReslice(len(s))
    if !ok {
        m = b.grow(len(s))
    }
    return copy(b.buf[m:], s), nil
}
```

切片在创建时并不会申请内存块，只有在往里写数据时才会申请，首次申请的大小即为写入数据的大小。如果写入的数据小于64字节，则按64字节申请。采用动态扩展`slice`的机制，字符串追加采用`copy`的方式将追加的部分拷贝到尾部，`copy`是内置的拷贝函数，可以减少内存分配。

但是在将`[]byte`转换为`string`类型依旧使用了标准类型，所以会发生内存分配：

```go
func (b *Buffer) String() string {
    if b == nil {
        // Special case, useful in debugging.
        return "<nil>"
    }
    return string(b.buf[b.off:])
}
```

5. **strings.join**：基于`strings.builder`来实现的，并且可以自定义分隔符。

```go
func Join(elems []string, sep string) string {
    switch len(elems) {
    case 0:
        return ""
    case 1:
        return elems[0]
    }
    n := len(sep) * (len(elems) - 1)
    for i := 0; i < len(elems); i++ {
        n += len(elems[i])
    }

    var b Builder
    b.Grow(n)
    b.WriteString(elems[0])
    for _, s := range elems[1:] {
        b.WriteString(sep)
        b.WriteString(s)
    }
    return b.String()
}
```

唯一不同在于在`Join`方法内调用了`b.Grow(n)`方法，这个是进行初步的容量分配，而前面计算的n的长度就是我们要拼接的slice的长度，因为我们传入切片长度固定，所以提前进行容量分配可以减少内存分配，很高效。

```go
func main(){
	a := []string{"a", "b", "c"}
	//方式1：
	ret := a[0] + a[1] + a[2]
	//方式2：
	ret := fmt.Sprintf("%s%s%s", a[0],a[1],a[2])
	//方式3：
	var sb strings.Builder
	sb.WriteString(a[0])
	sb.WriteString(a[1])
	sb.WriteString(a[2])
	ret := sb.String()
	//方式4：
	buf := new(bytes.Buffer)
	buf.Write(a[0])
	buf.Write(a[1])
	buf.Write(a[2])
	ret := buf.String()
	//方式5：
	ret := strings.Join(a,"")
}
```

**总结**：**strings.Join ≈ strings.Builder > bytes.Buffer > "+" > fmt.Sprintf**

## 如何知道一个对象是分配在栈上还是堆上？逃逸分析

**逃逸分析：通过指针的动态范围决定一个变量究竟是分配在栈上还是应该分配在堆上。**

我们知道栈区是可以自动清理的，所以栈区的效率很高，但是不可能把所有的对象都申请在栈上面，而且栈空间也是有限的。但如果所有的对象都分配在堆区的话，堆又不像栈那样可以自动清理，因此会频繁造成垃圾回收，从而降低运行效率。

Go 局部变量会进行**逃逸分析**，把一次性对象分配到栈区，如果后续还有引用，那么就放到堆区：

- 如果变量离开作用域后没有被引用，那么优先分配到栈中（如果申请的内存过大，栈区存不下，会分配到堆）。
- 如果变量离开作用域后外部还有引用，那么必定分配到堆中。

那么如何判断是否发生了逃逸呢？`go build -gcflags '-m -m -l' xxx.go`。

## 2 个 interface 可以比较吗 ？

Go 语言中，interface 的内部实现包含了 2 个字段，类型 `T` 和 值 `V`，interface 可以使用 `==` 或 `!=` 比较。2 个 interface 相等有以下 2 种情况

1. 两个 interface 均等于 nil（此时 V 和 T 都处于 unset 状态）
2. 类型 T 相同，且对应的值 V 相等。

看下面的例子：

```go
type Stu struct {
     Name string
}

type StuInt interface{}

func main() {
     var stu1, stu2 StuInt = &Stu{"Tom"}, &Stu{"Tom"}
     var stu3, stu4 StuInt = Stu{"Tom"}, Stu{"Tom"}
     fmt.Println(stu1 == stu2) // false
     fmt.Println(stu3 == stu4) // true
}
```

`stu1` 和 `stu2` 对应的类型是 `*Stu`，值是 Stu 结构体的地址，两个地址不同，因此结果为 false。
`stu3` 和 `stu4` 对应的类型是 `Stu`，值是 Stu 结构体，且各字段相等，因此结果为 true。

## 2 个 nil 可能不相等吗？

可能不等。interface 在运行时绑定值，只有值为 nil 接口值才为 nil，但是与指针的 nil 不相等。

举个例子：

```Go
var p *int = nil
var i interface{} = nil
if(p == i){
	fmt.Println("Equal")
}
```

两者并不相同。总结：**两个nil只有在类型相同时才相等**。

## Go 语言GC(垃圾回收)的工作原理

Go1.3采用**标记清除法**，该算法需要在执行期间需要暂停应用程序(STW)，无法满足并发程序的实时性。 Go1.5采用**三色标记法**，Go1.8采用**三色标记法+混合写屏障**，通过混合写屏障技术保证了Go并发执行GC时内存中对象的三色一致性（**这里的并发指的是GC和应用程序的goroutine能同时执行**）。

Go GC有**四**个阶段:

- STW，开启混合写屏障，扫描栈对象。
- 将所有对象加入白色集合，从根对象开始，将其放入灰色集合。每次从灰色集合取出一个对象标记为黑色，然后遍历其子对象，标记为灰色，放入灰色集合。如此循环直到灰色集合为空。剩余的白色对象就是需要清理的对象。
- STW，关闭混合写屏障。
- 在后台进行GC（并发）。

标记阶段会减少程序的性能，而清理阶段是不会对程序有影响的。

**标记清除法**

分为两个阶段：标记和清除

- 标记阶段：暂停应用程序的执行，从根对象触发查找并标记堆中所有存活的对象；
- 清除阶段：遍历堆中的全部对象，回收未被标记的垃圾对象并将回收的内存加入空闲链表，恢复应用程序的执行；

缺点是需要暂停程序STW。

**三色标记法**：

将对象标记为白色，灰色或黑色。

- 白色：潜在的垃圾，其内存可能会被垃圾收集器回收；
- 黑色：活跃的对象，包括不存在任何引用外部指针的对象以及从根对象可达的对象，垃圾回收器不会扫描这些对象的子对象；
- 灰色：活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象；

具体实现：

- 在进入GC的三色标记阶段的一开始，所有对象都是白色的。
- 遍历根节点集合里的所有根对象（栈上的对象或者堆上的全局变量），把根对象引用的对象标记为灰色，从白色集合放入灰色集合。
- 遍历灰色集合，将灰色对象引用的对象从白色集合放入灰色集合，之后将此灰色对象放入黑色集合。
- 重复上一步骤，直到灰色集合中无任何对象。
- 回收白色集合里的所有对象，本次垃圾回收结束。

三色标记清除算法本身是不可以并发或者增量执行的，它仍然需要 STW。

想要在并发或者增量的标记算法中保证正确性，我们需要达成以下两种三色不变性中的任意一种：

- 强三色不变性——黑色对象不会指向白色对象，只会指向灰色对象或者黑色对象。
- 弱三色不变性——黑色对象指向的白色对象必须包含一条从灰色对象经由多个白色对象的可达路径。

**屏障技术**

垃圾收集中的屏障技术更像是一个钩子方法，它是在用户程序读取对象、创建新对象以及更新对象指针时执行的一段代码，根据操作类型的不同，我们可以将它们分成读屏障和写屏障两种，因为读屏障需要在读操作中加入代码片段，对用户程序的性能影响很大，所以编程语言往往都会采用写屏障保证三色不变性。

- 插入写屏障：当一个对象引用另外一个对象时，将另外一个对象标记为灰色，以此满足强三色不变性，不会存在黑色对象引用白色对象。
- 删除写屏障：在灰色对象删除对白色对象的引用时，将白色对象置为灰色，其实就是快照保存旧的引用关系，这叫STAB(snapshot-at-the-beginning)，以此满足弱三色不变性。
- **混合写屏障**：v1.8版本之前，运行时会使用插入写屏障保证强三色不变性；在v1.8中，组合插入写屏障和删除写屏障构成了混合写屏障，保证弱三色不变性。

基于插入写屏障和删除写屏障在结束时需要STW来重新扫描栈，带来性能瓶颈。

**混合写屏障**将垃圾收集的时间缩短至 0.5ms 以内，整体几乎不需要STW，效率高。分为以下四步：

1. GC开始时，将栈上的全部对象标记为黑色（不需要二次扫描，无需STW）；
2. GC期间，任何栈上创建的新对象均为黑色
3. 被删除引用的对象标记为灰色
4. 被添加引用的对象标记为灰色

总而言之就是确保黑色对象不能引用白色对象。

## 函数返回局部变量的指针是否安全？

在Go里面返回局部变量的指针是安全的。因为Go会进行**逃逸分析**，如果发现局部变量的作用域超过该函数则会**把指针分配到堆区**，避免内存泄漏。

## 非接口的任意类型 T() 都能够调用 `*T` 的方法吗？反过来呢？

一个T类型的值可以调用*T类型声明的方法，当且仅当T是**可寻址的**。

反之：*T 可以调用T()的方法，因为指针可以解引用。

## Go slice是怎么扩容的？

如果当前容量小于1024，则判断所需容量是否大于原来容量2倍，如果大于，当前容量加上所需容量；否则当前容量乘2。

如果当前容量大于1024，则每次按照1.25倍速度递增容量，也就是每次加上cap/4。

## 无缓冲的 channel 和有缓冲的 channel 的区别？

对于无缓冲区channel：发送的数据如果没有被接收方接收，那么**发送方阻塞；**如果一直接收不到发送方的数据，**接收方阻塞**。

有缓冲的channel：发送方在缓冲区满的时候阻塞，接收方不阻塞；接收方在缓冲区为空的时候阻塞，发送方不阻塞。

## 为什么有协程泄露(Goroutine Leak)？

协程泄漏是指协程创建之后没有得到释放。主要原因有：

1. 缺少接收器，导致发送阻塞
2. 缺少发送器，导致接收阻塞
3. 死锁。多个协程由于竞争资源导致死锁。
4. WaitGroup Add()和Done()不相等，前者更大。

## Go 可以限制运行时操作系统线程的数量吗？ 常见的goroutine操作函数有哪些？

可以，使用runtime.GOMAXPROCS(num int)可以设置线程数目。该值默认为CPU逻辑核数，如果设的太大，会引起频繁的线程切换，降低性能。

runtime.Gosched()，用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其它等待的任务运行，并在下次某个时候从该位置恢复执行。

runtime.Goexit()，调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。请注意千万别在主函数调用runtime.Goexit，因为会引发panic。

## 如何控制协程数目。

可以使用环境变量 `GOMAXPROCS` 或 `runtime.GOMAXPROCS(num int)` 设置，例如：

```Go
runtime.GOMAXPROCS(1) // 限制同时执行Go代码的操作系统线程数为 1
```

从官方文档的解释可以看到，`GOMAXPROCS` 限制的是同时执行用户态 Go 代码的操作系统线程的数量，但是对于被系统调用阻塞的线程数量是没有限制的。`GOMAXPROCS` 的默认值等于 CPU 的逻辑核数，同一时间，一个核只能绑定一个线程，然后运行被调度的协程。因此对于 CPU 密集型的任务，若该值过大，例如设置为 CPU 逻辑核数的 2 倍，会增加线程切换的开销，降低性能。对于 I/O 密集型应用，适当地调大该值，可以提高 I/O 吞吐率。

另外对于协程，可以用带缓冲区的channel来控制，下面的例子是协程数为1024的例子

```Go
var wg sync.WaitGroup
ch := make(chan struct{}, 1024)
for i:=0; i<20000; i++{
	wg.Add(1)
	ch<-struct{}{}
	go func(){
		defer wg.Done()
		<-ch
	}
}
wg.Wait()
```

此外还可以用**协程池**：其原理无外乎是将上述代码中通道和协程函数解耦，并封装成单独的结构体。

## Go面向对象是如何实现的？

Go实现面向对象的两个关键是struct和interface。

封装：对于同一个包，对象对包内的文件可见；对不同的包，需要将对象以大写开头才是可见的。

继承：继承是编译时特征，在struct内加入所需要继承的类即可：

```Go
type A struct{}
type B struct{
    A
}
```

多态：多态是运行时特征，Go多态通过interface来实现。类型和接口是松耦合的，某个类型的实例可以赋给它所实现的任意接口类型的变量。

Go支持多重继承，就是在类型中嵌入所有必要的父类型。

## Go 内存管理机制

golang内存管理基本是参考tcmalloc来进行的。Go内存管理本质上是一个内存池，只不过内部做了很多优化：**自动伸缩内存池大小，合理的切割内存块**。

Golang 的程序在启动之初，会一次性从操作系统那里申请一大块内存作为内存池。这块内存空间会放在一个叫 `mheap` 的 `struct` 中管理，mheap 负责将这一整块内存切割成不同的区域，并将其中一部分的内存切割成合适的大小，分配给用户使用。

**重要概念**：

- **`page`**：内存页，一块 `8K` 大小的内存空间。Go 与操作系统之间的内存申请和释放，都是以 `page` 为单位的。
- **`span`**：内存块，**一个或多个连续的** `page` 组成一个 `span`。
- **`sizeclass`**：空间规格，每个 `span` 都带有一个 `sizeclass`，标记着该 `span` 中的 `page` 应该如何使用。
- **`object`**：对象，用来存储一个变量数据内存空间，一个 `span` 在初始化时，会被切割成一堆**等大**的 `object`。假设 `object` 的大小是 `16B`，`span` 大小是 `8K`，那么就会把 `span` 中的 `page` 就会被初始化 `8K / 16B = 512` 个 `object`。所谓内存分配，就是分配一个 `object` 出去。

**内部的整体内存布局**：

![img](https://upload-images.jianshu.io/upload_images/11662994-356f568da2987e54.png?imageMogr2/auto-orient/strip|imageView2/2/format/webp)

- `mheap.spans`：用来存储 `page` 和 `span` 信息，比如一个 span 的起始地址是多少，有几个 page，已使用了多大等等。
- `mheap.bitmap`：存储着各个 `span` 中对象的标记信息，比如对象是否可回收等等。
- `mheap.arena_start`：将要分配给应用程序使用的空间。

**mcentral**：

- **用途相同**的 `span` 会以链表的形式组织在一起。 这里的用途用 `sizeclass` 来表示，就是指该 `span` 用来存储哪种大小的对象。
- 找到合适的 `span` 后，会从中取一个 `object` 返回给上层使用。这些 `span` 被放在一个叫做 mcentral 的结构中管理。

**mcache**：

- 为了提高内存并发申请效率，加入缓存层 mcache。
- 每一个 mcache 和处理器 P 对应。Go 申请内存首先从 P 的 mcache 中分配，如果没有可用的 span 再从 mcentral 中获取。
- 从 mcache 上分配内存空间是不需要加锁的，因为在同一时间里，一个 P 只有一个线程在其上面运行，不可能出现竞争。没有了锁的限制，大大加速了内存分配。

**整体的内存分配模型**：

![img](https://upload-images.jianshu.io/upload_images/11662994-e6d7200368ec06b6.png?imageMogr2/auto-orient/strip|imageView2/2/w/696/format/webp)

**其他优化**：

- #### zero size：有一些对象所需的内存大小是0，比如 `[0]int`, `struct{}`，这种类型的数据根本就不需要内存，所以没必要走上面那么复杂的逻辑。系统会直接返回一个固定的内存地址。

- #### Tiny 对象：像 `int32`, `byte`, `bool` 以及小字符串等常用的微小对象，都会使用 `sizeclass=1` 的 span，但分配给他们 `8B` 的空间，大部分是用不上的。并且这些类型使用频率非常高，就会导致出现大量的内部碎片。所以 Go 尽量不使用 `sizeclass=1` 的 span， 而是将 `< 16B` 的对象为统一视为 tiny 对象(tinysize)。分配时，从 `sizeclass=2` 的 span 中获取一个 `16B` 的 object 用以分配。如果存储的对象小于 `16B`，这个空间会被暂时保存起来 (`mcache.tiny` 字段)，下次分配时会复用这个空间，直到这个 object 用完为止。平均会节省 `20%` 左右的内存。但如果要存储的数据里有**指针**，即使 `<= 8B` 也不会作为 tiny 对象对待，而是正常使用 `sizeclass=1` 的 `span`。

- 大对象：最大的 sizeclass 最大只能存放 `32K` 的对象。如果一次性申请超过 `32K` 的内存，系统会直接绕过 mcache 和 mcentral，直接从 mheap 上获取，mheap 中有一个 `freelarge` 字段管理着超大 span。

内存的释放过程，就是分配的返过程，当 mcache 中存在较多空闲 span 时，会归还给 mcentral；而 mcentral 中存在较多空闲 span 时，会归还给 mheap；mheap 再归还给操作系统。

这种设计之所以快，主要有以下几个优势：

1. 内存分配大多时候都是在用户态完成的，不需要频繁进入内核态。
2. 每个 P 都有独立的 span cache，多个 CPU 不会并发读写同一块内存，进而减少 CPU L1 cache 的 cacheline 出现 dirty 情况，增大 cpu cache 命中率。
3. 内存碎片的问题，Go 是自己在用户态管理的，在 OS 层面看是没有碎片的，使得操作系统层面对碎片的管理压力也会降低。
4. mcache 的存在使得内存分配不需要加锁。

当然这不是没有代价的，Go 需要预申请大块内存，这必然会出现一定的浪费。

## mutex 有几种模式？

mutex有两种模式：**normal** 和 **starvation**

- 正常模式：所有goroutine按照FIFO的顺序进行锁获取，被唤醒的goroutine和新请求锁的goroutine同时进行锁获取，通常**新请求锁的goroutine更容易获取锁**(持续占有cpu)，被唤醒的goroutine则不容易获取到锁。公平性：否。
- 饥饿模式：所有尝试获取锁的goroutine进行等待排队，**新请求锁的goroutine不会进行锁获取**(禁用自旋)，而是加入队列尾部等待获取锁。公平性：是。

## Go 如何进行调度的。GMP中状态流转。

Go 里面 GMP 分别代表：G：goroutine，M：线程（真正在CPU上跑的），P：调度器。

调度器 P 是 M 和 G 之间桥梁。

Go 进行调度过程：

- 某个线程尝试创建一个新的 G，那么这个 G 就会被安排到这个线程的 G 本地队列 LRQ 中，如果 LRQ 满了，就会分配到全局队列 GRQ 中；
- 尝试获取当前线程的 M，如果无法获取，就会从空闲的 M 列表中找一个，如果空闲列表也没有，那么就创建一个 M，然后绑定 G 与 P 运行。
- 进入调度循环：
  - 找到一个合适的 G
  - 执行 G，完成以后退出


## Go 什么时候发生阻塞？阻塞时，调度器会怎么做。

- 用于**原子、互斥量或通道**操作导致 goroutine 阻塞，调度器将把当前阻塞的 goroutine 从本地运行队列 **LRQ换出**，并重新调度其它 goroutine；
- 由于**网络请求**和 **IO** 导致的阻塞，Go 提供了网络轮询器（Netpoller）来处理，后台用 epoll 等技术实现 IO多路复用。

其它回答：

- **channel阻塞**：当 goroutine 读写 channel 发生阻塞时，会调用 gopark 函数，该 G 脱离当前的 M 和 P，调度器将新的 G 放入当前 M。
- **系统调用**：当某个 G 由于系统调用陷入内核态，该 P 就会脱离当前 M，此时 P 会更新自己的状态为 Psyscall，M 与 G 相互绑定，进行系统调用。结束以后，若该 P 状态还是 Psyscall，则直接关联该 M 和 G，否则使用闲置的处理器处理该 G。
- **系统监控**：当某个 G 在 P 上运行的时间超过 10ms 时候，或者 P 处于 Psyscall 状态过长等情况就会调用 retake 函数，触发新的调度。
- **主动让出**：由于是协作式调度，该 G 会主动让出当前的 P（通过 GoSched），更新状态为 Grunnable，该 P 会调度队列中的 G 运行。

## Go中GMP有哪些状态？

G的状态：

![img](https://static.sitestack.cn/projects/qcrao-Go-Questions/e42490ec98b6897bae32eb4e4d9d637c.png)

- 空闲中(_Gidle)：表示G刚刚新建, 仍未初始化
- 待运行(_Grunnable)：表示G在运行队列中, 等待M取出并运行
- 运行中(_Grunning)：表示M正在运行这个G, 这时候M会拥有一个P
- 系统调用中(_Gsyscall)：表示M正在运行这个G发起的系统调用, 这时候M并不拥有P
- 等待中(_Gwaiting)：表示G在等待某些条件完成, 这时候G不在运行也不在运行队列中(可能在channel的等待队列中)
- 已中止(_Gdead)：表示G未被使用, 可能已执行完毕(并在freelist中等待下次复用)
- 栈复制中(_Gcopystack)：表示G正在获取一个新的栈空间并把原来的内容复制过去(用于防止GC扫描)

P的状态：

![P 的状态流转图](https://static.sitestack.cn/projects/qcrao-Go-Questions/d9472d7b0c758b90d51190d6e595f591.png)

- 空闲中(_Pidle)：当M发现无待运行的G时会进入休眠, 这时M拥有的P会变为空闲并加到空闲P链表中
- 运行中(_Prunning)：当M拥有了一个P后, 这个P的状态就会变为运行中, M运行G会使用这个P中的资源
- 系统调用中(_Psyscall)：当Go调用原生代码, 原生代码又反过来调用Go代码时, 使用的P会变为此状态
- GC停止中(_Pgcstop)：当gc停止了整个世界(STW)时, P会变为此状态
- 已中止(_Pdead)：当P的数量在运行时改变, 且数量减少时多余的P会变为此状态

M的状态：

![M 的状态流转图](https://static.sitestack.cn/projects/qcrao-Go-Questions/2718ac731f8fe5e12b24fd1656672955.png)

- **自旋线程**：处于运行状态但是没有可执行goroutine的线程，数量最多为GOMAXPROC，若是数量大于GOMAXPROC就会进入休眠。
- **非自旋线程**：处于运行状态有可执行goroutine的线程。

## GMP 能不能去掉 P 层？会怎么样？

去掉 P 会导致，当 G 进行**系统调用时候，会一直阻塞**，其它 G 无法获得 M。

## 如果有一个G一直占用资源怎么办。

如果有个goroutine一直占用资源，那么GMP模型会**从正常模式转变为饥饿模式**（类似于mutex），允许其它goroutine抢占（禁用自旋锁）。

## 若干线程一个线程发生OOM(Out of memory)会怎么办。

对于线程而言：发生内存溢出的线程会被kill，其它线程不受影响。

## goroutine什么情况会发生内存泄漏？如何避免。

在Go中内存泄露分为暂时性内存泄露和永久性内存泄露。

**暂时性内存泄露**：

- 获取长字符串中的一段导致长字符串未释放
- 获取长slice中的一段导致长slice未释放
- 在长slice新建slice导致泄漏

string相比切片少了一个容量的cap字段，可以把string当成一个只读的切片类型。获取长string或者切片中的一段内容，由于新生成的对象和老的string或者切片共用一个内存空间，会导致老的string和切片资源暂时得不到释放，造成短暂的内存泄漏。

**永久性内存泄露**：

- goroutine永久阻塞而导致泄漏
- time.Ticker未关闭导致泄漏
- 不正确使用Finalizer导致泄漏

## Go 竞态条件

所谓竞态竞争，就是当**两个或以上的goroutine访问相同资源时候，对资源进行读/写。**

比如`var a int = 0`，有两个协程分别对a+=1，我们发现最后a不一定为2.这就是竞态竞争。

通常我们可以用`go run -race xx.go`来进行检测。

解决方法是，对临界区资源上锁，或者使用原子操作(atomics)，原子操作的开销小于上锁。

## goroutine panic

### 如果若干个goroutine，有一个panic会怎么做？

有一个panic，那么剩余goroutine也会退出。

### defer可以捕获goroutine的子goroutine吗？

不可以。它们处于不同的调度器P中。对于子goroutine，正确的做法是：

1. 必须通过 defer 关键字来调用 recover()。
2. 当通过 goroutine 调用某个方法，一定要确保内部有 recover() 机制。

## gRPC是什么？

基于Go的**远程过程调用**。RPC 框架的目标就是让远程服务调用更加简单、透明，RPC 框架负责屏蔽底层的传输方式（TCP 或者 UDP）、序列化方式（XML/Json/ 二进制）和通信细节。服务调用者可以像调用本地接口一样调用远程的服务提供者，而不需要关心底层通信细节和调用过程。

## 微服务

微服务是一种开发软件的架构和组织方法，其中软件由通过明确定义的 API 进行通信的小型独立服务组成。微服务架构使应用程序更易于扩展和更快地开发，从而加速创新并缩短新功能的上市时间。

## 服务发现

主要有两种服务发现机制：**客户端发现**和**服务端发现**。

- **客户端发现模式**：当我们使用客户端发现的时候，客户端负责决定可用服务实例的网络地址并且在集群中对请求负载均衡, 客户端访问**服务登记表**，也就是一个可用服务的数据库，然后客户端使用一种**负载均衡算法**选择一个可用的服务实例然后发起请求。
- **服务端发现模式**：客户端通过**负载均衡器**向某个服务提出请求，负载均衡器查询服务注册表，并将请求转发到可用的服务实例。如同客户端发现，服务实例在服务注册表中注册或注销。

## ETCD

**etcd**是一个**高度一致**的**分布式键值存储**，它提供了一种可靠的方式来存储需要由分布式系统或机器集群访问的数据。它可以优雅地处理网络分区期间的领导者**选举**，即使在领导者节点中也可以容忍机器故障。

etcd 是用**Go语言**编写的，它具有出色的跨平台支持，小的二进制文件和强大的社区。etcd机器之间的通信通过**Raft共识算法**处理。

## GIN怎么做参数校验？

go采用validator作参数校验。

它具有以下独特功能：

- 使用验证tag或自定义validator进行跨字段Field和跨结构体验证。
- 允许切片、数组和哈希表，多维字段的任何或所有级别进行校验。
- 能够对哈希表key和value进行验证
- 通过在验证之前确定它的基础类型来处理类型接口。
- 别名验证标签，允许将多个验证映射到单个标签，以便更轻松地定义结构体上的验证
- gin web 框架的默认验证器；

## 中间件用过吗？

Middleware是Web的重要组成部分，中间件（通常）是一小段代码，它们接受一个请求，对其进行处理，每个中间件只处理一件事情，完成后将其传递给另一个中间件或最终处理程序，这样就做到了程序的解耦。

## Go解析Tag是怎么实现的？

Go解析tag采用的是**反射**。

具体来说使用reflect.ValueOf方法获取其反射值，然后获取其Type属性，之后再通过Field(i)获取第i+1个field，再.Tag获得Tag。

反射实现的原理在: `src/reflect/type.go`中。

## 你项目有优雅的启停吗？

所谓「优雅」启停就是在启动退出服务时要满足以下几个条件：

- **不可以关闭现有连接**（进程）
- 新的进程启动并「**接管**」旧进程
- 连接要**随时响应用户请求**，不可以出现拒绝请求的情况
- 停止的时候，必须**处理完既有连接**，并且**停止接收新的连接**。

为此我们必须引用**信号**来完成这些目的：

启动：

- 监听SIGHUP（在用户终端连接(正常或非正常)结束时发出）；
- 收到信号后将服务监听的文件描述符传递给新的子进程，此时新老进程同时接收请求；

退出：

- 监听SIGINT和SIGSTP和SIGQUIT等。
- 父进程停止接收新请求，等待旧请求完成（或超时）；
- 父进程退出。

实现：go1.8采用Http.Server内置的Shutdown方法支持优雅关机。 然后[fvbock/endless](https://link.zhihu.com/?target=http%3A//github.com/fvbock/endless)可以实现优雅重启。

## 持久化怎么做的？

所谓持久化就是将要保存的字符串写到硬盘等设备。

- 最简单的方式就是采用ioutil的WriteFile()方法将字符串写到磁盘上，这种方法面临**格式化**方面的问题。
- 更好的做法是将数据按照**固定协议**进行组织再进行读写，比如JSON，XML，Gob，csv等。
- 如果要考虑**高并发**和**高可用**，必须把数据放入到数据库中，比如MySQL，PostgreDB，MongoDB等。

## atomic 底层怎么实现的.

atomic源码位于`sync\atomic`。通过阅读源码可知，atomic采用**CAS**（CompareAndSwap）的方式实现的。所谓CAS就是使用了CPU中的原子性操作。在操作共享变量的时候，CAS不需要对其进行加锁，而是通过类似于乐观锁的方式进行检测，总是假设被操作的值未曾改变（即与旧值相等），并一旦确认这个假设的真实性就立即进行值替换。本质上是**不断占用CPU资源来避免加锁的开销**。

## 参考

1. [Go常见面试题【由浅入深】2022版 | 迹寒](https://zhuanlan.zhihu.com/p/471490292)
2. [字符串拼接性能及原理 | Go 语言高性能编程 | 极客兔兔](https://geektutu.com/post/hpg-string-concat.html)
3. [Go 语言中的变量究竟是分配在栈上、还是分配在堆上？逃逸分析告诉你答案 | 古明地盆](https://www.cnblogs.com/traditional/p/11505189.html)
4. [一文搞懂Go gc垃圾回收原理 | yi个俗人](https://juejin.cn/post/7111515970669117447)
5. [GC垃圾回收机制设计原理 | wx602bc012c01dd](https://blog.51cto.com/u_15107299/4309453)
6. [Go 语言内存管理（二）：Go 内存管理 | 达菲格](https://www.jianshu.com/p/7405b4e11ee2)
7. [golang高并发探究之协程一 | 深秋鸟](https://juejin.cn/post/6844903872067010573)