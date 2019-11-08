package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		character   string
		repeatCount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "character repeated 5 times",
			args: args{
				character:   "a",
				repeatCount: 5,
			},
			want: "aaaaa",
		},
		{
			name: "character repeated 1 times",
			args: args{
				character:   "a",
				repeatCount: 1,
			},
			want: "a",
		},
		{
			name: "character repeated 10 times",
			args: args{
				character:   "a",
				repeatCount: 10,
			},
			want: "aaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.character, tt.args.repeatCount); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
