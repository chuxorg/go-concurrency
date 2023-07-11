package workgroups

import (
	"strings"
	"testing"
)
func Test_printSomething(t *testing.T){
	
	workgroupExample()
	
	// // get a reference to the original standard output
	// stdOut := os.Stdout
	// // create a pipe for reading and writing
	// r, w, _ := os.Pipe()
	// // set the standard output to the pipe
	// os.Stdout = w
	// // create a wait group
	// var wg sync.WaitGroup
	// // add 1 to the wait group
	// wg.Add(1)
    // // call the printSomething function
	// go printSomething("Omicron", &wg)
    // // wait for the wait group to finish
	// wg.Wait()
	// // close the pipe
	// _ = w.Close()
    // // read the output from the pipe
	// result, _ := io.ReadAll(r)
	// // convert the output to a string
	// output := string(result)

	// os.Stdout = stdOut

	if !strings.Contains(output, "Omicron") {
		t.Errorf("printSomething() = %v, want %v", output, "Omicron")
	}

}