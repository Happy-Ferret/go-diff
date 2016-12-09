package diff

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func TestStruct(t *testing.T) {
	type B struct {
		X string
		x string
	}
	type A struct {
		Name  string
		Value int
		B     *B
	}

	var a = A{Name: "hello w1", B: &B{X: "hello", x: "12312"}}
	var b = A{Name: "hello", B: &B{}}

	var diffM = Diff{}
	diffM.DiffStructs(b, a)
	for k, v := range diffM.ChangedItems() {
		log.Printf("Key: %s , Value: %t", k, v.Changed)
	}

}

func TestStructWithEmbedPtr(t *testing.T) {
	type B struct {
		X string
	}
	type A struct {
		Name  string
		Value int
		B     B
	}

	var a = A{Name: "hello w1", B: B{X: "hello"}}
	var b = A{Name: "hello"}

	var diffM = Diff{}
	diffM.DiffStructs(b, a)
	for k, v := range diffM.ChangedItems() {
		log.Printf("Key: %s , Value: %t", k, v.Changed)
	}
}

func TestMap(t *testing.T) {
	type M map[string]interface{}
	var a = M{
		"int_n": 123, "int_c": 123,
		"str_n": "hello", "str_c": "hello",

		"array_n": []string{"1231"}, "array_c": []string{"12313"},
	}
	var b = M{
		"int_n": 123, "int_c": 1231,
		"str_n": "hello", "str_c": "hello world",

		"array_n": []string{"1231"}, "array_c": []interface{}{"12313", 123},
	}

	for k, v := range DiffMaps(a, b).ChangedItems() {
		log.Printf("Key: %s , Value: %t", k, v.Changed)
	}
}
