package main

import (
	"os"
)

// 问题：当周围函数返回时，defer会调度一个函数调用。
// 在这种情况下，延迟调用不是在每次循环迭代期间执行，而是在readFiles函数返回时执行。
// 如果readFiles不返回，文件描述符将永远保持打开状态，从而导致泄漏
func readFiles1(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		// Do something with file
	}
	return nil
}

// 解决方案1：围绕defer创建另一个在每次迭代期间调用的周围函数。
func readFiles2(ch <-chan string) error {
	for path := range ch {
		if err := readFile(path); err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	// Do something with file
	return nil
}

// 解决方案2：在每次迭代中添加另一个周围函数（闭包）来执行延迟调用。
func readFiles3(ch <-chan string) error {
	for path := range ch {
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()

			// Do something with file
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
