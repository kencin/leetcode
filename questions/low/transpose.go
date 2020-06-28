// 867.转置矩阵
// 2020.06.28 21:53 by Kencin

// 2020.06.28 22:27 Done

package low

import "fmt"

func Transpose(A [][]int) [][]int {
	lenOfColumns := len(A)
	lenOfRows := len(A[0])

	fmt.Println("len of columns: ", lenOfColumns)
	fmt.Println("len of rows: ", lenOfRows)

	var newA [][] int

	for i := 0; i < lenOfRows; i ++{
		aOfRows := make([]int, 0, lenOfRows)
		for j:= 0; j < lenOfColumns; j ++{
			aOfRows = append(aOfRows, A[j][i])
		}
		newA = append(newA, aOfRows)
	}

	return newA
}

//func main(){
//
//	testA := [][]int {{1,2,3}, {4,5,6}}
//
//	fmt.Println(low.Transpose(testA))
//}
