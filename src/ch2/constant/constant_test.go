package constant

import "testing"

const(
	Monday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1 << iota
	Writable
	Executable
)


func TestConstant(t *testing.T){
	t.Log(Monday, Tuesday)
}

func TestDingf(t *testing.T){
	a := 7 //0111
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Executable==Executable)
}