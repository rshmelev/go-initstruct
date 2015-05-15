package initstruct

import (
	"fmt"
	"testing"
)

type Test3 struct {
	o string `init:"vasya"`
	E *Test2 `init:"yes"`
}

type Test2 struct {
	o string
	X string `init:"roma"`
}

type TestStruct struct {
	AA int     `init:"3"`
	A  int     `init:"3"`
	B  string  `init:"lala"`
	C  float32 `init:"3.14"`
	D  string
	E  *Test2 `init:"yes"`
	F  *Test2
	X  Test3 `init:"recurse"`
}

func TestSomething(t *testing.T) {
	s := &TestStruct{}
	InitZeroFields(s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", s.E)
	fmt.Printf("%+v\n", s.X)
}
