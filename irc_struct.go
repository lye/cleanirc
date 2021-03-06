// Copyright 2009 Thomas Jager <mail@jager.no>  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package irc

import (
	"net"
	"time"
	"crypto/tls"
)

type IRCConnection struct {
	Error chan error
	Log chan string
	Password string
	SSLConfig *tls.Config

	socket net.Conn
	pread, pwrite chan string
	syncreader, syncwriter chan bool
	reconnecting bool

	nick string //The nickname we want.
	nickcurrent string //The nickname we currently have.
	user string
	registered bool
	server string
	events map[string][]func(*IRCEvent)

	lastMessage time.Time
	ticker <-chan time.Time
	ticker2 <-chan time.Time

	VerboseCallbackHandler bool

	quitting bool
}

type IRCEvent struct {
	Code string
	Message string
	Raw string
	Nick string //<nick>
	Host string //<nick>!<usr>@<host>
	Source string //<host>
	User string //<usr>

	Arguments []string
}
