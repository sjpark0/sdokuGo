package main

import (
	"Sdoku/Sdoku"
	"fmt"
	"time"
)

func main() {
	sdoku := make([]int, Sdoku.NUM_X*Sdoku.NUM_Y*Sdoku.NUM_X*Sdoku.NUM_Y)
	sdokuOriginal := make([]int, Sdoku.NUM_X*Sdoku.NUM_Y*Sdoku.NUM_X*Sdoku.NUM_Y)
	sdoku[1] = 9
	sdoku[5] = 6
	sdoku[8] = 5
	sdoku[10] = 3
	sdoku[12] = 4
	sdoku[13] = 5
	sdoku[16] = 8
	sdoku[18] = 4
	sdoku[23] = 2
	sdoku[32] = 4
	sdoku[36] = 3
	sdoku[39] = 7
	sdoku[40] = 9
	sdoku[44] = 2
	sdoku[46] = 8
	sdoku[51] = 1
	sdoku[54] = 7
	sdoku[57] = 5
	sdoku[58] = 3
	sdoku[62] = 9
	sdoku[67] = 6
	sdoku[74] = 9
	sdoku[79] = 2

	copy(sdokuOriginal, sdoku)

	var start time.Time
	var end time.Duration
	var naive Sdoku.NaiveSolver
	var naive_0 Sdoku.NaiveSolver_0
	var fast Sdoku.FastSolver
	var fast_0 Sdoku.FastSolver_0
	var fast1 Sdoku.FastSolver1
	var fast2 Sdoku.FastSolver2
	var fast3 Sdoku.FastSolver3
	var newFast Sdoku.NewFastSolver
	var newFast1 Sdoku.NewFastSolver1
	var newFast2 Sdoku.NewFastSolver2
	var newFast3 Sdoku.NewFastSolver3

	copy(sdoku, sdokuOriginal)
	//game = Sdoku.NewNaiveSolver()
	naive.PrintSdoku(sdoku)
	start = time.Now()
	naive.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("NaiveSolver => Time measured : ", end)
	naive.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	naive_0.PrintSdoku(sdoku)
	start = time.Now()
	naive_0.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("NaiveSolver_0 => Time measured : ", end)
	naive_0.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	fast.PrintSdoku(sdoku)
	start = time.Now()
	fast.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("FastSolver => Time measured : ", end)
	fast.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	fast_0.PrintSdoku(sdoku)
	start = time.Now()
	fast_0.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("FastSolver_0 => Time measured : ", end)
	fast_0.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	fast1.PrintSdoku(sdoku)
	start = time.Now()
	fast1.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("FastSolver1 => Time measured : ", end)
	fast1.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	fast2.PrintSdoku(sdoku)
	start = time.Now()
	fast2.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("FastSolver2 => Time measured : ", end)
	fast2.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	fast3.PrintSdoku(sdoku)
	start = time.Now()
	fast3.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("FastSolver3 => Time measured : ", end)
	fast3.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	newFast.PrintSdoku(sdoku)
	start = time.Now()
	newFast.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("NewFastSolver => Time measured : ", end)
	newFast.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	newFast1.PrintSdoku(sdoku)
	start = time.Now()
	newFast1.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("NewFastSolver1 => Time measured : ", end)
	newFast1.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	newFast2.PrintSdoku(sdoku)
	start = time.Now()
	newFast2.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("NewFastSolver2 => Time measured : ", end)
	newFast2.PrintSdoku(sdoku)
	fmt.Println()

	copy(sdoku, sdokuOriginal)
	//game1 := Sdoku.NewFastSolver()
	newFast3.PrintSdoku(sdoku)
	start = time.Now()
	newFast3.SolveSdoku(sdoku)
	end = time.Since(start)
	fmt.Println("NewFastSolver3 => Time measured : ", end)
	newFast3.PrintSdoku(sdoku)
	fmt.Println()

}
