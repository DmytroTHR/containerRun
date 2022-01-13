package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"sync"
)

var projectName = "containersrun"
var projectNetwork = "dockernet"
var serviceToRun = "service_b"
var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Application A run..")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runAnotherServiceInDocker(strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println("Application A end.")
}

func runAnotherServiceInDocker(i string) {
	defer wg.Done()
	command := exec.Command("docker",
		"run",
		"--detach",
		"--rm",
		"--network", projectName+"_"+projectNetwork,
		"--env", "SCOOTER_ID=11"+i,
		"-p", "808"+i+":8080",
		projectName+"_"+serviceToRun)
	containerID, err := command.Output()
	if err != nil {
		log.Panicf("failed to create new container for ID <%s>: %v", i, err)
	}
	renameCommand := exec.Command("docker",
		"rename",
		string(containerID),
		serviceToRun+"_"+i)
	err = renameCommand.Run()
	if err != nil {
		log.Panicf("failed to rename container <%s>: %v", containerID, err)
	}
}
