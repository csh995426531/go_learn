package main

import (
	"reflect"
	"regexp"
)

type controllerInfo struct {
	regex			*regexp.Regexp
	params			map[int]string
	controllerType	reflect.Type
}
//
//type ControllerRegistor struct {
//	routers 		[]*controllerInfo
//	Application 	*App
//}
//
//func (p *ControllerRegistor) Add(pattern string, c ControllerInterface){
//
//}