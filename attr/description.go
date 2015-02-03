// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package attr

import (
	"code.wolfmud.org/WolfMUD-mini.git/has"
)

type description struct {
	attribute
	description string
}

// Some interfaces we want to make sure we implement
var (
	_ has.Attribute   = Description()
	_ has.Description = Description()
)

func Description() *description {
	return nil
}

func (*description) New(d string) *description {
	return &description{attribute{}, d}
}

func (*description) Find(t has.Thing) has.Description {
	for _, a := range t.Attrs() {
		if a, ok := a.(has.Description); ok {
			return a
		}
	}
	return nil
}

func (*description) FindAll(t has.Thing) (matches []has.Description) {
	for _, a := range t.Attrs() {
		if a, ok := a.(has.Description); ok {
			matches = append(matches, a)
		}
	}
	return
}

func (d *description) Dump() []string {
	return []string{DumpFmt("%p %[1]T %q", d, d.description)}
}

func (d *description) Description() string {
	return d.description
}
