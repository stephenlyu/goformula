package function

import (
	"fmt"
	"reflect"
	"strings"
)

type ValueCache struct {
	data map[string]any
}

func NewValueCache() *ValueCache {
	return &ValueCache{
		data: make(map[string]any),
	}
}

func (vc *ValueCache) Get(key string) any {
	ret, ok := vc.data[key]
	if ok {
		return ret
	}
	return nil
}

func (vc *ValueCache) Set(key string, value any) {
	vc.data[key] = value
}

func (vc *ValueCache) Del(key string) {
	delete(vc.data, key)
}

var GlobalValueCache = NewValueCache()

func ValueCacheBuildKey[T float32 | float64](values []Value[T]) string {
	var parts []string
	for _, v := range values {
		if v.IsScalar() {
			parts = append(parts, fmt.Sprintf("%v", v.Get(0)))
		} else {
			va := reflect.ValueOf(v)
			ptr := va.Pointer()
			parts = append(parts, fmt.Sprintf("%x", ptr))
		}
	}
	return strings.Join(parts, "_")
}
