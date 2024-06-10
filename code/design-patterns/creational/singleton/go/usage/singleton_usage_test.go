package usage

import (
	// Standard library packages
	"testing"

	// External package
	impl "tkh/code/design-patterns/creational/singleton/go/implementation"
)

func TestSingleton(t *testing.T) {
	singleton1 := impl.GetInstance()
	singleton2 := impl.GetInstance()

	// Ensure both variables point to the same instance
	if singleton1 != singleton2 {
		t.Errorf("Expected singleton1 and singleton2 to be the same instance, but they are not")
	}

	// Ensure the initial value is correct
	if singleton1.GetValue() != 42 || singleton2.GetValue() != 42 {
		t.Errorf("Expected initial value to be 42, but got %d and %d", singleton1.GetValue(), singleton2.GetValue())
	}

	// Set a new value using one instance
	singleton1.SetValue(100)

	// Ensure both instances reflect the new value
	if singleton1.GetValue() != 100 || singleton2.GetValue() != 100 {
		t.Errorf("Expected value to be updated to 100, but got %d and %d", singleton1.GetValue(), singleton2.GetValue())
	}

	// log that we are done testing
	t.Log("Done testing singleton")
}
