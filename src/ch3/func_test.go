package ch3

import "testing"

const (
	re int = 1<<iota; //0001
	wr;//0010
	ex;//0100
)

func Test(t *testing.T){
	a:=1;

	t.Log(a&re==re, a&wr, a&ex)

	c:=7 //0111

	t.Log(c&^2);

	d:=2; //0010
	e:=1; //0001

	t.Log(d|e, d&e, d^e, d<<1, d<<2, d>>1)
}


