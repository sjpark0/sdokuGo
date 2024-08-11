package Sdoku

type FastSolver_0 struct {
	Solver
}

func (s *FastSolver_0) SolveSdoku(sdoku []int) {
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

	if len(s.m_solved) != 0 {
		copy(sdoku, s.m_solved[0])
	}
}

func (s *FastSolver_0) SolveSdoku1(sdoku []int, emptyList []COORD1) int {

	availableList := make([]int, NUM_X*NUM_Y)
	var assignList []COORD1 = nil
	pos := 0
	for pos < len(emptyList) {
		numList := s.Solver.GetAvailableNumber(sdoku, emptyList[pos].y, emptyList[pos].x, availableList)
		if numList == 0 {
			for i := 0; i < len(assignList); i++ {
				emptyList = append(emptyList, assignList[i])
				sdoku[assignList[i].x+assignList[i].y*NUM_X*NUM_Y] = 0
			}
			return 0
		}
		if numList == 1 {
			sdoku[emptyList[pos].x+emptyList[pos].y*NUM_X*NUM_Y] = availableList[0]
			assignList = append(assignList, emptyList[pos])
			emptyList = s.Solver.Remove(emptyList, pos)
			pos = 0
		} else {
			pos++
		}
	}

	if len(emptyList) == 0 {
		sdokuTemp := make([]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
		copy(sdokuTemp, sdoku)

		s.m_solved = append(s.m_solved, sdokuTemp)

		for i := 0; i < len(assignList); i++ {
			emptyList = append(emptyList, assignList[i])
			sdoku[assignList[i].x+assignList[i].y*NUM_X*NUM_Y] = 0
		}
		return 1
	}

	result := 0

	numList := s.Solver.GetAvailableNumber(sdoku, emptyList[0].y, emptyList[0].x, availableList)
	tmp := emptyList[0]
	emptyList = emptyList[1:]
	assignList = append(assignList, tmp)
	for i := 0; i < numList; i++ {
		sdoku[tmp.x+tmp.y*NUM_X*NUM_Y] = availableList[i]
		tempResult := s.SolveSdoku1(sdoku, emptyList)
		if tempResult > 1 {
			result = 2
			break
		}
		result += tempResult
	}
	for i := 0; i < len(assignList); i++ {
		emptyList = append(emptyList, assignList[i])
		sdoku[assignList[i].x+assignList[i].y*NUM_X*NUM_Y] = 0
	}
	return result
}
