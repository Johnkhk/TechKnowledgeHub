from decorator import (
    EmailNotifier,
    SMSNotifier,
    FacebookNotifier,
    SlackNotifier,
    SimpleNotifier,
)

# Wrap the basic notifier with multiple decorators
notifier = SimpleNotifier()
notifier = EmailNotifier(SMSNotifier(FacebookNotifier(SlackNotifier(notifier))))

notifier.send()
