
export let API_REGISTER = '/auth/register'
export let API_LOGIN = '/base/login'
export let API_AUTH = '/auth'
export let API_LOGOUT = '/auth/logout'
export let API_GET_NOTEBOOK = '/notebooks'
export let API_ADD_NOTEBOOK = '/notebooks' /** POST */
export let API_UPDATE_NOTEBOOK = '/notebooks/:notebookId' /** PATCH */
export let API_DELETE_NOTEBOOK = '/notebooks/:notebookId' /** DELETE */
export let API_ADD_NOTE = '/notes/to/:notebookId' /** POST */
export let API_GET_NOTE = '/notes/from/:notebookId'
export let API_GET_NOTE_BY_ID = '/notes/:noteId'
export let API_DELETE_NOTE = '/notes/:noteId' /** DELETE */
export let API_UPDATE_NOTE = '/notes/:noteId' /** UPDATE */
export let API_CONFIRM_DELETE_NOTE = '/notes/confirm/:noteId' /** DELETE */
export let API_REVERT_NOTE = '/notes/revert/:noteId' /** PATCH */
export let API_GET_TRASH = '/notes/trash'
