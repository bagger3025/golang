package mydicts

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("there are no such word")
var errWordExists = errors.New("the word already exists")

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if ok {
		return val, nil
	} else {
		return "", errNotFound
	}
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	/*if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil*/
}
