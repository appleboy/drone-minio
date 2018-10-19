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

func TestPlugin_rmCommand(t *testing.T) {
	type fields struct {
		Config Config
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Remove a file.",
			fields: fields{
				Config: Config{
					Path: "1999/old-backup.tgz",
				},
			},
			want: []string{
				"mc",
				"rm",
				ALIAS + "/1999/old-backup.tgz",
			},
		},
		{
			name: "Remove all objects recursively from bucket 'jazz-songs' matching 'louis' prefix.",
			fields: fields{
				Config: Config{
					IsRecursive: true,
					Path:        "s3/jazz-songs/louis/",
				},
			},
			want: []string{
				"mc",
				"rm",
				"--recursive",
				ALIAS + "/s3/jazz-songs/louis/",
			},
		},
		{
			name: "Remove all objects older than '90' days recursively from bucket 'jazz-songs' that match 'louis' prefix.",
			fields: fields{
				Config: Config{
					IsRecursive: true,
					OlderThan:   90,
					Path:        "s3/jazz-songs/louis/",
				},
			},
			want: []string{
				"mc",
				"rm",
				"--recursive",
				"--older-than",
				"90",
				ALIAS + "/s3/jazz-songs/louis/",
			},
		},
		{
			name: "Remove all objects newer than 7 days recursively from bucket 'pop-songs'",
			fields: fields{
				Config: Config{
					IsRecursive: true,
					NewerThan:   7,
					IsDangerous: true,
					IsForce:     true,
					Path:        "s3/pop-songs/",
				},
			},
			want: []string{
				"mc",
				"rm",
				"--recursive",
				"--dangerous",
				"--force",
				"--newer-than",
				"7",
				ALIAS + "/s3/pop-songs/",
			},
		},
		{
			name: "Remove an encrypted object from s3.",
			fields: fields{
				Config: Config{
					EncryptKey: "s3/ferenginar/=32byteslongsecretkeymustbegiven1",
					Path:       "s3/ferenginar/1999/old-backup.tgz",
				},
			},
			want: []string{
				"mc",
				"rm",
				"--encrypt-key",
				`"s3/ferenginar/=32byteslongsecretkeymustbegiven1"`,
				ALIAS + "/s3/ferenginar/1999/old-backup.tgz",
			},
		},
		{
			name: "Drop all incomplete uploads on 'jazz-songs' bucket.",
			fields: fields{
				Config: Config{
					IsIncomplete: true,
					IsRecursive:  true,
					Path:         "s3/jazz-songs/",
				},
			},
			want: []string{
				"mc",
				"rm",
				"--recursive",
				`--incomplete`,
				ALIAS + "/s3/jazz-songs/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{
				Config: tt.fields.Config,
			}
			if got := p.rmCommand().Args; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plugin.rmCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
