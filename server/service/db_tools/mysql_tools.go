package db_tools

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type TableInfoService struct {
}

func (TableInfoService *TableInfoService) GetTableList(dbInfo db_tools.DbInfo) (list []tableInfos.TableInfosModel, err error) {
	db, flag := GetDB(strconv.FormatUint(uint64(dbInfo.ID), 10) + dbInfo.DbName)
	if flag {
		rows, err := db.Query("SELECT TABLE_NAME, TABLE_COMMENT, TABLE_SCHEMA, TABLE_TYPE, ENGINE,TABLE_ROWS, DATA_LENGTH, CREATE_TIME FROM information_schema.tables WHERE TABLE_SCHEMA = " + "'" + dbInfo.DbName + "';")
		if err != nil {
			// handle error
			return nil, err
		}
		for rows.Next() {
			var table tableInfos.TableInfosModel
			table.DbId = dbInfo.ID
			err := rows.Scan(&table.TableName, &table.TableComment, &table.TableSchema, &table.TableType, &table.Engine, &table.TableRows, &table.DataLength, &table.CreateTime)
			if err != nil {
				// handle error
				return nil, err
			}
			list = append(list, table)
		}
		defer rows.Close()
		if err := rows.Err(); err != nil {
			// handle error
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unable to connect to database")
	}
	return list, nil
}
