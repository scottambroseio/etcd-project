package main

import "testing"
import "github.com/coreos/etcd/clientv3"
import "time"
import "context"
import "sync"
import "math/rand"
import "strconv"

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

func TestNonQuorumReads(t *testing.T) {
	client1, err := createClient()
	client2, err := createClient()
	client3, err := createClient()

	if err != nil {
		t.Fatal(err)
	}

	defer client1.Close()
	defer client2.Close()
	defer client3.Close()

	var wg sync.WaitGroup
	wg.Add(3)

	ch := make(chan string, 3)
	rand.Seed(time.Now().UTC().UnixNano())
	key := strconv.Itoa(rand.Int())

	_, err = client1.Put(context.Background(), key, "value")

	if err != nil {
		t.Fatal(err)
	}

	go getKeyNonQuorum(key, &wg, ch, client1)
	go getKeyNonQuorum(key, &wg, ch, client2)
	go getKeyNonQuorum(key, &wg, ch, client3)

	wg.Wait()
}

func getKeyNonQuorum(key string, wg *sync.WaitGroup, ch chan string, client *clientv3.Client) {
	defer wg.Done()

	resp, err := client.Get(context.Background(), key, clientv3.WithSerializable())

	if err != nil {
		panic(err)
	}

	if len(resp.Kvs) > 0 {
		ch <- string(resp.Kvs[0].Value)
	} else {
		ch <- string("no value")
	}
}
