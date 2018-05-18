package main

import "testing"
import "github.com/coreos/etcd/clientv3"
import "os"
import "time"
import "context"

func TestCanConnectToETCD(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD")},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()
}

func TestCanCreateKey(t *testing.T) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD")},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		t.Fatal(err)
	}

	defer client.Close()

	_, err = client.Put(context.Background(), "key", "value")

	if err != nil {
		t.Fatal(err)
	}

}
