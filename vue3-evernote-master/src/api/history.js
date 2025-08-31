import service from '@/utils/request'

// @Summary 获取笔记历史版本
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /history/:id [get]
export const GetHistories = (data) => {
    return service({
        url: '/history/' + data.id,
        method: 'get',
    })
}

// @Summary 恢复笔记历史版本
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router history/recover [post]
export const RecoverHistory = (data) => {
    return service({
        url: 'history/recover',
        method: 'post',
        data
    })
}