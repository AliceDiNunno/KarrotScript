package src

type Parser struct {
	feeder ScriptFeeder

	content        []byte
	positionInFile int
}

type Token struct {
	Type        TokenType
	Content     []byte
	Line        int
	LineContent []byte
	Column      int
}

func (parser *Parser) Parse(fileName string) error {
	file, err := parser.feeder.ScriptRequired(fileName)

	if err != nil {
		return err
	}

	parser.content = file

	for parser.canRead() {
		if parser.positionInFile >= len(parser.content) {
			break
		}

		parser.ParseToken()

		parser.skipSpaces()
	}

	return nil
}

func (parser *Parser) parseIdentifier() {
	for {
		char := parser.readChar()

		if !isAlphaNumeric(char) {
			parser.positionInFile--
			break
		}

		print(string(char))
	}
}

func (parser *Parser) ParseToken() {
	char := parser.readChar()

	if isAlpha(char) {
		parser.parseIdentifier()
	} else if isStringDelimiter(char) {
		print("String")
	} else {
		print("Token")
	}

	print(string(char))
}

func (parser *Parser) readChar() byte {
	if parser.positionInFile >= len(parser.content) {
		return 0
	}

	char := parser.content[parser.positionInFile]

	parser.positionInFile++

	return char
}

func (parser *Parser) skipSpaces() {
	for {
		char := parser.readChar()

		if !isSpace(char) && char != 0 {
			parser.positionInFile--
			break
		}
	}
}

func (parser *Parser) canRead() bool {
	return parser.positionInFile < len(parser.content)
}

func NewParser(feeder ScriptFeeder) *Parser {
	return &Parser{
		feeder:         feeder,
		positionInFile: 0,
	}
}
