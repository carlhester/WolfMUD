// Copyright 2012 Andrew 'Diddymus' Rolfe. All rights resolved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

// Package world holds references to all of the locations in the world and
// accepts new client connections.
package world

import (
	"fmt"
	"log"
	"net"
	"wolfmud.org/client"
	"wolfmud.org/entities/location"
	"wolfmud.org/entities/mobile/player"
	"wolfmud.org/entities/thing"
	"wolfmud.org/utils/stats"
)

// greeting is displayed when a new client connects.
//
// TODO: Soft code with rest of settings.
const (
	greeting = `

WolfMUD © 2012 Andrew 'Diddymus' Rolfe

    World
    Of
    Living
    Fantasy

`
)

// World represents a single game world. It has references to all of the
// locations available.
type World struct {
	locations []location.Interface
}

// Create brings a new world into existance and returns a reference to it.
func Create() *World {
	return &World{}
}

// Genesis starts the world - what else? :)
func (w *World) Genesis() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile);

	log.Println("Starting WolfMUD server...")

	addr, err := net.ResolveTCPAddr("tcp", "localhost:4001")
	if err != nil {
		log.Printf("Error resolving TCP address, %s\nServer will now exit.\n", err)
		return
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Printf("Error setting up listener, %s\nServer will now exit.\n", err)
		return
	}

	log.Printf("Accepting connections on: %s\n", addr)

	stats.Start()

	for {
		if conn, err := listener.AcceptTCP(); err != nil {
			log.Printf("Error accepting connection: %s\nServer will now exit.\n", err)
			return
		} else {
			log.Printf("Connection from %s.\n", conn.RemoteAddr().String())
			w.startPlayer(conn)
		}
	}
}

func (w *World) startPlayer(conn *net.TCPConn) {
	c := client.New(conn)
	p := player.New(w)

	p.AttachClient(c)

	c.SendWithoutPrompt(greeting)
	w.locations[0].Add(p)
	p.Parse("LOOK")
	w.locations[0].Broadcast([]thing.Interface{p}, "There is a puff of smoke and %s appears spluttering and coughing.", p.Name())

	log.Printf("Connection %s allocated %s, %d players online.\n", conn.RemoteAddr().String(), p.Name(), player.PlayerList.Length())

	go c.Start()
}

func (w *World) AddLocation(l location.Interface) {
	w.locations = append(w.locations, l)
}

func (w *World) Broadcast(ommit []thing.Interface, format string, any ...interface{}) {

	msg := fmt.Sprintf("\n"+format, any...)

OMMIT:
	for _, p := range player.PlayerList.List() {
		for _, o := range ommit {
			if o.IsAlso(p) {
				continue OMMIT
			}
		}
		p.Respond(msg)
	}
}