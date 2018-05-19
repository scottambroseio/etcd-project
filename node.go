package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

// Node represents the running instance
type Node struct {
	*clientv3.Client
	ID string
}

// NewNode creates an instance of a node with the given id
func NewNode(id string, client *clientv3.Client) *Node {
	return &Node{
		Client: client,
		ID:     id,
	}
}

// RegisterAsWorker registers the node as an availible worker
func (node *Node) RegisterAsWorker() error {
	// todo: proper use of contexts

	// create lease with ttl
	resp, err := node.Client.Grant(context.TODO(), 60)

	// todo: handler errs rather than just return
	if err != nil {
		return err
	}

	key := fmt.Sprintf("workers/%s", node.ID)

	// create leased key
	_, err = node.Client.Put(context.TODO(), key, node.ID, clientv3.WithLease(resp.ID))

	// todo: handler errs rather than just return
	if err != nil {
		return err
	}

	// the lease will be kept as long as this node is active
	// basically an empheral node in zookeeper terms
	ch, err := node.Client.KeepAlive(context.TODO(), resp.ID)

	// todo: handler errs rather than just return
	if err != nil {
		return err
	}

	// drain ch as documented in the docs

	// KeepAlive keeps the given lease alive forever. If the keepalive response
	// posted to the channel is not consumed immediately, the lease client will
	// continue sending keep alive requests to the etcd server at least every
	// second until latest response is consumed.
	<-ch

	return nil
}
