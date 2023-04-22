package Sdoku

type FastSolver struct {
	Solver
}

func NewFastSolver() *FastSolver {
	s := new(FastSolver)
	s.m_solved = nil

	return s
}

func (s *FastSolver) SolveSdoku(sdoku []int) {
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

func (s *FastSolver) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	sdokuTemp := make([]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	copy(sdokuTemp, sdoku)

	availableList := make([]int, NUM_X*NUM_Y)
	emptyListTemp := make([]COORD1, len(emptyList))
	copy(emptyListTemp, emptyList)

	pos := 0
	for pos < len(emptyListTemp) {
		numList := s.Solver.GetAvailableNumber(sdokuTemp, emptyListTemp[pos].y, emptyListTemp[pos].x, availableList)
		if numList == 0 {
			return 0
		}
		if numList == 1 {
			sdokuTemp[emptyListTemp[pos].x+emptyListTemp[pos].y*NUM_X*NUM_Y] = availableList[0]
			//emptyListTemp = emptyListTemp[1:]
			emptyListTemp = s.Solver.Remove(emptyListTemp, pos)
			pos = 0
		} else {
			pos++
		}
	}

	if len(emptyListTemp) == 0 {
		s.m_solved = append(s.m_solved, sdokuTemp)
		return 1
	}

	result := 0

	numList := s.Solver.GetAvailableNumber(sdokuTemp, emptyListTemp[0].y, emptyListTemp[0].x, availableList)
	tmp := emptyListTemp[0]
	emptyListTemp = emptyListTemp[1:]

	for i := 0; i < numList; i++ {
		sdokuTemp[tmp.x+tmp.y*NUM_X*NUM_Y] = availableList[i]
		tempResult := s.SolveSdoku1(sdokuTemp, emptyListTemp)
		if tempResult > 1 {
			result = 2
			break
		}
		result += tempResult
	}

	return result
}
