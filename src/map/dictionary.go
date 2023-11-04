package main

// mapは参照型のため常に元の値を指すことができる
// そのため空のmapを定義しているとnilマップへの参照になりパニックがおきる
// var m map[string]stringで値の指定がないなど
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)で作成する
type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// mapからはそのキーの値と値があるかどうかのboolが変える
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

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

func (d Dictionary) Delete(word string) {
	// 存在しない値を削除しても何もエラーなど発生しないのでエラーハンドリング不要
	delete(d, word)
}
