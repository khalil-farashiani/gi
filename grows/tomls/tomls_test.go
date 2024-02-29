// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tomls

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	A string
	B float32
}

func TestTOML(t *testing.T) {
	tpath := filepath.Join("testdata", "test.toml")

	s := &testStruct{A: "aaa", B: 3.14}
	assert.NoError(t, Save(s, tpath))
	b, err := WriteBytes(s)
	assert.NoError(t, err)

	a := &testStruct{}
	assert.NoError(t, Open(a, tpath))
	if *a != *s {
		t.Errorf("Open failed to read same data as saved: wanted %v != got %v", s, a)
	}

	c := &testStruct{}
	assert.NoError(t, ReadBytes(c, b))
	if *c != *s {
		t.Errorf("ReadBytes or WriteBytes failed to read same data as saved: wanted %v != got %v", s, c)
	}
}

type sliceItem struct {
	F string
}

type testSliceStruct struct {
	Slice []sliceItem
}

func TestSlice(t *testing.T) {
	spath := filepath.Join("testdata", "slice.toml")
	sl := &testSliceStruct{[]sliceItem{{F: "a"}, {F: "b"}}}
	assert.NoError(t, Save(sl, spath))
	assert.NoError(t, Open(sl, spath))
	assert.NoError(t, Open(sl, spath))
	if len(sl.Slice) != 2 {
		t.Errorf("expected slice to be length 2) but got length %d", len(sl.Slice))
	}
}
