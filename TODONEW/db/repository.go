package db

import (
	"context"

	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/arijitnayak92/taskAfford/TODONEW/utils"
)

const keyRepository = "Repository"

type Repository interface {
	Insert(todo *schema.Todo) (int, *utils.APIError)
	Delete(id int) *utils.APIError
	Update(id int, todo *schema.Todo) *utils.APIError
	MarkAsDone(id int, status bool) *utils.APIError
	GetAll() ([]schema.Todo, *utils.APIError)
	GetOne(id int) (schema.Todo, *utils.APIError)
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

func Insert(ctx context.Context, todo *schema.Todo) (int, *utils.APIError) {
	return getRepository(ctx).Insert(todo)
}

func Delete(ctx context.Context, id int) *utils.APIError {
	return getRepository(ctx).Delete(id)
}

func Update(ctx context.Context, id int, todo *schema.Todo) *utils.APIError {
	return getRepository(ctx).Update(id, todo)
}

func MarkAsDone(ctx context.Context, id int, status bool) *utils.APIError {
	return getRepository(ctx).MarkAsDone(id, status)
}

func GetAll(ctx context.Context) ([]schema.Todo, *utils.APIError) {
	return getRepository(ctx).GetAll()
}

func GetOne(ctx context.Context, id int) (schema.Todo, *utils.APIError) {
	return getRepository(ctx).GetOne(id)
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
}
