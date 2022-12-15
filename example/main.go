package example

import (
	"fmt"
	"github.com/godcong/co"
)

func init() {
	str.New("abc")
}

func main() {
	str.New("bcd")
	fmt.Println(str.Data())
	//output:
	//abc
}

var str co.CreateOnce[string]
