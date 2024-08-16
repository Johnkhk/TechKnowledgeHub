package main

func main() {
	smsOTP := &SMSOTP{}
	smsOTP.Process()

	emailOTP := &EmailOTP{}
	emailOTP.Process()

	authenticatorOTP := &AuthenticatorOTP{}
	authenticatorOTP.Process()
}
