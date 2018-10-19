package main

import (
	"reflect"
	"testing"
)

func TestPlugin_versionCommand(t *testing.T) {
	type fields struct {
		Config Config
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name:   "get version command",
			fields: fields{},
			want:   []string{"mc", "version"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{
				Config: tt.fields.Config,
			}
			if got := p.versionCommand().Args; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plugin.versionCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlugin_addConfigCommand(t *testing.T) {
	type fields struct {
		Config Config
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "get version command",
			fields: fields{
				Config: Config{
					URL:       "http://example.com",
					AccessKey: "1234",
					SecretKey: "5678",
				},
			},
			want: []string{
				"mc",
				"config",
				"host",
				"add",
				"minio",
				"http://example.com",
				"1234",
				"5678",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{
				Config: tt.fields.Config,
			}
			if got := p.addConfigCommand().Args; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plugin.addConfigCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
