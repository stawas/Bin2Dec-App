package main

import (
	"errors"
	"strconv"
	"testing"
)

func TestInputNot0or1ShouldError(t *testing.T) {
	testName := "binaryToDecimal(12345)"
	want := errors.New(ERROR_NUMBER_NOT_0_OR_1)

	_, err := binaryToDecimal(12345)
	if err == nil || err.Error() != want.Error() {
		t.Fatalf("%v\n expected: %v\n actual: %v", testName, want.Error(), err.Error())
	}

}

func TestInputOver8lensShouldError(t *testing.T) {
	testName := "binaryToDecimal(101011011)"
	want := errors.New(ERROR_NUMBER_OVER_8_LENGTH)

	_, err := binaryToDecimal(101011011)
	if err == nil || err.Error() != want.Error() {
		t.Fatalf("%v\n expected: %v\n actual: %v", testName, want.Error(), err.Error())
	}

}

var testCases = []struct {
	input int
	output int
}{
	{0, 0},
	{1, 1},
	{10, 2},
	{1011, 11},
	{110101, 53},
	{1101011, 107},
	{11010110, 214},
	{11111111, 255},
	{00000000, 0},
}

func TestConvertBinaryToDecimal(t *testing.T) {
	for _, tt := range testCases {
		t.Run(strconv.Itoa(tt.input), func(t *testing.T) {
			decimal, err := binaryToDecimal(tt.input)

			if err != nil {
				t.Fatalf("expected: %v\n actual: %v", tt.output, err)
			}

			if decimal != tt.output {
				t.Fatalf("expected: %v\n actual: %v", tt.output, decimal)
			}
		})
	}
}
