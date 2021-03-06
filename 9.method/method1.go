package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAvgGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func main1() {
	s1 := Student{"Azizxon", []int{70, 80, 49, 55}, 19}
	s2 := Student{"Jamshid", []int{85, 65, 78, 99}, 21}

	s1.setAge(7)
	fmt.Println(s1.getAge())
	fmt.Println(s1.getAvgGrade())
	fmt.Println(s2.getAvgGrade())
	fmt.Println(s2.getMaxGrade())
}
