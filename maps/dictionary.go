package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

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
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}
