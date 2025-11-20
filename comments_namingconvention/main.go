package main

//**COMMENTS88**//

// This is a single line comment

/*
	This is a single block comment
	a :=10
	fmt.Println(a)
	}
*/

var name = "john Wick" //inline comment

//** NAMING CONVENTION **//

var s string
var i int
var num int
var msg string
var v string
var err error
var done bool

//use camel case instead of snake_case

var maxValue = 100  //recommended (camel case)
var max_Value = 100 //not recommended (snake_case)

//recommended
func writeToFile() {

}

//not recommended
func write_to_file() {

}

//write acronyms in all caps
var writeToDB = true // recommended
var writeToD = true  // not recommended

func main() {

	// use fewer letters dont be too verbose

	var packetsreceived int //Not ok, too verbose
	var n int               // ok

	_, _ = packetsreceived, n

	// an uppercase first letter has special significance in go (it will be exported in other packages)

}
