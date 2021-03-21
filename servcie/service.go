package service

var strs = make([]string, 0)

func Register(str string) {
	strs = append(strs, str)
}

func GetStrs() []string {
	return strs
}
