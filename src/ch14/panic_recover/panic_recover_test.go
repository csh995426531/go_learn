package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicExit(t *testing.T) {

	//defer func() {
	//	fmt.Println("Finally!")
	//}()
	defer func() {
		if err:=recover(); err !=nil{
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("start")
	//os.Exit(-1)
	panic(errors.New("Something wrong!"))
}