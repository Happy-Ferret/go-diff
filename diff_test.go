package diff

import (
	log "github.com/irupsen/logrus"
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
	var a = map[string]interface{}{
		"int_n": 123, "int_c": 123,
		"str_n": "hello", "str_c": "hello",
		"a": map[string]interface{}{
			"name": 123,
		},
		"array_n": []string{"1231"}, "array_c": []string{"12313"},
	}
	var b = map[string]interface{}{
		"int_n": 123, "int_c": 1231,
		"str_n": "hello", "str_c": "hello world",
		"a": map[string]interface{}{
			"name": 123123,
		},
		"array_n": []string{"1231"}, "array_c": []interface{}{"12313", 123},
	}

	for k, v := range DiffMaps(a, b).ChangedItems() {
		log.Printf("Key: %s , Value: %t", k, v.Changed)
	}
}

func TestMapWithEmpty(t *testing.T) {
	var a = map[string]interface{}{
		"int_n": 123, "int_c": 123,
		"str_n": "hello", "str_c": "hello",
		"a": map[string]interface{}{
			"name": 123,
		},
		"array_n": []string{"1231"}, "array_c": []string{"12313"},
	}
	var b map[string]interface{}
	for k, v := range DiffMaps(a, b).ChangedItems() {
		log.Printf("Key: %s , Value: %t", k, v.Changed)
	}
}

func TestStructWithEmpty(t *testing.T) {
	type X struct {
		Name string
		Age  int
	}
	var a = &X{"123", 12}
	var b *X
	var d = Diff{}
	d.DiffStructs(a, b)
	for k, v := range d.ChangedItems() {
		log.Printf("Key: %s , Value: %t", k, v.Changed)
	}
}
