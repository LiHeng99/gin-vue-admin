package canal

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func SaveDataBaseFunc(info db_tools.DbInfo) {
	// 定义导出文件的路径和文件名
	// 获取当前项目路径
	projectPath, err := os.Getwd()
	if err != nil {
		fmt.Println("获取项目路径失败:", err)
		return
	}
	currentTime := time.Now().Unix()
	// 定义导出文件的相对路径和文件名
	fileName := info.DbName + "-" + strconv.FormatInt(currentTime, 10) + ".sql"
	outputPath := filepath.Join(projectPath, "sql")
	outputFile := filepath.Join(outputPath, fileName)

	// 创建输出文件所在的目录
	err = os.MkdirAll(outputPath, 0755)
	if err != nil {
		fmt.Println("创建导出目录失败:", err)
		return
	}

	// 定义 mysqldump 命令及参数
	cmd := exec.Command("mysqldump", "-u", info.DbUserName, "-p"+info.DbPassword, "-h", "127.0.0.1", info.DbName)

	// 将输出重定向到文件
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("创建导出文件失败:", err)
		return
	}
	defer output.Close()
	cmd.Stdout = output

	// 执行导出命令
	err = cmd.Run()
	if err != nil {
		fmt.Println("导出数据库失败:", err)
		return
	}

	fmt.Println("数据库导出完成")
}
