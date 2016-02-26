// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package attr

import (
	"code.wolfmud.org/WolfMUD.git/has"

	"strings"
)

// TODO: Currently vetoes can only be applied to the Thing they are vetoing
// for. This means, for example, a guard could not veto the get of items at a
// location it is guarding. Also a Veto is static and unconditional.
//
// TODO: Currently a veto cannot be dynamically added or removed.

// Vetoes implement an attribute for lists of Veto preventing commands for a
// Thing that would otherwise be valid. For example you could Veto the drop
// command if a very sticky item is picked up :)
type Vetoes struct {
	Attribute
	vetoes map[string]has.Veto
}

// Some interfaces we want to make sure we implement. If we don't we'll throw
// compile time errors.
var (
	_ has.Vetoes = &Vetoes{}
)

// NewVetoes returns a new Vetoes attribute initialised with the specified
// Vetos.
func NewVetoes(veto ...has.Veto) *Vetoes {
	vetoes := &Vetoes{Attribute{}, make(map[string]has.Veto)}
	for _, v := range veto {
		vetoes.vetoes[v.Command()] = v
	}
	return vetoes
}

// FindVetoes searches the attributes of the specified Thing for attributes
// that implement has.Vetoes returning the first match it finds or a *Vetoes
// typed nil otherwise.
func FindVetoes(t has.Thing) has.Vetoes {
	for _, a := range t.Attrs() {
		if a, ok := a.(has.Vetoes); ok {
			return a
		}
	}
	return (*Vetoes)(nil)
}

func (v *Vetoes) Dump() (buff []string) {
	buff = append(buff, DumpFmt("%p %[1]T %d vetoes:", v, len(v.vetoes)))
	for _, veto := range v.vetoes {
		for _, line := range veto.Dump() {
			buff = append(buff, DumpFmt("%s", line))
		}
	}
	return buff
}

// Check checks if any of the passed commands are vetoed. The first matching
// Veto found is returned otherwise nil is returned.
func (v *Vetoes) Check(cmd ...string) has.Veto {
	if v == nil {
		return nil
	}

	// For single checks we can take a shortcut
	if len(cmd) == 1 {
		veto, _ := v.vetoes[cmd[0]]
		return veto
	}

	// For multiple checks return the first that is vetoed
	for _, cmd := range cmd {
		if veto, _ := v.vetoes[cmd]; veto != nil {
			return veto
		}
	}
	return nil
}

// Veto implements a veto for a specific command. Veto need to be added to a
// Vetoes list using NewVetoes.
type Veto struct {
	cmd string
	msg string
}

// Some interfaces we want to make sure we implement. If we don't we'll throw
// compile time errors.
var (
	_ has.Veto = &Veto{}
)

// NewVeto returns a new Veto attribute initialised for the specified command
// with the specified message text. The command is a normal command such as GET
// and DROP and will automatically be uppercased. The message text should
// indicate why the command was vetoed such as "You can't drop the sword. It
// seems to be cursed". Referring to specific items - such as the sword in the
// example - is valid as a Veto is for a specific known Thing.
func NewVeto(cmd string, msg string) *Veto {
	return &Veto{strings.ToUpper(cmd), msg}
}

func (v *Veto) Dump() (buff []string) {
	return append(buff, DumpFmt("%p %[1]T %q:%q", v, v.Command(), v.Message()))
}

// Command returns the command associated with the Veto.
func (v *Veto) Command() string {
	return v.cmd
}

// Message returns the message associated with the Veto.
func (v *Veto) Message() string {
	return v.msg
}
