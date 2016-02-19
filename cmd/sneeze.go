// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package cmd

import (
	"code.wolfmud.org/WolfMUD.git/attr"
)

// Syntax: SNEEZE
func init() {
	AddHandler(Sneeze, "SNEEZE")
}

func Sneeze(s *state) {

	// Try locking locations within a distance of 2 moves
	lockAdded := false
	locs := attr.FindExits(s.where.Parent()).Within(2)
	for _, e1 := range locs {
		for _, e2 := range e1 {
			if !s.CanLock(e2) {
				s.AddLock(e2)
				lockAdded = true
			}
		}
	}

	// If we added any locks return to the parser so we can relock
	if lockAdded {
		return
	}

	// Notify actor
	s.msg.actor.WriteString("You sneeze. Aaahhhccchhhooo!")

	// Notify observers in same location
	who := attr.FindName(s.actor).Name("Someone")
	s.msg.observer.WriteJoin("You see ", who, " sneeze.")

	// Notify observers in near by locations
	for _, e := range locs[1] {
		s.msg.observers[e].WriteString("You hear a loud sneeze.")
	}

	// Notify observers in further out locations
	for _, e := range locs[2] {
		s.msg.observers[e].WriteString("You hear a sneeze.")
	}

	s.ok = true
}
