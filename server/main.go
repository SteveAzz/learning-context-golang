package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SteveAzz/context-lesson/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(42), int64(100))

	log.Println(ctx, "Handler stated")
	defer log.Println(ctx, "Handler ended")

	fmt.Printf("value for foo is: %s", ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
