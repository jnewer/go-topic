package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3,but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30,but %d got", ans)
	}
}

type calCase struct {
	A, B, Expected int
}

func createMulTestCase(t *testing.T, c *calCase) {
	//t.Helper()//报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息
	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d excepted %d, but %d got", c.A, c.B, c.Expected, ans)
	}
}
func TestMul(t *testing.T) {
	if ans := Mul(1, 2); ans != 2 {
		t.Errorf("1 * 2 expected be 2,but %d got", ans)
	}

	if ans := Mul(-10, -20); ans != 200 {
		t.Errorf("-10 * -20 expected be 200,but %d got", ans)
	}

	//子测试
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})

	//cases := []struct {
	//	Name           string
	//	A, B, Expected int
	//}{
	//	{"pos", 2, 3, 6},
	//	{"neg", 2, -3, -6},
	//	{"zero", 2, 0, 0},
	//}
	//
	//for _, c := range cases {
	//	t.Run(c.Name, func(t *testing.T) {
	//		if ans := Mul(c.A, c.B); ans != c.Expected {
	//			t.Fatalf("%d * %d excepted %d, but %d got", c.A, c.B, c.Expected, ans)
	//		}
	//	})
	//}

	createMulTestCase(t, &calCase{2, 3, 6})
	createMulTestCase(t, &calCase{2, -3, -6})
	createMulTestCase(t, &calCase{2, 0, 1})
}
