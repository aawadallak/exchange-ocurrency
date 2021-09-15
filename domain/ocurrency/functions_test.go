package ocurrency

import (
	"testing"
)

func Test_validateFromAndTo(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Shouldn't return an error because field are valids",
			args: args{
				from: "BRL",
				to:   "USD",
			},
			wantErr: false,
		},
		{
			name: "Should return an error because field are invalid",
			args: args{
				from: "BRL",
				to:   "JPY",
			},
			wantErr: true,
		},
		{
			name: "Should return an error because field are invalid",
			args: args{
				from: "LTC",
				to:   "BRL",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateFromAndTo(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("validateFromAndTo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
