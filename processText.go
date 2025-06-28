package main

import (
	"strings"
)

func get(text string, i *int) byte {
	if *i >= len(text) {
		return 0
	}
	return text[*i]
}
func getS(text string, i *int) string {
	if *i >= len(text) {
		return ""
	}
	return string(text[*i])
}
func advance(i *int) {
	(*i)++
}
func endLine(text string, i *int) bool {
	return *i >= len(text)
}
func readArray(text string, i *int) (list []string) {
	var value string
	c := get(text, i)
	if c != '[' {
		return list
	}
	advance(i)
	for {
		c = get(text, i)
		if endLine(text, i) {
			break
		}
		if c == ']' {
			list = append(list, value)
			value = ""
			advance(i)
			break
		}
		if c == ',' {
			list = append(list, value)
			value = ""
		} else {
			value += getS(text, i)
		}
		advance(i)
	}
	return
}
func readString(text string, i *int) string {
	value := ""
	for {
		if endLine(text, i) {
			break
		}
		c := get(text, i)
		if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
			value += getS(text, i)
		} else {
			break
		}
		advance(i)
	}
	return value
}
func processLine(text string, i *int) (inputs []string, outputs []string, nodeName string) {
	inputs = readArray(text, i)
	nodeName = readString(text, i)
	outputs = readArray(text, i)
	return
}

type NodeConnection struct {
	inputs, outputs []string
	nodeName        string
}

func processText(text string) []NodeConnection {
	connections := []NodeConnection{}
	text = strings.ReplaceAll(text, " ", "")
	for i := 0; i < len(text); i++ {
		inputs, outputs, nodeName := processLine(text, &i)
		if endLine(text, &i) {
			connections = append(connections,
				NodeConnection{
					inputs:   inputs,
					outputs:  outputs,
					nodeName: nodeName,
				})
			break
		}
		if text[i] == '\n' {
			connections = append(connections,
				NodeConnection{
					inputs:   inputs,
					outputs:  outputs,
					nodeName: nodeName,
				})
			i++
		}
	}
	return connections
}
