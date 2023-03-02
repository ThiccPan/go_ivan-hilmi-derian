package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type Student struct {
	name   []string
	score  []uint
	minIdx uint
	maxIdx uint
	ave    uint
}

func (s *Student) addStudent(name string, score uint) error {
	if score > 100 || score < 0 {
		fmt.Println("score input is not valid, try again")
		return errors.New("invalid score input")
	}
	s.name = append(s.name, name)
	s.score = append(s.score, score)

	if score > s.score[s.maxIdx] {
		s.maxIdx = uint(len(s.score) - 1)
	}

	if score < s.score[s.minIdx] {
		s.minIdx = uint(len(s.score) - 1)
	}

	return nil
}

func (s *Student) CalcAve() {
	for _, v := range s.score {
		s.ave += v
		fmt.Println(v)
	}
	s.ave /= uint(len(s.score))
}

func (s *Student) PrintMax() {
	fmt.Printf("Max score of Students: %v (%v)\n", s.name[s.maxIdx], s.score[s.maxIdx])
}

func (s *Student) PrintMin() {
	fmt.Printf("Min score of Students: %v (%v)\n", s.name[s.minIdx], s.score[s.minIdx])
}

func (s *Student) PrintAve() {
	fmt.Println("Average score:", s.ave)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sList := Student{}
	i := 1
	for {
		fmt.Print("input ", i, " student's name: ")
		sc.Scan()

		sName := sc.Text()
		if sName == "" {
			break
		}

		fmt.Print("input ", i, " student's score: ")
		sc.Scan()

		inScore := sc.Text()
		var sScore uint8

		if _, err := fmt.Sscan(inScore, &sScore); err != nil {
			fmt.Println(err)
			continue
		}

		err := sList.addStudent(sName, uint(sScore))
		if err != nil {
			fmt.Println(err)
			continue
		}
		i++
	}
	sList.CalcAve()
	sList.PrintAve()
	sList.PrintMin()
	sList.PrintMax()
}
