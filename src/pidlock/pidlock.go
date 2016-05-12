package pidlock

//package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	//"time"
)

// Choose whether to exit in case of a lockfile creation error
var Exit_on_file_error bool = false

// Choose whether to be silent or verbose
var Verbose = false

func isError(err error, mesg string, hard bool) bool {
	if err != nil {
		if Verbose {
			fmt.Printf("%s:- %v\n", mesg, err)
		}

		if hard {
			os.Exit(1)
		} else {
			return true
		}
	}
	return false
}

func isPID(pid int) bool {
	out, _ := exec.Command("kill", "-s", "0", strconv.Itoa(pid)).CombinedOutput()
	//isError(err, "PID error",false)

	if string(out) == "" {
		return true // pid exists
	}
	return false
}

func getLockingPID(lockfile string) int {
	var opid int
	fd, _ := os.Open(lockfile)
	defer fd.Close()

	//if isError(err, "File can't be read", false){
	//	return -1
	//}

	_, _ = fmt.Fscanf(fd, "%d", &opid)
	//opid, err = strconv.Atoi(opid)

	if opid <= 1 {
		return -1
	} else {
		return opid
	}

}

func writePID(lockfile string, pid int) bool {
	fd, err := os.Create(lockfile)
	defer fd.Close()

	if isError(err, "File write open Error", Exit_on_file_error) {
		return false
	}

	_, err = fmt.Fprintf(fd, "%d", pid)

	if isError(err, "File write Error", Exit_on_file_error) {
		return false
	}
	return true
}

func GetLock(lockfile string) bool {
	pid := os.Getpid()
	opid := getLockingPID(lockfile)

	if opid == -1 {
		if Verbose {
			fmt.Println("PID file missing or scrambled, getting Lock.")
		}

		if !writePID(lockfile, pid) {
			return false
		}
	} else {
		if isPID(opid) {
			if Verbose {
				fmt.Println("Can't get lock.. another instance running")
			}
			return false
		} else {
			//fmt.Println("OPID doesn't exist, getting Lock.")
			if !writePID(lockfile, pid) {
				return false
			}
		}

	}

	return true
}

func ReleaseLock(lockfile string) bool {
	err := os.Remove(lockfile)
	isError(err, "Lockfile removal Error", true)
	return true
}

/*
// Example Usage
func main() {
        var pidfile string = "/tmp/pidlocktest.pid"

        if pidlock.GetLock(pidfile){
                defer pidlock.ReleaseLock(pidfile)
                fmt.Println("This is a Simple Test of Rajani's pidlock module")
                fmt.Println("You see this message because I have obtained the lock!!")
                fmt.Println("Let's wait for 10 secs...")
                time.Sleep(time.Second * 10)

        }else {
                fmt.Println("Couldn't obtain lock.. another instance must be running.")
                fmt.Println("Quitting gracefully")
        }
}
*/
