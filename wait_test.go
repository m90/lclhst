package lclhst

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

var (
	okHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	errorHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusInternalServerError)
	})
)

const (
	port = 9786
)

func TestWait_OK(t *testing.T) {
	srv := http.Server{
		Handler: okHandler,
		Addr:    fmt.Sprintf(":%d", port),
	}
	go func() {
		time.Sleep(time.Second * 2)
		srv.ListenAndServe()
	}()

	err := wait(context.Background(), port)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	srv.Close()
}
func TestWait_Error(t *testing.T) {
	srv := http.Server{
		Handler: errorHandler,
		Addr:    fmt.Sprintf(":%d", port),
	}
	go func() {
		time.Sleep(time.Second * 2)
		srv.ListenAndServe()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := wait(ctx, port)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	srv.Close()
}

func TestWait_Timeout(t *testing.T) {
	srv := http.Server{
		Handler: okHandler,
		Addr:    fmt.Sprintf(":%d", port),
	}
	go func() {
		time.Sleep(time.Second * 25)
		srv.ListenAndServe()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := wait(ctx, port)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	srv.Close()
}
