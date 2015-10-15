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
		var m string
		if is_int(a) && is_int(b) {
			if fmt.Sprintf("%d", a) == fmt.Sprintf("%d", b) {
				return true
			}
			m = fmt.Sprintf("numbers don't match %v(%T) != %v(%T)", a, a, b, b)
		} else {
			m = fmt.Sprintf("%v(%T) != %v(%T)", a, a, b, b)
		}
		s := fmt.Sprintf("%s\n%s\n", m, stacktrace.StackTrace(0))
		if t != nil {
			t.Fatalf(s)
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
		s := fmt.Sprintf("%v(%T) != %v(%T)\n%s\n", a, a, b, b, stacktrace.StackTrace(0))
		if t != nil {
			t.Fatalf(s)
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
		s := fmt.Sprintf("%v(%T) != nil\n%s\n", a, a, stacktrace.StackTrace(0))
		if t != nil {
			t.Fatalf(s)
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
		s := fmt.Sprintf("%v(%T) == nil\n%s\n", a, a, stacktrace.StackTrace(0))
		if t != nil {
			t.Fatalf(s)
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
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	default:
	}
	return false
}
