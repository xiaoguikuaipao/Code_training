package main

import (
	"log"
	"net/http"
)

//var mu sync.Mutex
//var count int

func main() {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//mu.Lock()
	//count++
	//mu.Unlock()
	//fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
	w.Header().Set("Content-Type", "image/gif")
	lissajous(w)
}

//func counter(w http.ResponseWriter, r *http.Request) {
//	mu.Lock()
//	fmt.Fprintf(w, "count %d\n", count)
//	mu.Unlock()
//}
