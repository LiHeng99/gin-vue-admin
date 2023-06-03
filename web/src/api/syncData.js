import service from '@/utils/request'

// @Tags SyncDataModel
// @Summary 创建SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SyncDataModel true "创建SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /syncDataTask/createSyncDataModel [post]
export const createSyncDataModel = (data) => {
  return service({
    url: '/syncDataTask/createSyncDataModel',
    method: 'post',
    data
  })
}

// @Tags SyncDataModel
// @Summary 删除SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SyncDataModel true "删除SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /syncDataTask/deleteSyncDataModel [delete]
export const deleteSyncDataModel = (data) => {
  return service({
    url: '/syncDataTask/deleteSyncDataModel',
    method: 'delete',
    data
  })
}

// @Tags SyncDataModel
// @Summary 删除SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /syncDataTask/deleteSyncDataModel [delete]
export const deleteSyncDataModelByIds = (data) => {
  return service({
    url: '/syncDataTask/deleteSyncDataModelByIds',
    method: 'delete',
    data
  })
}

// @Tags SyncDataModel
// @Summary 更新SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SyncDataModel true "更新SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /syncDataTask/updateSyncDataModel [put]
export const updateSyncDataModel = (data) => {
  return service({
    url: '/syncDataTask/updateSyncDataModel',
    method: 'put',
    data
  })
}

// @Tags SyncDataModel
// @Summary 用id查询SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SyncDataModel true "用id查询SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /syncDataTask/findSyncDataModel [get]
export const findSyncDataModel = (params) => {
  return service({
    url: '/syncDataTask/findSyncDataModel',
    method: 'get',
    params
  })
}

// @Tags SyncDataModel
// @Summary 分页获取SyncDataModel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SyncDataModel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /syncDataTask/getSyncDataModelList [get]
export const getSyncDataModelList = (params) => {
  return service({
    url: '/syncDataTask/getSyncDataModelList',
    method: 'get',
    params
  })
}
