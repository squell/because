/*

Algorithm for computing multiplicative inverses in a field, due to
Colin Plumb: ftp://ftp.csc.fi/index/crypt/math/inverses-modulo-n.txt

Translated to Go

*/

package main

import "fmt"

type datatype int

func mult_inverse(a datatype, b datatype) datatype {
	var t0, t1, c, q datatype

	t1 = 1

	if b == 1 {
		return t1
	}

	t0 = a / b
	c = a % b

	for c != 1 {
		q = b / c
		b %= c
		t1 += q * t0
		if b == 1 {
			return t1
		}
		q = c / b
		c %= b
		t0 += q * t1
	}
	return a - t0
}

func main() {
	fmt.Println(mult_inverse(65537, 37))
}
