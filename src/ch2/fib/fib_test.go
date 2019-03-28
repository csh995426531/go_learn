package fib

import (
	"testing"
)

func TestFibTry(t *testing.T){

	var (
		a int = 0
		b int = 1
	)

	for i:=0;i<5;i++ {
		t.Log("", b)
		tmp := a
		a = b
		b = a+tmp
	}
}

func TestExchange(t *testing.T){
	a := 2
	b := 3
	a,b = b,a

	t.Log(a, b)
}