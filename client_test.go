package alldebrid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		key     string
		appname string
	}
	tests := []struct {
		name      string
		args      args
		want      Client
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "two empty args",
			args: args{
				appname: "",
				key:     "",
			},
			assertion: assert.Error,
		},
		{
			name: "only appname",
			args: args{
				appname: "myclient",
				key:     "",
			},
			assertion: assert.NoError,
		},
		{
			name: "no appname",
			args: args{
				appname: "",
				key:     "123456abcdef",
			},
			assertion: assert.Error,
		},
		{
			name: "no empty args",
			args: args{
				appname: "myclient",
				key:     "123456abcdef",
			},
			want: Client{
				ic: &innerClient{
					appName: "myclient",
					apikey:  "123456abcdef",
				},
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.key, tt.args.appname)

			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
