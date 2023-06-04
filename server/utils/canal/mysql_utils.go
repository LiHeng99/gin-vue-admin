package canal

import (
	"database/sql"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	dbToolsService "github.com/flipped-aurora/gin-vue-admin/server/service/db_tools"
	"strconv"
)

// 查询创建表sql
func SelectCreateSql(dbInfo db_tools.DbInfo, tableName string) (createSql string) {
	db, _ := dbToolsService.GetDB(strconv.FormatUint(uint64(dbInfo.ID), 10) + dbInfo.DbName)
	// 执行查询语句
	query := fmt.Sprintf("SHOW CREATE TABLE %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	defer rows.Close()

	// 解析查询结果
	var createStatement string
	for rows.Next() {
		err = rows.Scan(&tableName, &createStatement)
		if err != nil {
			fmt.Println("扫描失败:", err)
			return
		}
		return createStatement
		// 输出表的创建语句
	}
	return ""
}

//查询表是否存在

func SelectTableExist(dbInfo db_tools.DbInfo, tableName string) (b bool) {
	// 执行查询语句
	query := fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s'", dbInfo.DbName, tableName)
	var result string
	db, _ := dbToolsService.GetDB(strconv.FormatUint(uint64(dbInfo.ID), 10) + dbInfo.DbName)
	err := db.QueryRow(query).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("表 %s 不存在\n", tableName)
		} else {
			fmt.Println("查询失败:", err)
		}
		return false
	}
	return true
	// 表存在
}

// 根据时间字段查询出最新数据
func SelectDataByTime(dbInfo db_tools.DbInfo, var1 string, tableName string) (result int) {
	query := fmt.Sprintf("SELECT MAX(%s) as result FROM %s", var1, tableName)
	//query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s DESC LIMIT 1", tableName, var1)
	db, _ := dbToolsService.GetDB(strconv.FormatUint(uint64(dbInfo.ID), 10) + dbInfo.DbName)
	row := db.QueryRow(query)
	if err := row.Scan(&result); err != nil {
		fmt.Println("查询失败:", err)
	}
	return
}
