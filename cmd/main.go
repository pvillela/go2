// Code generated by go2go; DO NOT EDIT.


//line ./cmd/main.go2:1
package main

//line ./cmd/main.go2:1
import (
//line ./cmd/main.go2:1
 "fmt"
//line ./cmd/main.go2:1
 "github.com/pvillela/go2/foo"
//line ./cmd/main.go2:1
)

//line ./cmd/main.go2:10
func main() {
//line ./cmd/main.go2:10
 instantiate୦foo୦PrintSlice୦int([]int{1, 2, 3})
				instantiate୦foo୦PrintSlice୦string([]string{"a", "b", "c"})
//line ./cmd/main.go2:13
}
//line ./cmd/main.go2:4
func instantiate୦foo୦PrintSlice୦int(s []int) {
//line ./cmd/main.go2:7
 for _, v := range s {
//line ./cmd/main.go2:10
  fmt.Println(v)
	}
//line ./cmd/main.go2:11
}
//line ./cmd/main.go2:4
func instantiate୦foo୦PrintSlice୦string(s []string) {
//line ./cmd/main.go2:7
 for _, v := range s {
//line ./cmd/main.go2:10
  fmt.Println(v)
	}
//line ./cmd/main.go2:11
}

//line ./cmd/main.go2:11
type Importable୦ int

//line ./cmd/main.go2:11
var _ = fmt.Errorf

//line ./cmd/main.go2:11
type _ foo.Importable୦
