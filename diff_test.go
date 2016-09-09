package diff

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func TestStruct(t *testing.T) {
	type B struct {
		X string
	}
	type A struct {
		Name  string
		Value int
		B     *B
	}

	var a = A{Name: "hello w1", B: &B{X: "hello"}}
	var b = A{Name: "hello", B: &B{}}

	var diffM = Diff{}
	diffM.DiffStructs(b, a)

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
}

func TestMap(t *testing.T) {
	type M map[string]interface{}
	var a = M{"name": "hello", "1": 123}
	var b = M{"1": 123123}
	var diffM = Diff{}
	diffM.DiffMaps(b, a)
}
