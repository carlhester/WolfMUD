// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package attr

import (
	"code.wolfmud.org/WolfMUD.git/attr/internal"
	"code.wolfmud.org/WolfMUD.git/has"
	"code.wolfmud.org/WolfMUD.git/recordjar/decode"
	"code.wolfmud.org/WolfMUD.git/recordjar/encode"
)

// Register marshaler for Description attribute.
func init() {
	internal.AddMarshaler((*Description)(nil), "description")
}

// Description implements an attribute for describing Things. Things can have
// multiple descriptions or other attributes that implement the has.Description
// interface to add additional information to descriptions.
type Description struct {
	Attribute
	description string
}

// Some interfaces we want to make sure we implement. If we don't we'll throw
// compile time errors.
var (
	_ has.Description = &Description{}
)

// NewDescription returns a new Description attribute initialised with the
// specified description.
func NewDescription(description string) *Description {
	return &Description{Attribute{}, description}
}

// FindAllDescription searches the attributes of the specified Thing for
// attributes that implement has.Description returning all that match. If no
// matches are found an empty slice will be returned.
func FindAllDescription(t has.Thing) (matches []has.Description) {
	for _, a := range t.Attrs() {
		if a, ok := a.(has.Description); ok {
			matches = append(matches, a)

			// If type is an actual *Description move it to the front of the slice as
			// we want main descriptions first and additional descriptions afterwards
			if _, ok := a.(*Description); ok {
				copy(matches[1:], matches[0:])
				matches[0] = a
			}
		}
	}
	return
}

// Is returns true if passed attribute implements a description else false.
func (*Description) Is(a has.Attribute) bool {
	_, ok := a.(has.Description)
	return ok
}

// Found returns false if the receiver is nil otherwise true.
func (d *Description) Found() bool {
	return d != nil
}

// Unmarshal is used to turn the passed data into a new Description attribute.
func (*Description) Unmarshal(data []byte) has.Attribute {
	return NewDescription(decode.String(data))
}

// Marshal returns a tag and []byte that represents the receiver.
func (d *Description) Marshal() (tag string, data []byte) {
	return "description", encode.Bytes([]byte(d.description))
}

func (d *Description) Dump() []string {
	return []string{DumpFmt("%p %[1]T: %q", d, d.description)}
}

// Description returns the descriptive string of the attribute.
func (d *Description) Description() string {
	return d.description
}

// Copy returns a copy of the Description receiver.
func (d *Description) Copy() has.Attribute {
	if d == nil {
		return (*Description)(nil)
	}
	return NewDescription(d.description)
}
