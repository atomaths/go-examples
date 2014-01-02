// 0 ~ 1초 사이의 Duration을 랜덤하게 뽑기
// Time은 나노세컨드 단위로 처리하기 때문에 1e9(1000000000)면 1초.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed value가 있어야 함. 이게 없으면 1을 seed 기본값으로 함.
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println(time.Duration(rand.Int63n(1e9)))
	}
	// start := time.Now()
	// time.Sleep(1 * time.Second)
	// fmt.Println(time.Now().Sub(start))
}
