// Copyright 2015-2017, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dbr_test

import (
	"fmt"
	"strings"

	"github.com/corestoreio/csfw/storage/dbr"
	"github.com/corestoreio/errors"
)

// Make sure that type categoryEntity implements interface
var _ dbr.ArgumentsAppender = (*categoryEntity)(nil)

// categoryEntity represents just a demo record.
type categoryEntity struct {
	EntityID       int64 // Auto Increment
	AttributeSetID int64
	ParentID       string
	Path           dbr.NullString
	// TeaserIDs contain a list of foreign primary keys which identifies special
	// teaser to be shown on the category page.
	TeaserIDs []string
}

func (pe *categoryEntity) AppendArguments(stmtType int, args dbr.Arguments, columns []string) (dbr.Arguments, error) {
	for _, c := range columns {
		switch c {
		case "entity_id":
			args = append(args, dbr.Int64(pe.EntityID))
		case "attribute_set_id":
			args = append(args, dbr.Int64(pe.AttributeSetID))
		case "parent_id":
			args = append(args, dbr.String(pe.ParentID))
		case "path":
			args = append(args, pe.Path)
		case "teaser_id_s":
			if stmtType&dbr.SQLPartSet != 0 {
				if pe.TeaserIDs == nil {
					args = append(args, nil)
				} else {
					args = append(args, dbr.String(strings.Join(pe.TeaserIDs, "|")))
				}
			} else {
				args = append(args, dbr.Strings(pe.TeaserIDs))
			}
		default:
			return nil, errors.NewNotFoundf("[dbr_test] Column %q not found", c)
		}
	}
	return args, nil
}

func ExampleUpdate_SetRecord() {

	ce := &categoryEntity{345, 6, "p123", dbr.MakeNullString("4/5/6/7"), []string{"saleAutumn", "saleShoe"}}

	// Updates all rows in the table
	u := dbr.NewUpdate("catalog_category_entity").
		AddColumns("attribute_set_id", "parent_id", "path", "teaser_id_s").
		SetRecord(ce)
	writeToSQLAndInterpolate(u)

	fmt.Print("\n\n")

	ce = &categoryEntity{678, 6, "p456", dbr.NullString{}, nil}

	// Updates only one row in the table. You can call SetRecord and Exec as
	// often as you like. Each call to Exec will reassemble the arguments.
	u = dbr.NewUpdate("catalog_category_entity").
		AddColumns("attribute_set_id", "parent_id", "path", "teaser_id_s").
		SetRecord(ce).
		Where(dbr.Column("entity_id").PlaceHolder()) // No Arguments in Int64s because we need a place holder.
	writeToSQLAndInterpolate(u)

	// Output:
	//Prepared Statement:
	//UPDATE `catalog_category_entity` SET `attribute_set_id`=?, `parent_id`=?,
	//`path`=?, `teaser_id_s`=?
	//Arguments: [6 p123 4/5/6/7 saleAutumn|saleShoe]
	//
	//Interpolated Statement:
	//UPDATE `catalog_category_entity` SET `attribute_set_id`=6, `parent_id`='p123',
	//`path`='4/5/6/7', `teaser_id_s`='saleAutumn|saleShoe'
	//
	//Prepared Statement:
	//UPDATE `catalog_category_entity` SET `attribute_set_id`=?, `parent_id`=?,
	//`path`=?, `teaser_id_s`=? WHERE (`entity_id` = ?)
	//Arguments: [6 p456 <nil> <nil> 678]
	//
	//Interpolated Statement:
	//UPDATE `catalog_category_entity` SET `attribute_set_id`=6, `parent_id`='p456',
	//`path`=NULL, `teaser_id_s`=NULL WHERE (`entity_id` = 678)
}
