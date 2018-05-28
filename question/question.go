package question

type Question struct {
	Question string `yaml:question`
	Answer   string `yaml:answer`
	Content  string `yaml:content`
}
