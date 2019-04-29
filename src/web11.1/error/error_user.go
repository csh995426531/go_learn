package error

import (
	//"encoding/json"
	//"fmt"
)

type SyntaxError struct {
	msg string // 错误描述
	Offset int64 // 错误发生的位置
}

func (e *SyntaxError) Error() string {return e.msg}

func main(){

	//if err := dec.Decode(&val); err != nil {
	//	if serr, of := err.(*json.SyntaxError); ok {
	//		line,col := findLine(f,serr,offset)
	//		return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
	//	}
	//	return err
	//}
}
