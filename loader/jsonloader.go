package loader

import (
	"encoding/json"
	"io/ioutil"
)

// JsonLoader Json格式配置加载器
type JsonLoader[T any] struct {
	config T
}

// SetString 输入Json格式内容，string类型
func (loader *JsonLoader[T]) SetString(content string) error {
	return loader.set([]byte(content))
}

// SetBytes 输入Json格式内容，bytes类型
func (loader *JsonLoader[T]) SetBytes(content []byte) error {
	return loader.set(content)
}

// SetFile 输入Json格式文件路径
func (loader *JsonLoader[T]) SetFile(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return loader.set(bytes)
}

func (loader *JsonLoader[T]) set(content []byte) error {
	return json.Unmarshal(content, &loader.config)
}

// Get 获取解析后的配置
func (loader *JsonLoader[T]) Get() T {
	return loader.config
}
