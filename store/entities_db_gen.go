// Code generated by codegen. DO NOT EDIT.
// Generated by sql/dmlgen. DO NOT EDIT.
// +build csall db

package store

import (
	"context"
	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/sql/ddl"
	"github.com/corestoreio/pkg/sql/dml"
)

const (
	TableNameStore        = "store"
	TableNameStoreGroup   = "store_group"
	TableNameStoreWebsite = "store_website"
)

// NewTables returns a goified version of the MySQL/MariaDB table schema for the
// tables:  store, store_group, store_website Auto generated by dmlgen.
func NewTables(ctx context.Context, opts ...ddl.TableOption) (tm *ddl.Tables, err error) {
	if tm, err = ddl.NewTables(
		append(opts, ddl.WithCreateTable(ctx, TableNameStore, "", TableNameStoreGroup, "", TableNameStoreWebsite, ""))...,
	); err != nil {
		return nil, errors.WithStack(err)
	}
	return tm, nil
}

// AssignLastInsertID updates the increment ID field with the last inserted ID
// from an INSERT operation. Implements dml.InsertIDAssigner. Auto generated.
func (e *Store) AssignLastInsertID(id int64) {
	e.StoreID = uint32(id)
}

// MapColumns implements interface ColumnMapper only partially. Auto generated.
func (e *Store) MapColumns(cm *dml.ColumnMap) error {
	if cm.Mode() == dml.ColumnMapEntityReadAll {
		return cm.Uint32(&e.StoreID).NullString(&e.Code).Uint32(&e.WebsiteID).Uint32(&e.GroupID).String(&e.Name).Uint32(&e.SortOrder).Bool(&e.IsActive).Err()
	}
	for cm.Next() {
		switch c := cm.Column(); c {
		case "store_id":
			cm.Uint32(&e.StoreID)
		case "code":
			cm.NullString(&e.Code)
		case "website_id":
			cm.Uint32(&e.WebsiteID)
		case "group_id":
			cm.Uint32(&e.GroupID)
		case "name":
			cm.String(&e.Name)
		case "sort_order":
			cm.Uint32(&e.SortOrder)
		case "is_active":
			cm.Bool(&e.IsActive)
		default:
			return errors.NotFound.Newf("[store] Store Column %q not found", c)
		}
	}
	return errors.WithStack(cm.Err())
}

// AssignLastInsertID traverses through the slice and sets a decrementing new ID
// to each entity.
func (cc *StoreCollection) AssignLastInsertID(id int64) {
	var j int64
	for i := len(cc.Data) - 1; i >= 0; i-- {
		cc.Data[i].AssignLastInsertID(id - j)
		j++
	}
}
func (cc *StoreCollection) scanColumns(cm *dml.ColumnMap, e *Store, idx uint64) error {
	if err := e.MapColumns(cm); err != nil {
		return errors.WithStack(err)
	}
	// this function might get extended.
	return nil
}

// MapColumns implements dml.ColumnMapper interface. Auto generated.
func (cc *StoreCollection) MapColumns(cm *dml.ColumnMap) error {
	switch m := cm.Mode(); m {
	case dml.ColumnMapEntityReadAll, dml.ColumnMapEntityReadSet:
		for i, e := range cc.Data {
			if err := cc.scanColumns(cm, e, uint64(i)); err != nil {
				return errors.WithStack(err)
			}
		}
	case dml.ColumnMapScan:
		if cm.Count == 0 {
			cc.Data = cc.Data[:0]
		}
		e := new(Store)
		if err := cc.scanColumns(cm, e, cm.Count); err != nil {
			return errors.WithStack(err)
		}
		cc.Data = append(cc.Data, e)
	case dml.ColumnMapCollectionReadSet:
		for cm.Next() {
			switch c := cm.Column(); c {
			case "store_id":
				cm = cm.Uint32s(cc.StoreIDs()...)
			case "code":
				cm = cm.NullStrings(cc.Codes()...)
			default:
				return errors.NotFound.Newf("[store] StoreCollection Column %q not found", c)
			}
		} // end for cm.Next

	default:
		return errors.NotSupported.Newf("[store] Unknown Mode: %q", string(m))
	}
	return cm.Err()
}

// AssignLastInsertID updates the increment ID field with the last inserted ID
// from an INSERT operation. Implements dml.InsertIDAssigner. Auto generated.
func (e *StoreGroup) AssignLastInsertID(id int64) {
	e.GroupID = uint32(id)
}

// MapColumns implements interface ColumnMapper only partially. Auto generated.
func (e *StoreGroup) MapColumns(cm *dml.ColumnMap) error {
	if cm.Mode() == dml.ColumnMapEntityReadAll {
		return cm.Uint32(&e.GroupID).Uint32(&e.WebsiteID).String(&e.Name).Uint32(&e.RootCategoryID).Uint32(&e.DefaultStoreID).NullString(&e.Code).Err()
	}
	for cm.Next() {
		switch c := cm.Column(); c {
		case "group_id":
			cm.Uint32(&e.GroupID)
		case "website_id":
			cm.Uint32(&e.WebsiteID)
		case "name":
			cm.String(&e.Name)
		case "root_category_id":
			cm.Uint32(&e.RootCategoryID)
		case "default_store_id":
			cm.Uint32(&e.DefaultStoreID)
		case "code":
			cm.NullString(&e.Code)
		default:
			return errors.NotFound.Newf("[store] StoreGroup Column %q not found", c)
		}
	}
	return errors.WithStack(cm.Err())
}

// AssignLastInsertID traverses through the slice and sets a decrementing new ID
// to each entity.
func (cc *StoreGroupCollection) AssignLastInsertID(id int64) {
	var j int64
	for i := len(cc.Data) - 1; i >= 0; i-- {
		cc.Data[i].AssignLastInsertID(id - j)
		j++
	}
}
func (cc *StoreGroupCollection) scanColumns(cm *dml.ColumnMap, e *StoreGroup, idx uint64) error {
	if err := e.MapColumns(cm); err != nil {
		return errors.WithStack(err)
	}
	// this function might get extended.
	return nil
}

// MapColumns implements dml.ColumnMapper interface. Auto generated.
func (cc *StoreGroupCollection) MapColumns(cm *dml.ColumnMap) error {
	switch m := cm.Mode(); m {
	case dml.ColumnMapEntityReadAll, dml.ColumnMapEntityReadSet:
		for i, e := range cc.Data {
			if err := cc.scanColumns(cm, e, uint64(i)); err != nil {
				return errors.WithStack(err)
			}
		}
	case dml.ColumnMapScan:
		if cm.Count == 0 {
			cc.Data = cc.Data[:0]
		}
		e := new(StoreGroup)
		if err := cc.scanColumns(cm, e, cm.Count); err != nil {
			return errors.WithStack(err)
		}
		cc.Data = append(cc.Data, e)
	case dml.ColumnMapCollectionReadSet:
		for cm.Next() {
			switch c := cm.Column(); c {
			case "group_id":
				cm = cm.Uint32s(cc.GroupIDs()...)
			case "code":
				cm = cm.NullStrings(cc.Codes()...)
			default:
				return errors.NotFound.Newf("[store] StoreGroupCollection Column %q not found", c)
			}
		} // end for cm.Next

	default:
		return errors.NotSupported.Newf("[store] Unknown Mode: %q", string(m))
	}
	return cm.Err()
}

// AssignLastInsertID updates the increment ID field with the last inserted ID
// from an INSERT operation. Implements dml.InsertIDAssigner. Auto generated.
func (e *StoreWebsite) AssignLastInsertID(id int64) {
	e.WebsiteID = uint32(id)
}

// MapColumns implements interface ColumnMapper only partially. Auto generated.
func (e *StoreWebsite) MapColumns(cm *dml.ColumnMap) error {
	if cm.Mode() == dml.ColumnMapEntityReadAll {
		return cm.Uint32(&e.WebsiteID).NullString(&e.Code).NullString(&e.Name).Uint32(&e.SortOrder).Uint32(&e.DefaultGroupID).Bool(&e.IsDefault).Err()
	}
	for cm.Next() {
		switch c := cm.Column(); c {
		case "website_id":
			cm.Uint32(&e.WebsiteID)
		case "code":
			cm.NullString(&e.Code)
		case "name":
			cm.NullString(&e.Name)
		case "sort_order":
			cm.Uint32(&e.SortOrder)
		case "default_group_id":
			cm.Uint32(&e.DefaultGroupID)
		case "is_default":
			cm.Bool(&e.IsDefault)
		default:
			return errors.NotFound.Newf("[store] StoreWebsite Column %q not found", c)
		}
	}
	return errors.WithStack(cm.Err())
}

// AssignLastInsertID traverses through the slice and sets a decrementing new ID
// to each entity.
func (cc *StoreWebsiteCollection) AssignLastInsertID(id int64) {
	var j int64
	for i := len(cc.Data) - 1; i >= 0; i-- {
		cc.Data[i].AssignLastInsertID(id - j)
		j++
	}
}
func (cc *StoreWebsiteCollection) scanColumns(cm *dml.ColumnMap, e *StoreWebsite, idx uint64) error {
	if err := e.MapColumns(cm); err != nil {
		return errors.WithStack(err)
	}
	// this function might get extended.
	return nil
}

// MapColumns implements dml.ColumnMapper interface. Auto generated.
func (cc *StoreWebsiteCollection) MapColumns(cm *dml.ColumnMap) error {
	switch m := cm.Mode(); m {
	case dml.ColumnMapEntityReadAll, dml.ColumnMapEntityReadSet:
		for i, e := range cc.Data {
			if err := cc.scanColumns(cm, e, uint64(i)); err != nil {
				return errors.WithStack(err)
			}
		}
	case dml.ColumnMapScan:
		if cm.Count == 0 {
			cc.Data = cc.Data[:0]
		}
		e := new(StoreWebsite)
		if err := cc.scanColumns(cm, e, cm.Count); err != nil {
			return errors.WithStack(err)
		}
		cc.Data = append(cc.Data, e)
	case dml.ColumnMapCollectionReadSet:
		for cm.Next() {
			switch c := cm.Column(); c {
			case "website_id":
				cm = cm.Uint32s(cc.WebsiteIDs()...)
			case "code":
				cm = cm.NullStrings(cc.Codes()...)
			default:
				return errors.NotFound.Newf("[store] StoreWebsiteCollection Column %q not found", c)
			}
		} // end for cm.Next

	default:
		return errors.NotSupported.Newf("[store] Unknown Mode: %q", string(m))
	}
	return cm.Err()
}
