package connectrpc

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sessionCookie(t *testing.T) {
	type args struct {
		h http.Header
	}
	tests := map[string]struct {
		args    args
		want    string
		wantErr bool
	}{
		"success_1": {
			args: args{
				h: map[string][]string{
					"Header1": {"key1=val1", "key2=val2"},
					"Cookie":  {"69=1336", "ory_kratos_session=kratos", "other=other"},
					"Header2": {"key21=val21", "key22=val22"},
				},
			},
			want:    "ory_kratos_session=kratos",
			wantErr: false,
		},
		"success_2": {
			args: args{
				h: map[string][]string{
					"Cookie": {"ory_kratos_session=kratos"},
				},
			},
			want:    "ory_kratos_session=kratos",
			wantErr: false,
		},

		"not_found": {
			args: args{
				h: map[string][]string{
					"Header1": {"key1=val1", "key2=val2"},
					"Header2": {"key21=val21", "key22=val22"},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := sessionCookie(tt.args.h)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
