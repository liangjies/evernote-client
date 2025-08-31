
export const note = {
  namespaced: true,
  state: {
    noteId: -1,
    notebookId: -1,
    currentNote: {},
    trashNote: {},
    allTrashNotes: [],
    allNotes: [],
    richText: '',
    allNoteBooks: [],
  },
  mutations: {
    setNoteId(state, noteId) {
      state.noteId = noteId
    },
    setNotebookId(state, notebookId) {
      state.notebookId = notebookId
    },
    setCurrentNote(state, note) {
      state.currentNote = note
    },
    setAllNotes(state, notes) {
      state.allNotes = notes
    },
    setTrashNote(state, note) {
      state.trashNote = note
    },
    setAllTrashNotes(state, notes) {
      state.allTrashNotes = notes
    },
    setRichText(state, rich) {
      state.richText = rich
    },
    setAllNoteBooks(state, notebooks) {
      state.allNoteBooks = notebooks
    },
  },
  getters: {
    noteId(state) {
      return state.noteId
    },
    notebookId(state) {
      return state.notebookId
    },
    currentNote(state) {
      return state.currentNote
    },
    trashNote(state) {
      return state.trashNote
    },
    allTrashNotes(state) {
      return state.allTrashNotes
    },
    allNotes(state) {
      return state.allNotes
    },
    richText(state) {
      return state.richText
    },
    allNoteBooks(state) {
      return state.allNoteBooks
    },
  },
}
