import service from '@/utils/request'

// @Summary 获取所有笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /notes/list/all [get]
export const GetAllNotes = (data) => {
    return service({
        url: '/notes/list/all',
        method: 'get',
        data
    })
}

// @Summary 获取笔记本笔记列表
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const GetNotesByNotebookID = (data) => {
    return service({
        url: '/notes/list/' + data.id,
        method: 'get',
    })
}


// @Summary 用户根据id获取笔记详情
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /notes/:id [get]
export const GetNoteById = (data) => {
    return service({
        url: '/notes/' + data.id,
        method: 'get',
    })
}

// @Summary 用户修改笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /notes/update [post]
export const UpdateNote = (data) => {
    return service({
        url: '/notes/update',
        method: 'post',
        data
    })
}

// @Summary 用户新建笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /notes/add [post]
export const CreateNote = (data) => {
    return service({
        url: '/notes/add',
        method: 'post',
        data
    })
}

// @Summary 用户删除笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /notes/add [post]
export const DeleteNote = (data) => {
    return service({
        url: '/notes/' + data.id,
        method: 'delete',
    })
}

// @Summary 用户搜索笔记
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /notes/search [post]
export const SearchNote = (data) => {
    return service({
        url: 'notes/search',
        method: 'post',
        data
    })
}