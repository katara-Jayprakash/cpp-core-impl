package threadpool

import "fmt"

// i got the job from user or by client what i have to do
var name string = "Jay prakash"

type jobs func() error

/*  for me the job can be anything, like something that can be executable so for that i have to make it function      take parameter and return error i convert this job into function so i use my custom structure
// and now i am using closure for making it function*/

var myJob jobs = func() error {
	fmt.Println("Job coming from outside", name)
	return nil
}

func main() {
	//creating a queue which can store these all jobs
	// i am not implementing queue data strucute since we got channel we are going to use it
	jobsQueue := make(chan jobs, 30)

}
