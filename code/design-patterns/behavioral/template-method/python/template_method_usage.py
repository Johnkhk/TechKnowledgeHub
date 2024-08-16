from template_method import SMSOTP, EmailOTP, AuthenticatorOTP

sms_otp = SMSOTP()
sms_otp.process()

email_otp = EmailOTP()
email_otp.process()

authenticator_otp = AuthenticatorOTP()
authenticator_otp.process()
