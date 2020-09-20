/*
Inspire by https://github.com/go-gorm/gorm/blob/b4166d9515c3a86da2a1c7a695bc73d83861737d/association.go
And specially thanks for Lin Pin
*/

package orm

import (
	"context"
	"reflect"

	"github.com/Masterminds/squirrel"

	"github.com/cathatino/notification-service/pkg/sql"
	"github.com/cathatino/notification-service/pkg/utils/reflectutil"
)

type ORM interface {
	Create(ctx context.Context, model Model) error
	Update(ctx context.Context, model Model) error
	Find(ctx context.Context, models *[]Model, pred interface{}, args ...interface{}) error
}

type orm struct {
	sql.Connector
}

// Create db record using the model object
// only the PrimaryKey will be updated for the other fields remain as it is
func (o *orm) Create(ctx context.Context, model Model) (err error) {
	if !reflectutil.IsPtr(model) {
		err = ErrModelObjIsNotPtr
		return
	}

	sqlCmd, args, err := squirrel.Insert(model.GetTableName()).
		Columns(model.GetColumns()...).
		Values(model.GetValues()...).
		ToSql()
	if err != nil {
		return
	}

	db, err := o.GetDB()
	if err != nil {
		return
	}

	result, err := db.Exec(sqlCmd, args)
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	model.SetPrimaryKey(id)

	return err
}

// Update db record using the model object
// note: the syntax will be
//     Update xxx_tab set col_a = val_a, ..., where id = model.GetPrimaryKey();
func (o *orm) Update(ctx context.Context, model Model) error {
	if !reflectutil.IsPtr(model) {
		return ErrModelObjIsNotPtr
	}

	cols, vals := model.GetColumns(), model.GetValues()
	if len(cols) != len(vals) {
		return ErrInvalidLengthBetweenColsAndVals
	}

	setMapContent := squirrel.Eq{}
	for idx := 0; idx < len(cols); idx++ {
		setMapContent[cols[idx]] = vals[idx]
	}

	_, err := squirrel.Update(model.GetTableName()).SetMap(setMapContent).ExecContext(ctx)
	return err
}

// Find db record using model object
func (o *orm) Find(ctx context.Context, models *[]Model, pred interface{}, args ...interface{}) error {
	if len(*models) != 0 {
		return ErrNonZeroSliceLength
	}

	modelElem := reflect.ValueOf(models).Elem().Type().Elem()
	tableName := reflect.New(modelElem).Interface().(Model)

	sqlCmd, args, err := squirrel.Select("*").From(tableName.GetTableName()).Where(pred, args).ToSql()
	if err != nil {
		return err
	}

	db, err := o.GetDB()
	if err != nil {
		return err
	}

	rows, err := db.QueryxContext(ctx, sqlCmd, args)
	if err != nil {
		return err
	}
	for rows.Next() {
		model := reflect.New(modelElem).Interface().(Model)
		err := rows.StructScan(model)
		if err != nil {
			return err
		}
		*models = append(*models, model)
	}
	return nil
}
