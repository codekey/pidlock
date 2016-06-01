package pidlock_test

import (
	//"fmt"
	"github.com/rajni4go/GO/pidlock"
        "testing"
)

var pidfile string = "/tmp/pidlocktest.pid"

func TestGetLock(t *testing.T) {
	if ! pidlock.GetLock(pidfile){
		t.Log("Couldn't obtain lock.")
	}
}
 
func TestReleaseLock(t *testing.T) {
	if ! pidlock.ReleaseLock(pidfile) {
		t.Log("Countldn't release lock.")
	}
}

