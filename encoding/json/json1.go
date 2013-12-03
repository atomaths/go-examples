// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
        "encoding/json"
        "fmt"
//        "reflect"
)

type Folder struct {
	Id string
	Title string
	Time int64
	Items []Item
}

type Item struct {
	Id string
	Content string
	Time int64
}

type Message struct {
        Name string
        Body string
        Time int64
}

func Encode() {
	item := Item{"", "테스트아이템", 1332751113}
	folder := Folder{"", "폴더제목", 1332751113, append([]Item{}, item)}
	//folder := Folder{"", "폴더제목", 1234, []Item{}}

	b, err := json.Marshal(folder)
	b2, err := json.MarshalIndent(folder, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", b2)


//	m := Message{"Alice", "Hello", 1294706395881547000}
//        b, err := json.Marshal(m)
//
//        if err != nil {
//                panic(err)
//        }
//
//        expected := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
//        if !reflect.DeepEqual(b, expected) {
//                log.Panicf("Error marshalling %q, expected %q, got %q.", m, expected, b)
//        }
}

func main() {
	Encode()
}
