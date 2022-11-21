package imagebuilder

// @link https://pkg.go.dev/github.com/fogleman/gg --> image manager
import (
	"fmt"
	"github.com/fogleman/gg"
	"log"
	"strconv"
	"time"
)

const (
	width        = 800
	height       = 600
	colourLength = 16
)

// BuildImage
// Main function for building the image based off what was the last image iteration
// Takes in the name of the current image should be a string representation of an integer
// Generates following increment
func BuildImage(oldName string) {
	iteration, err := strconv.Atoi(oldName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(iteration)

	// width(x), height(y)
	current := gg.NewContext(width, height)

	// Mess with this value to see outcomes
	start := time.Now()

	// @todo is there a way to generate faster through dividing the iteration number?
	// @todo dec to hex conversion algorithm will save time
	newColourPoints := generateIteration(100)
	fmt.Printf("%v", newColourPoints)

	hex := strconv.FormatInt(100, 16)
	fmt.Println(hex)

	end := time.Since(start)
	fmt.Println(end)
	fmt.Println("")

	x, y := 0, 0
	for i, num := range newColourPoints {
		// X, Y, r
		fmt.Println(i)
		// https://pkg.go.dev/github.com/fogleman/gg#SetPixel
		colour := getColourFromList(num)
		fmt.Println(colour)
		// We need to define the colour before we set the pixel as the pixel takes the current colour
		current.SetHexColor(colour)
		current.SetPixel(x, y)
		x++
		if x == width {
			x = 0
			y++
		}
	}
	newName := fileNameBuilder(iteration + 1)
	current.SavePNG(newName)
}

// Gets the colour from the list based off a pointer
func getColourFromList(pointer int) string {
	// Greyscale
	// @todo create actual list of colours
	colourList := []string{
		"000",
		"111",
		"222",
		"333",
		"444",
		"555",
		"666",
		"777",
		"888",
		"999",
		"AAA",
		"BBB",
		"CCC",
		"DDD",
		"EEE",
		"FFF",
	}
	return colourList[pointer]
}

// Takes in the current iteration and increases by 1.
// Then returns the full filename for what the new current image would be
func fileNameBuilder(iteration int) string {
	fileName := strconv.Itoa(iteration)
	return "./current/" + fileName + ".png"
}

// This output an array of points to be printed
// This will iterate through the range
func generateIteration(iteration int) []int {
	var newColours []int
	for i := 0; i <= iteration; i++ {
		// add first point
		// increment until next position needs to increase
		// go back to start and increment

		// on first iteration we want to make a new point
		if len(newColours) > 0 {
			newColours[0]++
			// The denominator should match the length of colours in getColourFromList
			if newColours[0] == colourLength {
				newColours[0] = 0
				newColours = increaseNextPointer(newColours, 0)
			}
		} else {
			// make first point
			newColours = insert(newColours, 0, 0)
		}
	}
	return newColours
}

// Will increase the next position of the array
// Either by incrementing the next point or creating the new position
func increaseNextPointer(newColours []int, currPointer int) []int {
	// if the length is greater than the current pointer we can iterate
	newPointer := currPointer + 1
	if len(newColours) >= newPointer {
		// if the length is the same amount as the next pointer then we should increase
		if len(newColours) == newPointer {
			newColours = insert(newColours, newPointer, 0)
		} else {
			// increment the next position
			newColours[newPointer]++
		}
		// if are at the "colour" threshold, then rest required on next position and time to recur
		if newColours[newPointer] == colourLength {
			newColours[newPointer] = 0
			newColours = increaseNextPointer(newColours, newPointer)
		}
	} else {
		// length not higher, then we need to start the next position
		newColours = insert(newColours, newPointer, 0)
	}
	return newColours
}

// Inserts record to next iteration of the array
// 0 <= index <= len(a)
func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
