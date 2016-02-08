// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package cmd

import (
	"code.wolfmud.org/WolfMUD.git/attr"
)

// Syntax: READ item
func init() {
	AddHandler(Read, "READ")
}

func Read(s *state) {

	if len(s.words) == 0 {
		s.msg.actor.WriteString("Did you want to read something specific?")
		return
	}

	var (
		name = s.words[0]

		writing string
	)

	// Try searching inventory where we are
	what := s.where.Search(name)

	// If we are somewhere and item still not found try searching narratives
	// where we are
	if what == nil && s.where != nil {
		if a := attr.FindNarrative(s.where.Parent()); a != nil {
			what = a.Search(name)
		}
	}

	// If item still not found try our own inventory
	if what == nil {
		what = attr.FindInventory(s.actor).Search(name)
	}

	// Was item to read found?
	if what == nil {
		s.msg.actor.WriteJoin("You see no '", name, "' to read.")
		return
	}

	// Get item's proper name
	name = attr.FindName(what).Name("something")

	// Find if item has writing
	if a := attr.FindWriting(what); a != nil {
		writing = a.Writing()
	}

	// Was writing found?
	if writing == "" {
		s.msg.actor.WriteJoin("You see no writing on ", name, " to read.")
		return
	}

	s.msg.actor.WriteJoin("You read the writing on ", name, ". It says: ", writing)
	s.ok = true
}
