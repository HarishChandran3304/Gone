// Currently valid lines:
// let five = 5;
// let ten = 10;
// let add = fn(x, y) {
// 	x + y;
// };
// let result = add(five, ten);


package token

type TokenType string	

type Token struct {
	Type	TokenType
	Literal	string
}

const (
	// Special
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 1343456

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)