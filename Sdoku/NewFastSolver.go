package Sdoku

type NewFastSolver struct {
	Solver
	m_assignList [][]COORD1
}

func (s *NewFastSolver) GetAvailableNumber(y int, x int, originalEmptyList []COORD2, assignList []COORD1, numList []int) int {
	numListTemp := make([]int, NUM_X*NUM_Y)

	count := 0
	pos := 0
	for i := 0; i < len(originalEmptyList); i++ {
		if originalEmptyList[i].x == x && originalEmptyList[i].y == y {
			for j := 0; j < NUM_X*NUM_Y; j++ {
				if originalEmptyList[i].availalbeList[j] > 0 {
					numListTemp[originalEmptyList[i].availalbeList[j]-1] = 1
				}
			}
			pos = i
			break
		}
	}

	for i := 0; i < len(assignList); i++ {
		if assignList[i].x == x {
			numListTemp[assignList[i].val-1] = 0
		} else if assignList[i].y == y {
			numListTemp[assignList[i].val-1] = 0
		} else if assignList[i].group == originalEmptyList[pos].group {
			numListTemp[assignList[i].val-1] = 0
		}
	}

	for i := 0; i < NUM_X*NUM_Y; i++ {
		numList[i] = 0
	}

	for i := 0; i < NUM_X*NUM_Y; i++ {
		if numListTemp[i] > 0 {
			numList[count] = i + 1
			count++
		}
	}
	return count
}
func (s *NewFastSolver) SolveSdoku(sdoku []int) {
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

func (s *NewFastSolver) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	assignList := make([]COORD1, 0)
	originalEmptyList := make([]COORD2, 0)

	var tmp COORD2
	for i := 0; i < len(emptyList); i++ {
		tmp.x = emptyList[i].x
		tmp.y = emptyList[i].y
		tmp.group = emptyList[i].group
		tmp.val = emptyList[i].val
		tmp.availalbeList = make([]int, NUM_X*NUM_Y)
		_ = s.Solver.GetAvailableNumber(sdoku, tmp.y, tmp.x, tmp.availalbeList)
		originalEmptyList = append(originalEmptyList, tmp)

	}

	_ = s.SolveSdokuR(originalEmptyList, emptyList, assignList)

	for i := 0; i < len(s.m_assignList[0]); i++ {
		index := s.m_assignList[0][i].x + s.m_assignList[0][i].y*NUM_X*NUM_Y
		sdoku[index] = s.m_assignList[0][i].val
	}
	return 0
}

func (s *NewFastSolver) SolveSdokuR(originalEmptyList []COORD2, emptyList []COORD1, assignList []COORD1) int {

	availableList := make([]int, NUM_X*NUM_Y)
	emptyListTemp := make([]COORD1, len(emptyList))
	assignListTemp := make([]COORD1, len(assignList))
	copy(emptyListTemp, emptyList)
	copy(assignListTemp, assignList)

	pos := 0
	for pos < len(emptyListTemp) {
		numList := s.GetAvailableNumber(emptyListTemp[pos].y, emptyListTemp[pos].x, originalEmptyList, assignListTemp, availableList)
		if numList == 0 {
			return 0
		}
		if numList == 1 {
			tmp := emptyListTemp[pos]
			tmp.val = availableList[0]
			emptyListTemp = s.Remove(emptyListTemp, pos)
			assignListTemp = append(assignListTemp, tmp)
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
	numList := s.GetAvailableNumber(emptyListTemp[pos].y, emptyListTemp[pos].x, originalEmptyList, assignListTemp, availableList)
	tmp := emptyListTemp[pos]
	emptyListTemp = s.Remove(emptyListTemp, pos)
	assignListTemp = append(assignListTemp, tmp)

	for i := 0; i < numList; i++ {
		assignListTemp[len(assignListTemp)-1].val = availableList[i]
		tempResult := s.SolveSdokuR(originalEmptyList, emptyListTemp, assignListTemp)
		if tempResult > 1 {
			result = 2
			break
		}
		result += tempResult
	}

	return result
}
