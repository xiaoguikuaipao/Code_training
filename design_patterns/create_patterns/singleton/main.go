package main

import (
	"fmt"
	"sync"
	"time"
)

type DB struct {
	Database int
}

var db *DB
var once sync.Once

func GetDB(d int) *DB {
	once.Do(func() {
		time.Sleep(1 * time.Second)
		db = &DB{
			Database: d,
		}
	})
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(100)
			db := GetDB(i)
			fmt.Println(i, db.Database)
			wg.Done()
		}()
	}
	wg.Wait()
}
