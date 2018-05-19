package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"testing"
)

func TestNewNode(t *testing.T) {
	id := "id"
	client := &clientv3.Client{}
	node := NewNode(id, client)

	if node.ID != id {
		t.Errorf("Expected id to be %s, got %s", id, node.ID)
	}

	if node.Client != client {
		t.Errorf("Expected client to be %v, got %v", client, node.Client)
	}
}

func TestRegisterAsWorker(t *testing.T) {
	client, err := createClient()

	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	node := &Node{
		ID:     "node",
		Client: client,
	}

	err = node.RegisterAsWorker()

	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Get(context.Background(), "workers/node")

	if err != nil {
		t.Fatal(err)
	}

	if resp.Count != 1 {
		t.Errorf("Expected 1 result, got %d", resp.Count)
	}

	if v := resp.Kvs[0].Value; string(v) != node.ID {
		t.Errorf("Expected keys value to be %s, got %s", node.ID, string(v))
	}
}
