// Copyright 2015 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

package stats

import (
	"code.wolfmud.org/WolfMUD.git/attr"
	"code.wolfmud.org/WolfMUD.git/has"

	"log"
	"sync"
)

// players represents a list of all players currently in the game world.
var players = struct {
	list []has.Thing
	sync.Mutex
}{}

// Add adds the specified player to the list of players.
func Add(player has.Thing) {
	players.Lock()
	defer players.Unlock()

	players.list = append(players.list, player)
}

// Remove removes the specified player from the list of players.
func Remove(player has.Thing) {
	players.Lock()
	defer players.Unlock()

	for i, p := range players.list {
		if p == player {
			players.list[i] = nil
			players.list = append(players.list[:i], players.list[i+1:]...)
			break
		}
	}

	// A tiny bit of housekeeping, in case we've had a lot of players recently
	// create a new, smaller capacity player list. The capacity of 10 is totally
	// arbitrary - maybe this should be config tweakable?
	if len(players.list) == 0 && cap(players.list) > 10 {
		log.Printf("Last one out reclaims the player list: %d slots reclaimed", cap(players.list)-10)
		players.list = make([]has.Thing, 0, 10)
	}
}

// List returns the names of all players in the player list. The omit parameter
// may be used to specify a player that should be omitted from the list.
func List(omit has.Thing) []string {
	players.Lock()
	defer players.Unlock()

	list := make([]string, 0, len(players.list))

	for _, player := range players.list {
		if player == omit {
			continue
		}
		if a := attr.FindName(player); a != nil {
			list = append(list, a.Name())
		}
	}

	return list
}

// Len returns the length of the player list.
func Len() int {
	players.Lock()
	defer players.Unlock()

	return len(players.list)
}
