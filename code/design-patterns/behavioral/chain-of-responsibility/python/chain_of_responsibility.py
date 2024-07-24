from abc import ABC, abstractmethod


class Handler(ABC):

    @abstractmethod
    def handle(self, patient):
        pass

    @abstractmethod
    def set_next(self, handler):
        pass


class Department(Handler):

    def __init__(self):
        self.next = None

    def set_next(self, handler):
        self.next = handler
        return handler

    def handle(self, patient):
        if self.next:
            return self.next.handle(patient)


class Reception(Department):

    def handle(self, patient):
        # Business logic
        if not patient.get("insurance"):
            print("Ending chain no insurance")
            return
        patient["log"] += "Reception -> "
        super().handle(patient)  # Call the next handler


class Doctor(Department):

    def handle(self, patient):
        patient["log"] += "Doctor -> "
        super().handle(patient)  # Call the next handler


class MedicalExaminationRoom(Department):

    def handle(self, patient):
        patient["log"] += "MedicalExaminationRoom -> "
        super().handle(patient)  # Call the next handler


class Cashier(Department):

    def handle(self, patient):
        if patient.get("broke"):
            print("Ending chain patient is broke")
            return
        patient["log"] += "Cashier -> "
        super().handle(patient)  # Call the next handler
