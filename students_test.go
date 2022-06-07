package coverage

import (
	"os"
	"reflect"
	"testing"
	"time"
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
		{
			name:   "first test: negative row",
			fields: fields{rows: 5, cols: 5, data: []int{1, 2, 3, 4, 5}},
			args:   args{row: -1},
			want:   false,
		},
		{
			name:   "second test: correct args",
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
		{
			name:   "first test: correct args",
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
		{
			name:   "first test: correct args",
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

func TestPeople_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
	}{
		{
			name: "first test: all correct",
			p:    []Person{{"a", "b", time.Unix(1, 0)}, {"aa", "bb", time.Unix(2, 0)}},
			args: args{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestPeople_Len(t *testing.T) {
	tests := []struct {
		name string
		p    People
		want int
	}{
		{
			name: "first test: all correct",
			p:    []Person{{"a", "b", time.Unix(1, 0)}, {"aa", "bb", time.Unix(2, 0)}},
			want: 2,
		},
		{
			name: "second test: empty",
			p:    []Person{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeople_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		p    People
		args args
		want bool
	}{
		{
			name: "first less",
			p:    People{{"a", "a", time.Unix(3, 0)}, {"b", "b", time.Unix(2, 0)}},
			args: args{0, 1},
			want: true,
		},
		{
			name: "both date equal",
			p:    People{{"a", "a", time.Unix(1, 0)}, {"b", "a", time.Unix(1, 0)}},
			args: args{0, 1},
			want: true,
		},
		{
			name: "both equal, same names",
			p:    People{{"b", "b", time.Unix(1, 0)}, {"b", "a", time.Unix(1, 0)}},
			args: args{0, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *Matrix
		wantErr bool
	}{
		{
			name:    "all correct",
			args:    args{"1 2\n3 4"},
			want:    &Matrix{2, 2, []int{1, 2, 3, 4}},
			wantErr: false,
		},
		{
			name:    "incorrect cols",
			args:    args{"1 2\n 5\n 3 4"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "not a number",
			args:    args{"1 2\na 4"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
