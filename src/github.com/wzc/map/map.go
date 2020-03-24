package mapdic

import "errors"

//Dictionary is a nickname of map
//永远不要声明一个空map变量。默认值nil会导致运行时错误
type Dictionary map[string]string

//ErrorNotFound is an error
var ErrorNotFound = errors.New("Couldn't find word in your dictionary")

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
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
