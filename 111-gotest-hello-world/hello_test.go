package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Danny", "")
	want := "Hello, Danny"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloAgain(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("James", "")
		want := "Hello, James"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("If language is spanish say Hola", func(t *testing.T) {
		got := Hello("James", spanish)
		want := "Hola, James"
		assertCorrectMessage(t, got, want)
	})
	t.Run("If language is french say Bonjour", func(t *testing.T) {
		got := Hello("James", french)
		want := "Bonjour, James"
		assertCorrectMessage(t, got, want)
	})
	TestHello(t)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
