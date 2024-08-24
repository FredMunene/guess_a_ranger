package src

import (
	"reflect"
	"testing"
)

func TestGetTetromino(t *testing.T) {
	type args struct {
		tetrimino []string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{"Valid tetros", args{[]string{"....", "....", ".##.", ".##.","","\n"}}, [][]string{{"....","....",".##.",".##."}}, false},
		{"Valid tetros", args{[]string{"....", "....", ".##.", ".##.",""}}, [][]string{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTetromino(tt.args.tetrimino)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTetromino() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTetromino() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateBoard(t *testing.T) {
	type args struct {
		height int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Valid 2*2 board", args{4}, []string{"....", "....", "....", "...."}},
		{"Empty board",args{0},[]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateBoard(tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
