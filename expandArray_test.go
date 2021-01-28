package coding_test_20210128

import (
	"reflect"
	"testing"
)

func Test_expandArray(t *testing.T) {
	type args struct {
		arrTwo [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_expandArray: ",
			args: args{
				arrTwo: [][]int{
					{1, 3, 10},
					{1, 2, 10, 21, 100},
					{4, 7},
					{12, 15, 18, 30},
					{35},
				},
			},
			want: []int{
				1, 1, 2, 3, 4, 7, 10, 10, 12, 15, 18, 21, 30, 35, 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expandArray(tt.args.arrTwo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
