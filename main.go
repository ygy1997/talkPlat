package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Rule struct {
	Entities map[string][]string `yaml:"实体"`
	Rules    []string            `yaml:"规则"`
}

func main() {
	// 解析YAML文件
	data, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		panic(err)
	}

	// 解析YAML
	var rule Rule
	err = yaml.Unmarshal(data, &rule)
	if err != nil {
		panic(err)
	}

	ac := AC.NewAhoCorasick()
	for role, entity := range rule.Entities {
		for _, e := range entity {
			ac.AddPattern(e)
		}
	}

	ac.BuildFailPointers()
	text := "我想看林俊杰和李小龙的猛龙过江"

	// 在文本中搜索匹配的规则
	matches := ac.Search(text)
	fmt.Println(matches)
}
