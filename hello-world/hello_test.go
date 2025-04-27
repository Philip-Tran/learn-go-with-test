package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Print when there is a name", func(t *testing.T) {
		got := Hello("Quyet", "")
		want := "Hello, Quyet"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Print when there is NOT a name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Spanish", func(t *testing.T) {
		got := Hello("Quyet", "Spanish")
		want := "Hola, Quyet"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("French", func(t *testing.T) {
		got := Hello("Quyet", "French")
		want := "Bonjour, Quyet"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}