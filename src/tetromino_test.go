package src

import (
	"reflect"
	"testing"
)

func TestResizeTetri(t *testing.T) {
	type args struct {
		cube []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Resizing required", args{[]string{"....", "....", ".##.", ".##."}}, []string{"##", "##"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResizeTetri(tt.args.cube); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResizeTetri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRowEmpty(t *testing.T) {
	type args struct {
		row string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty Row", args{"...."}, true},
		{"Non-empty Row", args{"...#"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRowEmpty(tt.args.row); got != tt.want {
				t.Errorf("isRowEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isColumnEmpty(t *testing.T) {
	type args struct {
		cube   []string
		colNum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty Row", args{[]string{"....", "....", ".##.", ".##."}, 0}, true},
		{"Non-empty Row", args{[]string{"....", "....", ".##.", "#.#."}, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isColumnEmpty(tt.args.cube, tt.args.colNum); got != tt.want {
				t.Errorf("isColumnEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
