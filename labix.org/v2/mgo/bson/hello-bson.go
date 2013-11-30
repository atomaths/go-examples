package main

import (
	"fmt"

	"labix.org/v2/mgo/bson"
)

func main() {
	bson, _ := bson.Marshal(bson.M{"hello": "bson"})
	fmt.Printf("%q\n", bson)
}
