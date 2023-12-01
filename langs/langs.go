// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package langs

import (
	"fmt"

	"goki.dev/fi"
)

var ParserBytes map[fi.Supported][]byte = make(map[fi.Supported][]byte)

func OpenParser(sl fi.Supported) ([]byte, error) {
	parserBytes, ok := ParserBytes[sl]
	if !ok {
		return nil, fmt.Errorf("langs.OpenParser: no parser bytes for %v", sl)
	}
	return parserBytes, nil
}
