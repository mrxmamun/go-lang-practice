package main
import "fmt"
func main() {
	//  var a  [5]int
	// fmt.Println("empty:", a)
	// b := [5]int{1, 2, 3, 4, 5} //size with array elememt
	// fmt.Println( b)
 	// nums:=[]int{2,4} //without fixed size
	// fmt.Println(nums)
	var twoD [2][3]int
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D", twoD)
}