package str

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T)  {
	s:="A:B:C"
	parts := strings.Split(s, ":")

	for _,v := range parts {
		t.Log(v)
	}

	t.Log(strings.Join(parts, "_"))
}

func TestStringConv(t *testing.T){
	s:=strconv.Itoa(10)
	t.Log("str"+s)
	if i,err := strconv.Atoi("10");err==nil{
		t.Log(i+10)
	}
}