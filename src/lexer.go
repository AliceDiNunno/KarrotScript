package src

type TokenType int

const (
	StringToken TokenType = iota
	NumberToken
	KeywordToken
)

func isSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

func isAlphaNumeric(char byte) bool {
	return isAlpha(char) || isDigit(char)
}

func isStringDelimiter(char byte) bool {
	return char == '"'
}
