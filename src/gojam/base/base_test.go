package base

import (
	"testing"
)

//
// Int Tests
//

func TestGetInt(t *testing.T) {
	if n := GetInt("1"); n != 1 {
		t.Errorf("Not 1")
	}
}

func TestGetIntList(t *testing.T) {
	if n := GetIntList("1 2 3"); n[0] != 1 || n[1] != 2 || n[2] != 3 {
		t.Errorf("get int list fail, %v", n)
	}
}

func TestReverseInts(t *testing.T) {
	k := []int{0, 2, 3}
	ReverseInts(k)
	if k[0] != 3 && k[1] != 2 && k[2] != 0 {
		t.Errorf("fail -- val: %v", k)
	}
	k = []int{0, 2, 3, 4}
	ReverseInts(k)
	if k[0] != 4 && k[1] != 3 && k[2] != 2 && k[3] != 0 {
		t.Errorf("fail -- val: %v", k)
	}
}

//
// Float Tests
//

func TestGetFloat(t *testing.T) {
	if n := GetFloat("1.23"); n != 1.23 {
		t.Errorf("Not 1.23")
	}
}

func TestGetFloatList(t *testing.T) {
	if n := GetFloatList("1.23 4.56"); n[0] != 1.23 || n[1] != 4.56 {
		t.Errorf("get int list fail, %v", n)
	}
}

func TestReverseFloats(t *testing.T) {
	k := []float64{0, 2.2, 3}
	ReverseFloats(k)
	if k[0] != 3 && k[1] != 2.2 && k[2] != 0 {
		t.Errorf("fail -- val: %v", k)
	}
	k = []float64{0, 2, 3, 4.1}
	ReverseFloats(k)
	if k[0] != 4.1 && k[1] != 3 && k[2] != 2 && k[3] != 0 {
		t.Errorf("fail -- val: %v", k)
	}
}
