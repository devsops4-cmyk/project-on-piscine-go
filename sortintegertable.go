package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if table[j] > table[j+1] {
				// Swap
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
