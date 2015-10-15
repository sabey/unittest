// Copyright 2015, JuanDeFu.ca. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
package unittest

import (
	"fmt"
	"testing"
)

type wat struct {
	i int
}

func Wat() *wat {
	return nil
}
func (self *wat) okay() bool {
	return true
}

func TestUnitTest(t *testing.T) {
	fmt.Println("TestUnitTest()")
	Equals(t, Equals(nil, 1, 1), true)
	Equals(t, Equals(nil, 1, int64(1)), true)
	Equals(t, Equals(nil, int8(1), int64(1)), true)
	Equals(t, Equals(nil, uint8(1), int64(1)), true)
	Equals(t, Equals(nil, int8(-1), int64(-1)), true)
	Equals(t, Equals(nil, int8(-1), int64(1)), false)
	Equals(t, EqualsExact(nil, uint8(1), int64(1)), false)
	Equals(t, IsNil(nil, nil), true)
	Equals(t, IsNil(nil, Wat()), true)
	Equals(t, NotNil(nil, &wat{}), true)
	Equals(t, NotNil(nil, Wat()), false)
}
