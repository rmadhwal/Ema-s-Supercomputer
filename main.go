package main

/**
Ema's supercomputer
https://www.hackerrank.com/challenges/two-pluses
Essentially, this solution revolves around parsing the entire grid into an internal structure and then searching for all "plus" structures inside the grid
After all pluses are found, they are sorted according to their size
Then, all the pluses are compared with each other for overlaps, a 2D array stores overlap information
Again, all the plus structures are copmpared against each other, if two plus do not overlap, their element size is multiplied with each other
and if their product is greater than the maximum product (initilized as 0), the maximum product is updated to their product.
Finally the maximum product i.e. the biggest possible product of two non-overlapping pluses in the grid is calculated.
Rohan Madhwal
**/
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//init strucutres
//basic element structure, holds the x and y coordinate of any element, enough tom model a coordinate
//The 'G' or 'B' value is redundant to store because if the element is being stored, it's already assumed it's a good value
type elem struct {
	x int
	y int
}

//Attempted model a of "plus" structure
//The numElements is essentially redudent and can be subsituted by a call to len(exampleplus.elems), but I find specifying a seperate int more elegant,
//It also makes debugging much easier hehe
type plus struct {
	numElements int
	elems       []elem
}

//Defined type of a slice of pluses, required since to sort a slice of pluses, go requires you to name them
//Also note definition of three pluses methods later which act as helper functions to sorting the interface of pluses
type pluses []plus

func main() {
	//length and width integers at top to store the first 2 values
	var length, width int
	//Read inputs
	scanner := bufio.NewScanner(os.Stdin)
	//Single scan since we only want to load in one line
	scanner.Scan()
	firstLine := scanner.Text()
	//Use of split function to break the first line into a slice of strings, using an empty space as delimiter to seperate elements
	ints := strings.Split(firstLine, " ")
	//Custom function since typing out the same function twice is redundant, refer to the function for details of functioning
	convertToIntAndAssignTo(ints[0], &length)
	convertToIntAndAssignTo(ints[1], &width)

	//Load in provided grid into a 2D boolean array
	//Making it a boolean array instead of character array improves computational speed
	//Usually anywhere that only 2 choices are plausible, boolean can be used, it's essentially 1 and 0, equating it to necessarily mean "true" and "false".
	//Impairs you of a lot of functionality, and hey! A single bit
	grid := make([][]bool, length)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, width)
	}
	for x := 0; x < length; x++ {
		scanner.Scan()
		line := scanner.Text()
		//Using go's rune functionality to automatically convert each individual character in the string to it's own elemt in a slice
		a := []rune(line)
		for y := 0; y < width; y++ {

			if a[y] == 'G' {
				grid[x][y] = true
			} else {
				grid[x][y] = false
			}
		}
	}
	//Custom function to well, the name pretty much says it all lol
	sliceOfPluses := checkGridForPlusesAndAddIntoSliceOfPluses(grid, length, width)
	//Sort function can be implemented because of helper functions that specify customisation of sort for this interface
	sort.Sort(sliceOfPluses)
	//Custom function to detect overlaps and load them into a 2D slice
	//The name is misleading, originally intended the function to remove overlaps but that was tedious and well not working
	//Instead, now rely on simply detecting overlaps and storing them.
	//However, not going to change the name yet since directly removing redundant and non-valulable slices to return the optimal
	//pluses as index 0 and 1 would be the more efficient solution instead of a 2 step process of first calculating the
	//overlaps and then calculating the optimal pair, however, this works for now, so I will leave this here until a better idea strikes me.
	sliceOfOverlaps := removeOverlaps(sliceOfPluses)

	var maxProd int
	//Calculating optimal pair using incremental sorting process (should probably abstract the sort into a function)
	//Golang allows functions to be passed as parameters however, at this point that seems more like a macro than a useful function
	//Might work on it later
	for i := 0; i < len(sliceOfPluses); i++ {
		for j := i + 1; j < len(sliceOfPluses); j++ {

			if sliceOfOverlaps[i][j] == false {
				if sliceOfPluses[j].numElements*sliceOfPluses[i].numElements > maxProd {
					maxProd = sliceOfPluses[j].numElements * sliceOfPluses[i].numElements

				}
			}
		}
	}

	fmt.Printf("%d", maxProd)
}

func removeOverlaps(slice pluses) [][]bool {
	sliceOfOverlaps := make([][]bool, len(slice))
	for i := 0; i < len(sliceOfOverlaps); i++ {
		sliceOfOverlaps[i] = make([]bool, len(slice))
	}
	for x := 0; x < len(sliceOfOverlaps); x++ {
		for y := 0; y < len(slice)-1; y++ {
			sliceOfOverlaps[x][y] = false
		}
	}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {

			if haveCommonElem(slice[i].elems, slice[j].elems) {
				sliceOfOverlaps[i][j] = true
			}
		}
	}

	return sliceOfOverlaps
}

func haveCommonElem(a []elem, b []elem) bool {
	for _, elemx := range a {
		for _, elemy := range b {
			if elemx == elemy {
				return true
			}
		}
	}
	return false
}

func convertToIntAndAssignTo(s string, container *int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	*container = i
}

func checkGridForPlusesAndAddIntoSliceOfPluses(grid [][]bool, length int, width int) pluses {
	sliceOfPluses := pluses{}
	for x := 0; x < length; x++ {
		for y := 0; y < width; y++ {
			if grid[x][y] == true {
				if x == 0 || y == 0 {
					var temp plus
					var tempelem elem
					tempelem.x = x
					tempelem.y = y
					temp.numElements = 1
					temp.elems = append(temp.elems, tempelem)
					sliceOfPluses = append(sliceOfPluses, temp)
				} else {
					//fmt.Printf("%d,%d:", x, y)
					top := numOnTop(grid, x, y)
					//fmt.Printf("top:%d ", top)
					bot := numOnBot(grid, x, y, length)
					//fmt.Printf("bot:%d ", bot)
					if bot > top {
						bot = top
					}
					left := numOnLeft(grid, x, y)
					//fmt.Printf("left:%d ", left)
					if left > bot {
						left = bot
					}
					right := numOnRight(grid, x, y, width)
					//fmt.Printf("right:%d\n", right)
					if right > left {
						right = left
					}
					var temp plus
					var tempelem elem
					tempelem.x = x
					tempelem.y = y
					temp.numElements = 1
					temp.elems = append(temp.elems, tempelem)
					sliceOfPluses = append(sliceOfPluses, temp)
					for x1 := 1; x1 < right+1; x1++ {

						var tempelem1 elem
						tempelem1.x = x + x1
						tempelem1.y = y
						temp.elems = append(temp.elems, tempelem1)

						var tempelem2 elem
						tempelem2.x = x - x1
						tempelem2.y = y
						temp.elems = append(temp.elems, tempelem2)

						var tempelem3 elem
						tempelem3.x = x
						tempelem3.y = y + x1
						temp.elems = append(temp.elems, tempelem3)

						var tempelem4 elem
						tempelem4.x = x
						tempelem4.y = y - x1
						temp.elems = append(temp.elems, tempelem4)
						temp.numElements = x1*4 + 1
						sliceOfPluses = append(sliceOfPluses, temp)
					}

				}
			}
		}
	}
	return sliceOfPluses
}

func numOnTop(grid [][]bool, x int, y int) int {
	num := 0
	for x > 0 {
		if grid[x-1][y] == true {
			num++
			x--
		} else {
			break
		}
	}
	return num
}

func numOnBot(grid [][]bool, x int, y int, length int) int {
	num := 0
	for x < length-1 {
		if grid[x+1][y] == true {
			num++
			x++
		} else {
			break
		}
	}
	return num
}

func numOnLeft(grid [][]bool, x int, y int) int {
	num := 0
	for y > 0 {
		if grid[x][y-1] == true {
			num++
			y--
		} else {
			break
		}
	}
	return num
}

func numOnRight(grid [][]bool, x int, y int, width int) int {
	num := 0
	for y < width-1 {
		if grid[x][y+1] == true {
			num++
			y++
		} else {
			break
		}
	}
	return num
}

func (slice pluses) Len() int {
	return len(slice)
}

func (slice pluses) Less(i, j int) bool {
	return slice[i].numElements > slice[j].numElements
}

func (slice pluses) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
