// From $GOROOT/src/pkg/go/parser/interface.go
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	b, err := readSource("type_switch.go", nil)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s\n", b)
}

func readSource(filename string, src interface{}) ([]byte, error) {
	if src != nil {
		switch s := src.(type) {
		case string:
			return []byte(s), nil
		case []byte:
			return s, nil
		case *bytes.Buffer:
			// is io.Reader, but src is already available in []byte form
			if s != nil {
				return s.Bytes(), nil
			}
		case io.Reader:
			var buf bytes.Buffer
			if _, err := io.Copy(&buf, s); err != nil {
				return nil, err
			}
			return buf.Bytes(), nil
		}
		return nil, errors.New("invalid source")
	}
	return ioutil.ReadFile(filename)
}
