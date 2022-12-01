package mydicts

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("there are no such word")
	errWordExists = errors.New("the word already exists")
	errCantUpdate = errors.New("cant update non-existing word")
	errCantDelete = errors.New("cant delete non-existing word")
)

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

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
