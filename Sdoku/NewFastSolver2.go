package Sdoku

type NewFastSolver2 struct {
	Solver
	m_assignList [][]COORD2
}

func (s *NewFastSolver2) Remove(slice []COORD2, i int) []COORD2 {
	return append(slice[:i], slice[i+1:]...)
}

func (s *NewFastSolver2) GetAvailableNumber(sdoku []int, i int, j int) []int {
	var index1 int
	var index2 int
	var isAvail bool
	var numList []int = nil

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
			numList = append(numList, aa)
		}
	}
	return numList
}
func (s *NewFastSolver2) SolveSdoku(sdoku []int) {
	var emptyList []COORD2 = nil
	s.m_solved = nil
	var tmpCoord COORD2

	for i := 0; i < NUM_X*NUM_Y; i++ {
		for j := 0; j < NUM_X*NUM_Y; j++ {
			if sdoku[j+i*NUM_X*NUM_Y] == 0 {
				tmpCoord.x = j
				tmpCoord.y = i
				tmpCoord.group = (j / NUM_X) + (i/NUM_Y)*NUM_Y
				tmpCoord.val = 0
				tmpCoord.availalbeList = s.GetAvailableNumber(sdoku, i, j)
				emptyList = append(emptyList, tmpCoord)
			}
		}
	}

	s.SolveSdoku1(sdoku, emptyList)
}

func (s *NewFastSolver2) SolveSdoku1(sdoku []int, emptyList []COORD2) int {
	assignList := make([]COORD2, 0)

	_ = s.SolveSdokuR(assignList, emptyList)
	for i := 0; i < len(s.m_assignList[0]); i++ {
		index := s.m_assignList[0][i].x + s.m_assignList[0][i].y*NUM_X*NUM_Y
		sdoku[index] = s.m_assignList[0][i].val
	}
	return 0
}

func (s *NewFastSolver2) SolveSdokuR(assignList []COORD2, emptyList []COORD2) int {

	emptyListTemp := make([]COORD2, len(emptyList))
	assignListTemp := make([]COORD2, len(assignList))

	copy(emptyListTemp, emptyList)
	copy(assignListTemp, assignList)

	pos := 0

	for pos < len(emptyListTemp) {
		numList := len(emptyListTemp[pos].availalbeList)
		if numList == 0 {
			return 0
		}
		if numList == 1 {
			tmp := emptyListTemp[pos]
			assignListTemp = append(assignListTemp, tmp)
			emptyListTemp = s.Remove(emptyListTemp, pos)
			s.AssignValue(assignListTemp, tmp.x, tmp.y, tmp.availalbeList[0], emptyListTemp)
			pos = 0
		} else {
			pos++
		}
	}
	if len(emptyListTemp) == 0 {
		s.m_assignList = append(s.m_assignList, assignListTemp)
		return 1
	}

	result := 0
	pos = 0
	numList := len(emptyListTemp[pos].availalbeList)
	tmp := emptyListTemp[pos]

	assignListTemp = append(assignListTemp, tmp)
	emptyListTemp = s.Remove(emptyListTemp, pos)

	emptyListTemp2 := make([]COORD2, len(emptyListTemp))
	copy(emptyListTemp2, emptyListTemp)

	for i := 0; i < numList; i++ {
		s.AssignValue(assignListTemp, tmp.x, tmp.y, tmp.availalbeList[i], emptyListTemp)

		tempResult := s.SolveSdokuR(assignListTemp, emptyListTemp)
		if tempResult > 1 {
			result = 2
			break
		}
		result += tempResult
		for m := 0; m < len(emptyListTemp); m++ {
			if emptyListTemp[m].x == tmp.x {
				emptyListTemp[m].availalbeList = make([]int, len(emptyListTemp2[m].availalbeList))
				copy(emptyListTemp[m].availalbeList, emptyListTemp2[m].availalbeList)
			} else if emptyListTemp[m].y == tmp.y {
				emptyListTemp[m].availalbeList = make([]int, len(emptyListTemp2[m].availalbeList))
				copy(emptyListTemp[m].availalbeList, emptyListTemp2[m].availalbeList)
			} else if emptyListTemp[m].group == (tmp.x/NUM_X + tmp.y/NUM_Y*NUM_Y) {
				emptyListTemp[m].availalbeList = make([]int, len(emptyListTemp2[m].availalbeList))
				copy(emptyListTemp[m].availalbeList, emptyListTemp2[m].availalbeList)
			}
		}
	}

	return result
}
func (s *NewFastSolver2) AssignValue(assignList []COORD2, x int, y int, val int, emptyList []COORD2) {
	assignList[len(assignList)-1].val = val
	for i := 0; i < len(emptyList); i++ {
		if emptyList[i].x == x {
			for m := 0; m < len(emptyList[i].availalbeList); m++ {
				if emptyList[i].availalbeList[m] == val {
					emptyList[i].availalbeList = append(emptyList[i].availalbeList[:m], emptyList[i].availalbeList[m+1:]...)
					break
				}
			}
		} else if emptyList[i].y == y {
			for m := 0; m < len(emptyList[i].availalbeList); m++ {
				if emptyList[i].availalbeList[m] == val {
					emptyList[i].availalbeList = append(emptyList[i].availalbeList[:m], emptyList[i].availalbeList[m+1:]...)
					break
				}
			}
		} else if emptyList[i].group == (x/NUM_X + y/NUM_Y*NUM_Y) {
			for m := 0; m < len(emptyList[i].availalbeList); m++ {
				if emptyList[i].availalbeList[m] == val {
					emptyList[i].availalbeList = append(emptyList[i].availalbeList[:m], emptyList[i].availalbeList[m+1:]...)
					break
				}
			}
		}
	}

}
