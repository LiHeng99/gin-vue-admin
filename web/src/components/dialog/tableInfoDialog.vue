<template>
   <el-dialog
      title="选择表信息"
      v-model="tableInfoDialogVisible"
      :before-close="closeDialog"
      width="50%"
    >
      <el-table
        :data="tableInfoList"
        style="width: 100%"
        @selection-change="tableInfosHandleSelectionChange"
      >
        <el-table-column type="selection" width="55"></el-table-column>
        <el-table-column prop="tableName" label="表名称"></el-table-column>
        <el-table-column prop="dbId" label="库连接ID"></el-table-column>
        <el-table-column prop="tableComment" label="表注释"></el-table-column>
        <el-table-column prop="tableSchema" label="表架构"></el-table-column>
        <el-table-column prop="tableType" label="表类型"></el-table-column>
        <el-table-column prop="engine" label="引擎"></el-table-column>
        <el-table-column prop="tableRows" label="行数"></el-table-column>
        <el-table-column prop="dataLength" label="数据长度"></el-table-column>
        <el-table-column prop="createTime" label="创建时间"></el-table-column>
        <!-- ... 你可以添加其他列 ... -->
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveTableInfoListFunc"
          >保存</el-button
        >
      </span>
    </el-dialog>
</template>

<script>
export default {
  name: "tableInfoDialog",
};
</script>

<script setup>
import {
  createTableInfosModel,
  deleteTableInfosModel,
  deleteTableInfosModelByIds,
  updateTableInfosModel,
  findTableInfosModel,
  getTableInfosModelList
} from '@/api/tableInfos'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTableInfosModelList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()
</script>

<style>

</style>