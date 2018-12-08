package main

import "os"
import "fmt"
import "strings"

func main() {

	/*argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]*/
	
	arg := len(os.Args)

	string1:="";
	string2:="";
	if(arg == 4) {
		string1=os.Args[1];
		string2=os.Args[2];
	} else {
		return
	}
	/* print the args */
	if (strings.EqualFold(string1, string2)) {
		fmt.Printf("%s is equal to %s \n", string1, string2)
	} else {
		fmt.Printf("%s is not equal to %s \n", string1, string2)
	}
}
