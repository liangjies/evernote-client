<template>
  <sider-bar></sider-bar>
  <div class="search-page" v-show="!detailShow">
    <div class="search-top">
      <div class="search-input-box">
        <input
          type="text"
          class="search-input-text"
          id="gwt-debug-searchViewSearchBox"
          placeholder="搜索笔记"
          autofocus
          v-model="searchKey"
          ref="search"
          @keydown.enter="onSearchFirst"
        />
        <el-icon class="search-input-icon" @click="clearSearch"
          ><circle-close-filled
        /></el-icon>
      </div>
    </div>
    <div class="search-input-footer">
      <div class="search-footer-box">
        <div class="search-input-footer-text">正在搜索&nbsp;</div>
        <div class="search-input-footer-tip">你的笔记本</div>
      </div>
    </div>
    <div class="search-input-footer-split"></div>
  </div>

  <div v-show="detailShow">
    <div class="search-page search-page-top">
      <div class="search-top">
        <div class="search-input-box">
          <el-input
            prefix-icon="Search"
            :clearable="true"
            type="text"
            class="search-input-text"
            id="gwt-debug-searchViewSearchBox"
            placeholder="搜索笔记"
            autofocus
            v-model="searchKey"
            ref="search"
            @keydown.enter="onSearch"
          />
          <!-- <el-icon class="search-input-icon" @click="clearSearch"
            ><circle-close-filled
          /></el-icon> -->
        </div>
      </div>
      <!-- <div class="search-input-footer">
        <div class="search-footer-box">
          <div class="search-input-footer-text">正在搜索&nbsp;</div>
          <div class="search-input-footer-tip">你的笔记本</div>
        </div>
      </div> -->
    </div>
    <!-- <div class="search-input-footer-split"></div> -->
    <search-side
      :searchValue="searchValue"
      v-on:noteChange="noteChange"
      ref="note"
    ></search-side>
    <div
      class="note-detail"
      v-show="this.id != -1"
      :class="{ isCollapse: isCollapse }"
    >
      <div class="note-header">
        <div class="note-operation">
          <div class="note-operation-left">
            <!--伸缩功能-->
            <span @click="totalCollapse">
              <el-icon class="note-operation-icon" v-if="isCollapse">
                <expand />
              </el-icon>
              <el-icon class="note-operation-icon" v-else>
                <fold />
              </el-icon>
            </span>
            <!---->
            <el-popover
              placement="bottom"
              :width="300"
              trigger="hover"
              :show-after="300"
            >
              <template #reference>
                <el-icon class="note-operation-icon" ref="upload">
                  <info-filled />
                </el-icon>
              </template>
              <el-table :data="gridData" :show-header="false">
                <el-table-column width="100" property="name" label="name" />
                <el-table-column width="170" property="time" label="time" />
              </el-table>
            </el-popover>
            <el-icon
              class="note-operation-icon"
              title="历史记录"
              @click="openHistory"
            >
              <clock />
            </el-icon>
            <span></span>
            <el-icon
              class="note-operation-icon"
              @click="doDelete()"
              title="删除"
            >
              <delete />
            </el-icon>
          </div>
          <div class="note-operation-right">
            <el-button
              v-if="id == -2"
              type="danger"
              class="note-operation-icon"
              @click="doCancel()"
              size="small"
            >
              取消
            </el-button>
            <el-button
              type="primary"
              class="note-operation-icon"
              @click="doUpdateNote()"
              size="small"
            >
              保存
            </el-button>
            <div class="note-notebook">
              <el-icon class="note-notebook-icon">
                <notebook />
              </el-icon>
              <!-- <div class="note-tags">
              <el-icon class="note-tags-icon">
                <price-tag />
              </el-icon>
              <span class="note-tag"> {{ notebook }} </span>
            </div> -->
              <el-dropdown trigger="click" style="margin-top: 1px">
                <span class="el-dropdown-link">
                  {{ notebook }}
                  <el-icon class="el-icon--right">
                    <arrow-down />
                  </el-icon>
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item
                      v-for="notebook in notebooks"
                      :key="notebook.id"
                      @click="updateNotebook(notebook.id)"
                      >{{ notebook.title }}</el-dropdown-item
                    >
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
              <div class="note-label-split"></div>
            </div>
          </div>
        </div>
        <div class="note-label"></div>
      </div>
      <div class="note-title">
        <input
          class="note-title-input"
          v-model="titleInput"
          placeholder="请输入标题"
        />
      </div>
      <!-- 编辑器容器 -->
      <div id="editor" v-if="noteType == 1">
        <note-editor
          :value="value"
          v-on:inputData="inputData"
          v-on:onClickEidtor="onClickEidtor"
        ></note-editor>
      </div>
      <div id="editor" v-else-if="noteType == 2">
        <note-md :value="value" v-on:inputData="inputData"></note-md>
      </div>
    </div>
  </div>
  <!-- 历史记录-->
  <history ref="historyRef"></history>
</template>
<script>
import SiderBar from "@/components/SiderBar.vue";
import NoteMd from "@/views/note/NoteMD.vue";
import History from "@/views/note/History.vue";
import { GetNoteById, UpdateNote, CreateNote, DeleteNote } from "@/api/note";
import { ElMessage, ElMessageBox } from "element-plus";
import NoteEditor from "@/views/note/NoteEditor.vue";
import SearchSide from "@/views/search/SearchSide.vue";
import { mapState } from "vuex";
import { getFullDate, isHtmlLabel } from "@/utils/util";

export default {
  name: "SearchDetail",
  components: {
    SiderBar,
    SearchSide,
    NoteEditor,
    NoteMd,
    History,
  },
  data() {
    return {
      detailShow: false,
      id: -1,
      noteType: 1,
      titleInput: "",
      title: "",
      value: "",
      searchKey: "",
      searchValue: {},
      content: "",
      notebook: "",
      notebookID: -1,
      refresh: false,
      showHistory: false,
      highlightArray: [],
      gridData: [
        {
          name: "创建时间",
          time: "",
        },
        {
          name: "更新时间",
          time: "",
        },
      ],
      isCollapse: false,
    };
  },
  created() {
    this.titleInput = "";
    this.id = -2;
  },
  computed: {
    ...mapState({
      notebooks: (state) => state.notebook.notebooks,
    }),
  },
  methods: {
    // 切换笔记
    noteChange: async function (childValue) {
      console.log("noteChange");
      // 未保存笔记自动保存
      console.log(this.value);
      console.log(this.content);
      console.log(this.title);
      console.log(this.titleInput);
      if (this.value !== this.content || this.title != this.titleInput) {
        console.log("save here");
        await this.doUpdateNote();
        // this.noteSave = this.id;
        this.$refs.note.noteSave(this.id);
      }
      // 请求新的笔记
      if (typeof childValue != "undefined") {
        console.log("请求新的笔记");
        console.log(childValue.id);
        // 清空富媒体编辑框
        this.clearEditor(1);
        // 请求新的笔记
        this.title = this.titleInput = childValue.title;
        this.id = childValue.id;
        this.notebookID = Number(childValue.notebookID);
        this.noteType = childValue.type;
        console.log;
        this.gridData[0].time = getFullDate(childValue.createdAt); // 创建时间
        this.gridData[1].time = getFullDate(childValue.updatedAt); // 更新时间
        await this.GetNote(childValue.id); // 发送请求
      } else {
        this.titleInput = "";
        this.value = "";
        this.noteType = childValue.type;
        this.notebookID = Number(this.$route.params.id);
        this.setNotebookTitle(this.$route.params.id); // 根据ID获取Name
      }
    },
    inputData: function (inputData) {
      // childValue就是子组件传过来的值
      this.content = inputData;
    },
    clearSearch() {
      this.searchKey = "";
      this.$refs.search.focus();
    },
    onSearchFirst() {
      this.detailShow = true;
      this.searchValue = { searchKey: this.searchKey };
    },
    onSearch() {
      this.highlightArray = [];
      this.searchValue = { searchKey: this.searchKey };
    },
    onClickEidtor: function () {
      // 取消高亮关键词
      this.value = this.highlightKeyCancel(this.value, this.searchKey);
      this.$refs.upload.$el.click();
      this.setNotebookTitle(this.$route.params.id);
    },
    // 高亮关键词
    highlightKey(text, keyWord) {
      // 排除html标签
      if (isHtmlLabel(text, keyWord)) {
        return text;
      }
      var a = new RegExp(keyWord, "gi");
      return text.replace(a, (value) => {
        // 使用箭头函数才能获取this
        let temp =
          '<span id="highlight" style="border-style: solid;border-width:1px;border-color: rgb(255, 131, 29);\
          background-color: rgba(255, 170, 0, 0.34);">' +
          value +
          "</span>";
        this.highlightArray.push(temp);
        return temp;
      });
    },
    // 取消高亮关键词
    highlightKeyCancel(text, keyWord) {
      this.highlightArray.forEach((element) => {
        // 使用正则
        var a = new RegExp(
          '<span id="highlight"(.*)>(' + keyWord + ")</span>",
          "i"
        );
        text = text.replace(element, element.match(a)[2]);
      });
      return text;
    },
    // 切换焦点
    onClickEidtor: function () {
      this.$refs.upload.$el.click();
      this.setNotebookTitle(this.$route.params.id);
    },
    // 获取笔记
    async GetNote(noteId) {
      console.log("GetNote(noteId))");
      if (noteId == -2) {
        console.log("noteId == -2");
        console.log(this.content);
        setTimeout(() => {
          this.content = "";
          this.value = "";
        }, 100);
        if (this.$route.params.id == 0) {
          this.notebook = "笔记本";
          this.notebookID = 0;
        } else {
          this.setNotebookTitle(this.$route.params.id);
          this.notebookID = Number(this.$route.params.id);
        }
      } else {
        const res = await GetNoteById({ id: noteId });
        if (res.code === 200) {
          this.value = this.content = res.data.list[0].content;
          this.noteType = res.data.list[0].type;
          this.notebookID = res.data.list[0].notebookId;
          this.setNotebookTitle(res.data.list[0].notebookId);
        }
      }
    },
    // 输入值
    inputData: function (inputData) {
      this.content = inputData;
    },
    // 保存或新建笔记
    async doUpdateNote() {
      console.log("doUpdateNote()");
      // 保存笔记
      if (this.id != -2) {
        const res = await UpdateNote({
          id: this.id,
          title: this.titleInput,
          content: this.content,
          notebookId: this.notebookID,
        });
        if (res.code === 200) {
          ElMessage({
            showClose: true,
            message: res.msg,
            type: "success",
          });
          // TODO 更新指定note
          this.value = this.content;
          this.title = this.titleInput;
          this.$refs.note.noteSave(this.id);
        }
      } else if (this.id == -2) {
        // 新建笔记
        const res = await CreateNote({
          title: this.titleInput,
          content: this.content,
          notebookId: this.notebookID,
          type: this.noteType,
        });
        if (res.code === 200) {
          this.id = res.data.id;
          this.value = this.content;
          this.title = this.titleInput;
          this.$refs.note.noteSave(this.id);

          ElMessage({
            showClose: true,
            message: res.msg,
            type: "success",
          });
          if (this.$route.params.id == "add") {
            await router.push({
              name: "NoteDetail",
              params: { addNoteState: true },
            });
          }
          this.$refs.note.refresh();
        }
      }
    },
    // 笔记移动笔记本
    async updateNotebook(id) {
      this.notebookID = id;
      this.setNotebookTitle(id);
    },
    // 根据ID显示笔记本名称
    setNotebookTitle(id) {
      this.notebooks.forEach((notebook) => {
        if (notebook.id == id) {
          this.notebook = notebook.title;
        }
      });
    },
    // 取消新建笔记
    doCancel() {
      if (this.$route.params.id == -1) {
        router.push({ name: "NoteDetail", params: { id: 0 } });
      }
      this.id = -1;
      this.noteType = 1;
      this.$refs.note.refresh();
    },
    // 删除笔记
    doDelete() {
      ElMessageBox.confirm("是否删除?", "提示", {
        confirmButtonText: "是",
        cancelButtonText: "否",
        type: "warning",
      }).then(async () => {
        const res = await DeleteNote({ id: this.id });
        if (res.code === 200) {
          //   this.$refs.note.refresh();
          ElMessage({
            showClose: true,
            message: res.msg,
            type: "success",
          });
          this.$refs.note.delNote(this.id);
          // 清空编辑器标题与内容
          this.clearEditor();
        }
      });
    },
    // 笔记历史
    openHistory() {
      this.$refs.historyRef.openHistory({ id: this.id });
    },
    // 笔记列表伸缩
    totalCollapse() {
      this.isCollapse = !this.isCollapse;
    },
    // 清空富媒体编辑框
    clearEditor(type) {
      this.titleInput = "";
      this.notebookID = -1;
      if (type == 1 && this.noteType == 1) {
        console.log("run here");
        this.value = "&nbsp;";
        this.content = "&nbsp;";
      } else {
        this.value = "";
        this.content = "";
      }
    },
  },
};
</script>

<style lang="less" scoped>
.search-page {
  margin-left: 73px;
  .search-top {
    transition: padding 0.3s ease-in-out;
    padding: 121px 64px 25px;
    .search-input-box {
      position: relative;
      height: 52px;
      line-height: 45px;
      .search-input-text {
        background-color: transparent;
        margin: 0;
        padding: 0;
        border: none;
        border-width: 0;
        color: #606060;
        height: 52px;
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        font-size: 30px;
        font-weight: 500;
        letter-spacing: -0.3px;
        &:focus {
          outline: none;
        }
      }
      .search-input-icon {
        font-size: 30px;
        opacity: 0.2;
        overflow: hidden;
        position: absolute;
        top: 12px;
        right: 18px;
        cursor: pointer;
      }
    }
  }
  .search-input-footer {
    .search-footer-box {
      margin: 0 0 13px 64px;
      transition: margin 0.3s ease-in-out;
      .search-input-footer-text {
        color: #878787;
        display: inline-block;
        vertical-align: bottom;
        font-size: 13px;
        font-weight: 500;
      }
      .search-input-footer-tip {
        color: #2dbe60;
        display: inline-block;
        max-width: 500px;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
        vertical-align: bottom;
        font-size: 13px;
        font-weight: 500;
      }
    }
  }
}
.search-input-footer-split {
  margin: 0 64px 46px;
  border-bottom: 1px solid #ececec;
  transition: margin 0.3s ease-in-out;
}
.search-page-top {
  height: 12vh;
  width: 350px;
  padding: 20px;
  .search-top {
    padding: 0;
  }
  .search-input-footer {
    .search-footer-box {
      margin: 0;
    }
  }
}
.search-input-footer-split {
  margin: 15px;
}

.note-detail {
  margin-left: 423px;
  background: white;
  border-left: 1px solid #ececec;
  bottom: 0;
  left: 0;
  position: absolute;
  right: 0;
  top: 0;
  min-width: 403px;
  .note-operation {
    position: relative;
    padding: 12px 0;
    .note-operation-left {
      display: inline-block;
      padding-top: 2px;
      padding-left: 18px;
      opacity: 1;
      transition: opacity 0.2s ease-in-out;
      position: relative;
      .note-operation-icon {
        cursor: pointer;
        opacity: 0.4;
        font-size: 23px;
        color: black;
        margin: 0 16px 0 0;
        &:hover {
          opacity: 1;
          color: #2dbe60;
        }
      }
    }
    .note-operation-right {
      float: right;
      margin-left: 20px;
      .note-notebook {
        margin-top: 2px;
        float: right;
        .note-notebook-icon {
          margin-top: 2px;
          vertical-align: top;
          opacity: 0.5;
          font-size: 14px;
        }
        .el-dropdown-link {
          margin-left: 5px;
          font-size: 13px;
          font-weight: 400;
          color: #878787;
          .el-icon--right {
            opacity: 0.5;
            margin-left: 0;
            top: 2px;
          }
        }
        .note-label-split {
          border-left: 1px solid #ececec;
          height: 19px;
          display: inline-block;
          vertical-align: top;
          margin: 0 8px 0;
        }
      }
      .note-tags {
        margin: 0 16px 0 0;
        float: right;
        .note-tags-icon {
          opacity: 0.5;
        }
        .note-tag {
          float: right;
          margin-left: 5px;
          color: #fff;
          border-color: #2dbe60;
          background-color: #2dbe60;
          vertical-align: top;
          padding: 1px 8px 1px 6px;
          border-radius: 4px;
          font-size: 12px;
          font-weight: 400;
          line-height: 17px;
          text-align: center;
        }
      }
      .note-operation-icon {
        margin-right: 8px;
        float: right;
      }
    }
    .note-label {
      position: relative;
      max-width: 960px;
      min-width: 240px;
      padding: 0 48px;
      margin: 0 auto 38px;
      height: 35px;
      border-bottom: 1px solid #ececec;
    }
  }
  .note-header {
    height: 8vh;
  }
  .note-title {
    height: 7vh;
    .note-title-input {
      width: 100%;
      padding: 10px 20px;
      font-size: 25px;
      font-weight: 700;
      outline: none;
      border: none;
    }
  }
}
.example-showcase .el-dropdown-link {
  cursor: pointer;
  color: var(--el-color-primary);
  display: flex;
  align-items: center;
}
:deep(.el-input__wrapper) {
  border-radius: 100px;
}
</style>>
