### 방법 1

/main.go
    pkg/
        pkg.go


이런 구조일 때 main.go에서 

    import "./pkg"


### 방법 2

/main.go
/pkg.go

이런 구조일 때는 main.go에서 build constrains를 사용.

    // +build ignore

    import "."

    func main() {
        pkg.Foo()
    }
