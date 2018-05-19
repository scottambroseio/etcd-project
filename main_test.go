package main

import "testing"
import "github.com/coreos/etcd/clientv3"
import "os"
import "time"
import "context"
import "strings"

func createClient() (*clientv3.Client, error) {
	cluster := os.Getenv("ETCD_CLUSTER")
	split := strings.Split(cluster, "|")

	return clientv3.New(clientv3.Config{
		Endpoints:   split,
		DialTimeout: 2 * time.Second,
	})
}

func TestCanConnectToETCD(t *testing.T) {
	client, err := createClient()

	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()
}

func TestCanCreateKey(t *testing.T) {
	client, err := createClient()

	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	_, err = client.Put(context.Background(), "key", "value")

	if err != nil {
		t.Fatal(err)
	}
}
