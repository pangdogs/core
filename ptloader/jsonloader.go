// Package ptloader 提供原型配置加载器
package ptloader

import (
	"encoding/json"
	"io/ioutil"
)

// JsonLoader Json格式原型配置加载器
type JsonLoader struct {
	servicePts ServicePts
}

// SetString 输入Json格式内容，string类型
func (loader *JsonLoader) SetString(content string) error {
	return loader.set([]byte(content))
}

// SetBytes 输入Json格式内容，bytes类型
func (loader *JsonLoader) SetBytes(content []byte) error {
	return loader.set(content)
}

// SetFile 输入Json格式文件路径
func (loader *JsonLoader) SetFile(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return loader.set(bytes)
}

func (loader *JsonLoader) set(content []byte) error {
	return json.Unmarshal(content, &loader.servicePts)
}

// Get 获取解析后的配置
func (loader *JsonLoader) Get() ServicePts {
	return loader.servicePts
}
