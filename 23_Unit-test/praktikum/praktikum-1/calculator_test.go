package praktikum1

import "testing"

var (
	x int
	y int
)

func TestAddition(t *testing.T) {
	x, y = 5, 4
	if Addition(x, y) != 9 {
		t.Error("Expected x (+) y to equal 9")
	}
	x, y = -5, -4
	if Addition(x, y) != -9 {
		t.Error("Expected x (+) y to equal -9")
	}
}

func TestSubtraction(t *testing.T) {
	x, y = 3, 2
	if Subtraction(x, y) != 1 {
		t.Error("Expected x (-) y to equal 1")
	}

	x, y = -3, 2
	if Subtraction(x, y) != -5 {
		t.Error("Expected x (-) y to equal 5")
	}
}

func TestMultiplication(t *testing.T) {
	x, y = 3, 2
	if Multiplication(x, y) != 6 {
		t.Error("Expected x (*) y to equal 6")
	}
}

func TestDivision(t *testing.T) {
	x, y = 4, 2
	if Division(x, y) != 2 {
		t.Error("Expected x (/) y to equal 2")
	}
}
