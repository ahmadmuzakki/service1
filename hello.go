package main

import(
	"fmt"
	"github.com/facebookgo/clock"
)

func main(){
	cl := clock.NewMock()
	
	fmt.Println("hello world",cl.Now())
}
