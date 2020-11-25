package service

import (
	"context"

	"github.com/arijitnayak92/taskAfford/TODONEW/db"
	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/arijitnayak92/taskAfford/TODONEW/utils"
)

func Insert(ctx context.Context, todo *schema.Todo) (int, *utils.APIError) {
	return db.Insert(ctx, todo)
}

func Delete(ctx context.Context, id int) *utils.APIError {
	return db.Delete(ctx, id)
}

func Update(ctx context.Context, id int, todo *schema.Todo) *utils.APIError {
	return db.Update(ctx, id, todo)
}

func MarkAsDone(ctx context.Context, id int, status bool) *utils.APIError {
	return db.MarkAsDone(ctx, id, status)
}

func GetAll(ctx context.Context) ([]schema.Todo, *utils.APIError) {
	return db.GetAll(ctx)
}

func GetOne(ctx context.Context, id int) (schema.Todo, *utils.APIError) {
	return db.GetOne(ctx, id)
}
