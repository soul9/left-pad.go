package leftpad

import (
	"fmt"
	"strings"
)

func LeftPad(str string, num int, ch string) string {
	str = fmt.Sprint(str)

	if ch == "" {
		ch = " "
	}

	num = num - len(str)

	str = fmt.Sprintf("%s%s", strings.Repeat(ch, num), str)

	return str
}
