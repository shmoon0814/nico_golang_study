package mydict

import "errors"

//Dictionary
type Dictionary map[string]string

var (
	errNotFound   = errors.New("not foud")
	errWordExists = errors.New("word is exist")
	errNotUpdate  = errors.New("Update fail")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	} else {
		return "", errNotFound
	}
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	if err != nil {
		d[word] = def
	} else {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotUpdate:
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
