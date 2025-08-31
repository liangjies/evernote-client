import service from '@/utils/request'

// @Summary 获取所有回收站笔记
// @Produce  application/json
// @Param data body 
// @Router /trash/all [get]
export const GetTrashs = (data) => {
    return service({
        url: '/trash/all',
        method: 'get',
        data
    })
}


// @Summary 用户根据id获取废纸篓笔记详情
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /trash/:id [get]
export const GetTrashById = (data) => {
    return service({
        url: '/trash/'+data.id,
        method: 'get',
    })
}

// @Summary 用户彻底删除笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /trash/confirm/:id [delete]
export const DeleteTrash = (data) => {
    return service({
        url: '/trash/confirm/'+data.id,
        method: 'delete',
    })
}

// @Summary 用户从废纸篓恢复笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /trash/confirm/:id [patch]
export const RevertNote = (data) => {
    return service({
        url: '/trash/revert/'+data.id,
        method: 'patch',
    })
}