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
func GetAllByPagination(c context.Context, offset, limit int, conds ...interface{}) (todos []Todo, err error) {
	query := database.Db.Limit(limit).Offset(offset)

	if result := query.Find(&todos, conds...); result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func GetAll(c context.Context, conds ...interface{}) (todos []Todo, err error) {
	if result := database.Db.Find(&todos, conds...); result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func First(c context.Context, conds ...interface{}) (todo Todo, err error) {
	if result := database.Db.First(&todo, conds...); result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

func Last(c context.Context, conds ...interface{}) (todo Todo, err error) {
	if result := database.Db.Last(&todo, conds...); result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

func Take(c context.Context, conds ...interface{}) (todo Todo, err error) {
	if result := database.Db.Take(&todo, conds...); result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

func Delete(c context.Context, value interface{}, conds ...interface{}) error {
	if result := database.Db.Delete(value, conds...); result.Error != nil {
		return result.Error
	}
	return nil
}
