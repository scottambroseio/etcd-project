package main

import (
	"github.com/coreos/etcd/clientv3"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	cluster := os.Getenv("ETCD_CLUSTER")
	split := strings.Split(cluster, "|")

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   split,
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// bring this in as a env var later on
	id := "node"
	node := NewNode(id, client)

	err = node.RegisterAsWorker()

	if err != nil {
		log.Fatal(err)
	}
}

func createClient() (*clientv3.Client, error) {
	cluster := os.Getenv("ETCD_CLUSTER")
	split := strings.Split(cluster, "|")

	return clientv3.New(clientv3.Config{
		Endpoints:   split,
		DialTimeout: 2 * time.Second,
	})
}
