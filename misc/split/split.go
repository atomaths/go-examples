// TODO: 다시 합치기
// 이것 말고 command line에서 파이프로 연결해서 나누고 합치는 방법이 있었는데
// 못 찾겠음.
package main

import (
	"fmt"
	"io"
	"os"
)

func SliceUpFile(name string, bytesPerFile int) error {
	src, err := os.Open(name)
	if err != nil {
		return err
	}
	defer src.Close()

	i := 0

	for {
		dst, err := os.Create(fmt.Sprint(name, ".", i))
		if err != nil {
			return err
		}
		_, err = io.CopyN(dst, src, int64(bytesPerFile))
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		i++
	}

	return nil
}

func main() {
	// TODO: flag로 받아야 함
	err := SliceUpFile(os.Args[1], 102400)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
