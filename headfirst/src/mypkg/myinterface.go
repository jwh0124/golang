package mypkg

import "fmt"

// MyInterface Interface
type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

// MyType integer
type MyType int

// MethodWithoutParameters function
func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

// MethodWithParameter function
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

// MethodWithReturnValue function
func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

// MethodNotInInterface function
func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
