from singleton_implementation import Singleton

# Example usage
singleton1 = Singleton()
singleton2 = Singleton()

# Ensure both variables point to the same instance
assert singleton1 is singleton2

# Ensure the initial value is correct
assert singleton1.get_value() == singleton2.get_value() == 42

# Set a new value using one instance
singleton1.set_value(100)

# Ensure both instances reflect the new value
assert singleton1.get_value() == singleton2.get_value() == 100

print("Singleton works correctly!")
