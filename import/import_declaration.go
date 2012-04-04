// http://weekly.golang.org/ref/spec#Import_declarations
package main

// 
//import "fmt"
//import "log" // 이렇게 개별 선언은 당연히 가능


//import "fmt"; "log"; // 불가능


import (
	"fmt"; "log";  // 가능
)

// 아래처럼 하면 패키지의 init()만 호출되고 pprof의 exported identifiers에도
// 접근할 수 없음. 이런 것을 "blank identifier"라고 함.
//   An import declaration declares a dependency relation between the importing
//   and imported package. It is illegal for a package to import itself or to 
//   import a package without referring to any of its exported identifiers.
//   To import a package solely for its side-effects (initialization), 
//   use the blank identifier as explicit package name:
import _ "net/http/pprof" // 패키지의 init()만 호출됨


func main() {
	fmt.Println("fmt")
	log.Println("log")
}
