package multiple_test

import (
	"testing"
)

var LOOP uint64 = 9999999999

func calN1(n uint64) uint64 {
	var res uint64 = 1
	var v1 uint64 = 1
	for v1 < n {
		res = v1 * res
		v1++
	}
	return res
}

func calN2(n uint64) uint64 {
	var res1 uint64 = 1
	var res2 uint64 = 1
	var res3 uint64 = 1
	var res4 uint64 = 1
	var v1 uint64 = 1
	for v1 < n {
		res1 = res1 * v1
		res2 = res2 * (v1 + 1)
		res3 = res3 * (v1 + 2)
		res4 = res4 * (v1 + 3)
		v1 += 4
	}
	return res1 + res2 + res3 + res4
}

func TestOne(t *testing.T) {
	calN1(LOOP)
}

func TestTwo(t *testing.T) {
	calN2(LOOP)
}

/*
From: https://mp.weixin.qq.com/s/LCvWnAdM4wUiGTqSoiBxrA
Result:

=== RUN   TestTwo
--- PASS: TestTwo (0.54s)
PASS
ok      command-line-arguments  0.542s

real    0m0.641s
user    0m0.682s
sys     0m0.173s
=== RUN   TestOne
--- PASS: TestOne (2.17s)
PASS
ok      command-line-arguments  2.167s

real    0m2.264s
user    0m2.336s
sys     0m0.133s
*/
