package Sdoku

type FastSolver2 struct {
	Solver
}

func (s *FastSolver2) GetAvailableNumber(sdoku []int, i int, j int) []int {
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
func (s *FastSolver2) SolveSdoku(sdoku []int) {
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

func (s *FastSolver2) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	sdokuTemp := make([]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	copy(sdokuTemp, sdoku)

	//availableList := make([]int, NUM_X*NUM_Y)
	var availableList []int = nil
	emptyListTemp := make([]COORD1, len(emptyList))
	copy(emptyListTemp, emptyList)

	pos := 0
	for pos < len(emptyListTemp) {
		availableList = s.GetAvailableNumber(sdokuTemp, emptyListTemp[pos].y, emptyListTemp[pos].x)
		numList := len(availableList)
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

	availableList = s.GetAvailableNumber(sdokuTemp, emptyListTemp[0].y, emptyListTemp[0].x)
	numList := len(availableList)
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
