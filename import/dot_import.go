package main

import   "lib/math" // math.Sin 로 사용
import M "lib/math" // M.Sin 로 사용
import . "lib/math" // Sin 으로 바로 접근

// 이런 것을 Dot imports라고 하는데, 패키지 내에 같은 이름의 indentifier가
// 있으면 충돌이 난다. 테스트 용도 외에는 사용을 권장하지 않음.

func main() {
}
