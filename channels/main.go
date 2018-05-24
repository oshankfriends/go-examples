package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var funcMap = make(map[string]func(context.Context) (string, error))

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(25)
}

func fun1(ctx context.Context) (string, error) {
	t0 := time.Now()
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("%v from fun1", ctx.Err())
	case <-time.After(time.Duration(random()) * time.Second):
		return fmt.Sprintf("fun1 in %v sec", time.Since(t0).Seconds()), nil
	}

}

func fun2(ctx context.Context) (string, error) {
	t0 := time.Now()
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("%v from fun2", ctx.Err())
	case <-time.After(time.Duration(random()) * time.Second):
		return fmt.Sprintf("fun1 in %v sec", time.Since(t0).Seconds()), nil
	}

}

func fun3(ctx context.Context) (string, error) {
	t0 := time.Now()
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("%v from fun3", ctx.Err())
	case <-time.After(time.Duration(random()) * time.Second):
		return fmt.Sprintf("fun1 in %v sec", time.Since(t0).Seconds()), nil
	}

}
func fun4(ctx context.Context) (string, error) {
	t0 := time.Now()
	select {
	case <-ctx.Done():
		return "", fmt.Errorf("%v from fun4", ctx.Err())
	case <-time.After(time.Duration(random()) * time.Second):
		return fmt.Sprintf("fun1 in %v sec", time.Since(t0).Seconds()), nil
	}

}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	respCh := make(chan string)
	sigChan := make(chan os.Signal)
	for name, fun := range funcMap {
		go func(name string, fun func(context.Context) (string, error)) {
			fmt.Printf("calling %s\n", name)
			resp, err := fun(ctx)
			if err == nil {
				respCh <- resp
				return
			}

		}(name, fun)
	}
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case resp := <-respCh:
		fmt.Println(resp)
	}
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println(<-sigChan)
}

func init() {
	funcMap["fun1"] = fun1
	funcMap["fun2"] = fun2
	funcMap["fun3"] = fun3
	funcMap["fun4"] = fun4
}
