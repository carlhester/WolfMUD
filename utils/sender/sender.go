// Copyright 2012 Andrew 'Diddymus' Rolfe. All rights reserved.
//
// Use of this source code is governed by the license in the LICENSE file
// included with the source code.

// Package sender provides the sender Interface. Senders should format outgoing
// messages. For an implementation see the Client type.
//
// Typically parsers and senders are paired together to process incomming
// (parser) and outgoing (sender) data over a network connection.
package sender

// Interface should be implemented by anything that wants to send data. This
// is typically anything a user will see such as responses to input, menus and
// messages.
type Interface interface {

	// Send is modelled after fmt.Sprintf and takes parameters in the same way.
	// Send should format the message, add any required prompt to the end and
	// then send the message over the network to the connecting client.
	Send(format string, any ...interface{})

	// SendWithoutPrompt is like Send but sould not append a prompt.
	SendWithoutPrompt(format string, any ...interface{})
}
