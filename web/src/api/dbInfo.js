import service from '@/utils/request'

// @Tags DbInfo
// @Summary 创建DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DbInfo true "创建DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dbInfo/createDbInfo [post]
export const createDbInfo = (data) => {
  return service({
    url: '/dbInfo/createDbInfo',
    method: 'post',
    data
  })
}

// @Tags DbInfo
// @Summary 删除DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DbInfo true "删除DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dbInfo/deleteDbInfo [delete]
export const deleteDbInfo = (data) => {
  return service({
    url: '/dbInfo/deleteDbInfo',
    method: 'delete',
    data
  })
}

// @Tags DbInfo
// @Summary 删除DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dbInfo/deleteDbInfo [delete]
export const deleteDbInfoByIds = (data) => {
  return service({
    url: '/dbInfo/deleteDbInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags DbInfo
// @Summary 更新DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DbInfo true "更新DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dbInfo/updateDbInfo [put]
export const updateDbInfo = (data) => {
  return service({
    url: '/dbInfo/updateDbInfo',
    method: 'put',
    data
  })
}

// @Tags DbInfo
// @Summary 用id查询DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DbInfo true "用id查询DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dbInfo/findDbInfo [get]
export const findDbInfo = (params) => {
  return service({
    url: '/dbInfo/findDbInfo',
    method: 'get',
    params
  })
}

// @Tags DbInfo
// @Summary 分页获取DbInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DbInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dbInfo/getDbInfoList [get]
export const getDbInfoList = (params) => {
  return service({
    url: '/dbInfo/getDbInfoList',
    method: 'get',
    params
  })
}

//连接数据库
export const linkDbUrl = (params) => {
  return service({
    url: '/dbInfo/linkDbUrl',
    method: 'get',
    params
  })
}