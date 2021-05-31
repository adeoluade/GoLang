//Declaration of main package
package main // This tells the Go compiler that the package should compile as an executable program instead of a shared library

//Importing packages
import (
	"bufio"   //This package allows the user provide input
	"fmt"     //This package implements formatted I/O (fmt means format)
	"math"    //This package provides basic constants and mathematical fuctions
	"os"      //This package provides a platform-independent interface to the OS functionality
	"strconv" //This package implements conversions to and from string representations of basic data types.
	"strings" // This package implements simple functions to manipulate UTF-8 encoded strings
	"time"    //This package provides the functionality for measuring and displaying time
)

//Main Function
func main() {
	fmt.Println("A simple calculator") // A simple header for the program
	time.Sleep(2 * time.Second)        // To delay next output
	fmt.Print("Value 1: ")             //
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)

}
