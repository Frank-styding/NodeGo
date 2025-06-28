package controllers

type Wire struct {
	Value       int
	Name        string
	InputNode   *Gate
	OutputsNode []*Gate
}