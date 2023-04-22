package Sdoku

import (
	"fmt"
	"math/rand"
)

const NUM_X int = 3
const NUM_Y int = 3

type COORD1 struct {
	x     int
	y     int
	group int
	val   int
}

type COORD2 struct {
	x             int
	y             int
	group         int
	val           int
	availalbeList []int
}

type Solver struct {
	m_solved [][]int
}

func NewSolver() *Solver {
	s := new(Solver)
	s.m_solved = nil
	return s
}
func (s *Solver) Remove(slice []COORD1, i int) []COORD1 {
	return append(slice[:i], slice[i+1:]...)
}

func (s *Solver) MakeSdoku(sdoku []int) {
	bSuccess := false
	totalNum := 0
	availableList := make([]int, NUM_X*NUM_Y)

	for bSuccess == false {
		totalNum++
		bSuccess = true
		for i := 0; i < NUM_X*NUM_Y; i++ {
			for j := 0; j < NUM_X*NUM_Y; j++ {
				sdoku[j+i*NUM_X*NUM_Y] = 0
			}
		}
		for i := 0; i < NUM_X*NUM_Y; i++ {
			if bSuccess == true {
				for j := 0; j < NUM_X*NUM_Y; j++ {
					numList := s.GetAvailableNumber(sdoku, i, j, availableList)
					if numList == 0 {
						bSuccess = false
						break
					}
					sdoku[j+i*NUM_X*NUM_Y] = availableList[rand.Intn(numList)]

				}
			}
		}
	}
}
func (s *Solver) GetAvailableNumber(sdoku []int, i int, j int, numList []int) int {
	var index1 int
	var index2 int
	count := 0
	var isAvail bool

	for m := 0; m < NUM_X*NUM_Y; m++ {
		numList[m] = 0
	}
	for aa := 1; aa < 1+NUM_X*NUM_Y; aa++ {
		isAvail = true
		for m := 0; m < NUM_X*NUM_Y; m++ {
			if sdoku[m+i*NUM_X*NUM_Y] == aa {
				isAvail = false
				break
			}
			if sdoku[j+m*NUM_X*NUM_Y] == aa {
				isAvail = false
				break
			}
		}

		if isAvail == true {
			index1 = (i / NUM_Y) * NUM_Y
			index2 = (j / NUM_X) * NUM_X
			for m := index1; m < index1+NUM_Y; m++ {
				if isAvail == true {
					for n := index2; n < index2+NUM_X; n++ {
						if sdoku[n+m*NUM_X*NUM_Y] == aa {
							isAvail = false
							break
						}
					}
				}
			}
		}
		if isAvail == true {
			numList[count] = aa
			count++
		}
	}
	return count
}

func (s *Solver) PrintSdoku(sdoku []int) {
	for i := 0; i < NUM_X*NUM_Y; i++ {
		for j := 0; j < NUM_X*NUM_Y; j++ {
			fmt.Printf("%d\t", sdoku[j+i*NUM_X*NUM_Y])
		}
		fmt.Println()
	}
}
func (s *Solver) Test() {
	fmt.Println("Test")
}
