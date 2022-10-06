package decoders

import (
	"testing"
)

func TestNotify(t *testing.T) {
	test := []struct {
		nameTest string
		nameCode string
	}{
		{
			nameTest: "Test Notify",
			nameCode: "adb",
		},
	}
	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			sd := &Morse{}
			sd.Notify(tt.nameCode)
		})
	}
}

func TestMorse(t *testing.T) {
	tests := []struct {
		nameTest string
		name     string
		want     string
	}{
		{
			nameTest: "Morse",
			name:     "abc",
			want:     ".--....-.-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.nameTest, func(t *testing.T) {
			codeMorse(tt.name)
		})
	}
}
