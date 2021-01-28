package coding_test_20210128

import (
	"reflect"
	"testing"
)

func Test_expandArrayWithBucket(t *testing.T) {
	type args struct {
		inputs [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test_expandArrayWithBucket",
			args: args{
				inputs: [][]int{
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
			if got := expandArrayWithBucket(tt.args.inputs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandArrayWithBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
