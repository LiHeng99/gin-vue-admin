<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="数据库名称:" prop="dbUserName">
          <el-input v-model="formData.dbUserName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="密码:" prop="dbPassword">
          <el-input v-model="formData.dbPassword" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="url:" prop="dbUrl">
          <el-input v-model="formData.dbUrl" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="数据库驱动:" prop="driverClassName">
          <el-input v-model="formData.driverClassName" :clearable="true" placeholder="请输入" />
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
  name: 'DbInfo'
}
</script>

<script setup>
import {
  createDbInfo,
  updateDbInfo,
  findDbInfo
} from '@/api/dbInfo'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            dbUserName: '',
            dbPassword: '',
            dbUrl: '',
            driverClassName: '',
        })
// 验证规则
const rule = reactive({
               dbUserName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dbPassword : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               dbUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               driverClassName : [{
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
      const res = await findDbInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redbInfo
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
               res = await createDbInfo(formData.value)
               break
             case 'update':
               res = await updateDbInfo(formData.value)
               break
             default:
               res = await createDbInfo(formData.value)
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
