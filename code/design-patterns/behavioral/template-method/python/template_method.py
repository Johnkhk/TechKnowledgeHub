from abc import ABC, abstractmethod
import random


class OTP(ABC):
    def process(self):
        otp = self.generate_otp()
        self.save_otp(otp)
        self.prepare_content()
        self.send_notification()

    def generate_otp(self, length=6):
        otp = "".join([str(random.randint(0, 9)) for _ in range(length)])
        print(f"Generated OTP: {otp}")
        return otp

    def save_otp(self, otp):
        # Simulate saving OTP in cache
        print(f"OTP {otp} saved in cache")

    @abstractmethod
    def prepare_content(self):
        pass

    @abstractmethod
    def send_notification(self):
        pass


class SMSOTP(OTP):
    def prepare_content(self):
        print("Preparing SMS content")

    def send_notification(self):
        print("Sending SMS notification")


class EmailOTP(OTP):
    def prepare_content(self):
        print("Preparing Email content")

    def send_notification(self):
        print("Sending Email notification")


class AuthenticatorOTP(OTP):
    def prepare_content(self):
        print("Preparing Authenticator App content")

    def send_notification(self):
        print("Sending Authenticator App notification")
