package diff

import (
	"fmt"
	"github.com/fatih/structs"
	"reflect"
	"strings"
)

type Diff struct {
	items  map[string]DiffItem
	noDiff bool
}

func NewDiff() Diff {
	return Diff{}
}

func (d *Diff) HadChanged(path ...string) (changed bool, err error) {
	if path == nil {
		return !d.noDiff, nil
	}

	// TODO： 如果这个值不存在，那么是否应该判断？
	if k, found := d.items[strings.Join(path, ".")]; !found {
		return false, fmt.Errorf("%s", "not found")
	} else {
		changed = k.Changed
		return changed, nil
	}
}

func (d *Diff) AddItem(item DiffItem) {
	if d.items == nil {
		d.items = make(map[string]DiffItem)
	}
	d.items[item.Path] = item
}

func (d *Diff) ChangedItems() map[string]DiffItem {
	var result = make(map[string]DiffItem)
	for k, v := range d.items {
		if v.Changed == false {
			continue
		}
		result[k] = v
	}
	return result
}

type DiffItem struct {
	Name     string
	Path     string
	OrgValue interface{}
	NewValue interface{}
	Changed  bool
}

func (d DiffItem) Get() (nw, od interface{}) {
	return d.NewValue, d.OrgValue
}

func (d *Diff) DiffStructs(newStruct, orgStruct interface{}, prefix ...string) {
	if reflect.DeepEqual(newStruct, orgStruct) {
		d.noDiff = true
		return
	}
	d.DiffMaps(structs.Map(newStruct), structs.Map(orgStruct), prefix...)
}

func DiffStructs(n, o map[string]interface{}) *Diff {
	d := NewDiff()
	d.DiffMaps(n, o)
	d.DiffStructs(n, o)
	return &d
}

func DiffMaps(n, o map[string]interface{}) *Diff {
	d := NewDiff()
	d.DiffMaps(n, o)
	d.DiffMaps(n, o)
	return &d
}

func (d *Diff) DiffMaps(newMap, orgMap map[string]interface{}, prefix ...string) {
	var _p = ""

	if prefix != nil {
		_p = prefix[0]
	}

	if reflect.DeepEqual(newMap, orgMap) {
		d.noDiff = true
		return
	}

	var tmpNames = make(map[string]bool)

	for k, _ := range newMap {
		tmpNames[k] = true
	}

	for k, _ := range orgMap {
		tmpNames[k] = true
	}

	for k, _ := range tmpNames {

		path := fmt.Sprintf("%s%s", _p, k)

		dv, ov := newMap[k], orgMap[k]

		var diffItem = DiffItem{
			Name:     k,
			Path:     path,
			OrgValue: ov,
			NewValue: dv,
			Changed:  false,
		}

		if !reflect.DeepEqual(dv, ov) {
			diffItem.Changed = true
			var kind reflect.Kind
			if dv == nil {
				kind = reflect.TypeOf(ov).Kind()
			} else {
				kind = reflect.TypeOf(dv).Kind()
			}

			switch kind {
			case reflect.Struct:
				d.DiffStructs(dv, ov, fmt.Sprintf("%s.", k))
			case reflect.Map:
				if dv == nil {
					dv = make(map[string]interface{})
				}
				if ov == nil {
					ov = make(map[string]interface{})
				}
				d.DiffMaps(dv.(map[string]interface{}), ov.(map[string]interface{}), fmt.Sprintf("%s.", k))
			default:
				diffItem.Changed = true
			}

		}
		d.AddItem(diffItem)
	}
}
