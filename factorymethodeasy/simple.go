package factorymethodeasy

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("hi,%s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("hello,%s", name)
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}

	return nil
}
