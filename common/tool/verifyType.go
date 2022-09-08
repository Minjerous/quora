package tool

var AnswerType = []string{"answer", "idea"}

func VerifyType(Type string) bool {
	for _, k := range AnswerType {
		if k == Type {
			return true
		}
	}
	return false
}
