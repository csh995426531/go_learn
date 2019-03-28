package fn__test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)
func returnMultiValues() (int,int) {
	return rand.Intn(10), rand.Intn(20)
}
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int)int {
		start:=time.Now()
		ret:= inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}
}
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	a,b := returnMultiValues()
	t.Log(a,b)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
func sum(ops ...int) int {

	ret:=0
	for _,v :=range ops {
		ret += v
	}
	return  ret
}
func TestVarParam(t *testing.T) {
	t.Log(sum(1,2,43))
}
func Clear(){
	fmt.Print("this is defer")
}
func TestDefer(t *testing.T) {
	defer Clear()

	fmt.Print("aaa")
	panic("fail")
}




