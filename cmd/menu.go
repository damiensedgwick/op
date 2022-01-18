package cmd

import "fmt"

func Menu() {
	fmt.Println("***************************************************************")
	fmt.Println("*                              OP                             *")
	fmt.Println("* ----------------------------------------------------------- *")
	fmt.Println("*                                                             *")
	fmt.Println("* In order to use op, you need to search in one of two ways.  *")
	fmt.Println("*                                                             *")
	fmt.Println("* You can search using the following commands.                *")
	fmt.Println("*                                                             *")
	fmt.Println("* 1) op <language|framework|library> - single word queries.   *")
	fmt.Println("* 2) op mdn your search query - single / multi word queries.  *")
	fmt.Println("*                                                             *")
	fmt.Println("* More ways to use op will be released soon.                  *")
	fmt.Println("*                                                             *")
	fmt.Println("***************************************************************")
}
