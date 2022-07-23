package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, def string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	if def != got {
		t.Errorf("got %q want %q", got, def)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defi := "this is just a test"
		err := dictionary.Add(word, defi)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, defi)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defi := "this is just a test"
		dictionary := Dictionary{word: defi}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, defi)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	defi := "this is just a test"
	dictionary := Dictionary{word: defi}
	newDefi := "new defi"

	dictionary.Update(word, newDefi)
	assertDefinition(t, dictionary, word, newDefi)
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defi := "this is just a test"
		dictionary := Dictionary{word: defi}
		newDefi := "new defi"

		err := dictionary.Update(word, newDefi)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefi)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		defi := "this is just a test"
		dictionary := Dictionary{}
		err := dictionary.Update(word, defi)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test defi"}

	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}
