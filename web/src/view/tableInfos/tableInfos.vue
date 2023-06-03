<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="表名称">
         <el-input v-model="searchInfo.tableName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="表注释">
         <el-input v-model="searchInfo.tableComment" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="数据库名称">
         <el-input v-model="searchInfo.tableSchema" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="表类型">
         <el-input v-model="searchInfo.tableType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="存储引擎">
         <el-input v-model="searchInfo.engine" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="行数">
            
             <el-input v-model.number="searchInfo.tableRows" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="数据长度">
            
             <el-input v-model.number="searchInfo.dataLength" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="创建时间">
         <el-input v-model="searchInfo.createTime" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="数据库id">
            
             <el-input v-model.number="searchInfo.dbId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="表名称" prop="tableName" width="120" />
        <el-table-column align="left" label="表注释" prop="tableComment" width="120" />
        <el-table-column align="left" label="数据库名称" prop="tableSchema" width="120" />
        <el-table-column align="left" label="表类型" prop="tableType" width="120" />
        <el-table-column align="left" label="存储引擎" prop="engine" width="120" />
        <el-table-column align="left" label="行数" prop="tableRows" width="120" />
        <el-table-column align="left" label="数据长度" prop="dataLength" width="120" />
        <el-table-column align="left" label="创建时间" prop="createTime" width="120" />
        <el-table-column align="left" label="数据库id" prop="dbId" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateTableInfosModelFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="表名称:"  prop="tableName" >
          <el-input v-model="formData.tableName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="表注释:"  prop="tableComment" >
          <el-input v-model="formData.tableComment" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据库名称:"  prop="tableSchema" >
          <el-input v-model="formData.tableSchema" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="表类型:"  prop="tableType" >
          <el-input v-model="formData.tableType" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="存储引擎:"  prop="engine" >
          <el-input v-model="formData.engine" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="行数:"  prop="tableRows" >
          <el-input v-model.number="formData.tableRows" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据长度:"  prop="dataLength" >
          <el-input v-model.number="formData.dataLength" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建时间:"  prop="createTime" >
          <el-input v-model="formData.createTime" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据库id:"  prop="dbId" >
          <el-input v-model.number="formData.dbId" :clearable="true" placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TableInfosModel'
}
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

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tableName: '',
        tableComment: '',
        tableSchema: '',
        tableType: '',
        engine: '',
        tableRows: 0,
        dataLength: 0,
        createTime: '',
        dbId: 0,
        })

// 验证规则
const rule = reactive({
               tableName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dbId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


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

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTableInfosModelFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteTableInfosModelByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTableInfosModelFunc = async(row) => {
    const res = await findTableInfosModel({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.retableInfoModel
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTableInfosModelFunc = async (row) => {
    const res = await deleteTableInfosModel({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        tableName: '',
        tableComment: '',
        tableSchema: '',
        tableType: '',
        engine: '',
        tableRows: 0,
        dataLength: 0,
        createTime: '',
        dbId: 0,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createTableInfosModel(formData.value)
                  break
                case 'update':
                  res = await updateTableInfosModel(formData.value)
                  break
                default:
                  res = await createTableInfosModel(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
