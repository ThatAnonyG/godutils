package godutils

import (
	"reflect"
	"testing"
	"time"
)

func TestGetDynamicField(t *testing.T) {
	testStruct := struct {
		IntField int64
	}{
		IntField: 1,
	}
	want := int64(1)
	got, err := GetDynamicField(&testStruct, "IntField", int64(0))
	if err != nil {
		t.Errorf("GetDynamicField() err: %v", err)
	}
	if got != want {
		t.Errorf("GetDynamicField() = %v, want %v", got, want)
	}
}

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

func TestInSameClock(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		t.Errorf("time.LoadLocation() err: %v", err)
	}
	t1 := time.Date(2024, 10, 25, 10, 30, 0, 0, loc)
	want := "2024-10-25 10:30:00 +0000 UTC"
	got := InSameClock(t1, time.UTC).String()
	if got != want {
		t.Errorf("InSameClock() = %v, want %v", got, want)
	}
}
