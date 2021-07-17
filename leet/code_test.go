package leet

import "testing"

func TestLeet_getNewData(t *testing.T) {
	tests := []struct {
		name string
		l    *Leet
		want string
	}{
		{
			name: "test1",
			l:    &Leet{},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Leet{}
			if got := l.getNewData(); got != tt.want {
				t.Errorf("Leet.getNewData() = %v, want %v", got, tt.want)
			}
		})
	}
}
