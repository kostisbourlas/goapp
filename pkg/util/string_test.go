package util

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRandStringWithPositiveLength(t *testing.T) {
	length := 10
	string := RandString(length)

	if len(string) != length {
		t.Errorf("Expected string length %d, got %d", length, len(string))
	}

	regex := regexp.MustCompile(`^[0-9A-F]+$`)
	if !regex.MatchString(string) {
		t.Errorf("Expected a hexadecimal value, got %s", string)
	}
}

func TestRandStringWithZeroLength(t *testing.T) {
	length := 0
	string := RandString(length)

	if len(string) != length {
		t.Errorf("Expected string length %d, got %d", length, len(string))
	}

	if string != "" {
		t.Errorf("Expected emptry string, got %s", string)
	}
}

func BenchmarkRandStringLength(b *testing.B) {
	lengths := []int{1000, 10000, 100000}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("Length=%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RandString(length)
			}
		})
	}
}
