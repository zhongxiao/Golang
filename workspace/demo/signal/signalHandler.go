/*
 *  @author zhongxiao.yzx
 *  @date   2017/12/10
 *  @description
 *    signal handling demo
 *    signal Notify will add rather than overwrite
 *    the signal set of handler(chan)
 */
// signalHandler.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var sigRecV1 chan os.Signal // = make(chan os.Signal, 1)
var sigRecV2 chan os.Signal // = make(chan os.Signal, 1)
var wg sync.WaitGroup

func RegisterSignalHandler() {
	sigRecV1 = make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("set notification for %s [sigRecV1]\n", sigs1)
	signal.Notify(sigRecV1, sigs1...)

	sigRecV2 = make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGQUIT}
	fmt.Printf("set notification for %s [sigRecV2]\n", sigs2)
	signal.Notify(sigRecV2, sigs2...)
}

func LoopForSignal() {
	wg.Add(2)
	go func() {
		for sig := range sigRecV1 {
			fmt.Printf("received a signal from sigRecV1: %s\n", sig)
		}
		fmt.Printf("End. [sigRecV1]\n")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecV2 {
			fmt.Printf("received a signal for recRecV2: %s\n", sig)
		}
		fmt.Printf("End. [sigRecV2]\n")
		wg.Done()
	}()
	fmt.Println("End. LoopForSignal")
}

func main() {
	RegisterSignalHandler()
	LoopForSignal()

	fmt.Println("wait 5s to add SIGINT handler for sigRecV2")
	time.Sleep(5 * time.Second)
	signal.Notify(sigRecV2, syscall.SIGINT)
	fmt.Println("add SIGINT handler for sigRecV2 complete")

	fmt.Println("wait 5s to add SIGQUIT handler for sigRecV2")
	time.Sleep(5 * time.Second)
	signal.Notify(sigRecV2, syscall.SIGQUIT)
	fmt.Println("add SIGQUIT handler for sigRecV2 complete")

	fmt.Println("wait 5s to stop sigRecV1")
	time.Sleep(5 * time.Second)
	signal.Stop(sigRecV1)
	close(sigRecV1)

	fmt.Println("wait 5s to stop sigRecV2")
	time.Sleep(5 * time.Second)
	signal.Stop(sigRecV2)
	close(sigRecV2)

	wg.Wait()
	fmt.Println("done")
}
