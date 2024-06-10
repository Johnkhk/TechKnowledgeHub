package implementation

import (
	"sync"
)

// lowercase singleton struct so it is not accessible outside the package (not exported)
type singleton struct {
	value int
}

// declare package level variables
var (
	instance *singleton // Pointer to the singleton instance (initialized to nil)
	once     sync.Once  // used to ensure the singleton is created only once (even with goroutines) (initialized to zero value of sync.Once)
)

// GetInstance returns the singleton instance of singleton
func GetInstance() *singleton { // returns a pointer to the singleton instance
	once.Do(func() {
		// instance is set to the address of a new singleton instance with value 42
		// singleton's constructor is private, so it can only be created within the package
		instance = &singleton{value: 42}
	})
	return instance
}

// GetValue returns the value of the singleton instance
func (s *singleton) GetValue() int {
	return s.value
}

// SetValue sets the value of the singleton instance
func (s *singleton) SetValue(value int) {
	s.value = value
}
