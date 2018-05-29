package question

type Question struct {
	Question string `yaml:question`
	Answer   string `yaml:answer`
	Filename string `yaml:filename`
	Content  string
}
