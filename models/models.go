package models

import (
	"context"

	"github.com/robertantonyjaikumar/hangover-common/database"
)

// CreateTodo adds a new Todo to the database
func Create(c context.Context, item interface{}) (err error) {
	if result := database.Db.Create(item); result.Error != nil {
		return result.Error
	}
	return
}

func Update(c context.Context, item interface{}, updates interface{}, conds ...interface{}) error {

	query := database.Db.Model(&item).Where(conds)
	if result := query.Updates(updates); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetTodosByPagination retrieves Todos with pagination and optional filters
func GetAllByPagination(c context.Context, item interface{}, offset, limit int, conds ...interface{}) (err error) {
	query := database.Db.Limit(limit).Offset(offset)

	if result := query.Find(item, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAll(c context.Context, item interface{}, conds ...interface{}) (err error) {
	if result := database.Db.Find(item, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}

func First(c context.Context, item interface{}, conds ...interface{}) (err error) {
	if result := database.Db.First(item, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}

func Last(c context.Context, item interface{}, conds ...interface{}) (err error) {
	if result := database.Db.Last(item, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}

func Take(c context.Context, item interface{}, conds ...interface{}) (err error) {
	if result := database.Db.Take(item, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}

func Delete(c context.Context, value interface{}, conds ...interface{}) error {
	if result := database.Db.Delete(value, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}
