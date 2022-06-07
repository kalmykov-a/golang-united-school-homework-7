package coverage

import (
	"os"
	"reflect"
	"testing"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestMatrix_Set(t *testing.T) {
	type fields struct {
		rows int
		cols int
		data []int
	}
	type args struct {
		row   int
		col   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "first test",
			fields: fields{rows: 5, cols: 5, data: []int{1, 2, 3, 4, 5}},
			args:   args{row: -1},
			want:   false,
		},
		{
			name:   "second test",
			fields: fields{rows: 5, cols: 5, data: []int{25: 0}},
			args:   args{row: 4, col: 4},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Set(tt.args.row, tt.args.col, tt.args.value); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cols(t *testing.T) {
	type fields struct {
		rows int
		cols int
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		// TODO: Add test cases.
		{
			name:   "First test",
			fields: fields{cols: 2, rows: 2, data: []int{1, 2, 3, 4}},
			want:   [][]int{{1, 3}, {2, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Cols(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cols() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Rows(t *testing.T) {
	type fields struct {
		rows int
		cols int
		data []int
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]int
	}{
		// TODO: Add test cases.
		{
			name:   "First test",
			fields: fields{cols: 2, rows: 2, data: []int{1, 2, 3, 4}},
			want:   [][]int{{1, 2}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Matrix{
				rows: tt.fields.rows,
				cols: tt.fields.cols,
				data: tt.fields.data,
			}
			if got := m.Rows(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}
