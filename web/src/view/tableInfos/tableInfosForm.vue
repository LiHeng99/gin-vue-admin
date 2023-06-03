<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="表名称:" prop="tableName">
          <el-input v-model="formData.tableName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="表注释:" prop="tableComment">
          <el-input v-model="formData.tableComment" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据库名称:" prop="tableSchema">
          <el-input v-model="formData.tableSchema" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="表类型:" prop="tableType">
          <el-input v-model="formData.tableType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="存储引擎:" prop="engine">
          <el-input v-model="formData.engine" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="行数:" prop="tableRows">
          <el-input v-model.number="formData.tableRows" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据长度:" prop="dataLength">
          <el-input v-model.number="formData.dataLength" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
          <el-input v-model="formData.createTime" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据库id:" prop="dbId">
          <el-input v-model.number="formData.dbId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateTableInfosModel,
  findTableInfosModel
} from '@/api/tableInfos'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTableInfosModel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retableInfoModel
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
