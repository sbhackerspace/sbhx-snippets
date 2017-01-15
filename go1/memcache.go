// Steve Phillips / elimisteve
// 2017.01.15

package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Usage: %s <key> [<value>]", path.Base(os.Args[0]))
	}

	mc := memcache.New("localhost:11211")

	key := os.Args[1]

	if len(os.Args) == 2 {
		item, err := mc.Get(key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Get %q -> %q\n", key, item.Value)
		return
	}

	val := os.Args[2]

	item := &memcache.Item{
		Key:   key,
		Value: []byte(val),
	}

	if err := mc.Set(item); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Set %q -> %q\n", key, val)
}
