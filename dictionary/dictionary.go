package dictionary

import "errors"

var errNotFound = errors.New("word not in the dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", errNotFound
	}

	return res, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
