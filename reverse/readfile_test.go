package reverse

import (
	"os"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 bool
	}{
		{
			name:  "Successful Read Unix Line Endings",
			args:  args{fileName: createTempFile(t, "line1\nline2\nline3")},
			want:  []string{"line1", "line2", "line3"},
			want1: false,
		},
		{
			name:  "File Not Found",
			args:  args{fileName: "nonexistentfile.txt"},
			want:  nil,
			want1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ReadFile(tt.args.fileName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

// Helper function to create a temporary file with the specified content
func createTempFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()
	return tmpFile.Name()
}
