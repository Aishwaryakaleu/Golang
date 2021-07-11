package main  //The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library.
                 The main function in the package “main” will be the entry point of our executable program
                 Program does not exicute if there is no package main in Golang//

import "fmt"  // fmt is short for format.
                 Package fmt implements formatted I/O with functions analogous to C's printf and scanf//

func main() {
	fmt.Println("Hello, world!")
}
