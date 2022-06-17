package src

type KarrotScriptInstance struct {
	Parser ScriptParser
}

type ScriptFeeder interface {
	ScriptRequired(filename string) ([]byte, error)
}

type ScriptParser interface {
	Parse(fileName string) error
}

func NewKarrotScript(feeder ScriptFeeder) *KarrotScriptInstance {
	return &KarrotScriptInstance{
		Parser: NewParser(feeder),
	}
}
