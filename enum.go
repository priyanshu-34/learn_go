package main

import "fmt"

//There is no enum in go but we can use const and iota to create enum like functionality

type Status int;

const (
	StatusPending Status = iota
	StatusRunning 
	StatusDone 
)

func handle(status Status){
	if(status == StatusPending){
		fmt.Println("waiting")
	}
	if(status == StatusRunning){
		fmt.Println("running")
	}
	if(status == StatusDone){
		fmt.Println("done")
	}
	if(status == 3){
		fmt.Println("invalid status")
	}
}

func enum() {
	fmt.Println("Enum.....")

	handle(StatusPending)
	handle(StatusDone)
	
}