package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Database struct {
	c *clientv3.Client
}

func (d *Database) Connect() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panic("Database is not reachable")
	}
	d.c = cli
}

func (d *Database) GetAllJobs() {
	res, err := d.c.Get(context.Background(), "job_", clientv3.WithPrefix())
	if err != nil {
		log.Panic("Can't read jobs from database")
	}
	for _, kv := range res.Kvs {
		log.Printf("key: %s value:\n", kv.Key)
		var obj map[string]interface{}
		json.Unmarshal([]byte(kv.Value), &obj)
		s, _ := colorjson.Marshal(obj)
		fmt.Println(string(s))
	}
}

func (d *Database) Close() {
	err := d.c.Close()
	if err != nil {
		log.Panic("Can't close database session")
	}
}
