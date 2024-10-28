package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func (t *testing.T)  {
		got := Hello("Chris", "")
		want := "Hello, Chris"
	
		assertCorrectMessage(t,got,want)
	})

	t.Run("say 'Hello, World' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want:= "Hello, World"

		assertCorrectMessage(t,got,want)
	})

	t.Run("in Portuguese",func(t *testing.T) {
		got:= Hello("Miguel","Portuguese")
		want := "Olá, Miguel"

		assertCorrectMessage(t,got,want)
	})

	t.Run("in Spanish",func(t *testing.T) {
		got:= Hello("Pedro", "Spanish")
		want:= "Hola, Pedro"

		assertCorrectMessage(t,got ,want) 
			

	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()

	if got != want{
		t.Errorf("got %q want %q",got,want)
	}
}