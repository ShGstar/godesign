package factorymethodeasy

import "testing"

func TestType1(t *testing.T) {
	api := NewAPI(1)
	say := api.Say("TOM")
	if say != "hi,TOM" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	say := api.Say("Tom")
	if say != "hello,Tom" {
		t.Fatal("Type2 test fail")
	}
}
