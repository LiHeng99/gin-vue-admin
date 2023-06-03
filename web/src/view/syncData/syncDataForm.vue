<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务名称:" prop="taskName">
          <el-input v-model="formData.taskName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="源数据库ID:" prop="sourceDbId">
          <el-input v-model.number="formData.sourceDbId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="目标数据ID:" prop="targetDbId">
          <el-input v-model.number="formData.targetDbId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="同步类型:" prop="syncType">
          <el-select v-model="formData.syncType" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in syncTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="同步方向:" prop="syncDirection">
          <el-select v-model="formData.syncDirection" placeholder="请选择" :clearable="true">
            <el-option v-for="(item,key) in SyncDirectionOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="描述:" prop="readme">
          <el-input v-model="formData.readme" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="表信息ID:" prop="tableInfoId">
          <el-input v-model="formData.tableInfoId" :clearable="true" placeholder="请输入" />
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
  name: 'SyncDataModel'
}
</script>

<script setup>
import {
  createSyncDataModel,
  updateSyncDataModel,
  findSyncDataModel
} from '@/api/syncData'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const syncTypeOptions = ref([])
const SyncDirectionOptions = ref([])
const formData = ref({
            taskName: '',
            sourceDbId: 0,
            targetDbId: 0,
            syncType: undefined,
            syncDirection: undefined,
            readme: '',
            tableInfoId: '',
        })
// 验证规则
const rule = reactive({
               taskName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sourceDbId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               targetDbId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               syncType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               syncDirection : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               tableInfoId : [{
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
      const res = await findSyncDataModel({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resyncDataTask
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    syncTypeOptions.value = await getDictFunc('syncType')
    SyncDirectionOptions.value = await getDictFunc('SyncDirection')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSyncDataModel(formData.value)
               break
             case 'update':
               res = await updateSyncDataModel(formData.value)
               break
             default:
               res = await createSyncDataModel(formData.value)
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
