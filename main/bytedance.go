package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	stack []string
}

func (s *Stack) Push(c string) {
	s.stack = append(s.stack, c)
}

func (s *Stack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() string {
	var c string
	if s.IsEmpty() {
		return c
	}
	c = s.stack[len(s.stack) - 1]
	s.stack = s.stack[:len(s.stack) - 1]
	return c
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) Print() {
	for _,c := range s.stack {
		fmt.Printf("%s", c)
	}
	fmt.Print("\n")
}

func NewStack() *Stack  {
	return &Stack{stack:make([]string, 0)}
}

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
	var v1, v2 string
	for index := 0; index < len(s); index++ {
		if s[index] == '*' {
			v1 = stack.Pop()
			stack.Push(Int2Str(Str2int(v1)*Str2int(string(s[index + 1]))))
			index++
		} else if s[index] == '+' && stack.Len() == 3 {
			v1 = stack.Pop()
			stack.Pop()
			v2 = stack.Pop()
			stack.Push(Int2Str(Str2int(v1) + Str2int(v2)))
			stack.Push(string(s[index]))
		} else {
			stack.Push(string(s[index]))
		}
	}
	if stack.Len() == 3 {
		v1 = stack.Pop()
		stack.Pop()
		v2 = stack.Pop()
		return Str2int(v1) + Str2int(v2)
	} else if stack.Len() == 1 {
		return Str2int(stack.Pop())
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
	var v1,v2,v3 string
	var result int
	for index := 0;index < len(s);index++ {
		if s[index] == '*' && s[index+1] != '('{
			v1 = stack.Pop()
			stack.Push(Int2Str(Str2int(v1)*Str2int(string(s[index+1]))))
			index++
		} else if s[index] == ')' {
			v3 = stack.Pop()
			v2 = stack.Pop()
			v1 = stack.Pop()
			if v2 == "(" {
				if v1 == "*" {
					stack.Push(Int2Str(Str2int(v3)*Str2int(stack.Pop())))
				} else {
					stack.Push(v1)
					stack.Push(v3)
				}
			} else {
				stack.Pop()
				result = Str2int(v3) + Str2int(v1)
				v1 = stack.Pop()
				if v1 == "*" {
					stack.Push(Int2Str(Str2int(stack.Pop())*result))
				} else if v1 == "+" && stack.Len() >= 3 {
					v3 = stack.Pop()
					v2 = stack.Pop()
					v1 = stack.Pop()
					if v2 == "+" {
						stack.Push(Int2Str(Str2int(v3) + Str2int(v1)))
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
				stack.Push(Int2Str(Str2int(v3) + Str2int(v1)))
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
		return Str2int(v3) + Str2int(v1)
	}
	return Str2int(stack.Pop())
}

func main() {
	//s := "1+3*4+5*8"
	s := "1+((3+4*2)+2)*5+8"
	fmt.Println(calculator2(s))
}