package service

import (
	"findings/dao"
	"testing"
)

func Test_repository_ExecuteScanner(t *testing.T) {
	type fields struct {
		repositoryDao dao.IRepositoryDao
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "Test Execute Scanner",
			fields: fields{
				repositoryDao: dao.NewRepositoryDao(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &repository{
				repositoryDao: tt.fields.repositoryDao,
			}
			svc.ExecuteScanner()
		})
	}
}
