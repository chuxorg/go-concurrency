package raceconditions

import (
	"testing"
)

func TestRaceCondition(t *testing.T) {
	// Run this test by running go test -run TestRaceCondition from the terminal in this directory
	// This test will pass. But! is it really passing?
	//
	// Now, run this test by running: 
	//    go test -run  TestRaceCondition --race 
	// from the terminal in this directory. After running this test, you should see an error indicating
	// a race condition. This is because the test is running two goroutines at the same time and they are
	// "racing" to update the msg variable. One go routine is going to get there first and the other will 
	// update the value but not return to this method. The state of "msg" is unknown because it is scoped at the
	// module level and not thread safe. This is where problems arise. A unit test passes and the code is deployed 
	// with a potential, almost certain, bug of a Race Condition. This is a simple example but it can be much more
	// complex and difficult to debug. When testing multi-threaded routines, always test with the  --race flag.

	msg = "Hello, World!"

	wg.Add(2)
	go updateMsg("Goodbye, cruel world!")
	go updateMsg("Goodbye, cruel world!")
	wg.Wait()
	
	if msg != "Goodbye, cruel world!" {
		t.Errorf("msg should be Hello, but got %s", msg)
	}
}

