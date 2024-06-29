package main

// Example usage
func main() {
	// Wrap the basic notifier with multiple decorators
	var notifier Notifier = &SimpleNotifier{}
	notifier = &EmailNotifier{notifier}
	notifier = &SMSNotifier{notifier}
	notifier = &FacebookNotifier{notifier}
	notifier = &SlackNotifier{notifier}
	notifier.Send()
}
