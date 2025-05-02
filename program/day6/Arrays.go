package main
import "fmt"

func main(){
	fmt.Println("this program is for understanding array concept in golang")
	//creating an array
	var arr [5]int
	//using len function to calculate size of an array
	var size =len(arr)
	fmt.Println("size of array is",size)
	//printing array
	fmt.Println(arr)

	for i:=0;i<size;i++ {
		fmt.Println("enter the ",i," element")
		fmt.Scan(&arr[i])
	}

	//reprinting an array
	fmt.Println(arr)

}