package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHellWorld() Code
}

type GoProgrammer struct {
}

func(g *GoProgrammer) WriteHellWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (j * JavaProgrammer) WriteHellWorld() Code{
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgram(p Programmer){
	fmt.Printf("%T %v\n", p, p.WriteHellWorld())
}

func TestPolymorphism(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}