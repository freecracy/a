package content

import (
	"reflect"
	"testing"
)

func TestGetAllFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "docs",
			args: args{
				name: "../docs",
			},
			want: []string{
				"README",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllFile(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
