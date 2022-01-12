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

	stopPatternService()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runAnotherServiceInDocker(strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println("Application A end.")
}

func stopPatternService() {
	_ = exec.Command("docker", "stop", serviceToRun).Run()
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
	log.Println(command.String())
	err := command.Run()
	if err != nil {
		log.Panic(err)
	}
}
