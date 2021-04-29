package main

import (
	bui "./builder"
	fac "./factorymethod"
	"fmt"
)

//TestFactorymethodeasy
func compute(factory fac.OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestFactorymethodeasy() {
	factory := fac.PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		fmt.Println("error with factory method pattern")
	}

	factory2 := fac.MinusOperatorFactory{}
	if compute(factory2, 4, 2) != 2 {
		fmt.Println("error with factory method pattern")
	}
}

func TestBuilder() {
	builder := &bui.Builder1{}
	director := bui.NewDirector(builder)
	director.Construct()
	result := builder.GetResult()
	if result != "123" {
		fmt.Println("not 123")
	}
	fmt.Println(result)
}

func main() {
	//TestFactorymethodeasy()

}
