from abc import ABC, abstractmethod


class Notifier(ABC):
    @abstractmethod
    def send(self):
        pass


class BaseNotifierDecorator(Notifier):
    def __init__(self, notifier: Notifier):
        self._wrapped_notifier = notifier

    def send(self):
        self._wrapped_notifier.send()


class EmailNotifier(BaseNotifierDecorator):
    def send(self):
        super().send()
        self._send_email()

    def _send_email(self):
        print("Sending email notification.")


class SMSNotifier(BaseNotifierDecorator):
    def send(self):
        super().send()
        self._send_sms()

    def _send_sms(self):
        print("Sending SMS notification.")


class FacebookNotifier(BaseNotifierDecorator):
    def send(self):
        super().send()
        self._send_facebook()

    def _send_facebook(self):
        print("Sending Facebook notification.")


class SlackNotifier(BaseNotifierDecorator):
    def send(self):
        super().send()
        self._send_slack()

    def _send_slack(self):
        print("Sending Slack notification.")


# Example usage
class SimpleNotifier(Notifier):
    def send(self):
        print("Basic notification.")
