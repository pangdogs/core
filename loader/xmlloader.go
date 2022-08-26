package loader

import (
	"encoding/xml"
	"io/ioutil"
)

// XmlLoader XML格式配置加载器
type XmlLoader[T any] struct {
	config T
}

// SetString 输入XML格式内容，string类型
func (loader *XmlLoader[T]) SetString(content string) error {
	return loader.set([]byte(content))
}

// SetBytes 输入XML格式内容，bytes类型
func (loader *XmlLoader[T]) SetBytes(content []byte) error {
	return loader.set(content)
}

// SetFile 输入XML格式文件路径
func (loader *XmlLoader[T]) SetFile(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return loader.set(bytes)
}

func (loader *XmlLoader[T]) set(content []byte) error {
	return xml.Unmarshal(content, &loader.config)
}

// Get 获取解析后的配置
func (loader *XmlLoader[T]) Get() T {
	return loader.config
}
