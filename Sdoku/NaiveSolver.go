package Sdoku

type NaiveSolver struct {
	Solver
}

func NewNaiveSolver() *NaiveSolver {
	s := new(NaiveSolver)
	s.m_solved = nil

	return s
}

func (s *NaiveSolver) SolveSdoku(sdoku []int) {
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

func (s *NaiveSolver) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	sdokuTemp := make([]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	copy(sdokuTemp, sdoku)

	if len(emptyList) == 0 {
		s.m_solved = append(s.m_solved, sdokuTemp)
		return 1
	}
	result := 0
	availableList := make([]int, NUM_X*NUM_Y)
	emptyListTemp := make([]COORD1, len(emptyList)-1)
	numList := s.Solver.GetAvailableNumber(sdoku, emptyList[0].y, emptyList[0].x, availableList)
	if numList == 0 {
		return 0
	} else {
		result := 0
		for i := 1; i < len(emptyList); i++ {
			emptyListTemp[i-1] = emptyList[i]
		}

		for i := 0; i < numList; i++ {
			sdokuTemp[emptyList[0].x+emptyList[0].y*NUM_X*NUM_Y] = availableList[i]
			tempResult := s.SolveSdoku1(sdokuTemp, emptyListTemp)
			if tempResult > 1 {
				result = 2
				break
			}
			result += tempResult
		}
	}

	return result
}
