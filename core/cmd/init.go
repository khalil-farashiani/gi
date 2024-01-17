// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"cogentcore.org/core/core/config"
	"cogentcore.org/core/grease"
)

// Init initializes the ".core" directory
// and a "config.toml" file inside it.
// The "config.toml" file has the given
// config info. Init also sets the config name
// to the current directory if it is unset.
func Init(c *config.Config) error { //gti:add
	err := os.MkdirAll(".core", 0750)
	if err != nil {
		return fmt.Errorf("error creating %q directory: %w", ".core", err)
	}
	err = grease.Save(c, ".core/config.toml")
	if err != nil {
		return fmt.Errorf("error writing to configuration file: %w", err)
	}
	return nil
}