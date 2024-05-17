package main

import (
	"reflect"
	"testing"
)

func TestGeneratePrimeNumbers(t *testing.T) {
	stop := make(chan struct{})
	primeChan := make(chan int)

	// test for generating primes up to 10
	expectedPrimesUpTo10 := []int{2, 3, 5, 7}
	go GeneratePrimeNumbers(stop, primeChan, 10)

	receivedPrimes := make([]int, 0)
	for prime := range primeChan {
		receivedPrimes = append(receivedPrimes, prime)
	}

	// close stop channel to terminate generatePrimeNumbers goroutine
	close(stop)

	// check if received primes match the expected primes up to 10
	if !reflect.DeepEqual(receivedPrimes, expectedPrimesUpTo10) {
		t.Errorf("Generated primes mismatch for N=10. Expected %v, got %v", expectedPrimesUpTo10, receivedPrimes)
	}
	// test for generating primes up to 20
	expectedPrimesUpTo20 := []int{2, 3, 5, 7, 11, 13, 17, 19}
	stop = make(chan struct{})
	primeChan = make(chan int)
	go GeneratePrimeNumbers(stop, primeChan, 20)

	receivedPrimes = []int{}
	for prime := range primeChan {
		receivedPrimes = append(receivedPrimes, prime)
	}

	// close stop channel to terminate generatePrimeNumbers goroutine
	close(stop)

	// check if received primes match the expected primes up to 20
	if !reflect.DeepEqual(receivedPrimes, expectedPrimesUpTo20) {
		t.Errorf("Generated primes mismatch for N=20. Expected %v, got %v", expectedPrimesUpTo20, receivedPrimes)
	}
}
