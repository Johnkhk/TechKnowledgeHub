package main

import (
	"fmt"
)

// Patient struct
type Patient struct {
	Insurance bool
	Broke     bool
	Log       string
}

// Handler interface
type Handler interface {
	Handle(patient *Patient)
	SetNext(handler Handler) Handler
}

// Department struct
type Department struct {
	next Handler
}

func (d *Department) SetNext(handler Handler) Handler {
	d.next = handler
	return handler
}

func (d *Department) Handle(patient *Patient) {
	if d.next != nil {
		d.next.Handle(patient)
	}
}

// Reception struct
type Reception struct {
	Department
}

func (r *Reception) Handle(patient *Patient) {
	if !patient.Insurance {
		fmt.Println("Ending chain - no insurance")
		return
	}
	patient.Log += "Reception -> "
	r.Department.Handle(patient)
}

// Doctor struct
type Doctor struct {
	Department
}

func (d *Doctor) Handle(patient *Patient) {
	patient.Log += "Doctor -> "
	d.Department.Handle(patient)
}

// MedicalExaminationRoom struct
type MedicalExaminationRoom struct {
	Department
}

func (m *MedicalExaminationRoom) Handle(patient *Patient) {
	patient.Log += "MedicalExaminationRoom -> "
	m.Department.Handle(patient)
}

// Cashier struct
type Cashier struct {
	Department
}

func (c *Cashier) Handle(patient *Patient) {
	if patient.Broke {
		fmt.Println("Ending chain - patient is broke")
		return
	}
	patient.Log += "Cashier -> "
	c.Department.Handle(patient)
}
