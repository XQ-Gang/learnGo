package main

import (
	"unicode"
)

// 机器人坐标问题
// 有一个机器人，给一串指令，L左转 R右转，F前进，B后退，问最后机器人的坐标。
// 最开始，机器人位于 0 0，方向为正Y。
// 可以输入重复指令n：比如 R2(LF) 这个等于指令 RLFLF。
// 问最后机器人的坐标是多少？

const (
	Left = iota
	Top
	Right
	Bottom
)

func move(cmd string, x0 int, y0 int, d0 int) (x, y, d int) {
	x, y, d = x0, y0, d0
	repeat := 0
	repeatCmd := ""
	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			repeat = repeat*10 + (int(s) - '0')
		case s == '(':
		case s == ')':
			for i := 0; i < repeat; i++ {
				x, y, d = move(repeatCmd, x, y, d)
			}
			repeat = 0
			repeatCmd = ""
		case repeat > 0:
			repeatCmd = repeatCmd + string(s)
		case s == 'L':
			d = (d - 1 + 4) % 4
		case s == 'R':
			d = (d + 1) % 4
		case s == 'F':
			switch {
			case d == Left || d == Right:
				x = x + d - 1
			case d == Top || d == Bottom:
				y = y - d + 2
			}
		case s == 'B':
			switch {
			case d == Left || d == Right:
				x = x - d + 1
			case d == Top || d == Bottom:
				y = y + d - 2
			}
		}
	}
	return
}

func main() {
	println(move("R2(LF)", 0, 0, Top))
}
