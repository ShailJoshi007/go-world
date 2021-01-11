package main
import "fmt"

//For loop
func main() {  
for i := 1; i <= 5; i++ {
fmt.Println(i)
    }

//IF-ELSE loop
a:=50
if a<100{
	fmt.Println("a is less then 100")
	}else{
		fmt.Println("a is greater then 100")
	}

//Multiple Else loop
var x = 100
    if x < 10 {
        //Executes if x is less than 10
        fmt.Println("x is less than 10")
    } else if x >= 10 && x <= 90 {
        //Executes if x >= 10 and x<=90
        fmt.Println("x is between 10 and 90")
    } else {
        //Executes if both above cases fail i.e x>90
        fmt.Println("x is greater than 90")
    }
//SWITCH loop
b,c := 2,1
    switch b+c {
    case 1:
        fmt.Println("Sum is 1")
    case 2:
        fmt.Println("Sum is 2")
    case 3:
        fmt.Println("Sum is 3")
    default:
        fmt.Println("Printing default")
    }

}