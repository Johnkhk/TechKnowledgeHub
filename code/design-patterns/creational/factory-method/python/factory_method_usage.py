from factory_method import (
    Notification,
    SMSNotification,
    EmailNotification,
    SMSNotificationFactory,
    EmailNotificationFactory,
)


email_factory = EmailNotificationFactory()
email_notification = email_factory.create_notification()
email_notification.send("Hello via Email!")

sms_factory = SMSNotificationFactory()
sms_notification = sms_factory.create_notification()
sms_notification.send("Hello via SMS!")

# tests
assert isinstance(email_notification, Notification)
assert isinstance(email_notification, EmailNotification)
assert isinstance(sms_notification, Notification)
assert isinstance(sms_notification, SMSNotification)
