package main

import (
	"fmt"
	"launchpad.net/mgo/bson"
)

func main() {
	bson, _ := bson.Marshal(bson.M{"hello": "world"})

	fmt.Printf("%q\n", bson)
}
