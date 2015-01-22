// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package attr

import (
	"code.wolfmud.org/WolfMUD-mini.git/has"
)

type name struct {
	attribute
	name string
}

// Some interfaces we want to make sure we implement
var _ has.Attribute = &name{}
var _ has.Name = &name{}

func NewName(n string) *name {
	return &name{attribute{}, n}
}

func (n *name) Dump() []string {
	return []string{DumpFmt("%p %[1]T %q", n, n.name)}
}

func FindName(t has.Thing) (n has.Name) {
	n, _ = t.Find(&n).(has.Name)
	return
}

func (n *name) Name() string {
	return n.name
}
