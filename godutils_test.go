package godutils

import (
	"reflect"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	str := "HelloWorld"
	want := "hello_world"
	got := ToSnakeCase(str)
	if got != want {
		t.Errorf("ToSnakeCase() = %v, want %v", got, want)
	}
}

func TestToPascalCase(t *testing.T) {
	str := "hello_world"
	want := "HelloWorld"
	got := ToPascalCase(str)
	if got != want {
		t.Errorf("ToPascalCase() = %v, want %v", got, want)
	}
}

func TestFilterList(t *testing.T) {
	list := []string{"a", "b", "c"}
	want := &[]string{"b"}
	got := FilterList(&list, func(item string) bool {
		return item == "b"
	})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("FilterList() = %v, want %v", got, want)
	}
}
