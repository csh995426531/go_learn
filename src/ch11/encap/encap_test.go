package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id		string
	Name 	string
	Age		string
}

func TestCreateEmployeeObg(t *testing.T) {
	e := Employee{"0","张三","23"}
	e1 := new(Employee) //返回指针
	e1.Id = "1"
	e1.Name = "李四"
	e1.Age = "18"
	e2 := Employee{Id:"2",Name:"王五",Age:"30"}

	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e1 is %T", e1)
	t.Logf("e2 is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s Name:%s Age:%s", e.Id, e.Name, e.Age)
}

func (e *Employee) String2() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s Name:%s Age:%s", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T){
	e := Employee{"5", "赵六", "24"}
	t.Log(e.String2())
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))

}