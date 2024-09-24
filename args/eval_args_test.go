package args

import (
	"testing"
)

// TestIsHelpFlag is a test function for the IsHelpFlag function
func TestIsHelpFlag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"test1", "--help", true},
		{"test2", "-help", false},
		{"test3", "-h", false},
		{"test4", "help", false},
		{"test5", "h", false},
		{"test6", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHelpFlag(tt.input); got != tt.want {
				t.Errorf("IsHelpFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestIsPathFlag is a test function for IsPathFlag function
func TestIsPathFlag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 bool
		want2 string
	}{
		{"test1", "-P=/home/Downloads", true, "/home/Downloads"},
		{"test2", "-P=/home", true, "/home"},
		{"test3", "-P=", false, ""},
		{"test4", "-P=...", true, "..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := IsPathFlag(tt.input)
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("IsPathFlag() = [%v %v], want [%v %v]", got1, got2, tt.want1, tt.want2)
			}
		})
	}
}

// TestInputFile is a test function for InputFile function
func TestInputFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 bool
		want2 string
	}{
		{"test1", "-i=file.txt", true, "file.txt"},
		{"test2", "-i=file", true, "file"},
		{"test3", "-i", false, ""},
		{"test4", "-i=.", false, ""},
		{"test5", "-i=..", false, ""},
		{"test6", "-i=...", true, "..."},
		{"test7", "-i=/", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := InputFile(tt.input)
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("InputFile() = [%v %v] , want [%v %v]", got1, got2, tt.want1, tt.want2)
			}
		})
	}
}

// TestIsOutputFlag is a test function for IsOutputFlag function
func TestIsOutputFlag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want1 bool
		want2 string
	}{
		{"test1", "-O=file.txt", true, "file.txt"},
		{"test2", "-O=file", true, "file"},
		{"test3", "-O=", false, ""},
		{"test4", "-o=file.txt", false, ""},
		{"test5", "-O=/", false, ""},
		{"test6", "-O=.", false, ""},
		{"test7", "-O=..", false, ""},
		{"test8", "-O=...", true, "..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := IsOutputFlag(tt.input)
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("IsOutputFlag() = got [%v %v], want [%v %v]", got1, got2, tt.want1, tt.want2)
			}
		})
	}
}
