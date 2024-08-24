package src

import "testing"

func TestValidTetrimino(t *testing.T) {
	type args struct {
		cube []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidTetrimino(tt.args.cube)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidTetrimino() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidTetrimino() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConnections(t *testing.T) {
	type args struct {
		cube []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Connections(tt.args.cube); got != tt.want {
				t.Errorf("Connections() = %v, want %v", got, tt.want)
			}
		})
	}
}
