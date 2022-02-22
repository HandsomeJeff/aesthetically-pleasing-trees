package calculations

func Solution(A []int) int {
	if len(A) < 3 {
		return -1
	}

	if isAesthetic(A) {
		return 0
	}

	cut := 0

	for i := 0; i < len(A)-1; i++ {
		lhalf := make([]int, i)
		copy(lhalf, A[:i])
		cutOne := append(lhalf, A[i+1:]...)
		if isAesthetic(cutOne) {
			cut++
		}
	}

	if isAesthetic(A[:len(A)-1]) {
		cut++
	}

	if cut == 0 {
		return -1
	}
	return cut
}

func isAesthetic(trees []int) bool {
	for i := 1; i < len(trees)-1; i++ {
		l, m, r := trees[i-1], trees[i], trees[i+1]
		if !isAestheticTriplet(l, m, r) {
			return false
		}
	}
	return true
}


func isAestheticTriplet(l, m, r int) bool {
	return l > m && m < r || l < m && m > r
}


//func CalcAesthetic(trees []int) int {
//	if len(trees) < 3 {
//		return -1
//	}
//	cut := 0
//	for i := 1; i < len(trees)-1; i++ {
//		l, m, r := trees[i-1], trees[i], trees[i+1]
//		if !isAestheticTriplet(l, m, r) {
//			if cut >= 1 {
//				return -1
//			}
//			cut++
//			l, m = cutTree(l, m, r)
//			r = trees[i+2]
//			if !isAestheticTriplet(l, m, r) {
//				return -1
//			}
//		}
//	}
//
//
//	return 0
//}
//
//func cutTree(l, m, r int) (left, right int){
//	if l == m {
//		return m, r
//	}
//	if m == r {
//		return l, m
//	}
//
//	return l, r
//}