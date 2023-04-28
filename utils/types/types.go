package types

import (
	"html/template"
)

// KeyVal 定义的通用的 key=>val的结构体，主要用于有序的配置输出
type KeyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SimpleId struct {
	Id  string `json:"id"`
	Ids string `json:"ids"`
}

type AutoKey struct {
	Where map[string]string `form:"where" json:"where"`
}

type HtmlShow struct {
	Html template.HTML
}
