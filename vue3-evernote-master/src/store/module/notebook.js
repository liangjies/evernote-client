import { GetNotebooks } from '@/api/notebook.js'
export const notebook = {
    namespaced: true,
    state: {
        notebooks: '',
    },
    mutations: {
        setNotebooks(state, notebooks) {
            state.notebooks = notebooks
        },
    },
    actions: {
        async GetNotebooksData() {
            const res = await GetNotebooks()
            console.log(res.code)
            if (res.code === 200) {
                this.commit('notebook/setNotebooks', res.data.list)
                return true
            }
        }
    },
    getters: {
        notebooks(state) {
            return state.notebooks
        },
    }
}