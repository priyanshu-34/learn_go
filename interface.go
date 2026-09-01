package main

import "fmt"

type Slack struct {}

func (s Slack) sendMessage(message string){
	fmt.Println("Sending message to Slack: ", message)
}

type Email struct {}

func (e Email) sendMessage(message string){
	fmt.Println("Sending message to Email: ", message)
}

func sendNotificationSlack(s Slack, message string){
	s.sendMessage(message)
}

func sendNotificationEmail(e Email, message string){
	e.sendMessage(message)
}

type Notifier interface {
	sendMessage(message string)
}

func sendNotification(n Notifier, message string){
	n.sendMessage(message)
}

func _interface(){
	fmt.Println("Interface.....")
	
	mySlack := Slack{}
	myEmail := Email{}

	// This is wrong way to use the sendNotification because we had to create two different functions for Slack and Email. Instead we can use interface to make it more generic.
	sendNotificationSlack(mySlack, "Hello from Slack!")
	sendNotificationEmail(myEmail, "Hello from Email!")

	// This is the right way to use the sendNotification function because we are using interface to make it more generic.
	sendNotification(mySlack, "Hello from Slack!")
	sendNotification(myEmail, "Hello from Email!")

}