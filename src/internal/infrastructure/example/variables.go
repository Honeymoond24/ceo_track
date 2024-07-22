package example

import "fmt"

func sayHello() {
	fmt.Println("Hello, World!")
}

func useFunctions() {
	// Call a function
	sayHello()

	// Define a function
	innerFunc := func() {
		fmt.Println("Hello, World!")
	}
	innerFunc()
}

func useVariables() {
	var a int
	a = 1
	fmt.Println(a)

	var b int = 2
	fmt.Println(b)

	c := 3
	fmt.Println(c)

	var d, e int = 4, 5
	fmt.Println(d, e)

	f, g := 6, 7
	fmt.Println(f, g)
}

type test struct {
	test1 string
	test2 int
}

type ITest interface {
	methodOne(o int, p string) (int, string)
	methodTwo() string
}

func (receiver test) structMethod() {
	fmt.Println("Test")
}

func (receiver test) methodOne(o int, p string) (int, string) {
	return 0, ""
}

func useStructs() {
	var t test
	t.test1 = "test"
	t.test2 = 1
	t.structMethod()
	fmt.Println(t)
}

func Run() {
	//useFunctions()
	//useVariables()
}
