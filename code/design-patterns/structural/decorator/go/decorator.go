package main

import (
	"fmt"
)

// Notifier is the interface that defines the send method.
type Notifier interface {
	Send()
}

// SimpleNotifier is a basic implementation of Notifier.
type SimpleNotifier struct{}

func (n *SimpleNotifier) Send() {
	fmt.Println("Basic notification.")
}

// EmailNotifier is a concrete decorator that adds email notification.
type EmailNotifier struct {
	notifier Notifier
}

func (n *EmailNotifier) Send() {
	n.notifier.Send()
	n.sendEmail()
}

func (n *EmailNotifier) sendEmail() {
	fmt.Println("Sending email notification.")
}

// SMSNotifier is a concrete decorator that adds SMS notification.
type SMSNotifier struct {
	notifier Notifier
}

func (n *SMSNotifier) Send() {
	n.notifier.Send()
	n.sendSMS()
}

func (n *SMSNotifier) sendSMS() {
	fmt.Println("Sending SMS notification.")
}

// FacebookNotifier is a concrete decorator that adds Facebook notification.
type FacebookNotifier struct {
	notifier Notifier
}

func (n *FacebookNotifier) Send() {
	n.notifier.Send()
	n.sendFacebook()
}

func (n *FacebookNotifier) sendFacebook() {
	fmt.Println("Sending Facebook notification.")
}

// SlackNotifier is a concrete decorator that adds Slack notification.
type SlackNotifier struct {
	notifier Notifier
}

func (n *SlackNotifier) Send() {
	n.notifier.Send()
	n.sendSlack()
}

func (n *SlackNotifier) sendSlack() {
	fmt.Println("Sending Slack notification.")
}
