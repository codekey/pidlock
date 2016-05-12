package main

import (
	"fmt"
	"pidlock"
	"time"
)

func main() {
	var pidfile string = "/tmp/pidlocktest.pid"

	if pidlock.GetLock(pidfile){
		defer pidlock.ReleaseLock(pidfile)
		fmt.Println("This is a Simple Test for Rajani's pidlock module")
		fmt.Println("You see this message because I have obtained the lock!!")
		fmt.Println("Let's wait for 10 secs...")
		time.Sleep(time.Second * 10)

	}else {
		fmt.Println("Couldn't obtain lock.. another instance must be running.")
		fmt.Println("Quitting gracefully")
	}



}
