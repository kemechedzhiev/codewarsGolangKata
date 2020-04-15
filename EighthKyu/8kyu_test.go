package EighthKyu

import "testing"

func TestXor(t *testing.T) {
	// false xor false == false // since both are false
	// true xor false == true // exactly one of the two expressions are true
	// false xor true == true // exactly one of the two expressions are true
	// true xor true == false
	a, b := true, false
	result := xor(b, b)
	if result != false {
		t.Error("Wrong answer! Expected false, got true")
	}
	result = xor(a, b)
	if result != true {
		t.Error("Wrong answer! Expected false, got true")
	}
	result = xor(b, a)
	if result != true {
		t.Error("Wrong answer! Expected false, got true")
	}
	result = xor(a, a)
	if result != false {
		t.Error("Wrong answer! Expected false, got true")
	}
}

func TestOddsCounter(t *testing.T) {
	// Test 1
	dataset := 7
	result := oddsCounter(dataset)
	if result != 3 {
		t.Error("Wrong answer! Expected 3, got ", result)
	}
	// Test 2
	dataset = 15
	result = oddsCounter(dataset)
	if result != 7 {
		t.Error("Wrong answer! Expected 7, got ", result)
	}
}

func TestDnaToRna(t *testing.T) {
	dataset := "GCAT"
	result := dnaToRna(dataset)
	if result != "GCAU" {
		t.Error("Wrong answer! Replacement didn't occur!")
	}
}

func TestIsThisMyTail(t *testing.T) {
	dataset1 := "donkey"
	var dataset2 rune
	dataset2 = 'y'
	result := isThisMyTale(dataset1, dataset2)
	if !result {
		t.Error("WTF? Donkey without a tail? Are you high?")
	}
}
