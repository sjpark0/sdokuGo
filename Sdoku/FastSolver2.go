package Sdoku

type FastSolver2 struct {
	Solver
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

func (s *FastSolver2) AssignValue(sdoku []int, x int, y int, val int, availableList [][]int, emptyList []COORD1) {
	index := x + y*NUM_X*NUM_Y
	sdoku[index] = val

	for i := 0; i < len(emptyList); i++ {
		tmpList := availableList[emptyList[i].x+emptyList[i].y*NUM_X*NUM_Y]

		if emptyList[i].x == x {
			for i := 0; i < len(tmpList); i++ {
				if tmpList[i] == val {
					tmpList = append(tmpList[:i], tmpList[i+1:]...)
					break
				}
			}
		} else if emptyList[i].y == y {
			for i := 0; i < len(tmpList); i++ {
				if tmpList[i] == val {
					tmpList = append(tmpList[:i], tmpList[i+1:]...)
					break
				}
			}
		} else if emptyList[i].group == (x/NUM_X + y/NUM_Y*NUM_Y) {
			for i := 0; i < len(tmpList); i++ {
				if tmpList[i] == val {
					tmpList = append(tmpList[:i], tmpList[i+1:]...)
					break
				}
			}
		}
	}
}
func (s *FastSolver2) SolveSdoku1(sdoku []int, emptyList []COORD1) int {
	availableList := make([][]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	for i := 0; i < NUM_X*NUM_Y*NUM_X*NUM_Y; i++ {
		availableList[i] = make([]int, NUM_X*NUM_Y)
	}
	for i := 0; i < len(emptyList); i++ {
		s.GetAvailableNumber(sdoku, emptyList[i].y, emptyList[i].x, availableList[emptyList[i].x+emptyList[i].y*NUM_X*NUM_Y])
	}

	result := s.SolveSdokuR(sdoku, availableList, emptyList)
	return result
}
func (s *FastSolver2) SolveSdokuR(sdoku []int, availableList [][]int, emptyList []COORD1) int {
	availableListTemp := make([][]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
    availableListTemp2 := make([][]int, NUM_X*NUM_Y*NUM_X*NUM_Y)
	for i:=0;i<NUM_X*NUM_Y*NUM_X*NUM_Y;i++{
		availableListTemp[i] = make([]int, NUM_X * NUM_Y)
		copy(availableListTemp[i], availableList[[i]])

		availableListTemp2[i] = make([]int, NUM_X * NUM_Y)
		copy(availableListTemp2[i], availableList[[i]])

	}
    vector<COORD1> emptyListTemp;
    vector<COORD1>::iterator iter;
    
    int numList;
    int *sdokuTemp = new int[NUM_X * NUM_Y * NUM_X * NUM_Y];
    memcpy(sdokuTemp, sdoku, NUM_X * NUM_Y * NUM_X * NUM_Y * sizeof(int));
    emptyListTemp.clear();
    for(iter = emptyList->begin();iter != emptyList->end();iter++){
        emptyListTemp.push_back(*iter);
        availableListTemp[iter->x + iter->y * NUM_X * NUM_Y] = availableList[iter->x + iter->y * NUM_X * NUM_Y];
    }
    
    iter = emptyListTemp.begin();
    while(iter != emptyListTemp.end()){
        numList = (int)availableListTemp[iter->x + iter->y * NUM_X * NUM_Y].size();
        if(numList == 0){
            delete []sdokuTemp;
            delete []availableListTemp;
            delete []availableListTemp2;
            return 0;
        }
        if(numList == 1){
            COORD1 tmp = (*iter);
            emptyListTemp.erase(iter);
            AssignValue(sdokuTemp, tmp.x, tmp.y, availableListTemp[tmp.x + tmp.y * NUM_X * NUM_Y][0], availableListTemp, &emptyListTemp);
            
            iter = emptyListTemp.begin();
        }
        else{
            iter++;
        }
    }
    
    if(emptyListTemp.size() == 0){
        m_solved.push_back(sdokuTemp);
        delete []availableListTemp;
        delete []availableListTemp2;
        return 1;
    }
    
    int result = 0;
    int tempResult;
    
    
    iter = emptyListTemp.begin();
    vector<int> tmpList = availableListTemp[iter->x + iter->y * NUM_X * NUM_Y];
    numList = (int)tmpList.size();
    COORD1 tmp = (*iter);
    emptyListTemp.erase(iter);
    result = 0;
    
    for(iter = emptyListTemp.begin();iter != emptyListTemp.end();iter++){
        availableListTemp2[iter->x + iter->y * NUM_X * NUM_Y] = availableListTemp[iter->x + iter->y * NUM_X * NUM_Y];
    }
    for(int i=0;i<numList;i++){
        AssignValue(sdokuTemp, tmp.x, tmp.y, tmpList[i], availableListTemp, &emptyListTemp);
        tempResult = SolveSdokuR(sdokuTemp, availableListTemp, &emptyListTemp);
        if(tempResult > 1){
            result = 2;
            break;
        }
        result += tempResult;
        for(iter = emptyListTemp.begin();iter != emptyListTemp.end();iter++){
            if(iter->x == tmp.x){
                availableListTemp[iter->x + iter->y * NUM_X * NUM_Y] = availableListTemp2[iter->x + iter->y * NUM_X * NUM_Y];
            }
            else if(iter->y == tmp.y){
                availableListTemp[iter->x + iter->y * NUM_X * NUM_Y] = availableListTemp2[iter->x + iter->y * NUM_X * NUM_Y];
            }
            /*else if((iter->x / NUM_X == tmp.x / NUM_X) && (iter->y / NUM_Y == tmp.y / NUM_Y)){
                availableListTemp[iter->x + iter->y * NUM_X * NUM_Y] = availableListTemp2[iter->x + iter->y * NUM_X * NUM_Y];
            }*/
            else if(iter->group == (tmp.x / NUM_X + tmp.y / NUM_Y * NUM_Y)){
                availableListTemp[iter->x + iter->y * NUM_X * NUM_Y] = availableListTemp2[iter->x + iter->y * NUM_X * NUM_Y];
            }
        }
    }
    delete []availableListTemp;
    delete []availableListTemp2;
    delete []sdokuTemp;
    return result;
}
