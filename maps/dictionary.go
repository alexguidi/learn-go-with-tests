/*
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps

In this section, we covered a lot. We made a full CRUD (Create, Read, Update and Delete) API for our dictionary. Throughout the process we learned how to:
- Create maps
- Search for items in maps
- Add new items to maps
- Update items in maps
- Delete items from a map
- Learned more about errors
  - How to create errors that are constants
  - Writing error wrappers
*/

package maps

//Dictionary is a type that store words in a map
type Dictionary map[string]string

//ErrNotFound is an error for when we did not find the word inside of a dictionary
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

//DictionaryErr make errors reusable and immutable
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search for words inside of a dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add a word to dictionary
func (d Dictionary) Add(word, definition string) error {
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

//Update a definition inside of a dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

//Delete a word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
