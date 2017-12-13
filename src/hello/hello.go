package main

import "fmt"
//import "math"
// import "time"


const s string = "constant"

func main(){
	// <<<<<<<>>>>>>>>	Lession 1: Variable
	//var a string = "hello";

	// var c int = 10;

	// fmt.Println("c +10: ", c+10);

	// var t = true;
	// fmt.Println(t);

	// <<<<<<<>>>>>>>>	Lession 2: Variable continute
	// const n = 500000000;
	// const d = 3e20/n;
	// fmt.Println("d before: " ,d);
	// fmt.Println("d After : " ,(int64 (d)))

	// fmt.Println("sin: ", math.Sin(n))


	// <<<<<<<>>>>>>>> Lession 3: Loop
	// for i:=1 ; i <= 10; i++{
		
	// 	if( i%2 == 0){
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	

	// <<<<<<<>>>>>>>> Lession 4: If/else
	
	// if(3%2 == 0){
	// 	fmt.Println("3 is even")
	// } else {
	// 	fmt.Println("3 is odd")
	// }


	// <<<<<<<>>>>>>>> Lession 5: Switch
	// i := 2;

	// fmt.Println("Write i: ", i, " as ");

	// switch i {
	// case 1:
	// 	fmt.Println("Hello one")
	// case 2:
	// 	fmt.Println("Hello two")
	// }

	// t := time.Now().Day();
	// fmt.Println("Now is: ", t)

	// whatAmI := func(i interface{}){
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a Boolean")
	// 	case int:
	// 		fmt.Println("I'm an Integer")
	// 	default: 
	// 		fmt.Println("Don't know type %T\n", t)		
	// 	}
		
	// }

	// whatAmI("true")


	// <<<<<<<>>>>>>>> Lession: slice
	s := make([] string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s , "e", "f")
	fmt.Println(s)
	fmt.Println("length current: ", len(s))

	c := make([] string, len(s))
	copy(c,s)
	fmt.Println("c: " ,c)
	fmt.Println("-----------------Slice-----------")
	l := s[2:5]
	fmt.Println("l: " ,l)

	l = s[:5]
	fmt.Println("sl2:" ,l)
	l = s[2:]
	fmt.Println("sl3:", l)
	
}