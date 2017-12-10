/*
 *  @author zhongxiao.yzx
 *  @date   2017/12/2
 *  @description
 *   extension type of HashSet for practices
 */
// set
package set

import (
	"bytes"
	"fmt"
)

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{}) bool
	Clear()
	Contains(e interface{}) bool
	Length() int
	Equal(rhs Set) bool
	Elements() []interface{}
	String() string
}

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) bool {
	if !set.m[e] {
		return false
	}
	delete(set.m, e)
	return true
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Length() int {
	return len(set.m)
}

func (set *HashSet) Equal(rhs *HashSet) bool {
	if rhs == nil || set.Length() != rhs.Length() {
		return false
	}
	for key := range rhs.m {
		if !set.Contains(key) {
			return false
		}
	}
	return false
}

func (set *HashSet) Elements() []interface{} {
	initLen := len(set.m)
	slices := make([]interface{}, initLen)
	index := 0
	for key := range set.m {
		if index < initLen {
			slices[index] = key
		} else {
			slices = append(slices, key)
		}
		index++
	}
	if index < initLen {
		slices = slices[:index]
	}
	return slices
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString(" }")
	return buf.String()
}
