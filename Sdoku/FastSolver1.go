package Sdoku

type FastSolver1 struct {
	Solver
}

func (s *FastSolver1) GetAvailableNumber(sdoku []int, i int, j int, numList []int) int {
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
			numList[aa-1] = 1
			count++
		}
	}
	return count
}

func (s *FastSolver1) SolveSdoku(sdoku []int) {
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

func (s *FastSolver1) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	sdokuTemp := make([]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	copy(sdokuTemp, sdoku)

	availableList := make([]int, NUM_X*NUM_Y)
	emptyListTemp := make([]COORD1, len(emptyList))
	copy(emptyListTemp, emptyList)

	pos := 0
	var idx int = 0

	for pos < len(emptyListTemp) {
		numList := s.GetAvailableNumber(sdokuTemp, emptyListTemp[pos].y, emptyListTemp[pos].x, availableList)
		if numList == 0 {
			return 0
		}
		if numList == 1 {
			for i := 0; i < NUM_X*NUM_Y; i++ {
				if availableList[i] == 1 {
					idx = i
					break
				}
			}
			sdokuTemp[emptyListTemp[pos].x+emptyListTemp[pos].y*NUM_X*NUM_Y] = idx + 1
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

	_ = s.GetAvailableNumber(sdokuTemp, emptyListTemp[0].y, emptyListTemp[0].x, availableList)
	tmp := emptyListTemp[0]
	emptyListTemp = emptyListTemp[1:]

	for i := 0; i < NUM_X*NUM_Y; i++ {
		if availableList[i] == 1 {
			sdokuTemp[tmp.x+tmp.y*NUM_X*NUM_Y] = i + 1
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
