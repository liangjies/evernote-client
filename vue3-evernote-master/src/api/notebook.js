import service from '@/utils/request'

// @Summary 获取笔记本
// @Produce  application/json
// @Param data
// @Router /notebooks/get [get]
export const GetNotebooks = (data) => {
    return service({
        url: '/notebook/get',
        method: 'get',
        data
    })
}


// @Summary 新建笔记本
// @Produce  application/json
// @Param data
// @Router /notebooks/add [post]
export const CreateNotebook = (data) => {
    return service({
        url: '/notebook/add',
        method: 'post',
        data
    })
}

// @Summary 删除笔记本
// @Produce  application/json
// @Param data
// @Router /notebooks/del [post]
export const DeleteNotebook = (data) => {
    return service({
        url: '/notebook/del/'+data.id,
        method: 'delete',
    })
}

// @Summary 用户修改笔记本
// @Produce  application/json
// @Param data
// @Router /notebook/:id [patch]
export const UpdateNotebook = (data) => {
    return service({
        url: '/notebook/'+data.id,
        method: 'patch',
        data
    })
}