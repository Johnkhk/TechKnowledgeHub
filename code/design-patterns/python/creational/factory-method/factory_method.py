from abc import ABC, abstractmethod


class Notification(ABC):
    @abstractmethod
    def send(self, content: str):
        pass


class EmailNotification(Notification):
    def send(self, content: str):
        print(f"Sending email with content: {content}")


class SMSNotification(Notification):
    def send(self, content: str):
        print(f"Sending SMS with content: {content}")


class NotificationFactory(ABC):
    @abstractmethod
    def create_notification() -> Notification:
        # can add complex shared logic here
        pass


class EmailNotificationFactory(NotificationFactory):
    def create_notification(self) -> Notification:
        # can add complex custom logic here
        return EmailNotification()


class SMSNotificationFactory(NotificationFactory):
    def create_notification(self) -> Notification:
        # can add complex custom logic here
        return SMSNotification()
