<template>
  <sider-bar></sider-bar>
  <div class="note-page" v-show="!isCollapse">
    <div class="note-header">
      <div class="note-header-title">回收站</div>
    </div>
    <div class="notes-view">
      <div class="notes-view-subheader">
        <div class="subheader-text">{{ noteNum }}条笔记</div>
        <div class="subheader-options">
          <el-dropdown trigger="click" placement="bottom-start">
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
                <el-dropdown-item>创建日期（最早优先）</el-dropdown-item>
                <el-dropdown-item>创建日期（最新优先）</el-dropdown-item>
                <el-dropdown-item>更新日期（最早优先）</el-dropdown-item>
                <el-dropdown-item>更新日期（最新优先）</el-dropdown-item>
                <el-dropdown-item>标题（升序排列）</el-dropdown-item>
                <el-dropdown-item>标题（降序排列）</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
      <div class="notes-view-ScrollWindow">
        <div
          v-for="(note, index) in allNotes"
          :key="note.id"
          @click="openTrash(note, index)"
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
            </div>
            <div class="note-operate">
              <button class="note-operate-button" @click="deleteTrash(note)">
                彻底删除
              </button>
              <button class="note-operate-button" @click="revertNote(note)">
                还原
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SiderBar from "@/components/SiderBar.vue";
import { ArrowDown } from "@element-plus/icons-vue";
import { GetTrashs } from "@/api/trash";
import { friendlyDate } from "@/utils/util";
import { DeleteTrash, RevertNote } from "@/api/trash";
import { ElMessage, ElMessageBox } from "element-plus";
export default {
  name: "TrashSide",
  emits: ["TrashValue"],
  props: {
    isCollapse: {
      type: Boolean,
      default: false,
    },
  },
  components: { SiderBar, ArrowDown },
  data() {
    return {
      noteNum: 0,
      allNotes: "",
      currentIndex: 0,
    };
  },
  created() {
    this.getTrash();
  },
  methods: {
    async getTrash() {
      const res = await GetTrashs();
      if (res.code === 200) {
        this.allNotes = res.data.list;
        this.noteNum = res.data.total;
        if (this.allNotes.length > 0 && this.currentIndex == 0) {
          this.$emit("TrashValue", this.allNotes[0]);
        }
      }
    },
    openTrash(note, index) {
      this.currentIndex = index;
      this.$emit("TrashValue", note);
    },
    deleteTrash(note) {
      ElMessageBox.confirm("是否彻底删除?", "提示", {
        confirmButtonText: "是",
        cancelButtonText: "否",
        type: "warning",
      }).then(async () => {
        const res = await DeleteTrash({ id: note.id });
        if (res.code === 200) {
          this._deleteNotes(note.id);
          ElMessage({
            showClose: true,
            message: res.msg,
            type: "success",
          });
        }
      });
    },
    async revertNote(note) {
      const res = await RevertNote({ id: note.id });
      if (res.code === 200) {
        this._deleteNotes(note.id);
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
      }
    },
    _deleteNotes(id) {
      let tempNotes = [];
      for (let i = 0; i < this.allNotes.length; i++) {
        if (this.allNotes[i].id != id) {
          tempNotes.push(this.allNotes[i]);
        }
      }
      this.allNotes = tempNotes;
      this.$emit("TrashValue", this.allNotes[0]);
    },
    _formateDate(dateStr) {
      if (dateStr == "") {
        return "";
      } else {
        return friendlyDate(dateStr);
      }
    },
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
    margin-bottom: -7px;
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
      height: 24px;
      margin-top: 32px;
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
        }
      }
      .note-operate {
        position: absolute;
        opacity: 1;
        top: 0;
        right: 0;
        .note-operate-button {
          font-size: 12px;
          font-weight: 400;
          line-height: 17px;
          background: none;
          border: 1px solid #fff;
          border-radius: 3px;
          color: #fff;
          cursor: pointer;
          display: inline-block;
          margin: 16px 16px 0 0;
          padding: 0 5px;
          vertical-align: top;
          outline: none;
          &:hover {
            color: #2dbe60;
            background-color: #fff;
            border-color: transparent;
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
</style>