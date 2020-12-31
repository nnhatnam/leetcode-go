package _32_Flipping_an_Image

func flipAndInvertArray(arr []int) {

	if len(arr) == 1 {
		arr[0] = arr[0] ^ 1
		return
	}

	for i, j  := 0, len(arr) - 1 ; i < j; i, j = i + 1, j - 1 {
		arr[i] , arr[j] = arr[j] ^ 1 , arr[i] ^ 1
		if i + 1 == j - 1 {
			arr[i + 1] ^= 1
		}
	}

}

func flipAndInvertImage(A [][]int) [][]int {
	for irow, _ := range A {
		flipAndInvertArray(A[irow])
	}
	return A
}
