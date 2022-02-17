package main

import (
	"context"
	"testing"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

func Test_req(t *testing.T) {

	endpoint := "opc.tcp://127.0.0.1:4840/opcua"
	nodeID := ua.NewStringNodeID(1, "ro_bool")

	ctx := context.Background()

	c := opcua.NewClient(endpoint)
	if err := c.Connect(ctx); err != nil {
		t.Fatal(err)
	}
	defer c.CloseWithContext(ctx)

	id, err := ua.ParseNodeID(nodeID.StringID())
	if err != nil {
		t.Fatalf("invalid node id: %v", err)
	}

	req := &ua.ReadRequest{
		MaxAge: 2000,
		NodesToRead: []*ua.ReadValueID{
			{NodeID: id},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	resp, err := c.ReadWithContext(ctx, req)
	if err != nil {
		t.Fatalf("Read failed: %s", err)
	}
	if resp.Results[0].Status != ua.StatusOK {
		t.Fatalf("Status not OK: %v", resp.Results[0].Status)
	}
	t.Logf("%#v", resp.Results[0].Value.Value())
}
