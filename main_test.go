package main

import "testing"
import "github.com/coreos/etcd/clientv3"
import "os"
import "time"

func TestCanConnectToETCD(t *testing.T) {
	_, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD")},
		DialTimeout: 2 * time.Second,
	})

	if err != nil {
		t.Error(err)
	}
}
