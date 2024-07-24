from chain_of_responsibility import Reception, Doctor, MedicalExaminationRoom, Cashier

reception = Reception()
doctor = Doctor()
medical_examination_room = MedicalExaminationRoom()
cashier = Cashier()

reception.set_next(doctor).set_next(medical_examination_room).set_next(cashier)

patient = {"insurance": True, "broke": False, "log": ""}
reception.handle(patient)
assert (
    patient.get("log") == "Reception -> Doctor -> MedicalExaminationRoom -> Cashier -> "
)

patient2 = {"insurance": False, "broke": False, "log": ""}
reception.handle(patient2)
print(patient2.get("log"))
assert patient2.get("log") == ""
