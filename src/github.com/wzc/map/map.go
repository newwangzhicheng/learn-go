package mapdic

//Dictionary is a nickname of map
//永远不要声明一个空map变量。默认值nil会导致运行时错误
type Dictionary map[string]string

//DictionaryErr string的别称，方便添加方法
type DictionaryErr string

const (
	//ErrorNotFound 不存在word返回的错误
	ErrorNotFound = DictionaryErr("Couldn't find word in your dictionary")
	//ErrWordExists Add失败返回的错误
	ErrWordExists = DictionaryErr("Cannot add word because it already existed")
	//ErrorWordNotExist update失败返回的错误
	ErrorWordNotExist = DictionaryErr("Cannot update word because it does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search 返回value
//这是Dictionary的方法
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		//errors.New("") 返回一个error类型的值，表示错误
		return "", ErrorNotFound
	}
	//go 允许返回多值
	return definition, nil
}

//Add add key-value pairs to the Dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

//Update 更新Dictionary的value
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

//Delete 是Dictionary删除key的一种方法
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
