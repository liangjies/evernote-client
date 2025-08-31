<template>
  <sider-bar v-on:addNoteValue="addNoteValue"></sider-bar>
  <div class="note-page" v-show="!isCollapse">
    <div class="note-header">
      <div class="note-header-title">{{ title }}</div>
      <div class="note-add">
        <el-dropdown trigger="click">
          <span>
            <div class="note-add-button">
              <el-icon class="note-add-icon">
                <plus />
              </el-icon>
              <span class="note-add-text">新建笔记</span>
            </div>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="addNoteValue()"
                ><img
                  class="dropdown-img"
                  src="/src/common/images/documents.png"
                />空白文档</el-dropdown-item
              >
              <el-dropdown-item @click="addNoteValue('md')"
                ><img
                  class="dropdown-img"
                  src="/src/common/images/md.png"
                />Markdown</el-dropdown-item
              >
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
    <div class="notes-view">
      <div class="notes-view-subheader">
        <div class="subheader-text">{{ noteNum }}条笔记</div>
        <div class="subheader-options">
          <el-dropdown
            trigger="click"
            placement="bottom-start"
            @command="orderCommand"
          >
            <span class="el-dropdown-link">
              选项
              <el-icon class="el-icon--right">
                <arrow-down />
              </el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item disabled style="font-size: 11px"
                  >排序方式</el-dropdown-item
                >
                <el-dropdown-item command="1"
                  >创建日期（最早优先）</el-dropdown-item
                >
                <el-dropdown-item command="2"
                  >创建日期（最新优先）</el-dropdown-item
                >
                <el-dropdown-item command="3"
                  >更新日期（最早优先）</el-dropdown-item
                >
                <el-dropdown-item command="4"
                  >更新日期（最新优先）</el-dropdown-item
                >
                <el-dropdown-item command="5"
                  >标题（升序排列）</el-dropdown-item
                >
                <el-dropdown-item command="6"
                  >标题（降序排列）</el-dropdown-item
                >
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
      <div class="notes-view-ScrollWindow">
        <div
          v-for="(note, index) in allNotes"
          :key="note.id"
          @click="openNote(note, index)"
        >
          <div
            class="notes-view-note"
            :class="{ 'notes-view-note-selected': currentIndex == index }"
          >
            <div class="note-snippet-divide" v-if="index > 0"></div>
            <div class="note-hover"></div>
            <div class="note-border"></div>
            <div class="note-snippetContent">
              <div class="note-title">{{ note.title }}</div>
              <div class="note-date">{{ _formateDate(note.updatedAt) }}</div>
              <div class="note-snippet">{{ note.snippet }}</div>
              <div class="note-markdown" v-if="note.type == 2">
                <img src="/src/common/images/markdown.png" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SiderBar from "@/components/SiderBar.vue";
import { GetAllNotes, GetNotesByNotebookID, GetNoteById } from "@/api/note";
import { friendlyDate } from "@/utils/util";
export default {
  name: "NoteList",
  emits: ["noteChange"],
  components: { SiderBar },
  props: {
    isCollapse: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      noteNum: 0,
      allNotes: "",
      currentIndex: 0,
      title: "笔记",
      addNoteState: false,
    };
  },
  created() {},
  mounted() {
    this._getNotes();
  },
  methods: {
    // 新增笔记
    addNoteValue: function (value) {
      if (this.addNoteState == false) {
        this.addNoteState = true;
        this.addNote(value);
      }
    },
    // 排序方式
    orderCommand(command) {
      if (command == 1) {
        this.allNotes.sort((a, b) => (a.CreatedAt < b.CreatedAt ? 1 : -1));
      } else if (command == 2) {
        this.allNotes.sort((a, b) => (a.CreatedAt > b.CreatedAt ? 1 : -1));
      } else if (command == 3) {
        this.allNotes.sort((a, b) => (a.UpdatedAt < b.UpdatedAt ? 1 : -1));
      } else if (command == 4) {
        this.allNotes.sort((a, b) => (a.UpdatedAt > b.UpdatedAt ? 1 : -1));
      } else if (command == 5) {
        this.allNotes.sort((a, b) => (a.title < b.title ? 1 : -1));
      } else if (command == 6) {
        this.allNotes.sort((a, b) => (a.title > b.title ? 1 : -1));
      }
    },
    // 获取所有笔记
    async getAllNotes() {
      console.log("getAllNotes");
      const res = await GetAllNotes();
      if (res.code === 200) {
        this.allNotes = res.data.list;
        this.noteNum = res.data.total;

        if (this.$route.params.id == -1 || this.addNoteState) {
          return;
        }
        if (this.currentIndex == 0) {
          this.$emit("noteChange", this.allNotes[0]);
        }
      }
    },
    // 根据笔记本ID获取笔记
    async getNotes(id) {
      console.log("getNotes(id)");
      const res = await GetNotesByNotebookID({ id: id });
      if (res.code === 200) {
        this.title = res.msg;
        this.allNotes = res.data.list;
        this.noteNum = res.data.total;
        if (this.currentIndex == 0) {
          this.$emit("noteChange", this.allNotes[0]);
        }
      }
    },
    // 获取笔记
    async _getNotes() {
      console.log("_getNotes()");
      console.log("this.$route.params.id", this.$route.params.id);
      if (this.$route.params.id == 0 || this.$route.params.id == undefined) {
        await this.getAllNotes();
        this.title = "笔记";
      } else if (this.$route.params.id == "add") {
        await this.getAllNotes();
        this.title = "笔记";
        this.addNote();
      } else {
        await this.getNotes(this.$route.params.id);
      }
    },
    // 点击笔记详情
    openNote(note, index) {
      console.log("openNote(note, index)");
      this.currentIndex = index;
      this.$emit("noteChange", note);
    },
    // 格式化日期
    _formateDate(dateStr) {
      if (dateStr == "") {
        return "";
      } else {
        return friendlyDate(dateStr);
      }
    },
    // 添加笔记
    addNote(type) {
      console.log("addNote()");
      if (type == "md") {
        this.allNotes.unshift({ id: -2, title: "", updatedAt: "", type: 2 });
        this.$emit("noteChange", {
          id: -2,
          title: "",
          notebookID: this.$route.params.id,
          type: 2,
        });
      } else {
        this.allNotes.unshift({ id: -2, title: "", updatedAt: "", type: 1 });
        this.$emit("noteChange", {
          id: -2,
          title: "",
          notebookID: this.$route.params.id,
          type: 1,
        });
      }
    },
    // 刷新笔记列表
    async refresh() {
      this.addNoteState = false;
      console.log("refresh()");
      await this._getNotes();
    },
    // 删除笔记
    async delNote(id) {
      let index;
      for (let i = 0; i < this.allNotes.length; i++) {
        if (this.allNotes[i].id == id) {
          index = i;
          break;
        }
      }
      this.allNotes.splice(index, 1);
      this.currentIndex = -1;
    },
    async noteSave(id) {
      const res = await GetNoteById({ id: id });
      let temp = this.allNotes;
      for (let i = 0; i < temp.length; i++) {
        if (temp[i].id == id) {
          temp[i] = res.data.list[0];
          break;
        }
      }
      this.allNotes = temp;
    },
  },
  watch: {
    "$route.params.id": async function () {
      console.log("watch $route.params.id");
      //
      if (this.$route.params.addNoteState) {
        return;
      }
      // 清空富媒体
      this.$emit("noteChange", {
        id: -2,
        title: "",
        notebookID: this.$route.params.id,
      });
      console.log("---------------");
      console.log(this.$route.params.id);
      if (this.$route.params.id == 0) {
        await this.getAllNotes();
        this.title = "笔记";
      } else if (this.$route.params.id > 0) {
        await this.getNotes(this.$route.params.id);
      } else if (this.$route.params.id == "add") {
        await this.addNote();
      }
    },
    // 监听保存事件
    // noteSave: {
    //   async handler(newVal, oldVal) {
    //     console.log("watch noteSave");
    //     console.log(newVal, oldVal);
    //     if (newVal == oldVal || newVal == 0 || newVal <= 0 || oldVal <= 0) {
    //       return;
    //     }
    //     const res = await GetNoteById({ id: newVal });
    //     let temp = this.allNotes;
    //     for (let i = 0; i < temp.length; i++) {
    //       if (temp[i].id == newVal) {
    //         temp[i] = res.data.list[0];
    //         break;
    //       }
    //     }
    //     this.allNotes = temp;
    //   },
    // },
  },
};
</script>

<style lang="less" scoped>
.note-page {
  margin-left: 73px;
  width: 350px;
  position: absolute;
  .note-header {
    height: 8vh;
    box-sizing: border-box;
    margin: 0 auto;
    padding: 24px 20px 0 24px;
    text-align: left;
    position: relative;
    .note-header-title {
      color: #878787;
      display: inline-block;
      vertical-align: top;
      font-size: 21px;
      letter-spacing: 0.5px;
      text-transform: uppercase;
      font-weight: 300;
      line-height: 26px;
    }
    .note-add {
      position: absolute;
      top: 24px;
      right: 25px;
      z-index: 1;
      .note-add-button {
        display: flex;
        border: 1px solid #2dbe60;
        border-radius: 3px;
        background-color: #2dbe60;
        font-size: 11px;
        color: white;
        line-height: 27px;
        cursor: pointer;
        .note-add-icon {
          margin: 8px 3px 0 5px;
        }
        .note-add-text {
          margin-right: 9px;
        }
      }
    }
  }
  .notes-view {
    height: 7vh;
    color: #878787;
    height: 100%;
    width: 350px;
    overflow: hidden;
    box-sizing: border-box;
    position: relative;
    .notes-view-subheader {
      position: relative;
      border-bottom: 1px solid #ececec;
      padding: 0 24px;
      width: 100%;
      height: 4vh;
      margin-top: 3vh;
      .subheader-text {
        font-size: 13px;
        font-weight: 400;
        color: #ababab;
      }
      .subheader-options {
        position: absolute;
        right: 24px;
        top: 0;
        .el-dropdown-link {
          font-size: 13px;
          font-weight: 400;
          color: #878787;
          .el-icon--right {
            margin-left: 0;
            top: 2px;
          }
        }
      }
    }
    .notes-view-ScrollWindow {
      height: 85vh;
      overflow: hidden scroll;
      .notes-view-note-selected {
        border: 3px solid #d9d9d9;
      }
      .notes-view-note {
        &:hover {
          background-color: rgba(43, 181, 92, 0.9);
          .note-snippetContent {
            .note-title,
            .note-date,
            .note-snippet {
              color: #fff;
            }
          }
        }
        .note-snippet-divide {
          border-top: 1px solid #ececec;
          left: 20px;
          right: 20px;
          top: 0;
          position: absolute;
        }
        height: 120px;
        cursor: pointer;
        margin: 0 auto;
        text-align: left;
        overflow: hidden;
        position: relative;
        .note-snippetContent {
          color: #878787;
          left: 24px;
          overflow: hidden;
          overflow-wrap: break-word;
          position: absolute;
          right: 24px;
          top: 12px;
          word-wrap: break-word;
          bottom: 15px;
          .note-title {
            font-size: 16px;
            font-weight: 400;
            color: #4a4a4a;
            margin-bottom: 4px;
            max-height: 40px;
            overflow: hidden;
            overflow-wrap: break-word;
            text-overflow: ellipsis;
            white-space: nowrap;
            word-wrap: break-word;
            line-height: 20px;
            width: 302px;
          }
          .note-date {
            font-size: 11px;
            font-weight: 400;
            text-transform: uppercase;
            letter-spacing: 1px;
            margin-bottom: 8px;
          }
          .note-snippet {
            font-size: 12px;
            font-weight: 400;
          }
          .note-markdown {
            position: absolute;
            top: 0px;
            right: 0px;
            img {
              width: 34px;
            }
          }
        }
      }
    }
  }
}
.example-showcase .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}

.dropdown-img {
  width: 20px;
  margin-right: 5px;
}
</style>