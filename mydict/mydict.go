package mydicts

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("there are no such word")

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if ok {
		return val, nil
	} else {
		return "", errNotFound
	}
}
