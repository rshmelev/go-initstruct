package initstruct

import (
	"reflect"
	"strconv"

	// . "github.com/davecgh/go-spew/spew"
)

var defaultInitializer StructInitializer

type StructInitializer struct {
	TagName string
}

func (si *StructInitializer) Init(someStruct interface{}, onlyFieldsWithDefaultValues bool, recursiveInitialization bool) {
	v := reflect.ValueOf(someStruct)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	tagName := si.TagName
	if tagName == "" {
		tagName = "init"
	}
	vtype := v.Type()
	for i := 0; i < vtype.NumField(); i++ {
		f := vtype.Field(i)
		withWhat := f.Tag.Get(tagName)
		if withWhat != "" {
			if vf := v.Field(i); vf.CanSet() {
				if vf.Kind() == reflect.Ptr {
					ptrtype := vf.Type().Elem()
					if ptrtype.Kind() == reflect.Struct {

						if !onlyFieldsWithDefaultValues || vf.IsNil() {

							newstruct := reflect.New(ptrtype)
							vf.Set(newstruct)
							//							si.Init(newstruct.Interface(), onlyFieldsWithDefaultValues)
							si.Init(v.Field(i).Interface(), onlyFieldsWithDefaultValues, recursiveInitialization)
						} else if !vf.IsNil() && recursiveInitialization {
							si.Init(vf.Interface(), onlyFieldsWithDefaultValues, recursiveInitialization)
						}
					}
				} else {
					si.InitValueWithStr(&vf, withWhat, onlyFieldsWithDefaultValues, recursiveInitialization)
				}
			}
		}
	}
}

func (si *StructInitializer) InitValueWithStr(v *reflect.Value, s string, onlyFieldsWithDefaultValues bool, recursiveInitialization bool) bool {
	if !v.CanSet() {
		return false
	}
	vk := v.Kind()
	if vk == reflect.String {
		if !onlyFieldsWithDefaultValues || v.String() == "" {
			v.SetString(s)
		}
		return true
	}
	if vk == reflect.Struct && v.CanAddr() && recursiveInitialization {
		p := v.Addr()
		si.Init(p.Interface(), onlyFieldsWithDefaultValues, recursiveInitialization)
		return true
	}
	if vk >= reflect.Int && vk <= reflect.Int64 {
		if !onlyFieldsWithDefaultValues || v.Int() == 0 {
			theint, _ := strconv.ParseInt(s, 10, 64)
			v.SetInt(theint)
		}
		return true
	}
	if vk == reflect.Bool {
		if !onlyFieldsWithDefaultValues || v.Bool() == false {
			thebool, _ := strconv.ParseBool(s)
			v.SetBool(thebool)
		}
		return true
	}
	if vk == reflect.Chan {
		if !onlyFieldsWithDefaultValues || v.IsNil() {
			cap, _ := strconv.ParseInt(s, 10, 64)
			vchan := reflect.MakeChan(v.Type(), int(cap)) //makeChannel(v.Type(), v.Type().ChanDir(), int(cap))
			v.Set(vchan)
		}
		return true
	}
	if vk == reflect.Slice {
		if !onlyFieldsWithDefaultValues || v.IsNil() {
			//			x := reflect.New(v.Type())
			x := reflect.MakeSlice(v.Type(), 0, 0)
			v.Set(x)
			//			println("dumping var interface! ")
			//			Dump(v.Interface())
		}
		return true
	}
	if vk == reflect.Map {
		if !onlyFieldsWithDefaultValues || v.IsNil() {
			m := reflect.MakeMap(v.Type())
			//x := reflect.New(v.Type())

			v.Set(m) //x.Elem().Interface())
			//			println("dumping var interface! ")
			//			Dump(v.Interface())
		}
		return true
	}
	if vk == reflect.Float32 || vk == reflect.Float64 {
		if !onlyFieldsWithDefaultValues || v.Float() == 0 {
			thefloat, _ := strconv.ParseFloat(s, 64)
			v.SetFloat(thefloat)
		}
		return true
	}
	if vk >= reflect.Uint && vk <= reflect.Uint64 {
		if !onlyFieldsWithDefaultValues || v.Uint() == 0 {
			theuint, _ := strconv.ParseUint(s, 10, 64)
			v.SetUint(theuint)
		}
		return true
	}

	return true
}

func InitZeroFieldsRecursively(someStruct interface{}) {
	initStruct(someStruct, true, true)
}
func ResetAllFields(someStruct interface{}) {
	initStruct(someStruct, false, false)
}

func initStruct(someStruct interface{}, onlyFieldsWithDefaultValues bool, recursiveInitialization bool) {
	defaultInitializer.Init(someStruct, onlyFieldsWithDefaultValues, recursiveInitialization)
}

//------------------

//func makeChannel(t reflect.Type, chanDir reflect.ChanDir, buffer int) reflect.Value {
//	ctype := reflect.ChanOf(chanDir, t)
//	return reflect.MakeChan(ctype, buffer)
//}
