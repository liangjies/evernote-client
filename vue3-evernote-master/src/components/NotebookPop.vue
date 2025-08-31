<template>
  <div class="notebook-pop">
    <div class="notebook-pop-header">
      <div class="header-title">笔记本</div>
      <el-tooltip
        class="box-item"
        effect="dark"
        content="创建笔记本"
        placement="bottom"
        ><el-icon class="add-icon" @click="addDialog"><folder-add /></el-icon
      ></el-tooltip>
      <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
        <input
          v-model="notebookInput"
          class="notebook-title-input"
          placeholder="给笔记本起个名称"
          autofocus="true"
        />
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="addNewNotebook()">确定</el-button>
          </span>
        </template>
      </el-dialog>

      <div class="search-bar">
        <div class="search-tip" v-if="!search">
          <el-icon class="search-icon">
            <search />
          </el-icon>
          <span class="search-tip-text">查找笔记本</span>
        </div>
        <input type="text" class="search-input" v-model="search" />
      </div>
    </div>
    <div class="notebook-lists-box">
      <div
        class="notebook-lists"
        v-for="notebook in queryPart"
        :key="notebook.id"
      >
        <div class="notebook-list" @click.stop="doClickNotebook(notebook.id)">
          <div class="notebook-title">
            {{ notebook.title }}
          </div>
          <div class="notebook-number">{{ notebook.noteCounts }} 条笔记</div>
          <div class="notebook-operate">
            <el-icon
              class="notebook-operate-icon"
              title="编辑"
              @click.stop="editDialog(notebook.id)"
              ><edit
            /></el-icon>
            <el-icon
              class="notebook-operate-icon"
              title="删除"
              @click.stop="deleteDialog(notebook.id)"
              ><delete-filled
            /></el-icon>
          </div>
          <div class="notebook-split"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { CreateNotebook, DeleteNotebook, UpdateNotebook } from "@/api/notebook";
import { mapActions } from "vuex";
import { ElMessage, ElMessageBox } from "element-plus";
import router from "@/router/index";
export default {
  name: "NotebookPop",
  emits: ["PopOpenValue"],
  data() {
    return {
      notebookInput: "",
      notebookID: -1,
      dialogFormVisible: false,
      open: true,
      dialogTitle: "新建笔记本",
      search:"",
    };
  },
  created() {
    // this.notebooks = this.$store.getters["notebook/notebooks"];
  },
  computed: {
    ...mapState({
      notebooks: (state) => state.notebook.notebooks,
    }),
    queryPart() {
      var search = this.search;
      return this.notebooks.filter((item) => {
        return String(item.title).toLowerCase().indexOf(search) > -1;
      });
      
      return this.item;
    },
  },
  methods: {
    ...mapActions("notebook", ["GetNotebooksData"]),
    async addNewNotebook() {
      if (this.dialogTitle == "新建笔记本") {
        const res = await CreateNotebook({ title: this.notebookInput });
        if (res.code === 200) {
          this.GetNotebooksData();
          ElMessage({
            showClose: true,
            message: res.msg,
            type: "success",
          });
        }
      } else {
        const res = await UpdateNotebook({
          id: this.notebookID,
          title: this.notebookInput,
        });
        if (res.code === 200) {
          this.GetNotebooksData();
          ElMessage({
            showClose: true,
            message: res.msg,
            type: "success",
          });
        }
      }
      this.dialogFormVisible = false;
    },
    addDialog() {
      this.dialogFormVisible = true;
      this.dialogTitle = "新建笔记本";
      this.notebookInput = "";
    },
    deleteDialog(notebookId) {
      ElMessageBox.confirm(
        `是否删除 ${this.getNotebookById(notebookId).title} ?`,
        "提示",
        {
          confirmButtonText: "是",
          cancelButtonText: "否",
          type: "warning",
        }
      )
        .then(async () => {
          const res = await DeleteNotebook({ id: notebookId });
          if (res.code === 200) {
            this.GetNotebooksData();
            ElMessage({
              showClose: true,
              message: res.msg,
              type: "success",
            });
          }
        })
        .catch((err) => {});
    },
    editDialog(notebookId) {
      this.dialogFormVisible = true;
      this.dialogTitle = "编辑笔记本";
      this.notebookInput = this.getNotebookById(notebookId).title;
      this.notebookID = notebookId;
    },
    getNotebookById(id) {
      let res = "";
      this.notebooks.some((notebook) => {
        if (notebook.id == id) {
          res = notebook;
        }
      });
      return res;
    },
    doClickNotebook(id) {
      router.push({ name: "NoteDetail", params: { id: id } });
      this.open = !this.open;
      this.$emit("PopOpenValue", this.open);
    },
  },
};
</script>

<style lang="less" scoped>
.notebook-pop {
  margin-left: 73px;
  width: 459px;
  background: #fff;
  bottom: 0;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  border-right: 3px solid rgba(236, 236, 236, 0.7);
  transition: transform 0.4s ease-in-out;
  .notebook-pop-header {
    border-bottom: 1px solid #ececec;
    .header-title {
      padding: 21px 24px 32px;
      color: #878787;
      font-size: 21px;
      font-weight: 300;
      text-transform: uppercase;
      letter-spacing: 0.5px;
    }
    .add-icon {
      opacity: 0.5;
      font-size: 20px;
      overflow: hidden;
      cursor: pointer;
      position: absolute;
      top: 21px;
      right: 24px;
    }
    .notebook-title-input {
      width: 100%;
      padding: 10px 20px;
      font-size: 25px;
      font-weight: 700;
      outline: none;
      border: none;
    }
    .search-bar {
      padding: 0 0 10px;
      margin: 0 24px;
      position: relative;
      .search-tip {
        position: absolute;
        left: 0;
        text-align: center;
        top: 8px;
        width: 100%;
        // height: 52px;
        cursor: text;
        .search-icon {
          right: 5px;
          font-size: 12px;
          vertical-align: baseline;
          text-align: center;
        }
        .search-tip-text {
          color: #ababab;
          font-size: 13px;
          font-weight: 400;
          vertical-align: baseline;
          text-align: center;
        }
      }
      .search-input {
        position: relative;
        background-color: transparent;
        border: 1px solid #e1e1e1;
        border-radius: 2px;
        padding: 10px 40px 9px 8px;
        outline: none;
        box-sizing: border-box;
        width: 100%;
        font-size: 13px;
        font-weight: 400;
      }
    }
  }
  .notebook-lists-box {
    top: 122px;
    position: absolute;
    left: 0;
    bottom: 0;
    width: 100%;
    .notebook-lists {
      cursor: pointer;
      .notebook-list {
        box-sizing: border-box;
        position: sticky;
        height: 66px;
        .notebook-title {
          white-space: nowrap;
          color: #4a4a4a;
          margin: 12px 80px 6px 24px;
          overflow: hidden;
          text-overflow: ellipsis;
          display: inline-block;
          width: 300px;
          height: 20px;
          font-size: 16px;
          font-weight: 400;
        }
        .notebook-number {
          color: #878787;
          margin: 0 0 12px 24px;
          height: 12px;
          font-size: 11px;
          font-weight: 500;
        }
        .notebook-operate {
          display: inline-block;
          position: absolute;
          right: 18px;
          top: 17px;
          z-index: 1000;
          // opacity: 0;
          .notebook-operate-icon {
            font-size: 20px;
            font-weight: 1000;
            color: #fff;
            display: inline-block;
            cursor: pointer;
            margin-left: 15px;
            vertical-align: top;
          }
        }
        .notebook-split {
          border-bottom: 1px solid #ececec;
          position: absolute;
          left: 12px;
          right: 12px;
        }
        &:hover {
          background-color: rgba(43, 181, 92, 0.9);
          .notebook-title,
          .notebook-number {
            color: #fff;
          }
        }
      }
    }
  }
}
</style>