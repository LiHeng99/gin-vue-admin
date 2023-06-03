<template>
  <div>
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        @keyup.enter="onSubmit"
      >
        <el-form-item label="创建时间">
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始时间"
          ></el-date-picker>
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束时间"
          ></el-date-picker>
        </el-form-item>
        <el-form-item label="数据库名称">
          <el-input v-model="searchInfo.dbName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="数据库名称">
          <el-input v-model="searchInfo.dbUserName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="密码">
          <el-input v-model="searchInfo.dbPassword" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="url">
          <el-input v-model="searchInfo.dbUrl" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="数据库驱动">
          <el-input
            v-model="searchInfo.driverClassName"
            placeholder="搜索条件"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit"
            >查询</el-button
          >
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog"
          >新增</el-button
        >
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button type="primary" link @click="deleteVisible = false"
              >取消</el-button
            >
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
              >删除</el-button
            >
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
          <template #default="scope">{{
            formatDate(scope.row.CreatedAt)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="数据库名称"
          prop="dbName"
          width="120"
        />
        <el-table-column
          align="left"
          label="数据库名称"
          prop="dbUserName"
          width="120"
        />
        <el-table-column
          align="left"
          label="密码"
          prop="dbPassword"
          width="120"
        />
        <el-table-column align="left" label="url" prop="dbUrl" width="500" />
        <el-table-column
          align="left"
          label="数据库驱动"
          prop="driverClassName"
          width="120"
        />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="link"
              class="table-button"
              @click="linkDbUrlFunc(scope.row)"
              >连接</el-button
            >
            <el-button
              type="primary"
              link
              icon="look"
              class="table-button"
              @click="getTableInfoListFunc(scope.row)"
              >查看</el-button
            >
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateDbInfoFunc(scope.row)"
              >变更</el-button
            >
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
              >删除</el-button
            >
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="弹窗操作"
    >
      <el-form
        :model="formData"
        label-position="right"
        ref="elFormRef"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item label="数据库名称:" prop="dbName">
          <el-input
            v-model="formData.dbName"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="用户名:" prop="dbUserName">
          <el-input
            v-model="formData.dbUserName"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="密码:" prop="dbPassword">
          <el-input
            v-model="formData.dbPassword"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="连接地址:" prop="dbUrl">
          <el-input
            v-model="formData.dbUrl"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item label="数据库驱动:" prop="driverClassName">
          <el-input
            v-model="formData.driverClassName"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">测试连接</el-button>
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- <el-button type="primary" @click="dialogVisible = true">打开弹框</el-button> -->
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
  </div>
</template>

<script>
export default {
  name: "DbInfo",
};
</script>

<script setup>
import {
  createDbInfo,
  deleteDbInfo,
  deleteDbInfoByIds,
  updateDbInfo,
  findDbInfo,
  linkDbUrl,
  getDbInfoList,
  getTableInfoList,
  saveTableInfoList
} from "@/api/dbInfo";

// 全量引入格式化工具 请按需保留
import {
  getDictFunc,
  formatDate,
  formatBoolean,
  filterDict,
} from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref, reactive } from "vue";

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  dbName: "",
  dbUserName: "",
  dbPassword: "",
  dbUrl: "",
  driverClassName: "",
});

// 验证规则
const rule = reactive({
  dbName: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  dbUserName: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  dbPassword: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  dbUrl: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
  driverClassName: [
    {
      required: true,
      message: "",
      trigger: ["input", "blur"],
    },
  ],
});

const elFormRef = ref();

// =========== 表格控制部分 ===========
const page = ref(1);
const total = ref(0);
const pageSize = ref(10);
const tableData = ref([]);
const tableInfoList = ref([]);
const searchInfo = ref({});

// 重置
const onReset = () => {
  searchInfo.value = {};
  getTableData();
};

// 搜索
const onSubmit = () => {
  page.value = 1;
  pageSize.value = 10;
  getTableData();
};

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val;
  getTableData();
};

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val;
  getTableData();
};

// 查询
const getTableData = async () => {
  const table = await getDbInfoList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  });
  if (table.code === 0) {
    tableData.value = table.data.list;
    total.value = table.data.total;
    page.value = table.data.page;
    pageSize.value = table.data.pageSize;
  }
};

getTableData();

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {};

// 获取需要的字典 可能为空 按需保留
setOptions();

// 多选数据
const multipleSelection = ref([]);

// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val;
};

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm("确定要删除吗?", "提示", {
    confirmButtonText: "确定",
    cancelButtonText: "取消",
    type: "warning",
  }).then(() => {
    deleteDbInfoFunc(row);
  });
};

// 批量删除控制标记
const deleteVisible = ref(false);
// 保存表信息多选数据
// 多选删除
const onDelete = async () => {
  const ids = [];
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: "warning",
      message: "请选择要删除的数据",
    });
    return;
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID);
    });
  const res = await deleteDbInfoByIds({ ids });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--;
    }
    deleteVisible.value = false;
    getTableData();
  }
};

// 行为控制标记（弹窗内部需要增还是改）
const type = ref("");
//连接数据库
const linkDbUrlFunc = async (row) => {
  const res = await linkDbUrl({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "连接数据库",
    });
  }else{
    ElMessage({
      type: "error",
      message: "连接数据库失败",
    });
  }
};

const tableInfosSelection = ref([]);
// 多选
const tableInfosHandleSelectionChange = (val) => {
  tableInfosSelection.value = val;
};

//保存表信息
const saveTableInfoListFunc = async () => {
  const res = await saveTableInfoList(tableInfosSelection.value);
  tableInfoDialogVisible.value = false;
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "保存表信息成功",
    });
  }else{
    ElMessage({
      type: "error",
      message: "保存表信息失败",
    });
  }
};
//查看
const getTableInfoListFunc = async (row) => {
  const res = await getTableInfoList({ ID: row.ID });
  if (res.code === 0) {
    tableInfoDialogVisible.value = true;
    console.log(res.data.list, "获取表信息");
    tableInfoList.value = res.data.list;
  }
};

// 更新行
const updateDbInfoFunc = async (row) => {
  const res = await findDbInfo({ ID: row.ID });
  type.value = "update";
  if (res.code === 0) {
    formData.value = res.data.redbInfo;
    dialogFormVisible.value = true;
  }
};

// 删除行
const deleteDbInfoFunc = async (row) => {
  const res = await deleteDbInfo({ ID: row.ID });
  if (res.code === 0) {
    ElMessage({
      type: "success",
      message: "删除成功",
    });
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    getTableData();
  }
};

// 弹窗控制标记
const dialogFormVisible = ref(false);
const tableInfoDialogVisible = ref(false);

// 打开弹窗
const openDialog = () => {
  type.value = "create";
  dialogFormVisible.value = true;
};

// 关闭弹窗
const closeDialog = () => {
  tableInfoDialogVisible.value = false;
  dialogFormVisible.value = false;
  formData.value = {
    dbName: "",
    dbUserName: "",
    dbPassword: "",
    dbUrl: "",
    driverClassName: "",
  };
};
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return;
    let res;
    switch (type.value) {
      case "create":
        res = await createDbInfo(formData.value);
        break;
      case "update":
        res = await updateDbInfo(formData.value);
        break;
      default:
        res = await createDbInfo(formData.value);
        break;
    }
    if (res.code === 0) {
      ElMessage({
        type: "success",
        message: "创建/更改成功",
      });
      closeDialog();
      getTableData();
    }
  });
};
</script>

<style>
</style>
