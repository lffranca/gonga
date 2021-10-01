package service

import (
	"context"
	"github.com/lffranca/gonga/core/entity"
	"reflect"
	"testing"
)

func str(s string) *string {
	return &s
}

func Test_service_List(t *testing.T) {
	mock := &dataMock{}
	ctx := context.Background()

	type fields struct {
		data Data
	}
	type args struct {
		ctx    context.Context
		offset *int
		limit  *int
		search *string
		sort   *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Service
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "Test service list",
			fields: fields{data: mock},
			args: args{
				ctx: ctx,
			},
			want: []entity.Service{
				{
					ID:   str("1"),
					Name: str("Name 1"),
				},
				{
					ID:   str("2"),
					Name: str("Name 2"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			module := &service{
				data: tt.fields.data,
			}
			got, err := module.List(tt.args.ctx, tt.args.offset, tt.args.limit, tt.args.search, tt.args.sort)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type dataMock struct{}

func (mod *dataMock) List(ctx context.Context, offset, limit *int, search, sort *string) ([]entity.Service, error) {
	return []entity.Service{
		{
			ID:   str("1"),
			Name: str("Name 1"),
		},
		{
			ID:   str("2"),
			Name: str("Name 2"),
		},
	}, nil
}

func (mod *dataMock) Create(ctx context.Context, service *entity.Service) (*entity.Service, error) {
	return nil, nil
}
