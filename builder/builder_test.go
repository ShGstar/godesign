package builder

import "testing"

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	result := builder.GetResult()
	t.Log(result)
	if result != "123" {
		t.Fatal("Builder1 fail expect 123 acture %s", result)
	}
	t.Log(result)
}

func TestBuilder2(t *testing.T) {
	builder2 := &Builder2{}
	director := NewDirector(builder2)
	director.Construct()
	result := builder2.GetResult()
	t.Log(result)
	if result != 6 {
		t.Fatal("Builder2 fail expect 6 acture %d", result)
	}
	t.Log(result)
}
