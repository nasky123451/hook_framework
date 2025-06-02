package nlp

import (
	"fmt"
	"regexp"
	"strings"
)

type Intent struct {
	Action string
	Params map[string]string
}

type Parser func(input string) (Intent, bool)

type NLP struct {
	parsers []Parser
}

func NewNLP() *NLP {
	return &NLP{}
}

// RegisterParser 註冊解析器
func (n *NLP) RegisterParser(parser Parser) {
	n.parsers = append(n.parsers, parser)
	fmt.Println("[NLP] Parser registered successfully.")
}

// ParseInput 解析輸入，返回匹配的 Intent
func (n *NLP) ParseInput(input string) Intent {
	input = standardizeInput(input) // 標準化輸入文本
	fmt.Printf("[NLP] Standardized input: %s\n", input)

	for _, parser := range n.parsers {
		intent, matched := parser(input)
		if matched {
			fmt.Printf("[NLP] Matched intent: %s with params: %+v\n", intent.Action, intent.Params)
			return intent
		}
	}

	fmt.Println("[NLP] No matching intent found.")
	return Intent{Action: "unknown", Params: map[string]string{}}
}

// standardizeInput 標準化輸入文本
func standardizeInput(input string) string {
	input = strings.ToLower(input)                                 // 轉換為小寫
	input = strings.TrimSpace(input)                               // 去除首尾空格
	input = regexp.MustCompile(`\s+`).ReplaceAllString(input, " ") // 替換多個空格為單個空格
	return input
}

// ExtractName 從輸入文本中提取全名
func ExtractName(input string) string {
	// 假設人名由一個或多個單詞組成，且每個單詞以大寫字母開頭
	nameRegex := regexp.MustCompile(`\b[A-Z][a-z]*(?:\s[A-Z][a-z]*)*\b`)
	matches := nameRegex.FindAllString(input, -1)

	// 返回第一個匹配的全名
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}
