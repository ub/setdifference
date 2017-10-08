package setdifference

import "sort"

func OfSortedInts(A, B []int) (out []int) {

	var ib, ia int = 0, 0
	var blen, alen = len(B), len(A)

	skipleft := func() bool {
		ib++
		return ib <= blen
	}
	skipboth := func() {
		ib++
		ia++
	}

	takeright := func() bool {
		var val = A[ia]
		ia++
		out = append(out, val)
		return ia <= alen
	}

	for {
		if ib == blen {
			break
		}
		if ia == alen {
			break
		}

		if B[ib] < A[ia] {
			if ! skipleft() {
				break
			}
		} else if A[ia] < B[ib] {
			if ! takeright() {
				break
			}
		} else {
			skipboth()
		}
	}
	for ia < alen {
		takeright()
	}
	return
}

func OfInts(A, B []int) (out []int) {
	a := append([]T(nil), A...)
	b := append([]T(nil), B...)
	sort.Ints(a)
	sort.Ints(b)
	return OfSortedInts(a,b)
}