package repository

import (
	"context"

	"github.com/ezio1119/fishapp-api-gateway/domain/user_grpc"
)

type UserRepository struct {
	Client user_grpc.UserServiceClient
}

func (r *UserRepository) GetByID(ctx context.Context, id *user_grpc.ID) (*user_grpc.User, error) {
	return r.Client.GetByID(ctx, id)
}

func (r *UserRepository) Create(ctx context.Context, req *user_grpc.CreateReq) (*user_grpc.UserWithToken, error) {
	return r.Client.Create(ctx, req)
}

func (r *UserRepository) Update(ctx context.Context, req *user_grpc.UpdateReq) (*user_grpc.User, error) {
	return r.Client.Update(ctx, req)
}

func (r *UserRepository) Delete(ctx context.Context, id *user_grpc.ID) (*user_grpc.DeleteRes, error) {
	return r.Client.Delete(ctx, id)
}

func (r *UserRepository) Login(ctx context.Context, req *user_grpc.LoginReq) (*user_grpc.UserWithToken, error) {
	return r.Client.Login(ctx, req)
}
