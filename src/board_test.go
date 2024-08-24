package src

import (
	"reflect"
	"testing"
)

func TestSolveBoard(t *testing.T) {
	type args struct {
		tetros [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Valid",args{[][]string{{"##","##"}}},[]string{"AA","AA"}},
		{"Valid",args{[][]string{{"####"},{"##","##"}}},[]string{"AAAA","BB..","BB..","...."}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveBoard(tt.args.tetros); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SolveBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveRecurs(t *testing.T) {
	type args struct {
		tetro [][]string
		board []string
		index int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Valid",args{[][]string{{"####"},{"##","##"}},[]string{"..",".."},0},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveRecurs(tt.args.tetro, tt.args.board, tt.args.index); got != tt.want {
				t.Errorf("solveRecurs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPlace(t *testing.T) {
	type args struct {
		tetro []string
		board []string
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Can PLace teromino is true",args{[]string{"##","##"},[]string{"..",".."},0,0},true},
		{"Cannot PLace teromino ",args{[]string{"##","##"},[]string{"##",".."},0,0},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlace(tt.args.tetro, tt.args.board, tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("canPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_place(t *testing.T) {
	type args struct {
		tetro []string
		board []string
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
	}{
		{"PLacing 1 teromino on a board",args{[]string{"##","##"},[]string{"..",".."},0,0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			place(tt.args.tetro, tt.args.board, tt.args.i, tt.args.j)
		})
	}
}

func Test_remove(t *testing.T) {
	type args struct {
		tetro []string
		board []string
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Removing 1 teromino from a board",args{[]string{"..",".."},[]string{"##","##"},0,0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			remove(tt.args.tetro, tt.args.board, tt.args.i, tt.args.j)
		})
	}
}

func Test_letterise(t *testing.T) {
	type args struct {
		tetro  []string
		letter int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Letterise the first tetromino",args{[]string{"##","##"},0},[]string{"AA","AA"}},
		{"Letterise the second tetromino",args{[]string{"####"},1},[]string{"BBBB"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterise(tt.args.tetro, tt.args.letter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterise() = %v, want %v", got, tt.want)
			}
		})
	}
}
