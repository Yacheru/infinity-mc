package websocket

import (
	"reflect"
	"testing"
)

func TestCutMessagePrefix(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "case1",
			args: args{
				bytes: []byte("Event PlayerChatEvent {\"message\":\"my message\",\"player\":{\"uuid\":\"0667138c-7e81-3d37-8bcf-d801adf345ae\",\"username\":\"yacheru\"}}"),
			},
			want: []byte("{\"message\":\"my message\",\"player\":{\"uuid\":\"0667138c-7e81-3d37-8bcf-d801adf345ae\",\"username\":\"yacheru\"}}"),
		},
		{
			name: "case2",
			args: args{
				bytes: []byte("{\"message\":\"my message\",\"player\":{\"uuid\":\"0667138c-7e81-3d37-8bcf-d801adf345ae\",\"username\":\"yacheru\"}}"),
			},
			want: []byte("{\"message\":\"my message\",\"player\":{\"uuid\":\"0667138c-7e81-3d37-8bcf-d801adf345ae\",\"username\":\"yacheru\"}}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutMessagePrefix(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CutMessagePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
