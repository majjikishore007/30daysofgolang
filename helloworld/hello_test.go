package main

import (
	"testing"
)

// Test 1
// func TestHello(t * testing.T)  {
// 	got := Hello()
// 	want := "Hello, World!"
// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

// func TestHello(t *testing.T) {
// 	got := Hello("Chris")
// 	want := "Hello, Chris"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }


func TestHello(t *testing.T)  {
	assertCorrectMessage:= func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		
	}
	t.Run("In Spanish",func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In French",func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})

	// t.Run("saying hello to people",func(t *testing.T) {
	// 	got := Hello("Chris")
	// 	want := "Hello, Chris"
	// 	assertCorrectMessage(t, got, want)
	// })
	// t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello, World!"
	// 	assertCorrectMessage(t, got, want)
	// })
}