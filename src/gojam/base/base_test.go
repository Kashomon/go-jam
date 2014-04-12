package base

import (
	"testing"
)

func TestGetInt(t *testing.T) {
	if n := GetInt("1"); n != 1 {
		t.Errorf("Not 1")
	}
}

func TestGetIntList(t *testing.T) {
	if n := GetIntList("1 2 3");
			n[0] != 1 || n[1] != 2 || n[2] != 3 {
		t.Errorf("get int list fail, %v", n)
	}
}

func TestGetFloat(t *testing.T) {
	if n := GetFloat("1.23"); n != 1.23 {
		t.Errorf("Not 1.23")
	}
}

func TestGetFloatList(t *testing.T) {
	if n := GetFloatList("1.23 4.56");
			n[0] != 1.23 || n[1] != 4.56 {
		t.Errorf("get int list fail, %v", n)
	}
}
