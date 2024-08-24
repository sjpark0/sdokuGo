package Sdoku

type NewFastSolver1 struct {
	Solver
	m_assignList [][]COORD1
}

func (s *NewFastSolver1) GetAvailableNumber(sdoku []int, i int, j int) []int {
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
func (s *NewFastSolver1) SolveSdoku(sdoku []int) {
	var emptyList []COORD1 = nil
	s.m_solved = nil
	var tmpCoord COORD1

	for i := 0; i < NUM_X*NUM_Y; i++ {
		for j := 0; j < NUM_X*NUM_Y; j++ {
			if sdoku[j+i*NUM_X*NUM_Y] == 0 {
				tmpCoord.x = j
				tmpCoord.y = i
				tmpCoord.group = (j / NUM_X) + (i/NUM_Y)*NUM_Y
				tmpCoord.val = 0
				emptyList = append(emptyList, tmpCoord)
			}
		}
	}

	s.SolveSdoku1(sdoku, emptyList)
}

func (s *NewFastSolver1) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	assignList := make([]COORD1, 0)
	availableList := make([][]int, NUM_X*NUM_Y*NUM_X*NUM_Y)

	for i := 0; i < len(emptyList); i++ {
		availableList[emptyList[i].x+emptyList[i].y*NUM_X*NUM_Y] = s.GetAvailableNumber(sdoku, emptyList[i].y, emptyList[i].x)
	}

	_ = s.SolveSdokuR(assignList, availableList, emptyList)
	for i := 0; i < len(s.m_assignList[0]); i++ {
		index := s.m_assignList[0][i].x + s.m_assignList[0][i].y*NUM_X*NUM_Y
		sdoku[index] = s.m_assignList[0][i].val
	}
	return 0
}

func (s *NewFastSolver1) SolveSdokuR(assignList []COORD1, availableList [][]int, emptyList []COORD1) int {

	availableListTemp := make([][]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	availableListTemp2 := make([][]int, NUM_X*NUM_Y*NUM_X*NUM_Y)

	emptyListTemp := make([]COORD1, len(emptyList))
	assignListTemp := make([]COORD1, len(assignList))
	copy(emptyListTemp, emptyList)
	copy(assignListTemp, assignList)

	for i := 0; i < len(emptyList); i++ {
		index2 := emptyList[i].x + emptyList[i].y*NUM_X*NUM_Y
		availableListTemp[index2] = make([]int, len(availableList[index2]))
		copy(availableListTemp[index2], availableList[index2])
	}

	pos := 0
	for pos < len(emptyListTemp) {
		index := emptyListTemp[pos].x + emptyListTemp[pos].y*NUM_X*NUM_Y
		numList := len(availableListTemp[index])

		if numList == 0 {
			return 0
		}
		if numList == 1 {
			tmp := emptyListTemp[pos]
			assignListTemp = append(assignListTemp, tmp)
			emptyListTemp = s.Remove(emptyListTemp, pos)
			s.AssignValue(assignListTemp, tmp.x, tmp.y, availableListTemp[tmp.x+tmp.y*NUM_X*NUM_Y][0], availableListTemp, emptyListTemp)
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
	index := emptyListTemp[pos].x + emptyListTemp[pos].y*NUM_X*NUM_Y
	tmpList := availableListTemp[index]
	numList := len(tmpList)
	tmp := emptyListTemp[pos]

	assignListTemp = append(assignListTemp, tmp)
	emptyListTemp = s.Remove(emptyListTemp, pos)

	for i := 0; i < len(emptyListTemp); i++ {
		index2 := emptyListTemp[i].x + emptyListTemp[i].y*NUM_X*NUM_Y
		availableListTemp2[index2] = make([]int, len(availableListTemp[index2]))
		copy(availableListTemp2[index2], availableListTemp[index2])
	}

	for i := 0; i < numList; i++ {
		s.AssignValue(assignListTemp, tmp.x, tmp.y, tmpList[i], availableListTemp, emptyListTemp)

		tempResult := s.SolveSdokuR(assignListTemp, availableListTemp, emptyListTemp)
		if tempResult > 1 {
			result = 2
			break
		}
		result += tempResult
		for m := 0; m < len(emptyListTemp); m++ {
			index = emptyListTemp[m].x + emptyListTemp[m].y*NUM_X*NUM_Y
			if emptyListTemp[m].x == tmp.x {
				availableListTemp[index] = make([]int, len(availableListTemp2[index]))
				copy(availableListTemp[index], availableListTemp2[index])
			} else if emptyListTemp[m].y == tmp.y {
				availableListTemp[index] = make([]int, len(availableListTemp2[index]))
				copy(availableListTemp[index], availableListTemp2[index])
			} else if emptyListTemp[m].group == (tmp.x/NUM_X + tmp.y/NUM_Y*NUM_Y) {
				availableListTemp[index] = make([]int, len(availableListTemp2[index]))
				copy(availableListTemp[index], availableListTemp2[index])
			}
		}

	}

	return result
}
func (s *NewFastSolver1) AssignValue(assignList []COORD1, x int, y int, val int, availableList [][]int, emptyList []COORD1) []int {
	var tmp COORD1
	tmp.x = x
	tmp.y = y
	tmp.group = (x/NUM_X + y/NUM_Y*NUM_Y)
	tmp.val = val
	res := make([]int, 0)
	//assignList = append(assignList, tmp)
	assignList[len(assignList)-1].val = val
	for i := 0; i < len(emptyList); i++ {
		index1 := emptyList[i].x + emptyList[i].y*NUM_X*NUM_Y
		tmpList := availableList[index1]
		pre_len := len(availableList[index1])

		if emptyList[i].x == x {
			for m := 0; m < len(tmpList); m++ {
				if tmpList[m] == val {
					availableList[index1] = append(tmpList[:m], tmpList[m+1:]...)
					break
				}
			}
		} else if emptyList[i].y == y {
			for m := 0; m < len(tmpList); m++ {
				if tmpList[m] == val {
					availableList[index1] = append(tmpList[:m], tmpList[m+1:]...)
					break
				}
			}
		} else if emptyList[i].group == (x/NUM_X + y/NUM_Y*NUM_Y) {
			for m := 0; m < len(tmpList); m++ {
				if tmpList[m] == val {
					availableList[index1] = append(tmpList[:m], tmpList[m+1:]...)
					break
				}
			}
		}
		post_len := len(availableList[index1])
		if pre_len != post_len {
			res = append(res, index1)
		}
	}
	return res
}
