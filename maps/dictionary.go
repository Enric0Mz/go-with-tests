package maps

import (
	"fmt"
)

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound = DictionaryErr("not found in dictionary")
	ErrConflict = DictionaryErr("word already exists, aborting insertion")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(w string) (string, error) {
	definition, ok := d[w]
	if !ok {
		return "", fmt.Errorf("word '%s' %w", w, ErrNotFound)
	}
	return definition, nil
}

func (d Dictionary) Add(w string, desc string) error {
	_, ok := d[w]
	if ok {
		return ErrConflict
	}
	d[w] = desc
	return nil
}
