/**
 * Auth :   liubo
 * Date :   2019/11/28 10:12
 * Comment: 将目录下的所有文件都变成只读的
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var r = flag.Bool("r", true, "")
var readonly = true

func main() {
	flag.Parse()

	readonly = *r

	if strings.HasSuffix(os.Args[0], "-w.exe") {
		readonly = false
	}

	cnt := procDir("./")
	fmt.Println("成功处理数目：", cnt)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func SetWritable(filepath string) error {
	err := os.Chmod(filepath, 0222)
	return err
}

func SetReadOnly(filepath string) error {
	err := os.Chmod(filepath, 0444)
	return err
}

func procDir(dir string) int {
	var count int
	srcOriginal := dir
	err := filepath.Walk(srcOriginal, func(src string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		// 注意:这里会walk所有得文件!(而不仅仅是当前层级得)
		if !f.IsDir() {
			var err2 error
			if readonly {
				err2 = SetReadOnly(src)
			} else {
				err2 = SetWritable(src)
			}
			if err2 != nil {
				fmt.Println("error:", err2.Error())
			} else {
				count++
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return count
}