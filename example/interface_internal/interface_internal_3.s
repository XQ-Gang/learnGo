"".T2.M1 STEXT size=16 args=0x18 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (example/interface_internal/interface_internal_3.go:10)	TEXT	"".T2.M1(SB), LEAF|NOFRAME|ABIInternal, $0-24
	0x0000 00000 (example/interface_internal/interface_internal_3.go:10)	FUNCDATA	ZR, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (example/interface_internal/interface_internal_3.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (example/interface_internal/interface_internal_3.go:10)	FUNCDATA	$5, "".T2.M1.arginfo1(SB)
	0x0000 00000 (example/interface_internal/interface_internal_3.go:10)	RET	(R30)
	0x0000 c0 03 5f d6 00 00 00 00 00 00 00 00 00 00 00 00  .._.............
"".T2.M2 STEXT size=16 args=0x18 locals=0x0 funcid=0x0 align=0x0 leaf
	0x0000 00000 (example/interface_internal/interface_internal_3.go:11)	TEXT	"".T2.M2(SB), LEAF|NOFRAME|ABIInternal, $0-24
	0x0000 00000 (example/interface_internal/interface_internal_3.go:11)	FUNCDATA	ZR, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (example/interface_internal/interface_internal_3.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (example/interface_internal/interface_internal_3.go:11)	FUNCDATA	$5, "".T2.M2.arginfo1(SB)
	0x0000 00000 (example/interface_internal/interface_internal_3.go:11)	RET	(R30)
	0x0000 c0 03 5f d6 00 00 00 00 00 00 00 00 00 00 00 00  .._.............
"".main STEXT size=272 args=0x0 locals=0x78 funcid=0x0 align=0x0
	0x0000 00000 (example/interface_internal/interface_internal_3.go:18)	TEXT	"".main(SB), ABIInternal, $128-0
	0x0000 00000 (example/interface_internal/interface_internal_3.go:18)	MOVD	16(g), R16
	0x0004 00004 (example/interface_internal/interface_internal_3.go:18)	PCDATA	$0, $-2
	0x0004 00004 (example/interface_internal/interface_internal_3.go:18)	MOVD	RSP, R17
	0x0008 00008 (example/interface_internal/interface_internal_3.go:18)	CMP	R16, R17
	0x000c 00012 (example/interface_internal/interface_internal_3.go:18)	BLS	260
	0x0010 00016 (example/interface_internal/interface_internal_3.go:18)	PCDATA	$0, $-1
	0x0010 00016 (example/interface_internal/interface_internal_3.go:18)	MOVD.W	R30, -128(RSP)
	0x0014 00020 (example/interface_internal/interface_internal_3.go:18)	MOVD	R29, -8(RSP)
	0x0018 00024 (example/interface_internal/interface_internal_3.go:18)	SUB	$8, RSP, R29
	0x001c 00028 (example/interface_internal/interface_internal_3.go:18)	FUNCDATA	ZR, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x001c 00028 (example/interface_internal/interface_internal_3.go:18)	FUNCDATA	$1, gclocals·c73a1255c0cd89e7a2b21dfb392a7db1(SB)
	0x001c 00028 (example/interface_internal/interface_internal_3.go:18)	FUNCDATA	$2, "".main.stkobj(SB)
	0x001c 00028 (example/interface_internal/interface_internal_3.go:24)	MOVD	$17, R2
	0x0020 00032 (example/interface_internal/interface_internal_3.go:24)	MOVD	R2, ""..autotmp_13-24(SP)
	0x0024 00036 (example/interface_internal/interface_internal_3.go:24)	MOVD	$go.string."hello, interface"(SB), R3
	0x002c 00044 (example/interface_internal/interface_internal_3.go:24)	MOVD	R3, ""..autotmp_13-16(SP)
	0x0030 00048 (example/interface_internal/interface_internal_3.go:24)	MOVD	$16, R4
	0x0034 00052 (example/interface_internal/interface_internal_3.go:24)	MOVD	R4, ""..autotmp_13-8(SP)
	0x0038 00056 (example/interface_internal/interface_internal_3.go:24)	MOVD	$type."".T2(SB), R0
	0x0040 00064 (example/interface_internal/interface_internal_3.go:24)	MOVD	$""..autotmp_13-24(SP), R1
	0x0044 00068 (example/interface_internal/interface_internal_3.go:24)	PCDATA	$1, ZR
	0x0044 00068 (example/interface_internal/interface_internal_3.go:24)	CALL	runtime.convT(SB)
	0x0048 00072 (example/interface_internal/interface_internal_3.go:24)	MOVD	R0, ""..autotmp_38-64(SP)
	0x004c 00076 (example/interface_internal/interface_internal_3.go:27)	MOVD	$17, R2
	0x0050 00080 (example/interface_internal/interface_internal_3.go:27)	MOVD	R2, ""..autotmp_13-24(SP)
	0x0054 00084 (example/interface_internal/interface_internal_3.go:27)	MOVD	$go.string."hello, interface"(SB), R2
	0x005c 00092 (example/interface_internal/interface_internal_3.go:27)	MOVD	R2, ""..autotmp_13-16(SP)
	0x0060 00096 (example/interface_internal/interface_internal_3.go:27)	MOVD	$16, R2
	0x0064 00100 (example/interface_internal/interface_internal_3.go:27)	MOVD	R2, ""..autotmp_13-8(SP)
	0x0068 00104 (example/interface_internal/interface_internal_3.go:27)	MOVD	$""..autotmp_13-24(SP), R1
	0x006c 00108 (example/interface_internal/interface_internal_3.go:27)	MOVD	$type."".T2(SB), R0
	0x0074 00116 (example/interface_internal/interface_internal_3.go:27)	PCDATA	$1, $1
	0x0074 00116 (example/interface_internal/interface_internal_3.go:27)	CALL	runtime.convT(SB)
	0x0078 00120 (example/interface_internal/interface_internal_3.go:27)	MOVD	R0, ""..autotmp_39-72(SP)
	0x007c 00124 (example/interface_internal/interface_internal_3.go:28)	STP	(ZR, ZR), ""..autotmp_19-40(SP)
	0x0080 00128 (example/interface_internal/interface_internal_3.go:28)	MOVD	$type."".T2(SB), R2
	0x0088 00136 (example/interface_internal/interface_internal_3.go:28)	MOVD	R2, ""..autotmp_19-40(SP)
	0x008c 00140 (example/interface_internal/interface_internal_3.go:28)	MOVD	""..autotmp_38-64(SP), R2
	0x0090 00144 (example/interface_internal/interface_internal_3.go:28)	MOVD	R2, ""..autotmp_19-32(SP)
	0x0094 00148 (<unknown line number>)	NOP
	0x0094 00148 (<unknown line number>)	PCDATA	$0, $-3
	0x0094 00148 ($GOROOT/src/fmt/print.go:274)	MOVD	os.Stdout(SB), R1
	0x00a0 00160 ($GOROOT/src/fmt/print.go:274)	PCDATA	$0, $-1
	0x00a0 00160 ($GOROOT/src/fmt/print.go:274)	MOVD	$""..autotmp_19-40(SP), R2
	0x00a4 00164 ($GOROOT/src/fmt/print.go:274)	MOVD	$1, R3
	0x00a8 00168 ($GOROOT/src/fmt/print.go:274)	MOVD	R3, R4
	0x00ac 00172 ($GOROOT/src/fmt/print.go:274)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x00b4 00180 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, $2
	0x00b4 00180 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00b8 00184 (example/interface_internal/interface_internal_3.go:29)	STP	(ZR, ZR), ""..autotmp_21-56(SP)
	0x00bc 00188 (example/interface_internal/interface_internal_3.go:29)	PCDATA	$0, $-4
	0x00bc 00188 (example/interface_internal/interface_internal_3.go:29)	MOVD	go.itab."".T2,"".NonEmptyInterface+8(SB), R2
	0x00c8 00200 (example/interface_internal/interface_internal_3.go:29)	PCDATA	$0, $-1
	0x00c8 00200 (example/interface_internal/interface_internal_3.go:29)	MOVD	R2, ""..autotmp_21-56(SP)
	0x00cc 00204 (example/interface_internal/interface_internal_3.go:29)	MOVD	""..autotmp_39-72(SP), R2
	0x00d0 00208 (example/interface_internal/interface_internal_3.go:29)	MOVD	R2, ""..autotmp_21-48(SP)
	0x00d4 00212 (<unknown line number>)	NOP
	0x00d4 00212 (<unknown line number>)	PCDATA	$0, $-3
	0x00d4 00212 ($GOROOT/src/fmt/print.go:274)	MOVD	os.Stdout(SB), R1
	0x00e0 00224 ($GOROOT/src/fmt/print.go:274)	PCDATA	$0, $-1
	0x00e0 00224 ($GOROOT/src/fmt/print.go:274)	MOVD	$go.itab.*os.File,io.Writer(SB), R0
	0x00e8 00232 ($GOROOT/src/fmt/print.go:274)	MOVD	$""..autotmp_21-56(SP), R2
	0x00ec 00236 ($GOROOT/src/fmt/print.go:274)	MOVD	$1, R3
	0x00f0 00240 ($GOROOT/src/fmt/print.go:274)	MOVD	R3, R4
	0x00f4 00244 ($GOROOT/src/fmt/print.go:274)	PCDATA	$1, ZR
	0x00f4 00244 ($GOROOT/src/fmt/print.go:274)	CALL	fmt.Fprintln(SB)
	0x00f8 00248 (example/interface_internal/interface_internal_3.go:30)	MOVD	-8(RSP), R29
	0x00fc 00252 (example/interface_internal/interface_internal_3.go:30)	MOVD.P	128(RSP), R30
	0x0100 00256 (example/interface_internal/interface_internal_3.go:30)	RET	(R30)
	0x0104 00260 (example/interface_internal/interface_internal_3.go:30)	NOP
	0x0104 00260 (example/interface_internal/interface_internal_3.go:18)	PCDATA	$1, $-1
	0x0104 00260 (example/interface_internal/interface_internal_3.go:18)	PCDATA	$0, $-2
	0x0104 00260 (example/interface_internal/interface_internal_3.go:18)	MOVD	R30, R3
	0x0108 00264 (example/interface_internal/interface_internal_3.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x010c 00268 (example/interface_internal/interface_internal_3.go:18)	PCDATA	$0, $-1
	0x010c 00268 (example/interface_internal/interface_internal_3.go:18)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb c9 07 00 54  ..@.....?......T
	0x0010 fe 0f 18 f8 fd 83 1f f8 fd 23 00 d1 22 02 80 d2  .........#.."...
	0x0020 e2 33 00 f9 03 00 00 90 63 00 00 91 e3 37 00 f9  .3......c....7..
	0x0030 e4 03 7c b2 e4 3b 00 f9 00 00 00 90 00 00 00 91  ..|..;..........
	0x0040 e1 83 01 91 00 00 00 94 e0 1f 00 f9 22 02 80 d2  ............"...
	0x0050 e2 33 00 f9 02 00 00 90 42 00 00 91 e2 37 00 f9  .3......B....7..
	0x0060 e2 03 7c b2 e2 3b 00 f9 e1 83 01 91 00 00 00 90  ..|..;..........
	0x0070 00 00 00 91 00 00 00 94 e0 1b 00 f9 ff 7f 05 a9  ................
	0x0080 02 00 00 90 42 00 00 91 e2 2b 00 f9 e2 1f 40 f9  ....B....+....@.
	0x0090 e2 2f 00 f9 1b 00 00 90 7b 03 00 91 61 03 40 f9  ./......{...a.@.
	0x00a0 e2 43 01 91 e3 03 40 b2 e4 03 03 aa 00 00 00 90  .C....@.........
	0x00b0 00 00 00 91 00 00 00 94 ff 7f 04 a9 1b 00 00 90  ................
	0x00c0 7b 03 00 91 62 03 40 f9 e2 23 00 f9 e2 1b 40 f9  {...b.@..#....@.
	0x00d0 e2 27 00 f9 1b 00 00 90 7b 03 00 91 61 03 40 f9  .'......{...a.@.
	0x00e0 00 00 00 90 00 00 00 91 e2 03 01 91 e3 03 40 b2  ..............@.
	0x00f0 e4 03 03 aa 00 00 00 94 fd 83 5f f8 fe 07 48 f8  .........._...H.
	0x0100 c0 03 5f d6 e3 03 1e aa 00 00 00 94 bd ff ff 17  .._.............
	rel 0+0 t=23 type."".T2+0
	rel 0+0 t=23 type."".T2+0
	rel 0+0 t=23 type.*os.File+0
	rel 0+0 t=23 type.*os.File+0
	rel 36+8 t=3 go.string."hello, interface"+0
	rel 56+8 t=3 type."".T2+0
	rel 68+4 t=9 runtime.convT+0
	rel 84+8 t=3 go.string."hello, interface"+0
	rel 108+8 t=3 type."".T2+0
	rel 116+4 t=9 runtime.convT+0
	rel 128+8 t=3 type."".T2+0
	rel 148+8 t=3 os.Stdout+0
	rel 172+8 t=3 go.itab.*os.File,io.Writer+0
	rel 180+4 t=9 fmt.Fprintln+0
	rel 188+8 t=3 go.itab."".T2,"".NonEmptyInterface+8
	rel 212+8 t=3 os.Stdout+0
	rel 224+8 t=3 go.itab.*os.File,io.Writer+0
	rel 244+4 t=9 fmt.Fprintln+0
	rel 264+4 t=9 runtime.morestack_noctxt+0
"".(*T2).M1 STEXT dupok size=112 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*T2).M1(SB), DUPOK|WRAPPER|ABIInternal, $16-8
	0x0000 00000 (<autogenerated>:1)	MOVD	16(g), R16
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	MOVD	RSP, R17
	0x0008 00008 (<autogenerated>:1)	CMP	R16, R17
	0x000c 00012 (<autogenerated>:1)	BLS	60
	0x0010 00016 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0010 00016 (<autogenerated>:1)	MOVD.W	R30, -16(RSP)
	0x0014 00020 (<autogenerated>:1)	MOVD	R29, -8(RSP)
	0x0018 00024 (<autogenerated>:1)	SUB	$8, RSP, R29
	0x001c 00028 (<autogenerated>:1)	MOVD	32(g), R16
	0x0020 00032 (<autogenerated>:1)	CBNZ	R16, 80
	0x0024 00036 (<autogenerated>:1)	NOP
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	ZR, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	$5, "".(*T2).M1.arginfo1(SB)
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	$6, "".(*T2).M1.argliveinfo(SB)
	0x0024 00036 (<autogenerated>:1)	PCDATA	$3, $1
	0x0024 00036 (<autogenerated>:1)	CBZ	R0, 52
	0x0028 00040 (<autogenerated>:1)	MOVD	-8(RSP), R29
	0x002c 00044 (<autogenerated>:1)	MOVD.P	16(RSP), R30
	0x0030 00048 (<autogenerated>:1)	RET	(R30)
	0x0034 00052 (<autogenerated>:1)	PCDATA	$1, $1
	0x0034 00052 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0038 00056 (<autogenerated>:1)	HINT	ZR
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVD	R0, 8(RSP)
	0x0040 00064 (<autogenerated>:1)	MOVD	R30, R3
	0x0044 00068 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0048 00072 (<autogenerated>:1)	MOVD	8(RSP), R0
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $-1
	0x004c 00076 (<autogenerated>:1)	JMP	0
	0x0050 00080 (<autogenerated>:1)	MOVD	(R16), R17
	0x0054 00084 (<autogenerated>:1)	ADD	$24, RSP, R20
	0x0058 00088 (<autogenerated>:1)	CMP	R17, R20
	0x005c 00092 (<autogenerated>:1)	BNE	36
	0x0060 00096 (<autogenerated>:1)	ADD	$8, RSP, R20
	0x0064 00100 (<autogenerated>:1)	MOVD	R20, (R16)
	0x0068 00104 (<autogenerated>:1)	JMP	36
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb 89 01 00 54  ..@.....?......T
	0x0010 fe 0f 1f f8 fd 83 1f f8 fd 23 00 d1 90 13 40 f9  .........#....@.
	0x0020 90 01 00 b5 80 00 00 b4 fd 83 5f f8 fe 07 41 f8  .........._...A.
	0x0030 c0 03 5f d6 00 00 00 94 1f 20 03 d5 e0 07 00 f9  .._...... ......
	0x0040 e3 03 1e aa 00 00 00 94 e0 07 40 f9 ed ff ff 17  ..........@.....
	0x0050 11 02 40 f9 f4 63 00 91 9f 02 11 eb 41 fe ff 54  ..@..c......A..T
	0x0060 f4 23 00 91 14 02 00 f9 ef ff ff 17 00 00 00 00  .#..............
	rel 52+4 t=9 runtime.panicwrap+0
	rel 68+4 t=9 runtime.morestack_noctxt+0
"".(*T2).M2 STEXT dupok size=112 args=0x8 locals=0x8 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*T2).M2(SB), DUPOK|WRAPPER|ABIInternal, $16-8
	0x0000 00000 (<autogenerated>:1)	MOVD	16(g), R16
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	MOVD	RSP, R17
	0x0008 00008 (<autogenerated>:1)	CMP	R16, R17
	0x000c 00012 (<autogenerated>:1)	BLS	60
	0x0010 00016 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0010 00016 (<autogenerated>:1)	MOVD.W	R30, -16(RSP)
	0x0014 00020 (<autogenerated>:1)	MOVD	R29, -8(RSP)
	0x0018 00024 (<autogenerated>:1)	SUB	$8, RSP, R29
	0x001c 00028 (<autogenerated>:1)	MOVD	32(g), R16
	0x0020 00032 (<autogenerated>:1)	CBNZ	R16, 80
	0x0024 00036 (<autogenerated>:1)	NOP
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	ZR, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	$5, "".(*T2).M2.arginfo1(SB)
	0x0024 00036 (<autogenerated>:1)	FUNCDATA	$6, "".(*T2).M2.argliveinfo(SB)
	0x0024 00036 (<autogenerated>:1)	PCDATA	$3, $1
	0x0024 00036 (<autogenerated>:1)	CBZ	R0, 52
	0x0028 00040 (<autogenerated>:1)	MOVD	-8(RSP), R29
	0x002c 00044 (<autogenerated>:1)	MOVD.P	16(RSP), R30
	0x0030 00048 (<autogenerated>:1)	RET	(R30)
	0x0034 00052 (<autogenerated>:1)	PCDATA	$1, $1
	0x0034 00052 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0038 00056 (<autogenerated>:1)	HINT	ZR
	0x003c 00060 (<autogenerated>:1)	NOP
	0x003c 00060 (<autogenerated>:1)	PCDATA	$1, $-1
	0x003c 00060 (<autogenerated>:1)	PCDATA	$0, $-2
	0x003c 00060 (<autogenerated>:1)	MOVD	R0, 8(RSP)
	0x0040 00064 (<autogenerated>:1)	MOVD	R30, R3
	0x0044 00068 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0048 00072 (<autogenerated>:1)	MOVD	8(RSP), R0
	0x004c 00076 (<autogenerated>:1)	PCDATA	$0, $-1
	0x004c 00076 (<autogenerated>:1)	JMP	0
	0x0050 00080 (<autogenerated>:1)	MOVD	(R16), R17
	0x0054 00084 (<autogenerated>:1)	ADD	$24, RSP, R20
	0x0058 00088 (<autogenerated>:1)	CMP	R17, R20
	0x005c 00092 (<autogenerated>:1)	BNE	36
	0x0060 00096 (<autogenerated>:1)	ADD	$8, RSP, R20
	0x0064 00100 (<autogenerated>:1)	MOVD	R20, (R16)
	0x0068 00104 (<autogenerated>:1)	JMP	36
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb 89 01 00 54  ..@.....?......T
	0x0010 fe 0f 1f f8 fd 83 1f f8 fd 23 00 d1 90 13 40 f9  .........#....@.
	0x0020 90 01 00 b5 80 00 00 b4 fd 83 5f f8 fe 07 41 f8  .........._...A.
	0x0030 c0 03 5f d6 00 00 00 94 1f 20 03 d5 e0 07 00 f9  .._...... ......
	0x0040 e3 03 1e aa 00 00 00 94 e0 07 40 f9 ed ff ff 17  ..........@.....
	0x0050 11 02 40 f9 f4 63 00 91 9f 02 11 eb 41 fe ff 54  ..@..c......A..T
	0x0060 f4 23 00 91 14 02 00 f9 ef ff ff 17 00 00 00 00  .#..............
	rel 52+4 t=9 runtime.panicwrap+0
	rel 68+4 t=9 runtime.morestack_noctxt+0
"".NonEmptyInterface.M1 STEXT dupok size=128 args=0x10 locals=0x18 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".NonEmptyInterface.M1(SB), DUPOK|WRAPPER|ABIInternal, $32-16
	0x0000 00000 (<autogenerated>:1)	MOVD	16(g), R16
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	MOVD	RSP, R17
	0x0008 00008 (<autogenerated>:1)	CMP	R16, R17
	0x000c 00012 (<autogenerated>:1)	BLS	68
	0x0010 00016 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0010 00016 (<autogenerated>:1)	MOVD.W	R30, -32(RSP)
	0x0014 00020 (<autogenerated>:1)	MOVD	R29, -8(RSP)
	0x0018 00024 (<autogenerated>:1)	SUB	$8, RSP, R29
	0x001c 00028 (<autogenerated>:1)	MOVD	32(g), R16
	0x0020 00032 (<autogenerated>:1)	CBNZ	R16, 96
	0x0024 00036 (<autogenerated>:1)	NOP
	0x0024 00036 (<autogenerated>:1)	MOVD	R0, ""..this(FP)
	0x0028 00040 (<autogenerated>:1)	MOVD	R1, ""..this+8(FP)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	ZR, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	$5, "".NonEmptyInterface.M1.arginfo1(SB)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	$6, "".NonEmptyInterface.M1.argliveinfo(SB)
	0x002c 00044 (<autogenerated>:1)	PCDATA	$3, $1
	0x002c 00044 (<autogenerated>:1)	MOVD	24(R0), R2
	0x0030 00048 (<autogenerated>:1)	MOVD	R1, R0
	0x0034 00052 (<autogenerated>:1)	PCDATA	$1, $1
	0x0034 00052 (<autogenerated>:1)	CALL	(R2)
	0x0038 00056 (<autogenerated>:1)	MOVD	-8(RSP), R29
	0x003c 00060 (<autogenerated>:1)	MOVD.P	32(RSP), R30
	0x0040 00064 (<autogenerated>:1)	RET	(R30)
	0x0044 00068 (<autogenerated>:1)	NOP
	0x0044 00068 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0044 00068 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0044 00068 (<autogenerated>:1)	MOVD	R0, 8(RSP)
	0x0048 00072 (<autogenerated>:1)	MOVD	R1, 16(RSP)
	0x004c 00076 (<autogenerated>:1)	MOVD	R30, R3
	0x0050 00080 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0054 00084 (<autogenerated>:1)	MOVD	8(RSP), R0
	0x0058 00088 (<autogenerated>:1)	MOVD	16(RSP), R1
	0x005c 00092 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005c 00092 (<autogenerated>:1)	JMP	0
	0x0060 00096 (<autogenerated>:1)	MOVD	(R16), R17
	0x0064 00100 (<autogenerated>:1)	ADD	$40, RSP, R20
	0x0068 00104 (<autogenerated>:1)	CMP	R17, R20
	0x006c 00108 (<autogenerated>:1)	BNE	36
	0x0070 00112 (<autogenerated>:1)	ADD	$8, RSP, R20
	0x0074 00116 (<autogenerated>:1)	MOVD	R20, (R16)
	0x0078 00120 (<autogenerated>:1)	JMP	36
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb c9 01 00 54  ..@.....?......T
	0x0010 fe 0f 1e f8 fd 83 1f f8 fd 23 00 d1 90 13 40 f9  .........#....@.
	0x0020 10 02 00 b5 e0 17 00 f9 e1 1b 00 f9 02 0c 40 f9  ..............@.
	0x0030 e0 03 01 aa 40 00 3f d6 fd 83 5f f8 fe 07 42 f8  ....@.?..._...B.
	0x0040 c0 03 5f d6 e0 07 00 f9 e1 0b 00 f9 e3 03 1e aa  .._.............
	0x0050 00 00 00 94 e0 07 40 f9 e1 0b 40 f9 e9 ff ff 17  ......@...@.....
	0x0060 11 02 40 f9 f4 a3 00 91 9f 02 11 eb c1 fd ff 54  ..@............T
	0x0070 f4 23 00 91 14 02 00 f9 eb ff ff 17 00 00 00 00  .#..............
	rel 0+0 t=24 type."".NonEmptyInterface+96
	rel 52+0 t=10 +0
	rel 80+4 t=9 runtime.morestack_noctxt+0
"".NonEmptyInterface.M2 STEXT dupok size=128 args=0x10 locals=0x18 funcid=0x15 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".NonEmptyInterface.M2(SB), DUPOK|WRAPPER|ABIInternal, $32-16
	0x0000 00000 (<autogenerated>:1)	MOVD	16(g), R16
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	MOVD	RSP, R17
	0x0008 00008 (<autogenerated>:1)	CMP	R16, R17
	0x000c 00012 (<autogenerated>:1)	BLS	68
	0x0010 00016 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0010 00016 (<autogenerated>:1)	MOVD.W	R30, -32(RSP)
	0x0014 00020 (<autogenerated>:1)	MOVD	R29, -8(RSP)
	0x0018 00024 (<autogenerated>:1)	SUB	$8, RSP, R29
	0x001c 00028 (<autogenerated>:1)	MOVD	32(g), R16
	0x0020 00032 (<autogenerated>:1)	CBNZ	R16, 96
	0x0024 00036 (<autogenerated>:1)	NOP
	0x0024 00036 (<autogenerated>:1)	MOVD	R0, ""..this(FP)
	0x0028 00040 (<autogenerated>:1)	MOVD	R1, ""..this+8(FP)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	ZR, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	$5, "".NonEmptyInterface.M2.arginfo1(SB)
	0x002c 00044 (<autogenerated>:1)	FUNCDATA	$6, "".NonEmptyInterface.M2.argliveinfo(SB)
	0x002c 00044 (<autogenerated>:1)	PCDATA	$3, $1
	0x002c 00044 (<autogenerated>:1)	MOVD	32(R0), R2
	0x0030 00048 (<autogenerated>:1)	MOVD	R1, R0
	0x0034 00052 (<autogenerated>:1)	PCDATA	$1, $1
	0x0034 00052 (<autogenerated>:1)	CALL	(R2)
	0x0038 00056 (<autogenerated>:1)	MOVD	-8(RSP), R29
	0x003c 00060 (<autogenerated>:1)	MOVD.P	32(RSP), R30
	0x0040 00064 (<autogenerated>:1)	RET	(R30)
	0x0044 00068 (<autogenerated>:1)	NOP
	0x0044 00068 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0044 00068 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0044 00068 (<autogenerated>:1)	MOVD	R0, 8(RSP)
	0x0048 00072 (<autogenerated>:1)	MOVD	R1, 16(RSP)
	0x004c 00076 (<autogenerated>:1)	MOVD	R30, R3
	0x0050 00080 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0054 00084 (<autogenerated>:1)	MOVD	8(RSP), R0
	0x0058 00088 (<autogenerated>:1)	MOVD	16(RSP), R1
	0x005c 00092 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005c 00092 (<autogenerated>:1)	JMP	0
	0x0060 00096 (<autogenerated>:1)	MOVD	(R16), R17
	0x0064 00100 (<autogenerated>:1)	ADD	$40, RSP, R20
	0x0068 00104 (<autogenerated>:1)	CMP	R17, R20
	0x006c 00108 (<autogenerated>:1)	BNE	36
	0x0070 00112 (<autogenerated>:1)	ADD	$8, RSP, R20
	0x0074 00116 (<autogenerated>:1)	MOVD	R20, (R16)
	0x0078 00120 (<autogenerated>:1)	JMP	36
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb c9 01 00 54  ..@.....?......T
	0x0010 fe 0f 1e f8 fd 83 1f f8 fd 23 00 d1 90 13 40 f9  .........#....@.
	0x0020 10 02 00 b5 e0 17 00 f9 e1 1b 00 f9 02 10 40 f9  ..............@.
	0x0030 e0 03 01 aa 40 00 3f d6 fd 83 5f f8 fe 07 42 f8  ....@.?..._...B.
	0x0040 c0 03 5f d6 e0 07 00 f9 e1 0b 00 f9 e3 03 1e aa  .._.............
	0x0050 00 00 00 94 e0 07 40 f9 e1 0b 40 f9 e9 ff ff 17  ......@...@.....
	0x0060 11 02 40 f9 f4 a3 00 91 9f 02 11 eb c1 fd ff 54  ..@............T
	0x0070 f4 23 00 91 14 02 00 f9 eb ff ff 17 00 00 00 00  .#..............
	rel 0+0 t=24 type."".NonEmptyInterface+104
	rel 52+0 t=10 +0
	rel 80+4 t=9 runtime.morestack_noctxt+0
type..eq."".T2 STEXT dupok size=128 args=0x10 locals=0x28 funcid=0x0 align=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq."".T2(SB), DUPOK|ABIInternal, $48-16
	0x0000 00000 (<autogenerated>:1)	MOVD	16(g), R16
	0x0004 00004 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0004 00004 (<autogenerated>:1)	MOVD	RSP, R17
	0x0008 00008 (<autogenerated>:1)	CMP	R16, R17
	0x000c 00012 (<autogenerated>:1)	BLS	96
	0x0010 00016 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0010 00016 (<autogenerated>:1)	MOVD.W	R30, -48(RSP)
	0x0014 00020 (<autogenerated>:1)	MOVD	R29, -8(RSP)
	0x0018 00024 (<autogenerated>:1)	SUB	$8, RSP, R29
	0x001c 00028 (<autogenerated>:1)	FUNCDATA	ZR, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x001c 00028 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001c 00028 (<autogenerated>:1)	FUNCDATA	$5, type..eq."".T2.arginfo1(SB)
	0x001c 00028 (<autogenerated>:1)	FUNCDATA	$6, type..eq."".T2.argliveinfo(SB)
	0x001c 00028 (<autogenerated>:1)	PCDATA	$3, $1
	0x001c 00028 (<autogenerated>:1)	MOVD	(R1), R3
	0x0020 00032 (<autogenerated>:1)	MOVD	(R0), R4
	0x0024 00036 (<autogenerated>:1)	CMP	R4, R3
	0x0028 00040 (<autogenerated>:1)	BNE	80
	0x002c 00044 (<autogenerated>:1)	MOVD	16(R0), R2
	0x0030 00048 (<autogenerated>:1)	MOVD	8(R1), R3
	0x0034 00052 (<autogenerated>:1)	MOVD	8(R0), R0
	0x0038 00056 (<autogenerated>:1)	MOVD	16(R1), R4
	0x003c 00060 (<autogenerated>:1)	CMP	R4, R2
	0x0040 00064 (<autogenerated>:1)	BNE	80
	0x0044 00068 (<autogenerated>:1)	MOVD	R3, R1
	0x0048 00072 (<autogenerated>:1)	PCDATA	$1, $1
	0x0048 00072 (<autogenerated>:1)	CALL	runtime.memequal(SB)
	0x004c 00076 (<autogenerated>:1)	JMP	84
	0x0050 00080 (<autogenerated>:1)	MOVD	ZR, R0
	0x0054 00084 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0054 00084 (<autogenerated>:1)	MOVD	-8(RSP), R29
	0x0058 00088 (<autogenerated>:1)	MOVD.P	48(RSP), R30
	0x005c 00092 (<autogenerated>:1)	RET	(R30)
	0x0060 00096 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0060 00096 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0060 00096 (<autogenerated>:1)	MOVD	R0, 8(RSP)
	0x0064 00100 (<autogenerated>:1)	MOVD	R1, 16(RSP)
	0x0068 00104 (<autogenerated>:1)	MOVD	R30, R3
	0x006c 00108 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0070 00112 (<autogenerated>:1)	MOVD	8(RSP), R0
	0x0074 00116 (<autogenerated>:1)	MOVD	16(RSP), R1
	0x0078 00120 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0078 00120 (<autogenerated>:1)	JMP	0
	0x0000 90 0b 40 f9 f1 03 00 91 3f 02 10 eb a9 02 00 54  ..@.....?......T
	0x0010 fe 0f 1d f8 fd 83 1f f8 fd 23 00 d1 23 00 40 f9  .........#..#.@.
	0x0020 04 00 40 f9 7f 00 04 eb 41 01 00 54 02 08 40 f9  ..@.....A..T..@.
	0x0030 23 04 40 f9 00 04 40 f9 24 08 40 f9 5f 00 04 eb  #.@...@.$.@._...
	0x0040 81 00 00 54 e1 03 03 aa 00 00 00 94 02 00 00 14  ...T............
	0x0050 00 00 80 d2 fd 83 5f f8 fe 07 43 f8 c0 03 5f d6  ......_...C..._.
	0x0060 e0 07 00 f9 e1 0b 00 f9 e3 03 1e aa 00 00 00 94  ................
	0x0070 e0 07 40 f9 e1 0b 40 f9 e2 ff ff 17 00 00 00 00  ..@...@.........
	rel 72+4 t=9 runtime.memequal+0
	rel 108+4 t=9 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
go.info.fmt.Println$abstract SDWARFABSFCN dupok size=42
	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
	rel 0+0 t=22 type.[]interface {}+0
	rel 0+0 t=22 type.error+0
	rel 0+0 t=22 type.int+0
	rel 19+4 t=31 go.info.[]interface {}+0
	rel 27+4 t=31 go.info.int+0
	rel 37+4 t=31 go.info.error+0
""..inittask SNOPTRDATA size=32
	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 fmt..inittask+0
go.string."hello, interface" SRODATA dupok size=16
	0x0000 68 65 6c 6c 6f 2c 20 69 6e 74 65 72 66 61 63 65  hello, interface
go.info."".T2.M1$abstract SDWARFABSFCN dupok size=11
	0x0000 05 2e 54 32 2e 4d 31 00 01 01 00                 ..T2.M1....
go.info."".T2.M2$abstract SDWARFABSFCN dupok size=11
	0x0000 05 2e 54 32 2e 4d 32 00 01 01 00                 ..T2.M2....
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*func()- SRODATA dupok size=9
	0x0000 00 07 2a 66 75 6e 63 28 29                       ..*func()
type.*func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b 90 75 1b 08 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func()-+0
	rel 44+4 t=-32763 type.*func()+0
runtime.interequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.interequal+0
type..namedata.*main.NonEmptyInterface. SRODATA dupok size=25
	0x0000 01 17 2a 6d 61 69 6e 2e 4e 6f 6e 45 6d 70 74 79  ..*main.NonEmpty
	0x0010 49 6e 74 65 72 66 61 63 65                       Interface
type.*"".NonEmptyInterface SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 af 20 d1 08 08 08 36 00 00 00 00 00 00 00 00  .. ....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.NonEmptyInterface.+0
	rel 48+8 t=1 type."".NonEmptyInterface+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type..namedata.M1. SRODATA dupok size=4
	0x0000 01 02 4d 31                                      ..M1
type..namedata.M2. SRODATA dupok size=4
	0x0000 01 02 4d 32                                      ..M2
type."".NonEmptyInterface SRODATA dupok size=112
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 4a 37 15 f3 07 08 08 14 00 00 00 00 00 00 00 00  J7..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.interequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.NonEmptyInterface.+0
	rel 44+4 t=5 type.*"".NonEmptyInterface+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".NonEmptyInterface+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+4 t=5 type..namedata.M1.+0
	rel 100+4 t=5 type.func()+0
	rel 104+4 t=5 type..namedata.M2.+0
	rel 108+4 t=5 type.func()+0
type..eqfunc."".T2 SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq."".T2+0
type..namedata.*main.T2. SRODATA dupok size=10
	0x0000 01 08 2a 6d 61 69 6e 2e 54 32                    ..*main.T2
type..namedata.*func(*main.T2)- SRODATA dupok size=17
	0x0000 00 0f 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 54 32  ..*func(*main.T2
	0x0010 29                                               )
type.*func(*"".T2) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 50 02 f4 c4 08 08 08 36 00 00 00 00 00 00 00 00  P......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.T2)-+0
	rel 48+8 t=1 type.func(*"".T2)+0
type.func(*"".T2) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 44 0d e1 9f 02 08 08 33 00 00 00 00 00 00 00 00  D......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.T2)-+0
	rel 44+4 t=-32763 type.*func(*"".T2)+0
	rel 56+8 t=1 type.*"".T2+0
type.*"".T2 SRODATA dupok size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 82 b1 e3 f3 09 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 02 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.T2.+0
	rel 48+8 t=1 type."".T2+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.M1.+0
	rel 76+4 t=26 type.func()+0
	rel 80+4 t=26 "".(*T2).M1+0
	rel 84+4 t=26 "".(*T2).M1+0
	rel 88+4 t=5 type..namedata.M2.+0
	rel 92+4 t=26 type.func()+0
	rel 96+4 t=26 "".(*T2).M2+0
	rel 100+4 t=26 "".(*T2).M2+0
type..namedata.*func(main.T2)- SRODATA dupok size=16
	0x0000 00 0e 2a 66 75 6e 63 28 6d 61 69 6e 2e 54 32 29  ..*func(main.T2)
type.*func("".T2) SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a0 0b cd f4 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.T2)-+0
	rel 48+8 t=1 type.func("".T2)+0
type.func("".T2) SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f7 a3 c3 30 02 08 08 33 00 00 00 00 00 00 00 00  ...0...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.T2)-+0
	rel 44+4 t=-32763 type.*func("".T2)+0
	rel 56+8 t=1 type."".T2+0
type..namedata.n- SRODATA dupok size=3
	0x0000 00 01 6e                                         ..n
type..namedata.s- SRODATA dupok size=3
	0x0000 00 01 73                                         ..s
type."".T2 SRODATA dupok size=176
	0x0000 18 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 c1 86 7b 79 07 08 08 19 00 00 00 00 00 00 00 00  ..{y............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 02 00 02 00 40 00 00 00 00 00 00 00  ........@.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 type..eqfunc."".T2+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*main.T2.+0
	rel 44+4 t=5 type.*"".T2+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".T2+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.n-+0
	rel 104+8 t=1 type.int+0
	rel 120+8 t=1 type..namedata.s-+0
	rel 128+8 t=1 type.string+0
	rel 144+4 t=5 type..namedata.M1.+0
	rel 148+4 t=26 type.func()+0
	rel 152+4 t=26 "".(*T2).M1+0
	rel 156+4 t=26 "".T2.M1+0
	rel 160+4 t=5 type..namedata.M2.+0
	rel 164+4 t=26 type.func()+0
	rel 168+4 t=26 "".(*T2).M2+0
	rel 172+4 t=26 "".T2.M2+0
go.itab."".T2,"".NonEmptyInterface SRODATA dupok size=40
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 c1 86 7b 79 00 00 00 00 00 00 00 00 00 00 00 00  ..{y............
	0x0020 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type."".NonEmptyInterface+0
	rel 8+8 t=1 type."".T2+0
	rel 24+8 t=-32767 "".(*T2).M1+0
	rel 32+8 t=-32767 "".(*T2).M2+0
go.itab.*os.File,io.Writer SRODATA dupok size=32
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 44 b5 f3 33 00 00 00 00 00 00 00 00 00 00 00 00  D..3............
	rel 0+8 t=1 type.io.Writer+0
	rel 8+8 t=1 type.*os.File+0
	rel 24+8 t=-32767 os.(*File).Write+0
runtime.nilinterequal·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.nilinterequal+0
type..namedata.*interface {}- SRODATA dupok size=15
	0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.nilinterequal·f+0
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=-32763 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=17
	0x0000 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20 7b  ..*[]interface {
	0x0010 7d                                               }
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 08 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=-32763 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..importpath.fmt. SRODATA dupok size=5
	0x0000 00 03 66 6d 74                                   ..fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
"".T2.M1.arginfo1 SRODATA static dupok size=11
	0x0000 fe 00 08 fe 08 08 10 08 fd fd ff                 ...........
"".T2.M2.arginfo1 SRODATA static dupok size=11
	0x0000 fe 00 08 fe 08 08 10 08 fd fd ff                 ...........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·c73a1255c0cd89e7a2b21dfb392a7db1 SRODATA dupok size=14
	0x0000 03 00 00 00 09 00 00 00 00 00 02 00 01 00        ..............
"".main.stkobj SRODATA static size=56
	0x0000 03 00 00 00 00 00 00 00 c8 ff ff ff 10 00 00 00  ................
	0x0010 10 00 00 00 00 00 00 00 d8 ff ff ff 10 00 00 00  ................
	0x0020 10 00 00 00 00 00 00 00 e8 ff ff ff 18 00 00 00  ................
	0x0030 10 00 00 00 00 00 00 00                          ........
	rel 20+4 t=5 runtime.gcbits.02+0
	rel 36+4 t=5 runtime.gcbits.02+0
	rel 52+4 t=5 runtime.gcbits.02+0
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
"".(*T2).M1.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*T2).M1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
"".(*T2).M2.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
"".(*T2).M2.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·09cf9819fc716118c209c2d2155a3632 SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 02 00                    ..........
"".NonEmptyInterface.M1.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".NonEmptyInterface.M1.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
"".NonEmptyInterface.M2.arginfo1 SRODATA static dupok size=7
	0x0000 fe 00 08 08 08 fd ff                             .......
"".NonEmptyInterface.M2.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
type..eq."".T2.arginfo1 SRODATA static dupok size=5
	0x0000 00 08 08 08 ff                                   .....
type..eq."".T2.argliveinfo SRODATA static dupok size=2
	0x0000 00 00                                            ..
