package main

import (
	"fmt"
	"math/rand"
	"time"
)

// OTP is the abstract base class
type OTP interface {
	Process()
	generateOTP(length int) string
	saveOTP(otp string)
	prepareContent()
	sendNotification()
}

// BaseOTP is the base struct that implements the common methods
type BaseOTP struct{}

// generateOTP generates a random OTP
func (b *BaseOTP) generateOTP(length int) string {
	rand.Seed(time.Now().UnixNano())
	otp := ""
	for i := 0; i < length; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}
	fmt.Println("Generated OTP:", otp)
	return otp
}

// saveOTP simulates saving the OTP to cache
func (b *BaseOTP) saveOTP(otp string) {
	fmt.Println("OTP", otp, "saved in cache")
}

// SMSOTP is a concrete class that implements OTP
type SMSOTP struct {
	BaseOTP
}

func (s *SMSOTP) Process() {
	otp := s.generateOTP(6)
	s.saveOTP(otp)
	s.prepareContent()
	s.sendNotification()
}

func (s *SMSOTP) prepareContent() {
	fmt.Println("Preparing SMS content")
}

func (s *SMSOTP) sendNotification() {
	fmt.Println("Sending SMS notification")
}

// EmailOTP is a concrete class that implements OTP
type EmailOTP struct {
	BaseOTP
}

func (e *EmailOTP) Process() {
	otp := e.generateOTP(6)
	e.saveOTP(otp)
	e.prepareContent()
	e.sendNotification()
}

func (e *EmailOTP) prepareContent() {
	fmt.Println("Preparing Email content")
}

func (e *EmailOTP) sendNotification() {
	fmt.Println("Sending Email notification")
}

// AuthenticatorOTP is a concrete class that implements OTP
type AuthenticatorOTP struct {
	BaseOTP
}

func (a *AuthenticatorOTP) Process() {
	otp := a.generateOTP(6)
	a.saveOTP(otp)
	a.prepareContent()
	a.sendNotification()
}

func (a *AuthenticatorOTP) prepareContent() {
	fmt.Println("Preparing Authenticator App content")
}

func (a *AuthenticatorOTP) sendNotification() {
	fmt.Println("Sending Authenticator App notification")
}
