package main

import "fmt"

// Main function
func main() {
	// Create the departments
	reception := &Reception{}
	doctor := &Doctor{}
	medicalExaminationRoom := &MedicalExaminationRoom{}
	cashier := &Cashier{}

	// Set up the chain of responsibility
	reception.SetNext(doctor).SetNext(medicalExaminationRoom).SetNext(cashier)

	// Create a patient
	patient := &Patient{
		Insurance: true,
		Broke:     false,
		Log:       "",
	}

	// Handle the patient through the chain of responsibility
	reception.Handle(patient)
	fmt.Println("Patient Log:", patient.Log)
	if patient.Log == "Reception -> Doctor -> MedicalExaminationRoom -> Cashier -> " {
		fmt.Println("Test passed for patient 1")
	} else {
		fmt.Println("Test failed for patient 1")
	}

	// Create another patient without insurance
	patient2 := &Patient{
		Insurance: false,
		Broke:     false,
		Log:       "",
	}

	// Handle the second patient through the chain of responsibility
	reception.Handle(patient2)
	fmt.Println("Patient2 Log:", patient2.Log)
	if patient2.Log == "" {
		fmt.Println("Test passed for patient 2")
	} else {
		fmt.Println("Test failed for patient 2")
	}
}
