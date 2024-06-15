package utils

// []int {9, 7, 5, 3, 1, 2, 4, 6, 8}
// []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
func BubbleSort(elements []int) {

	keepRuning := true

	for keepRuning {
		keepRuning = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRuning = true
			}
		}
	}
}
