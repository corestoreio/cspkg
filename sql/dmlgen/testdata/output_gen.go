// Auto generated via github.com/corestoreio/pkg/sql/dmlgen

package testdata

import (
	"context"
	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/sql/ddl"
	"github.com/corestoreio/pkg/sql/dml"
	"github.com/corestoreio/pkg/storage/null"
	"sort"
	"time"
)

const (
	TableNameCoreConfigData = "core_config_data"
	TableNameCustomerEntity = "customer_entity"
	TableNameDmlgenTypes    = "dmlgen_types"
)

// NewTables returns a goified version of the MySQL/MariaDB table schema for the
// tables: [core_config_data customer_entity dmlgen_types]
// Auto generated by dmlgen.
func NewTables(ctx context.Context, db dml.QueryExecPreparer, opts ...ddl.TableOption) (tm *ddl.Tables, err error) {
	if tm, err = ddl.NewTables(
		ddl.WithCreateTable(ctx, db,
			TableNameCoreConfigData, "",
			TableNameCustomerEntity, "",
			TableNameDmlgenTypes, ""),
		ddl.WithDB(db),
	); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := tm.Options(opts...); err != nil {
		return nil, errors.WithStack(err)
	}
	return tm, nil
}

// CoreConfigData represents a single row for DB table `core_config_data`.
// Auto generated.
//easyjson:json
type CoreConfigData struct {
	ConfigID  uint32      `json:"config_id,omitempty" max_len:"10"`  // config_id int(10) unsigned NOT NULL PRI  auto_increment "Id"
	Scope     string      `json:"scope,omitempty" max_len:"8"`       // scope varchar(8) NOT NULL MUL DEFAULT ''default''  "Scope"
	ScopeID   int32       `json:"scope_id" xml:"scope_id"`           // scope_id int(11) NOT NULL  DEFAULT '0'  "Scope Id"
	Expires   null.Time   `json:"expires,omitempty" `                // expires datetime NULL  DEFAULT 'NULL'  "Value expiration time"
	Path      string      `json:"x_path" xml:"y_path" max_len:"255"` // path varchar(255) NOT NULL  DEFAULT ''general''  "Config Path overwritten"
	Value     null.String `json:"value,omitempty" max_len:"65535"`   // value text NULL  DEFAULT 'NULL'  "Value"
	VersionTs time.Time   `json:"version_ts,omitempty" `             // version_ts timestamp(6) NOT NULL    "Timestamp Start Versioning"
	VersionTe time.Time   `json:"version_te,omitempty" `             // version_te timestamp(6) NOT NULL PRI   "Timestamp End Versioning"
}

// AssignLastInsertID updates the increment ID field with the last inserted ID
// from an INSERT operation. Implements dml.InsertIDAssigner. Auto generated.
func (e *CoreConfigData) AssignLastInsertID(id int64) {
	e.ConfigID = uint32(id)
}

// MapColumns implements interface ColumnMapper only partially. Auto generated.
func (e *CoreConfigData) MapColumns(cm *dml.ColumnMap) error {
	if cm.Mode() == dml.ColumnMapEntityReadAll {
		return cm.Uint32(&e.ConfigID).String(&e.Scope).Int32(&e.ScopeID).NullTime(&e.Expires).String(&e.Path).NullString(&e.Value).Time(&e.VersionTs).Time(&e.VersionTe).Err()
	}
	for cm.Next() {
		switch c := cm.Column(); c {
		case "config_id":
			cm.Uint32(&e.ConfigID)
		case "scope":
			cm.String(&e.Scope)
		case "scope_id":
			cm.Int32(&e.ScopeID)
		case "expires":
			cm.NullTime(&e.Expires)
		case "path", "storage_location", "config_directory":
			cm.String(&e.Path)
		case "value":
			cm.NullString(&e.Value)
		case "version_ts":
			cm.Time(&e.VersionTs)
		case "version_te":
			cm.Time(&e.VersionTe)
		default:
			return errors.NotFound.Newf("[testdata] CoreConfigData Column %q not found", c)
		}
	}
	return errors.WithStack(cm.Err())
}

// Empty empties all the fields of the current object. Also known as Reset.
func (e *CoreConfigData) Empty() *CoreConfigData { *e = CoreConfigData{}; return e }

// CoreConfigDataCollection represents a collection type for DB table core_config_data
// Not thread safe. Auto generated.
//easyjson:json
type CoreConfigDataCollection struct {
	Data             []*CoreConfigData                   `json:"data,omitempty"`
	BeforeMapColumns func(uint64, *CoreConfigData) error `json:"-"`
	AfterMapColumns  func(uint64, *CoreConfigData) error `json:"-"`
}

// NewCoreConfigDataCollection creates a new initialized collection. Auto generated.
func NewCoreConfigDataCollection() *CoreConfigDataCollection {
	return &CoreConfigDataCollection{
		Data: make([]*CoreConfigData, 0, 5),
	}
}

func (cc *CoreConfigDataCollection) scanColumns(cm *dml.ColumnMap, e *CoreConfigData, idx uint64) error {
	if err := cc.BeforeMapColumns(idx, e); err != nil {
		return errors.WithStack(err)
	}
	if err := e.MapColumns(cm); err != nil {
		return errors.WithStack(err)
	}
	if err := cc.AfterMapColumns(idx, e); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// MapColumns implements dml.ColumnMapper interface. Auto generated.
func (cc *CoreConfigDataCollection) MapColumns(cm *dml.ColumnMap) error {
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
		e := new(CoreConfigData)
		if err := cc.scanColumns(cm, e, cm.Count); err != nil {
			return errors.WithStack(err)
		}
		cc.Data = append(cc.Data, e)
	case dml.ColumnMapCollectionReadSet:
		for cm.Next() {
			switch c := cm.Column(); c {
			case "config_id":
				cm = cm.Uint32s(cc.ConfigIDs()...)

			default:
				return errors.NotFound.Newf("[testdata] CoreConfigDataCollection Column %q not found", c)
			}
		}
	default:
		return errors.NotSupported.Newf("[testdata] Unknown Mode: %q", string(m))
	}
	return cm.Err()
}

// ConfigIDs returns a slice or appends to a slice all values.
// Auto generated.
func (cc *CoreConfigDataCollection) ConfigIDs(ret ...uint32) []uint32 {
	if ret == nil {
		ret = make([]uint32, 0, len(cc.Data))
	}
	for _, e := range cc.Data {
		ret = append(ret, e.ConfigID)
	}
	return ret
}

// Paths belongs to the column `path`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *CoreConfigDataCollection) UniquePaths(ret ...string) []string {
	if ret == nil {
		ret = make([]string, 0, len(cc.Data))
	}

	dupCheck := make(map[string]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.Path] {
			ret = append(ret, e.Path)
			dupCheck[e.Path] = true
		}
	}
	return ret
}

// FilterThis filters the current slice by predicate f without memory allocation.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Filter(f func(*CoreConfigData) bool) *CoreConfigDataCollection {
	b := cc.Data[:0]
	for _, e := range cc.Data {
		if f(e) {
			b = append(b, e)
		}
	}
	cc.Data = b
	return cc
}

// Each will run function f on all items in []*CoreConfigData.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Each(f func(*CoreConfigData)) *CoreConfigDataCollection {
	for i := range cc.Data {
		f(cc.Data[i])
	}
	return cc
}

func (cc *CoreConfigDataCollection) SortByConfigID() {
	sort.Slice(cc.Data, func(i, j int) bool {
		return cc.Data[i].ConfigID < cc.Data[j].ConfigID
	})
}

// Cut will remove items i through j-1.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Cut(i, j int) *CoreConfigDataCollection {
	z := cc.Data // copy slice header
	copy(z[i:], z[j:])
	for k, n := len(z)-j+i, len(z); k < n; k++ {
		z[k] = nil // this should avoid the memory leak
	}
	z = z[:len(z)-j+i]
	cc.Data = z
	return cc
}

// Swap will satisfy the sort.Interface.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Swap(i, j int) { cc.Data[i], cc.Data[j] = cc.Data[j], cc.Data[i] }

// Delete will remove an item from the slice.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Delete(i int) *CoreConfigDataCollection {
	z := cc.Data // copy the slice header
	end := len(z) - 1
	cc.Swap(i, end)
	copy(z[i:], z[i+1:])
	z[end] = nil // this should avoid the memory leak
	z = z[:end]
	cc.Data = z
	return cc
}

// Insert will place a new item at position i.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Insert(n *CoreConfigData, i int) *CoreConfigDataCollection {
	z := cc.Data // copy the slice header
	z = append(z, &CoreConfigData{})
	copy(z[i+1:], z[i:])
	z[i] = n
	cc.Data = z
	return cc
}

// Append will add a new item at the end of *CoreConfigDataCollection.
// Auto generated via dmlgen.
func (cc *CoreConfigDataCollection) Append(n ...*CoreConfigData) *CoreConfigDataCollection {
	cc.Data = append(cc.Data, n...)
	return cc
}

// CustomerEntity represents a single row for DB table `customer_entity`.
// Auto generated.
type CustomerEntity struct {
	EntityID               uint32      `max_len:"10"`  // entity_id int(10) unsigned NOT NULL PRI  auto_increment "Entity ID"
	WebsiteID              null.Uint32 `max_len:"5"`   // website_id smallint(5) unsigned NULL MUL DEFAULT 'NULL'  "Website ID"
	Email                  null.String `max_len:"255"` // email varchar(255) NULL MUL DEFAULT 'NULL'  "Email"
	GroupID                uint32      `max_len:"5"`   // group_id smallint(5) unsigned NOT NULL  DEFAULT '0'  "Group ID"
	IncrementID            null.String `max_len:"50"`  // increment_id varchar(50) NULL  DEFAULT 'NULL'  "Increment Id"
	StoreID                null.Uint32 `max_len:"5"`   // store_id smallint(5) unsigned NULL MUL DEFAULT '0'  "Store ID"
	CreatedAt              time.Time   // created_at timestamp NOT NULL  DEFAULT 'current_timestamp()'  "Created At"
	UpdatedAt              time.Time   // updated_at timestamp NOT NULL  DEFAULT 'current_timestamp()' on update current_timestamp() "Updated At"
	IsActive               bool        `max_len:"5"`   // is_active smallint(5) unsigned NOT NULL  DEFAULT '1'  "Is Active"
	DisableAutoGroupChange uint32      `max_len:"5"`   // disable_auto_group_change smallint(5) unsigned NOT NULL  DEFAULT '0'  "Disable automatic group change based on VAT ID"
	CreatedIn              null.String `max_len:"255"` // created_in varchar(255) NULL  DEFAULT 'NULL'  "Created From"
	Prefix                 null.String `max_len:"40"`  // prefix varchar(40) NULL  DEFAULT 'NULL'  "Name Prefix"
	Firstname              null.String `max_len:"255"` // firstname varchar(255) NULL MUL DEFAULT 'NULL'  "First Name"
	Middlename             null.String `max_len:"255"` // middlename varchar(255) NULL  DEFAULT 'NULL'  "Middle Name/Initial"
	Lastname               null.String `max_len:"255"` // lastname varchar(255) NULL MUL DEFAULT 'NULL'  "Last Name"
	Suffix                 null.String `max_len:"40"`  // suffix varchar(40) NULL  DEFAULT 'NULL'  "Name Suffix"
	Dob                    null.Time   // dob date NULL  DEFAULT 'NULL'  "Date of Birth"
	PasswordHash           null.String `max_len:"128"` // password_hash varchar(128) NULL  DEFAULT 'NULL'  "Password_hash"
	RpToken                null.String `max_len:"128"` // rp_token varchar(128) NULL  DEFAULT 'NULL'  "Reset password token"
	RpTokenCreatedAt       null.Time   // rp_token_created_at datetime NULL  DEFAULT 'NULL'  "Reset password token creation time"
	DefaultBilling         null.Uint32 `max_len:"10"` // default_billing int(10) unsigned NULL  DEFAULT 'NULL'  "Default Billing Address"
	DefaultShipping        null.Uint32 `max_len:"10"` // default_shipping int(10) unsigned NULL  DEFAULT 'NULL'  "Default Shipping Address"
	Taxvat                 null.String `max_len:"50"` // taxvat varchar(50) NULL  DEFAULT 'NULL'  "Tax/VAT Number"
	Confirmation           null.String `max_len:"64"` // confirmation varchar(64) NULL  DEFAULT 'NULL'  "Is Confirmed"
	Gender                 null.Uint32 `max_len:"5"`  // gender smallint(5) unsigned NULL  DEFAULT 'NULL'  "Gender"
	FailuresNum            null.Int32  `max_len:"5"`  // failures_num smallint(6) NULL  DEFAULT '0'  "Failure Number"
	FirstFailure           null.Time   // first_failure timestamp NULL  DEFAULT 'NULL'  "First Failure"
	LockExpires            null.Time   // lock_expires timestamp NULL  DEFAULT 'NULL'  "Lock Expiration Date"
}

// AssignLastInsertID updates the increment ID field with the last inserted ID
// from an INSERT operation. Implements dml.InsertIDAssigner. Auto generated.
func (e *CustomerEntity) AssignLastInsertID(id int64) {
	e.EntityID = uint32(id)
}

// MapColumns implements interface ColumnMapper only partially. Auto generated.
func (e *CustomerEntity) MapColumns(cm *dml.ColumnMap) error {
	if cm.Mode() == dml.ColumnMapEntityReadAll {
		return cm.Uint32(&e.EntityID).NullUint32(&e.WebsiteID).NullString(&e.Email).Uint32(&e.GroupID).NullString(&e.IncrementID).NullUint32(&e.StoreID).Time(&e.CreatedAt).Time(&e.UpdatedAt).Bool(&e.IsActive).Uint32(&e.DisableAutoGroupChange).NullString(&e.CreatedIn).NullString(&e.Prefix).NullString(&e.Firstname).NullString(&e.Middlename).NullString(&e.Lastname).NullString(&e.Suffix).NullTime(&e.Dob).NullString(&e.PasswordHash).NullString(&e.RpToken).NullTime(&e.RpTokenCreatedAt).NullUint32(&e.DefaultBilling).NullUint32(&e.DefaultShipping).NullString(&e.Taxvat).NullString(&e.Confirmation).NullUint32(&e.Gender).NullInt32(&e.FailuresNum).NullTime(&e.FirstFailure).NullTime(&e.LockExpires).Err()
	}
	for cm.Next() {
		switch c := cm.Column(); c {
		case "entity_id", "parent_id":
			cm.Uint32(&e.EntityID)
		case "website_id":
			cm.NullUint32(&e.WebsiteID)
		case "email":
			cm.NullString(&e.Email)
		case "group_id":
			cm.Uint32(&e.GroupID)
		case "increment_id":
			cm.NullString(&e.IncrementID)
		case "store_id":
			cm.NullUint32(&e.StoreID)
		case "created_at":
			cm.Time(&e.CreatedAt)
		case "updated_at":
			cm.Time(&e.UpdatedAt)
		case "is_active":
			cm.Bool(&e.IsActive)
		case "disable_auto_group_change":
			cm.Uint32(&e.DisableAutoGroupChange)
		case "created_in":
			cm.NullString(&e.CreatedIn)
		case "prefix":
			cm.NullString(&e.Prefix)
		case "firstname":
			cm.NullString(&e.Firstname)
		case "middlename":
			cm.NullString(&e.Middlename)
		case "lastname":
			cm.NullString(&e.Lastname)
		case "suffix":
			cm.NullString(&e.Suffix)
		case "dob":
			cm.NullTime(&e.Dob)
		case "password_hash":
			cm.NullString(&e.PasswordHash)
		case "rp_token":
			cm.NullString(&e.RpToken)
		case "rp_token_created_at":
			cm.NullTime(&e.RpTokenCreatedAt)
		case "default_billing":
			cm.NullUint32(&e.DefaultBilling)
		case "default_shipping":
			cm.NullUint32(&e.DefaultShipping)
		case "taxvat":
			cm.NullString(&e.Taxvat)
		case "confirmation":
			cm.NullString(&e.Confirmation)
		case "gender":
			cm.NullUint32(&e.Gender)
		case "failures_num":
			cm.NullInt32(&e.FailuresNum)
		case "first_failure":
			cm.NullTime(&e.FirstFailure)
		case "lock_expires":
			cm.NullTime(&e.LockExpires)
		default:
			return errors.NotFound.Newf("[testdata] CustomerEntity Column %q not found", c)
		}
	}
	return errors.WithStack(cm.Err())
}

// Empty empties all the fields of the current object. Also known as Reset.
func (e *CustomerEntity) Empty() *CustomerEntity { *e = CustomerEntity{}; return e }

// CustomerEntityCollection represents a collection type for DB table customer_entity
// Not thread safe. Auto generated.
type CustomerEntityCollection struct {
	Data             []*CustomerEntity                   `json:"data,omitempty"`
	BeforeMapColumns func(uint64, *CustomerEntity) error `json:"-"`
	AfterMapColumns  func(uint64, *CustomerEntity) error `json:"-"`
}

// NewCustomerEntityCollection creates a new initialized collection. Auto generated.
func NewCustomerEntityCollection() *CustomerEntityCollection {
	return &CustomerEntityCollection{
		Data: make([]*CustomerEntity, 0, 5),
	}
}

func (cc *CustomerEntityCollection) scanColumns(cm *dml.ColumnMap, e *CustomerEntity, idx uint64) error {
	if err := cc.BeforeMapColumns(idx, e); err != nil {
		return errors.WithStack(err)
	}
	if err := e.MapColumns(cm); err != nil {
		return errors.WithStack(err)
	}
	if err := cc.AfterMapColumns(idx, e); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// MapColumns implements dml.ColumnMapper interface. Auto generated.
func (cc *CustomerEntityCollection) MapColumns(cm *dml.ColumnMap) error {
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
		e := new(CustomerEntity)
		if err := cc.scanColumns(cm, e, cm.Count); err != nil {
			return errors.WithStack(err)
		}
		cc.Data = append(cc.Data, e)
	case dml.ColumnMapCollectionReadSet:
		for cm.Next() {
			switch c := cm.Column(); c {
			case "entity_id", "parent_id":
				cm = cm.Uint32s(cc.EntityIDs()...)

			default:
				return errors.NotFound.Newf("[testdata] CustomerEntityCollection Column %q not found", c)
			}
		}
	default:
		return errors.NotSupported.Newf("[testdata] Unknown Mode: %q", string(m))
	}
	return cm.Err()
}

// EntityIDs returns a slice or appends to a slice all values.
// Auto generated.
func (cc *CustomerEntityCollection) EntityIDs(ret ...uint32) []uint32 {
	if ret == nil {
		ret = make([]uint32, 0, len(cc.Data))
	}
	for _, e := range cc.Data {
		ret = append(ret, e.EntityID)
	}
	return ret
}

// FilterThis filters the current slice by predicate f without memory allocation.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Filter(f func(*CustomerEntity) bool) *CustomerEntityCollection {
	b := cc.Data[:0]
	for _, e := range cc.Data {
		if f(e) {
			b = append(b, e)
		}
	}
	cc.Data = b
	return cc
}

// Each will run function f on all items in []*CustomerEntity.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Each(f func(*CustomerEntity)) *CustomerEntityCollection {
	for i := range cc.Data {
		f(cc.Data[i])
	}
	return cc
}

func (cc *CustomerEntityCollection) SortByEntityID() {
	sort.Slice(cc.Data, func(i, j int) bool {
		return cc.Data[i].EntityID < cc.Data[j].EntityID
	})
}

// Cut will remove items i through j-1.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Cut(i, j int) *CustomerEntityCollection {
	z := cc.Data // copy slice header
	copy(z[i:], z[j:])
	for k, n := len(z)-j+i, len(z); k < n; k++ {
		z[k] = nil // this should avoid the memory leak
	}
	z = z[:len(z)-j+i]
	cc.Data = z
	return cc
}

// Swap will satisfy the sort.Interface.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Swap(i, j int) { cc.Data[i], cc.Data[j] = cc.Data[j], cc.Data[i] }

// Delete will remove an item from the slice.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Delete(i int) *CustomerEntityCollection {
	z := cc.Data // copy the slice header
	end := len(z) - 1
	cc.Swap(i, end)
	copy(z[i:], z[i+1:])
	z[end] = nil // this should avoid the memory leak
	z = z[:end]
	cc.Data = z
	return cc
}

// Insert will place a new item at position i.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Insert(n *CustomerEntity, i int) *CustomerEntityCollection {
	z := cc.Data // copy the slice header
	z = append(z, &CustomerEntity{})
	copy(z[i+1:], z[i:])
	z[i] = n
	cc.Data = z
	return cc
}

// Append will add a new item at the end of *CustomerEntityCollection.
// Auto generated via dmlgen.
func (cc *CustomerEntityCollection) Append(n ...*CustomerEntity) *CustomerEntityCollection {
	cc.Data = append(cc.Data, n...)
	return cc
}

// DmlgenTypes represents a single row for DB table `dmlgen_types`.
// Auto generated.
// Just another comment.
//easyjson:json
type DmlgenTypes struct {
	ID             int32        `json:"id,omitempty"  max_len:"10"`                     // id int(11) NOT NULL PRI  auto_increment ""
	ColBigint1     null.Int64   `json:"col_bigint_1,omitempty"  max_len:"19"`           // col_bigint_1 bigint(20) NULL  DEFAULT 'NULL'  ""
	ColBigint2     int64        `json:"col_bigint_2,omitempty"  max_len:"19"`           // col_bigint_2 bigint(20) NOT NULL  DEFAULT '0'  ""
	ColBigint3     null.Uint64  `json:"col_bigint_3,omitempty"  max_len:"20"`           // col_bigint_3 bigint(20) unsigned NULL  DEFAULT 'NULL'  ""
	ColBigint4     uint64       `json:"col_bigint_4,omitempty"  max_len:"20"`           // col_bigint_4 bigint(20) unsigned NOT NULL  DEFAULT '0'  ""
	ColBlob        []byte       `json:"col_blob,omitempty"  max_len:"65535"`            // col_blob blob NULL  DEFAULT 'NULL'  ""
	ColDate1       null.Time    `json:"col_date_1,omitempty"  `                         // col_date_1 date NULL  DEFAULT 'NULL'  ""
	ColDate2       time.Time    `json:"col_date_2,omitempty"  `                         // col_date_2 date NOT NULL  DEFAULT ''0000-00-00''  ""
	ColDatetime1   null.Time    `json:"col_datetime_1,omitempty"  `                     // col_datetime_1 datetime NULL  DEFAULT 'NULL'  ""
	ColDatetime2   time.Time    `json:"col_datetime_2,omitempty"  `                     // col_datetime_2 datetime NOT NULL  DEFAULT ''0000-00-00 00:00:00''  ""
	ColDecimal101  null.Decimal `json:"col_decimal_10_1,omitempty"  max_len:"10"`       // col_decimal_10_1 decimal(10,1) unsigned NULL  DEFAULT 'NULL'  ""
	ColDecimal124  null.Decimal `json:"col_decimal_12_4,omitempty"  max_len:"12"`       // col_decimal_12_4 decimal(12,4) NULL  DEFAULT 'NULL'  ""
	Price124a      null.Decimal `json:"price_12_4a,omitempty"  max_len:"12"`            // price_12_4a decimal(12,4) NULL  DEFAULT 'NULL'  ""
	Price124b      null.Decimal `json:"price_12_4b,omitempty"  max_len:"12"`            // price_12_4b decimal(12,4) NOT NULL  DEFAULT '0.0000'  ""
	ColDecimal123  null.Decimal `json:"col_decimal_12_3,omitempty"  max_len:"12"`       // col_decimal_12_3 decimal(12,3) NOT NULL  DEFAULT '0.000'  ""
	ColDecimal206  null.Decimal `json:"col_decimal_20_6,omitempty"  max_len:"20"`       // col_decimal_20_6 decimal(20,6) NOT NULL  DEFAULT '0.000000'  ""
	ColDecimal2412 null.Decimal `json:"col_decimal_24_12,omitempty"  max_len:"24"`      // col_decimal_24_12 decimal(24,12) NOT NULL  DEFAULT '0.000000000000'  ""
	ColInt1        null.Int32   `json:"col_int_1,omitempty"  max_len:"10"`              // col_int_1 int(10) NULL  DEFAULT 'NULL'  ""
	ColInt2        int32        `json:"col_int_2,omitempty"  max_len:"10"`              // col_int_2 int(10) NOT NULL  DEFAULT '0'  ""
	ColInt3        null.Uint32  `json:"col_int_3,omitempty"  max_len:"10"`              // col_int_3 int(10) unsigned NULL  DEFAULT 'NULL'  ""
	ColInt4        uint32       `json:"col_int_4,omitempty"  max_len:"10"`              // col_int_4 int(10) unsigned NOT NULL  DEFAULT '0'  ""
	ColLongtext1   null.String  `json:"col_longtext_1,omitempty"  max_len:"4294967295"` // col_longtext_1 longtext NULL  DEFAULT 'NULL'  ""
	ColLongtext2   string       `json:"col_longtext_2,omitempty"  max_len:"4294967295"` // col_longtext_2 longtext NOT NULL  DEFAULT ''''  ""
	ColMediumblob  []byte       `json:"col_mediumblob,omitempty"  max_len:"16777215"`   // col_mediumblob mediumblob NULL  DEFAULT 'NULL'  ""
	ColMediumtext1 null.String  `json:"col_mediumtext_1,omitempty"  max_len:"16777215"` // col_mediumtext_1 mediumtext NULL  DEFAULT 'NULL'  ""
	ColMediumtext2 string       `json:"col_mediumtext_2,omitempty"  max_len:"16777215"` // col_mediumtext_2 mediumtext NOT NULL  DEFAULT ''''  ""
	ColSmallint1   null.Int32   `json:"col_smallint_1,omitempty"  max_len:"5"`          // col_smallint_1 smallint(5) NULL  DEFAULT 'NULL'  ""
	ColSmallint2   int32        `json:"col_smallint_2,omitempty"  max_len:"5"`          // col_smallint_2 smallint(5) NOT NULL  DEFAULT '0'  ""
	ColSmallint3   null.Uint32  `json:"col_smallint_3,omitempty"  max_len:"5"`          // col_smallint_3 smallint(5) unsigned NULL  DEFAULT 'NULL'  ""
	ColSmallint4   uint32       `json:"col_smallint_4,omitempty"  max_len:"5"`          // col_smallint_4 smallint(5) unsigned NOT NULL  DEFAULT '0'  ""
	HasSmallint5   bool         `json:"has_smallint_5,omitempty"  max_len:"5"`          // has_smallint_5 smallint(5) unsigned NOT NULL  DEFAULT '0'  ""
	IsSmallint5    null.Bool    `json:"is_smallint_5,omitempty"  max_len:"5"`           // is_smallint_5 smallint(5) NULL  DEFAULT 'NULL'  ""
	ColText        null.String  `json:"col_text,omitempty"  max_len:"65535"`            // col_text text NULL  DEFAULT 'NULL'  ""
	ColTimestamp1  time.Time    `json:"col_timestamp_1,omitempty"  `                    // col_timestamp_1 timestamp NOT NULL  DEFAULT 'current_timestamp()'  ""
	ColTimestamp2  null.Time    `json:"col_timestamp_2,omitempty"  `                    // col_timestamp_2 timestamp NULL  DEFAULT 'NULL'  ""
	ColTinyint1    int32        `json:"col_tinyint_1,omitempty"  max_len:"3"`           // col_tinyint_1 tinyint(1) NOT NULL  DEFAULT '0'  ""
	ColVarchar1    string       `json:"col_varchar_1,omitempty"  max_len:"1"`           // col_varchar_1 varchar(1) NOT NULL  DEFAULT ''0''  ""
	ColVarchar100  null.String  `json:"col_varchar_100,omitempty"  max_len:"100"`       // col_varchar_100 varchar(100) NULL  DEFAULT 'NULL'  ""
	ColVarchar16   string       `json:"col_varchar_16,omitempty"  max_len:"16"`         // col_varchar_16 varchar(16) NOT NULL  DEFAULT ''de_DE''  ""
	ColChar1       null.String  `json:"col_char_1,omitempty"  max_len:"21"`             // col_char_1 char(21) NULL  DEFAULT 'NULL'  ""
	ColChar2       string       `json:"col_char_2,omitempty"  max_len:"17"`             // col_char_2 char(17) NOT NULL  DEFAULT ''xchar''  ""
}

// AssignLastInsertID updates the increment ID field with the last inserted ID
// from an INSERT operation. Implements dml.InsertIDAssigner. Auto generated.
func (e *DmlgenTypes) AssignLastInsertID(id int64) {
	e.ID = int32(id)
}

// MapColumns implements interface ColumnMapper only partially. Auto generated.
func (e *DmlgenTypes) MapColumns(cm *dml.ColumnMap) error {
	if cm.Mode() == dml.ColumnMapEntityReadAll {
		return cm.Int32(&e.ID).NullInt64(&e.ColBigint1).Int64(&e.ColBigint2).NullUint64(&e.ColBigint3).Uint64(&e.ColBigint4).Byte(&e.ColBlob).NullTime(&e.ColDate1).Time(&e.ColDate2).NullTime(&e.ColDatetime1).Time(&e.ColDatetime2).Decimal(&e.ColDecimal101).Decimal(&e.ColDecimal124).Decimal(&e.Price124a).Decimal(&e.Price124b).Decimal(&e.ColDecimal123).Decimal(&e.ColDecimal206).Decimal(&e.ColDecimal2412).NullInt32(&e.ColInt1).Int32(&e.ColInt2).NullUint32(&e.ColInt3).Uint32(&e.ColInt4).NullString(&e.ColLongtext1).String(&e.ColLongtext2).Byte(&e.ColMediumblob).NullString(&e.ColMediumtext1).String(&e.ColMediumtext2).NullInt32(&e.ColSmallint1).Int32(&e.ColSmallint2).NullUint32(&e.ColSmallint3).Uint32(&e.ColSmallint4).Bool(&e.HasSmallint5).NullBool(&e.IsSmallint5).NullString(&e.ColText).Time(&e.ColTimestamp1).NullTime(&e.ColTimestamp2).Int32(&e.ColTinyint1).String(&e.ColVarchar1).NullString(&e.ColVarchar100).String(&e.ColVarchar16).NullString(&e.ColChar1).String(&e.ColChar2).Err()
	}
	for cm.Next() {
		switch c := cm.Column(); c {
		case "id":
			cm.Int32(&e.ID)
		case "col_bigint_1":
			cm.NullInt64(&e.ColBigint1)
		case "col_bigint_2":
			cm.Int64(&e.ColBigint2)
		case "col_bigint_3":
			cm.NullUint64(&e.ColBigint3)
		case "col_bigint_4":
			cm.Uint64(&e.ColBigint4)
		case "col_blob":
			cm.Byte(&e.ColBlob)
		case "col_date_1":
			cm.NullTime(&e.ColDate1)
		case "col_date_2":
			cm.Time(&e.ColDate2)
		case "col_datetime_1":
			cm.NullTime(&e.ColDatetime1)
		case "col_datetime_2":
			cm.Time(&e.ColDatetime2)
		case "col_decimal_10_1":
			cm.Decimal(&e.ColDecimal101)
		case "col_decimal_12_4":
			cm.Decimal(&e.ColDecimal124)
		case "price_12_4a":
			cm.Decimal(&e.Price124a)
		case "price_12_4b":
			cm.Decimal(&e.Price124b)
		case "col_decimal_12_3":
			cm.Decimal(&e.ColDecimal123)
		case "col_decimal_20_6":
			cm.Decimal(&e.ColDecimal206)
		case "col_decimal_24_12":
			cm.Decimal(&e.ColDecimal2412)
		case "col_int_1":
			cm.NullInt32(&e.ColInt1)
		case "col_int_2":
			cm.Int32(&e.ColInt2)
		case "col_int_3":
			cm.NullUint32(&e.ColInt3)
		case "col_int_4":
			cm.Uint32(&e.ColInt4)
		case "col_longtext_1":
			cm.NullString(&e.ColLongtext1)
		case "col_longtext_2":
			cm.String(&e.ColLongtext2)
		case "col_mediumblob":
			cm.Byte(&e.ColMediumblob)
		case "col_mediumtext_1":
			cm.NullString(&e.ColMediumtext1)
		case "col_mediumtext_2":
			cm.String(&e.ColMediumtext2)
		case "col_smallint_1":
			cm.NullInt32(&e.ColSmallint1)
		case "col_smallint_2":
			cm.Int32(&e.ColSmallint2)
		case "col_smallint_3":
			cm.NullUint32(&e.ColSmallint3)
		case "col_smallint_4":
			cm.Uint32(&e.ColSmallint4)
		case "has_smallint_5":
			cm.Bool(&e.HasSmallint5)
		case "is_smallint_5":
			cm.NullBool(&e.IsSmallint5)
		case "col_text":
			cm.NullString(&e.ColText)
		case "col_timestamp_1":
			cm.Time(&e.ColTimestamp1)
		case "col_timestamp_2":
			cm.NullTime(&e.ColTimestamp2)
		case "col_tinyint_1":
			cm.Int32(&e.ColTinyint1)
		case "col_varchar_1":
			cm.String(&e.ColVarchar1)
		case "col_varchar_100":
			cm.NullString(&e.ColVarchar100)
		case "col_varchar_16":
			cm.String(&e.ColVarchar16)
		case "col_char_1":
			cm.NullString(&e.ColChar1)
		case "col_char_2":
			cm.String(&e.ColChar2)
		default:
			return errors.NotFound.Newf("[testdata] DmlgenTypes Column %q not found", c)
		}
	}
	return errors.WithStack(cm.Err())
}

// Empty empties all the fields of the current object. Also known as Reset.
func (e *DmlgenTypes) Empty() *DmlgenTypes { *e = DmlgenTypes{}; return e }

// DmlgenTypesCollection represents a collection type for DB table dmlgen_types
// Not thread safe. Auto generated.
// Just another comment.
//easyjson:json
type DmlgenTypesCollection struct {
	Data             []*DmlgenTypes                   `json:"data,omitempty"`
	BeforeMapColumns func(uint64, *DmlgenTypes) error `json:"-"`
	AfterMapColumns  func(uint64, *DmlgenTypes) error `json:"-"`
}

// NewDmlgenTypesCollection creates a new initialized collection. Auto generated.
func NewDmlgenTypesCollection() *DmlgenTypesCollection {
	return &DmlgenTypesCollection{
		Data: make([]*DmlgenTypes, 0, 5),
	}
}

func (cc *DmlgenTypesCollection) scanColumns(cm *dml.ColumnMap, e *DmlgenTypes, idx uint64) error {
	if err := cc.BeforeMapColumns(idx, e); err != nil {
		return errors.WithStack(err)
	}
	if err := e.MapColumns(cm); err != nil {
		return errors.WithStack(err)
	}
	if err := cc.AfterMapColumns(idx, e); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// MapColumns implements dml.ColumnMapper interface. Auto generated.
func (cc *DmlgenTypesCollection) MapColumns(cm *dml.ColumnMap) error {
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
		e := new(DmlgenTypes)
		if err := cc.scanColumns(cm, e, cm.Count); err != nil {
			return errors.WithStack(err)
		}
		cc.Data = append(cc.Data, e)
	case dml.ColumnMapCollectionReadSet:
		for cm.Next() {
			switch c := cm.Column(); c {
			case "id":
				cm = cm.Int32s(cc.IDs()...)

			default:
				return errors.NotFound.Newf("[testdata] DmlgenTypesCollection Column %q not found", c)
			}
		}
	default:
		return errors.NotSupported.Newf("[testdata] Unknown Mode: %q", string(m))
	}
	return cm.Err()
}

// IDs returns a slice or appends to a slice all values.
// Auto generated.
func (cc *DmlgenTypesCollection) IDs(ret ...int32) []int32 {
	if ret == nil {
		ret = make([]int32, 0, len(cc.Data))
	}
	for _, e := range cc.Data {
		ret = append(ret, e.ID)
	}
	return ret
}

// ColDate2s belongs to the column `col_date_2`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *DmlgenTypesCollection) UniqueColDate2s(ret ...time.Time) []time.Time {
	if ret == nil {
		ret = make([]time.Time, 0, len(cc.Data))
	}

	dupCheck := make(map[time.Time]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.ColDate2] {
			ret = append(ret, e.ColDate2)
			dupCheck[e.ColDate2] = true
		}
	}
	return ret
}

// Price124as belongs to the column `price_12_4a`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *DmlgenTypesCollection) UniquePrice124as(ret ...null.Decimal) []null.Decimal {
	if ret == nil {
		ret = make([]null.Decimal, 0, len(cc.Data))
	}

	dupCheck := make(map[null.Decimal]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.Price124a] {
			ret = append(ret, e.Price124a)
			dupCheck[e.Price124a] = true
		}
	}
	return ret
}

// ColInt1s belongs to the column `col_int_1`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *DmlgenTypesCollection) UniqueColInt1s(ret ...int32) []int32 {
	if ret == nil {
		ret = make([]int32, 0, len(cc.Data))
	}

	dupCheck := make(map[int32]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.ColInt1.Int32] {
			ret = append(ret, e.ColInt1.Int32)
			dupCheck[e.ColInt1.Int32] = true
		}
	}
	return ret
}

// ColInt2s belongs to the column `col_int_2`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *DmlgenTypesCollection) UniqueColInt2s(ret ...int32) []int32 {
	if ret == nil {
		ret = make([]int32, 0, len(cc.Data))
	}

	dupCheck := make(map[int32]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.ColInt2] {
			ret = append(ret, e.ColInt2)
			dupCheck[e.ColInt2] = true
		}
	}
	return ret
}

// HasSmallint5s belongs to the column `has_smallint_5`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *DmlgenTypesCollection) UniqueHasSmallint5s(ret ...bool) []bool {
	if ret == nil {
		ret = make([]bool, 0, len(cc.Data))
	}

	dupCheck := make(map[bool]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.HasSmallint5] {
			ret = append(ret, e.HasSmallint5)
			dupCheck[e.HasSmallint5] = true
		}
	}
	return ret
}

// ColVarchar100s belongs to the column `col_varchar_100`
// and returns a slice or appends to a slice only unique values of that column.
// The values will be filtered internally in a Go map. No DB query gets
// executed. Auto generated.
func (cc *DmlgenTypesCollection) UniqueColVarchar100s(ret ...string) []string {
	if ret == nil {
		ret = make([]string, 0, len(cc.Data))
	}

	dupCheck := make(map[string]bool, len(cc.Data))
	for _, e := range cc.Data {
		if !dupCheck[e.ColVarchar100.String] {
			ret = append(ret, e.ColVarchar100.String)
			dupCheck[e.ColVarchar100.String] = true
		}
	}
	return ret
}

// FilterThis filters the current slice by predicate f without memory allocation.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Filter(f func(*DmlgenTypes) bool) *DmlgenTypesCollection {
	b := cc.Data[:0]
	for _, e := range cc.Data {
		if f(e) {
			b = append(b, e)
		}
	}
	cc.Data = b
	return cc
}

// Each will run function f on all items in []*DmlgenTypes.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Each(f func(*DmlgenTypes)) *DmlgenTypesCollection {
	for i := range cc.Data {
		f(cc.Data[i])
	}
	return cc
}

func (cc *DmlgenTypesCollection) SortByID() {
	sort.Slice(cc.Data, func(i, j int) bool {
		return cc.Data[i].ID < cc.Data[j].ID
	})
}

// Cut will remove items i through j-1.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Cut(i, j int) *DmlgenTypesCollection {
	z := cc.Data // copy slice header
	copy(z[i:], z[j:])
	for k, n := len(z)-j+i, len(z); k < n; k++ {
		z[k] = nil // this should avoid the memory leak
	}
	z = z[:len(z)-j+i]
	cc.Data = z
	return cc
}

// Swap will satisfy the sort.Interface.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Swap(i, j int) { cc.Data[i], cc.Data[j] = cc.Data[j], cc.Data[i] }

// Delete will remove an item from the slice.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Delete(i int) *DmlgenTypesCollection {
	z := cc.Data // copy the slice header
	end := len(z) - 1
	cc.Swap(i, end)
	copy(z[i:], z[i+1:])
	z[end] = nil // this should avoid the memory leak
	z = z[:end]
	cc.Data = z
	return cc
}

// Insert will place a new item at position i.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Insert(n *DmlgenTypes, i int) *DmlgenTypesCollection {
	z := cc.Data // copy the slice header
	z = append(z, &DmlgenTypes{})
	copy(z[i+1:], z[i:])
	z[i] = n
	cc.Data = z
	return cc
}

// Append will add a new item at the end of *DmlgenTypesCollection.
// Auto generated via dmlgen.
func (cc *DmlgenTypesCollection) Append(n ...*DmlgenTypes) *DmlgenTypesCollection {
	cc.Data = append(cc.Data, n...)
	return cc
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (cc *DmlgenTypesCollection) UnmarshalBinary(data []byte) error {
	return cc.Unmarshal(data) // Implemented via github.com/gogo/protobuf
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (cc *DmlgenTypesCollection) MarshalBinary() (data []byte, err error) {
	return cc.Marshal() // Implemented via github.com/gogo/protobuf
}
