package iteration

func Repeat(char string, itr int) string {
	var temp string
	for i := 0; i < itr; i++ {
		temp += char
	}
	return temp
}
