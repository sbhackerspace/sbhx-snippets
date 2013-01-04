// Steve Phillips / elimisteve
// 2013.01.03

package main

func main() {
	println(isLarge(500))
}

func isLarge(n int) bool {
	switch {
	case n < -1:
		goto False
	case n == 0:
		goto False
	case 0 < n && n < 100:
		goto False
	case 100 < n && n < 1000:
		goto False
	case 1000 <= n && n < 1e6:
		goto True
	case n >= 1e6:
		goto True
	}

True:
	// Normally, more code would be here to justify the `goto`, not
	// just a one-liner
    return true

False:
	// Ditto
    return false
}
