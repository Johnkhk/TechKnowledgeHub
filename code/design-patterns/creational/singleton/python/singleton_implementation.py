class Singleton:
    _instance = None

    def __new__(cls):
        if cls._instance is None:  # check if the static instance already exists
            # calls __new__ method of the parent (object) class
            # basically calling object.__new__(Singleton)
            cls._instance = super(Singleton, cls).__new__(cls)
        return cls._instance

    def __init__(self):
        if not hasattr(self, "initialized"):
            self.initialized = True  # Avoid reinitialization
            self.value = 42

    # Public methods (business logic)
    def get_value(self):
        return self.value

    def set_value(self, value):
        self.value = value
