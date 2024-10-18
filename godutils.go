package godutils

import (
	"database/sql"
	"fmt"
	"math"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func FindInList[T any](list *[]T, fn func(T) bool) *T {
	for i := range *list {
		ele := &(*list)[i]
		if fn(*ele) {
			return ele
		}
	}
	return nil
}

func FindInListP[T any](list *[]*T, fn func(*T) bool) *T {
	for i := range *list {
		ele := (*list)[i]
		if fn(ele) {
			return ele
		}
	}
	return nil
}

func GetOrigin(originHeader string) (origin string) {
	originDomain, err := url.Parse(originHeader)
	if err != nil {
		return
	}
	originHost := originDomain.Hostname()
	origin = strings.Split(originHost, ".")[0]
	return
}

func OffsetTime(t time.Time, offset int) (withOffset time.Time) {
	tCopy := t
	d := time.Duration(offset) * time.Minute
	withOffset = tCopy.Add(d)
	return
}

func OffsetToLoc(offset int) (loc *time.Location) {
	sign := "+"
	if math.Signbit(float64(offset)) {
		sign = "-"
	}

	mins := time.Duration(math.Abs(float64(offset))) * time.Minute

	hours := mins / time.Hour
	mins -= hours * time.Hour
	mins = mins / time.Minute

	zoneName := fmt.Sprintf("UTC%s%02d:%02d", sign, hours, mins)
	loc = time.FixedZone(zoneName, offset*60)

	return
}

func GetDynamicField[T any](structPtr interface{}, fieldName string, vType T) (value T, err error) {
	rv := reflect.ValueOf(structPtr)
	if rv.Kind() != reflect.Ptr {
		err = fmt.Errorf("structPtr must be a pointer")
		return
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		err = fmt.Errorf("structPtr must be a pointer to a struct")
		return
	}

	fv := rv.FieldByName(fieldName)
	if !fv.IsValid() {
		err = fmt.Errorf("field %s not found", fieldName)
		return
	}
	if !fv.CanInterface() {
		err = fmt.Errorf("field %s cannot be accessed", fieldName)
		return
	}
	if fv.Kind() != reflect.TypeOf(vType).Kind() {
		err = fmt.Errorf("field %s is not of the requested type", fieldName)
		return
	}

	value = fv.Interface().(T)

	return
}

func SetDynamicField[T any](structPtr interface{}, fieldName string, value T) (err error) {
	rv := reflect.ValueOf(structPtr)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("structPtr must be a pointer")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("structPtr must be a pointer to a struct")
	}

	fv := rv.FieldByName(fieldName)
	if !fv.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}
	if !fv.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}
	if fv.Kind() != reflect.TypeOf(value).Kind() {
		return fmt.Errorf("field %s is not of the same type as value", fieldName)
	}

	fv.Set(reflect.ValueOf(value))

	return
}

func GetNullInt32(v int) (nullInt32 sql.NullInt32) {
	nullInt32 = sql.NullInt32{Valid: false}
	if v == 0 {
		return
	}
	nullInt32 = sql.NullInt32{Valid: true, Int32: int32(v)}
	return
}

func GetOrCreate[K comparable, V any](mapObj map[K]*V, key K) *V {
	v, exists := mapObj[key]
	if !exists {
		typeOfV := reflect.TypeOf(v)
		isPtr := typeOfV.Kind() == reflect.Ptr
		if isPtr {
			typeOfV = typeOfV.Elem()
		}
		newVal := reflect.New(typeOfV).Elem().Interface().(V)
		v = &newVal
		mapObj[key] = v
	}
	return v
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToPascalCase(str string) string {
	// Remove all underscores
	str = strings.ReplaceAll(str, "_", " ")
	// Convert to title case
	str = strings.Title(str)
	// Remove all spaces
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func FilterList[T any](list *[]T, fn func(T) bool) *[]T {
	filtered := make([]T, 0)
	for _, ele := range *list {
		if fn(ele) {
			filtered = append(filtered, ele)
		}
	}
	return &filtered
}

func FilterListP[T any](list *[]*T, fn func(*T) bool) *[]*T {
	filtered := make([]*T, 0)
	for _, ele := range *list {
		if fn(ele) {
			filtered = append(filtered, ele)
		}
	}
	return &filtered
}
