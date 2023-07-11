package main 
import(
	"testing"
	"sync"
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

func TestWithMutex(t *testing.T) {
	// Run this test by running go test -run TestWithMutex from the terminal in this directory
	// This test will pass. But! is it really passing?
	// Now, run this test by running: 
	//    go test -run  TestWithMutex --race 
	// This test will pass because we are using a mutex to lock the msg variable. This ensures that only one
	// goroutine can access the msg variable at a time. 
	
	var mutex sync.Mutex
	msg = "Hello, World!"

	wg.Add(2)
	go updateMsgWithMutex("Goodbye, cruel world!", &mutex)
	go updateMsgWithMutex("I'm Back!", &mutex)
	wg.Wait()
	
	
	// The reason msg is equal to "I'm Back!" is because the goroutine that updates msg to "I'm Back!" is
	// the last goroutine to run. If you run this test multiple times, you will see that msg is sometimes
	// equal to "Goodbye, cruel world!" and sometimes equal to "I'm Back!". This is because the goroutines
	// are racing to update the msg variable. The last goroutine to update the msg variable, in this case, is the one that
	// will be returned. But you can't rely on this behavior. goroutines are non-deterministic and scheduled by the Go runtime.
	// whichever goroutine gets to the msg variable first will be the one that updates it. 
	// This is not a good assertion to check.
	if msg != "I'm Back!" {
		t.Errorf("msg should be Hello, but got %s", msg)
	}

	if msg != "I'm Back!" || msg != "Goodbye, cruel world!" {
		t.Errorf("msg should be Goodbye, cruel world! or I'm Back!, but got %s", msg)
	}

}