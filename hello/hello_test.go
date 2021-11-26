package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}

	}

	t.Run("sayHelloToPeople", func(t *testing.T) {
		got := Hello("jiejie", "")
		want := "Hello, jiejie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world' ", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("inChinese", func(t *testing.T) {
		got := Hello("jiejie", "zh")
		want := "你好, jiejie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("inFrench", func(t *testing.T) {
		got := Hello("jiejie", "fr")
		want := "Bonjour, jiejie"
		assertCorrectMessage(t, got, want)
	})

}
