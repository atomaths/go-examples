package main

import (
        "fmt"
        "strconv"
        "time"
)

func main() {
        m := make(map[int]string)

	start := time.Now()

        for i := 0; i < 100000; i++ {
                m[i] = "value-" + strconv.Itoa(i)
        }

	elapsed := time.Since(start)
	fmt.Println("make map: ", elapsed)

	start = time.Now()
	value := m[2345]
	fmt.Println("find time: ", time.Since(start))
	fmt.Println("value: ", value)
}
