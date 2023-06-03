import service from '@/utils/request'

// @Tags TableInfosModel
// @Summary 创建TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TableInfosModel true "创建TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tableInfoModel/createTableInfosModel [post]
export const createTableInfosModel = (data) => {
  return service({
    url: '/tableInfoModel/createTableInfosModel',
    method: 'post',
    data
  })
}

// @Tags TableInfosModel
// @Summary 删除TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TableInfosModel true "删除TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tableInfoModel/deleteTableInfosModel [delete]
export const deleteTableInfosModel = (data) => {
  return service({
    url: '/tableInfoModel/deleteTableInfosModel',
    method: 'delete',
    data
  })
}

// @Tags TableInfosModel
// @Summary 删除TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tableInfoModel/deleteTableInfosModel [delete]
export const deleteTableInfosModelByIds = (data) => {
  return service({
    url: '/tableInfoModel/deleteTableInfosModelByIds',
    method: 'delete',
    data
  })
}

// @Tags TableInfosModel
// @Summary 更新TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TableInfosModel true "更新TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tableInfoModel/updateTableInfosModel [put]
export const updateTableInfosModel = (data) => {
  return service({
    url: '/tableInfoModel/updateTableInfosModel',
    method: 'put',
    data
  })
}

// @Tags TableInfosModel
// @Summary 用id查询TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TableInfosModel true "用id查询TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tableInfoModel/findTableInfosModel [get]
export const findTableInfosModel = (params) => {
  return service({
    url: '/tableInfoModel/findTableInfosModel',
    method: 'get',
    params
  })
}

// @Tags TableInfosModel
// @Summary 分页获取TableInfosModel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TableInfosModel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tableInfoModel/getTableInfosModelList [get]
export const getTableInfosModelList = (params) => {
  return service({
    url: '/tableInfoModel/getTableInfosModelList',
    method: 'get',
    params
  })
}
