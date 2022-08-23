package ptloader

import (
	"encoding/xml"
	"io/ioutil"
)

// XmlLoader XML格式原型配置加载器
type XmlLoader struct {
	servicePts ServicePts
}

// SetString 输入XML格式内容，string类型
func (loader *XmlLoader) SetString(content string) error {
	return loader.set([]byte(content))
}

// SetBytes 输入XML格式内容，bytes类型
func (loader *XmlLoader) SetBytes(content []byte) error {
	return loader.set(content)
}

// SetFile 输入XML格式文件路径
func (loader *XmlLoader) SetFile(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return loader.set(bytes)
}

func (loader *XmlLoader) set(content []byte) error {
	return xml.Unmarshal(content, &loader.servicePts)
}

// Get 获取解析后的配置
func (loader *XmlLoader) Get() ServicePts {
	return loader.servicePts
}
