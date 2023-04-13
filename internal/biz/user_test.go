package biz

import (
	"context"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func TestHttpUsecase_Login(t *testing.T) {
	type fields struct {
		repo Repo
		log  *log.Helper
	}
	type args struct {
		ctx context.Context
		u   *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1.LoginResponse
		wantErr bool
	}{
		{
			name: "login_success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* uc := &HttpUsecase{
				repo: tt.fields.repo,
				log:  tt.fields.log,
			}
			got, err := uc.Login(tt.args.ctx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpUsecase.Login() = %v, want %v", got, tt.want)
			} */
		})
	}
}
