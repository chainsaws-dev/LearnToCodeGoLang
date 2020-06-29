package main

import "testing"

func Test_sqrt(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{name: "Positive nums test",
			args:    args{f: 10},
			want:    42,
			wantErr: false},

		{name: "Negative nums test",
			args:    args{f: -10},
			want:    -1,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sqrt(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("sqrt() test `%v` failed = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sqrt() test `%v` failed = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
