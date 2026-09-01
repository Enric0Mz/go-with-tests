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

func (d Dictionary) Update(w, desc string) error {
	err := searchWord(d, w)
	if err != nil {
		return err
	}
	d[w] = desc
	return nil
}

func (d Dictionary) Delete(w string) error {
	err := searchWord(d, w)
	if err != nil {
		return err
	}
	delete(d, w)
	return nil
}

func searchWord(d Dictionary, w string) error {
	_, err := d.Search(w)
	if err != nil {
		return fmt.Errorf("cannot perform operation: %w", err)
	}
	return nil
}
