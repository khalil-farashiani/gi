// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xmls

import (
	"path/filepath"
	"testing"

	"cogentcore.org/core/grr"
)

type testStruct struct {
	A string
	B float32
}

func TestXML(t *testing.T) {
	tpath := filepath.Join("testdata", "test.xml")
	tipath := filepath.Join("testdata", "test-indent.xml")

	s := &testStruct{A: "aaa", B: 3.14}
	grr.Test(t, Save(s, tpath))
	b, err := WriteBytes(s)
	grr.Test(t, err)

	a := &testStruct{}
	grr.Test(t, Open(a, tpath))
	if *a != *s {
		t.Errorf("Open failed to read same data as saved: wanted %v != got %v", s, a)
	}

	c := &testStruct{}
	grr.Test(t, ReadBytes(c, b))
	if *c != *s {
		t.Errorf("ReadBytes or WriteBytes failed to read same data as saved: wanted %v != got %v", s, c)
	}

	grr.Test(t, SaveIndent(s, tipath))
}