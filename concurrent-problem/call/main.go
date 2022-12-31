package main

import (
	"os/exec"
	"sync"
)

func task_call_server(wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		cmd := exec.Command("curl", "http://localhost:3000")
		cmd.Run()
		wg.Done()
	}

}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(100000)

	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)
	go task_call_server(&waitGroup)

	waitGroup.Wait()

}
