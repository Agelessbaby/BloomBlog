package errno

import (
	"errors"
	"github.com/Agelessbaby/BloomBlog/util/errno"
	code "github.com/a76yyyy/ErrnoCode"
	"reflect"
	"testing"
)

func TestConvertErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want errno.ErrNo
	}{
		// TODO: Add test cases.
		{
			name: "TestConvertErr",
			args: args{err: errors.New("")},
			want: errno.ErrNo{
				ErrCode: code.ErrBind,
				ErrMsg:  "Binding Error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errno.ConvertErr(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
