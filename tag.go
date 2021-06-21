package godegen

import (
	"reflect"
	"strings"
)

type Tag string

func (t Tag) Has(key string) bool {
	_, ok := reflect.StructTag(t).Lookup(key)

	return ok
}

func (t Tag) Get(key string) string {
	value := reflect.StructTag(t).Get(key)
	ss := strings.Split(value, ",")

	return ss[0]
}

func (t Tag) HasFlag(key, flag string) bool {
	value := reflect.StructTag(t).Get(key)
	ss := strings.Split(value, ",")

	flag = strings.ToLower(flag)
	for _, v := range ss[1:] {
		if strings.ToLower(v) == flag {
			return true
		}
	}

	return false
}
