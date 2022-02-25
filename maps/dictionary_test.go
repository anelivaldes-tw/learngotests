package maps

import "testing"

/*Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types.
The first is the key type, which is written inside the []. The second is the value type,
which goes right after the [].
The key type is special. It can only be a comparable type because without
the ability to tell if 2 keys are equal, we have no way to ensure that we are getting the correct value.*/

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrNotFound)
	})
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
