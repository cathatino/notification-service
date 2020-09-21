/*
Inspire by https://github.com/go-gorm/gorm/blob/b4166d9515c3a86da2a1c7a695bc73d83861737d/association.go
And specially thanks for Lin Pin
*/

package orm

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Masterminds/squirrel"

	"github.com/cathatino/notification-service/pkg/sql/connector"
	"github.com/cathatino/notification-service/pkg/utils/reflectutil"
)

type ORM interface {
	Create(ctx context.Context, model Model) error
	Update(ctx context.Context, model Model) error
	Find(ctx context.Context, models interface{}, pred interface{}, args ...interface{}) error
}

type orm struct {
	connector.Connector
}

func NewOrm(con connector.Connector) ORM {
	return &orm{con}
}

// Create db record using the model object
// only the PrimaryKey will be updated for the other fields remain as it is
func (o *orm) Create(ctx context.Context, model Model) (err error) {
	if !reflectutil.IsPtr(model) {
		err = ErrModelObjIsNotPtr
		return
	}

	withPrimaryKey := false
	primaryKeyName, primaryKeyValue := model.GetPrimaryKey()
	sqlCmd, sqlArgs, err := squirrel.Insert(model.GetTableName()).
		Columns(model.GetColumns(withPrimaryKey)...).
		Values(model.GetValues(withPrimaryKey)...).
		Suffix(fmt.Sprintf("RETURNING %s", primaryKeyName)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return
	}

	db, err := o.GetDB()
	if err != nil {
		return
	}

	err = db.QueryRowxContext(ctx, sqlCmd, sqlArgs...).Scan(&primaryKeyValue)
	if err != nil {
		return err
	}

	model.SetPrimaryKey(primaryKeyValue)

	return err
}

// Update db record using the model object
// note: the syntax will be
//     Update xxx_tab set col_a = val_a, ..., where id = model.GetPrimaryKey();
func (o *orm) Update(ctx context.Context, model Model) error {
	if !reflectutil.IsPtr(model) {
		return ErrModelObjIsNotPtr
	}

	withPrimaryKey := false
	cols, vals := model.GetColumns(withPrimaryKey), model.GetValues(withPrimaryKey)
	if len(cols) != len(vals) {
		return ErrInvalidLengthBetweenColsAndVals
	}

	setMapContent := squirrel.Eq{}
	for idx := 0; idx < len(cols); idx++ {
		setMapContent[cols[idx]] = vals[idx]
	}

	primaryKeyCol, primaryKeyVal := model.GetPrimaryKey()
	sqlCmd, args, err := squirrel.Update(model.GetTableName()).
		SetMap(setMapContent).
		Where(squirrel.Eq{
			primaryKeyCol: primaryKeyVal,
		}).ToSql()
	if err != nil {
		return err
	}

	db, err := o.GetDB()
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlCmd, args)
	if err != nil {
		return err
	}

	return err
}

// Find db record using model object
func (o *orm) Find(ctx context.Context, modelsPtr interface{}, pred interface{}, args ...interface{}) error {
	if !reflectutil.IsPtr(modelsPtr) {
		return ErrModelObjIsNotPtr
	}
	if reflect.TypeOf(modelsPtr).Elem().Kind() != reflect.Slice {
		return ErrPtrIsNotSlice
	}

	models := reflect.ValueOf(modelsPtr).Elem()
	elemType := models.Type().Elem()

	if reflect.TypeOf(reflect.New(elemType)).Kind() == reflect.Ptr {
		return ErrModelObjIsPtr
	}

	withPrimaryKey := true
	columns := reflect.New(elemType).Interface().(Model).GetColumns(withPrimaryKey)
	tableName := reflect.New(elemType).Interface().(Model).GetTableName()

	sqlCmd, sqlArgs, err := squirrel.Select(columns...).From(tableName).Where(pred, args).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return err
	}

	db, err := o.GetDB()
	if err != nil {
		return err
	}
	rows, err := db.QueryxContext(ctx, sqlCmd, sqlArgs...)
	if err != nil {
		return err
	}
	for rows.Next() {
		model := reflect.New(elemType)
		if err := rows.StructScan(model.Interface()); err != nil {
			return err
		}
		models.Set(reflect.Append(models, model.Elem()))
	}
	return nil
}
