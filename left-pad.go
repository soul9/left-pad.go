package leftpad

import "fmt"

func LeftPad(str string, num int, ch string) string {
	str = fmt.Sprint(str)

	if ch == "" {
		ch = " "
	}

	num = num - len(str)

	for i := 0; i < num; i++ {
		str = fmt.Sprintf("%s%s", ch, str)
	}

	return str
}
