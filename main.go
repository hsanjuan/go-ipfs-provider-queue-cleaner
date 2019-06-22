package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ipfs/go-datastore/query"
	badger "github.com/ipfs/go-ds-badger"
)

func main() {
	folder := os.Args[1]
	opts := &badger.DefaultOptions
	opts.ReadOnly = true
	ds, err := badger.NewDatastore(folder+"/badgerds", opts)
	if err != nil {
		log.Fatal(err)
	}

	q := query.Query{
		Prefix:   "/provider-v1/queue",
		KeysOnly: true,
	}

	results, err := ds.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	all, err := results.Rest()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(all), " provider records pending")
}
