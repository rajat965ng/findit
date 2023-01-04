package service

import (
	"findings/model"
	"reflect"
	"testing"
)

func Test_gitService_GitClone(t *testing.T) {
	type fields struct {
		repository *model.Repository
		target     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "go-secret-scanner",
			fields: fields{
				&model.Repository{Name: "go-secret-scanner", Url: "https://github.com/rajat965ng/go-secret-scanner"}, "./dir",
			},
			want: "./dir/go-secret-scanner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &gitService{
				repository: tt.fields.repository,
				target:     tt.fields.target,
			}
			got, err := svc.GitClone()
			if (err != nil) != tt.wantErr {
				t.Errorf("gitService.GitClone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gitService.GitClone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gitService_GrepText(t *testing.T) {
	type fields struct {
		repository *model.Repository
		target     string
	}
	type args struct {
		patterns []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "go-secret-scanner",
			fields: fields{
				&model.Repository{Name: "go-secret-scanner", Url: "https://github.com/rajat965ng/go-secret-scanner"}, "./dir",
			},
			args: args{
				patterns: []string{"private-key", "public-key"},
			},
			want: []string{"findings.txt:3", "findings.txt:6", "findings.txt:8", "findings.txt:10", "private-key.txt:3", "private-key.txt:6", "private-key.txt:8", "private-key.txt:10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &gitService{
				repository: tt.fields.repository,
				target:     tt.fields.target,
			}
			got, err := svc.GrepText(tt.args.patterns)
			if (err != nil) != tt.wantErr {
				t.Errorf("gitService.GrepText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gitService.GrepText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gitService_CleanUp(t *testing.T) {
	type fields struct {
		repository *model.Repository
		target     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "go-secret-scanner",
			fields: fields{
				&model.Repository{Name: "go-secret-scanner", Url: "https://github.com/rajat965ng/go-secret-scanner"}, "./dir",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &gitService{
				repository: tt.fields.repository,
				target:     tt.fields.target,
			}
			if err := svc.CleanUp(); (err != nil) != tt.wantErr {
				t.Errorf("gitService.CleanUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
