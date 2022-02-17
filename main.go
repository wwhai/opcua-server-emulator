// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"log"

	"github.com/gopcua/opcua/uacp"
)

var conns []*uacp.Conn = make([]*uacp.Conn, 10)

func main() {
	endpoint := "opc.tcp://127.0.0.1:4840/opcua"
	ctx := context.Background()

	log.Printf("Listening on %s", endpoint)
	listener, err := uacp.Listen(endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept(ctx)
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte("qwertyuiop"))
		log.Printf("conn %d: connection from %s", conn.ID(), conn.RemoteAddr())
		conns = append(conns, conn)
	}

}
