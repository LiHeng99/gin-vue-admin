import service from '@/utils/request'

// // @Tags TableInfo
// // @Summary 创建TableInfo
// // @Security ApiKeyAuth
// // @accept application/json
// // @Produce application/json
// // @Param data body model.TableInfo true "创建TableInfo"
// // @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// // @Router /table_info/createTableInfo [post]
// export const createTableInfo = (data) => {
//   return service({
//     url: '/table_info/createTableInfo',
//     method: 'post',
//     data
//   })
// }

// // @Tags TableInfo
// // @Summary 删除TableInfo
// // @Security ApiKeyAuth
// // @accept application/json
// // @Produce application/json
// // @Param data body model.TableInfo true "删除TableInfo"
// // @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// // @Router /table_info/deleteTableInfo [delete]
// export const deleteTableInfo = (data) => {
//   return service({
//     url: '/table_info/deleteTableInfo',
//     method: 'delete',
//     data
//   })
// }

// // @Tags TableInfo
// // @Summary 删除TableInfo
// // @Security ApiKeyAuth
// // @accept application/json
// // @Produce application/json
// // @Param data body request.IdsReq true "批量删除TableInfo"
// // @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// // @Router /table_info/deleteTableInfo [delete]
// export const deleteTableInfoByIds = (data) => {
//   return service({
//     url: '/table_info/deleteTableInfoByIds',
//     method: 'delete',
//     data
//   })
// }

// // @Tags TableInfo
// // @Summary 更新TableInfo
// // @Security ApiKeyAuth
// // @accept application/json
// // @Produce application/json
// // @Param data body model.TableInfo true "更新TableInfo"
// // @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// // @Router /table_info/updateTableInfo [put]
// export const updateTableInfo = (data) => {
//   return service({
//     url: '/table_info/updateTableInfo',
//     method: 'put',
//     data
//   })
// }

// // @Tags TableInfo
// // @Summary 用id查询TableInfo
// // @Security ApiKeyAuth
// // @accept application/json
// // @Produce application/json
// // @Param data query model.TableInfo true "用id查询TableInfo"
// // @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// // @Router /table_info/findTableInfo [get]
// export const findTableInfo = (params) => {
//   return service({
//     url: '/table_info/findTableInfo',
//     method: 'get',
//     params
//   })
// }

// @Tags TableInfo
// @Summary 分页获取TableInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取TableInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /table_info/getTableInfoList [get]
export const getTableInfoList = (params) => {
  return service({
    url: '/tableInfo/getTableInfoList',
    method: 'get',
    params
  })
}
