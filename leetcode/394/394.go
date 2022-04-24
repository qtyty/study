package _94

import "fmt"

func decodeString(s string) string {
	numstack, strstack := make([]int, 0), make([]string, 0)
	num, str := 0, ""
	for _, c := range s {
		if '0' <= c && c <= '9' {
			num = num*10 + int(c-'0')
		} else if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
			str += string(c)
		} else if c == '[' {
			numstack = append(numstack, num)
			strstack = append(strstack, str)
			num = 0
			str = ""
		} else {
			repeat := numstack[len(numstack)-1]
			item := strstack[len(strstack)-1]

			numstack = numstack[:len(numstack)-1]
			strstack = strstack[:len(strstack)-1]

			for ; repeat > 0; repeat-- {
				item += str
			}
			str = item
		}
	}
	return str
}

func main() {
	str := "3[a2[c]]"
	s := decodeString(str)
	fmt.Println(s)
}
