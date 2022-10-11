package main

import (
	"fmt"
)

//declaring variables outside of a function
var m int;
var n int = 10;
var o = 10;

func main() {
	// Variables are declared using the var keyword
	var x int32 = 10;

	// Variables can be declared without an initial variable name
	var student = "Jhon";

	// type is inferred
	y := 10.5;

	// Variables can be declared without an initial value
	var a string
	var b int
	var c bool

	//Value Assignment After Declaration
	var student1 string
	student1 = "Aftab";

	//Multiple Variable Declaration
	var student2, student3 string = "Aftab", "Jhon";

	//Multiple Variable Declaration in a block
	var (
		student4 string = "Aftab-IV"
		student5 string = "Jhon-V"
		sl 	 int    = 1
		sl2  int   	= 2
	)

	//Multi-Word Variable Names - Camel Case
	var myName = "Aftab-I";

	//Multi-Word Variable Names - Pascal Case
	var MyName = "Aftab-II";

	//Multi-Word Variable Names - Snake Case
	var my_name = "Aftab-III";


	fmt.Println(x);
	fmt.Println(student);
	fmt.Println(y);

	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	
	fmt.Println(student1);

	m = 11;
	fmt.Println(m);
	fmt.Println(n);
	fmt.Println(o);

	fmt.Println(student2);
	fmt.Println(student3);

	fmt.Println(student4);
	fmt.Println(student5);
	fmt.Println(sl);
	fmt.Println(sl2);

	fmt.Println(myName);
	fmt.Println(MyName);
	fmt.Println(my_name);
}