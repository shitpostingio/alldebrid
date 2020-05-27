package alldebrid

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		key     string
		appname string
	}
	tests := []struct {
		name    string
		args    args
		want    Client
		wantErr bool
	}{
		{
			name: "two empty args",
			args: args{
				appname: "",
				key:     "",
			},
			wantErr: true,
		},
		{
			name: "one empty arg",
			args: args{
				appname: "myclient",
				key:     "",
			},
			wantErr: true,
		},
		{
			name: "one empty arg",
			args: args{
				appname: "",
				key:     "123456abcdef",
			},
			wantErr: true,
		},
		{
			name: "no empty args",
			args: args{
				appname: "myclient",
				key:     "123456abcdef",
			},
			wantErr: false,
			want: Client{
				appName: "myclient",
				apikey:  "123456abcdef",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.key, tt.args.appname)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
