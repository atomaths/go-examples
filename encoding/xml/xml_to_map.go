package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

var xmlData = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<entry>
  <vars>
    <var key="foo">bar</var>
    <var key="foo2"><var key="hello">world</var></var>
  </vars>
</entry>
`)

type Var struct {
	Key      string `xml:"key,attr"`
	Value    string `xml:",chardata"`
	Children []Var  `xml:"var"`
}

type S struct {
	XMLName xml.Name `xml:"entry"`
	Vars    []Var    `xml:"vars>var"`
}

func toMap(vars ...Var) map[string]interface{} {
	m := make(map[string]interface{})
	for _, v := range vars {
		if len(v.Children) > 0 {
			m[v.Key] = toMap(v.Children...)
		} else {
			m[v.Key] = v.Value
		}
	}
	return m
}

func main() {
	s := &S{}
	if err := xml.Unmarshal(xmlData, s); err != nil {
		panic(err)
	}
	fmt.Printf("s: %#v\n", *s)

	m := toMap(s.Vars...)
	fmt.Printf("map: %v\n", m)

	js, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("json: %s", js)

}
