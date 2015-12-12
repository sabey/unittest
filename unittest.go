// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package unittest

import (
	"fmt"
	"github.com/sabey/stacktrace"
	"log"
	"reflect"
	"testing"
)

func Equals(
	t *testing.T,
	a interface{},
	b interface{},
) bool {
	if a != b {
		if is_int(a) && is_int(b) {
			if fmt.Sprintf("%d", a) == fmt.Sprintf("%d", b) {
				return true
			}
		}
		s := "`"
		s += print(a)
		s += fmt.Sprintf("` (%T) != `", a)
		s += print(b)
		s += fmt.Sprintf("` (%T)\n", b)
		s += stacktrace.StackTrace(0)
		s += "\n"
		if t != nil {
			t.Fatal(s)
		} else {
			log.Print(s)
		}
		return false
	}
	return true
}
func EqualsExact(
	t *testing.T,
	a interface{},
	b interface{},
) bool {
	if a != b {
		s := "`"
		s += print(a)
		s += fmt.Sprintf("` (%T) != `", a)
		s += print(b)
		s += fmt.Sprintf("` (%T)\n", b)
		s += stacktrace.StackTrace(0)
		s += "\n"
		if t != nil {
			t.Fatal(s)
		} else {
			log.Print(s)
		}
		return false
	}
	return true
}
func IsNil(
	t *testing.T,
	a interface{},
) bool {
	if !isNil(a) {
		s := "`"
		s += print(a)
		s += fmt.Sprintf("` (%T) != nil\n", a)
		s += stacktrace.StackTrace(0)
		s += "\n"
		if t != nil {
			t.Fatal(s)
		} else {
			log.Print(s)
		}
		return false
	}
	return true
}
func NotNil(
	t *testing.T,
	a interface{},
) bool {
	if isNil(a) {
		s := "`"
		s += print(a)
		s += fmt.Sprintf("` (%T) == nil\n", a)
		s += stacktrace.StackTrace(0)
		s += "\n"
		if t != nil {
			t.Fatal(s)
		} else {
			log.Print(s)
		}
		return false
	}
	return true
}
func isNil(
	i interface{},
) bool {
	if i == nil {
		return true
	}
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}
func is_int(
	t interface{},
) bool {
	switch t.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}
func print(
	t interface{},
) string {
	if is_int(t) {
		// there is no way to tell the difference between a byte and uint8, or rune and int32
		// reflect has no way to tell the difference between an alias
		// I rather just print them as integers than their character
		return fmt.Sprintf("%d", t)
	}
	switch v := t.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	}
	return fmt.Sprintf("%v", t)
}
