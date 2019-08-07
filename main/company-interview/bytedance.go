package main

import (
	"fmt"
	"strconv"
	. "LeetCode/main/container"
)

func Str2int(s string) int {
	ret,_ := strconv.Atoi(s)
	return ret
}

func Int2Str(i int) string {
	return strconv.Itoa(i)
}

/*
1+3*4+5*8
 */
func calculator(s string) int {
	if s == "" {
		return 0
	}
	stack := NewStack()
	var v1, v2 interface{}
	for index := 0; index < len(s); index++ {
		if s[index] == '*' {
			v1 = stack.Pop()
			stack.Push(Int2Str(Str2int(v1.(string))*Str2int(string(s[index + 1]))))
			index++
		} else if s[index] == '+' && stack.Len() == 3 {
			v1 = stack.Pop()
			stack.Pop()
			v2 = stack.Pop()
			stack.Push(Int2Str(Str2int(v1.(string)) + Str2int(v2.(string))))
			stack.Push(string(s[index]))
		} else {
			stack.Push(string(s[index]))
		}
	}
	if stack.Len() == 3 {
		v1 = stack.Pop()
		stack.Pop()
		v2 = stack.Pop()
		return Str2int(v1.(string)) + Str2int(v2.(string))
	} else if stack.Len() == 1 {
		return Str2int(stack.Pop().(string))
	}
	return 0
}

/*
1+(3+5)*4+5*8
 */
func calculator2(s string) int {
	if s == "" {
		return 0
	}
	stack := NewStack()
	var v1,v2,v3 interface{}
	var result int
	for index := 0;index < len(s);index++ {
		if s[index] == '*' && s[index+1] != '('{
			v1 = stack.Pop()
			stack.Push(Int2Str(Str2int(v1.(string))*Str2int(string(s[index+1]))))
			index++
		} else if s[index] == ')' {
			v3 = stack.Pop()
			v2 = stack.Pop()
			v1 = stack.Pop()
			if v2 == "(" {
				if v1 == "*" {
					stack.Push(Int2Str(Str2int(v3.(string))*Str2int(stack.Pop().(string))))
				} else {
					stack.Push(v1)
					stack.Push(v3)
				}
			} else {
				stack.Pop()
				result = Str2int(v3.(string)) + Str2int(v1.(string))
				v1 = stack.Pop()
				if v1 == "*" {
					stack.Push(Int2Str(Str2int(stack.Pop().(string))*result))
				} else if v1 == "+" && stack.Len() >= 3 {
					v3 = stack.Pop()
					v2 = stack.Pop()
					v1 = stack.Pop()
					if v2 == "+" {
						stack.Push(Int2Str(Str2int(v3.(string)) + Str2int(v1.(string))))
					} else {
						stack.Push(v1)
						stack.Push(v2)
						stack.Push(v3)
					}
					stack.Push("+")
					stack.Push(Int2Str(result))
				} else {
					stack.Push(v1)
					stack.Push(Int2Str(result))
				}
			}
		} else if s[index] == '+' && stack.Len() >= 3 {
			v3 = stack.Pop()
			v2 = stack.Pop()
			v1 = stack.Pop()
			if v2 == "+" {
				stack.Push(Int2Str(Str2int(v3.(string)) + Str2int(v1.(string))))
			} else {
				stack.Push(v1)
				stack.Push(v2)
				stack.Push(v3)
			}
			stack.Push(string(s[index]))
		} else {
			stack.Push(string(s[index]))
		}
	}
	if stack.Len() == 3 {
		v3 = stack.Pop()
		stack.Pop()
		v1 = stack.Pop()
		return Str2int(v3.(string)) + Str2int(v1.(string))
	}
	return Str2int(stack.Pop().(string))
}

func main() {
	//s := "1+3*4+5*8"
	s := "1+((3+4*2)+2)*5+8"
	fmt.Println(calculator2(s))
}