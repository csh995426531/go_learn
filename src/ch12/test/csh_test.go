package test

import (
	"fmt"
	"testing"
)

type Dog struct {
	name string
	distance int8
	speed int8
}

type TaiDi struct {
	Dog

}

func (d *Dog) going()  {
	d.distance += d.speed
	fmt.Printf("%s 跑了:%d 米", d.name, d.distance)
}

func (d *Dog) goTo() {
	d.going()
}

func (d *TaiDi) going()  {
	d.distance += d.speed
	fmt.Printf("%s 跑了:%d 公里", d.name, d.distance)
}

func TestCsh(t *testing.T) {
	td := new(TaiDi)

	td.name = "泰迪"
	td.speed = 10
	td.goTo()
}
