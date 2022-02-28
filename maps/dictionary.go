package maps

type Dictionary map[string]string

/*
	We made the errors constant; this required us to create our
	own DictionaryErr type which implements the error interface.
	Simply put, it makes the errors
	more reusable and immutable.
*/

type DictionaryErr string

const ErrNotFound = DictionaryErr("could not find the word you were looking for")
const ErrWordExists = DictionaryErr("cannot add word because it already exists")

func (e DictionaryErr) Error() string {
	return string(e)
}

// IMPORTANT
//A gotcha with maps is that they can be a nil value.
//A nil map behaves like an empty map when reading,
//but attempts to write to a nil map will cause a runtime panic.
//Therefore, you should never initialize an empty map variable:
// var m map[string]string NOOOOOOO
// Instead, you can initialize an empty map like
//we were doing above, or use the make keyword
//to create a map for you:
// var dictionary = map[string]string{}
//
//OR
//
//var dictionary = make(map[string]string)
// Both approaches create an empty hash map and point dictionary at it.
//Which ensures that you will never get a runtime panic.

func (d Dictionary) Search(word string) (string, error) {

	/*	A two-value assignment tests for the existence of a key:

		i, ok := m["route"]
		In this statement, the first value (i) is assigned the value stored
		under the key "route". If that key doesn’t exist, i is the value type’s
		zero value (0). The second value (ok) is a bool that is true if the key
		exists in the map, and false if not.

		To test for a key without retrieving the value,
		use an underscore in place of the first value:

		_, ok := m["route"]
	*/

	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word string, definition string) error {
	// Only add the word if it does not exist
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) {
	d[word] = definition
}
