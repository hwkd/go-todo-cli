package parser

type ASTNode struct {
}

type OptionNode struct {
	Option string
	Value  *ValueNode
}

type ValueNode struct {
	Literal string
}
