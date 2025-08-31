<template>
  <div class="sider-bar" id="sidebar">
    <div class="sidebar-con">
      <div class="login">
        <img src="/src/common/images/yinxiang.png" alt="" class="user-image" />
        <div class="icons">
          <el-tooltip
            class="box-item"
            effect="dark"
            content="新建"
            placement="right-start"
          >
            <div class="addNote show">
              <el-popover
                placement="right"
                :width="120"
                trigger="click"
                :hide-after="0"
              >
                <template #reference>
                  <el-icon>
                    <plus />
                  </el-icon>
                </template>
                <ul class="file-type-list">
                  <li class="can-disabled" @click="addNote">
                    <img
                      class="dropdown-img"
                      src="/src/common/images/documents.png"
                    /><span>空白文档</span>
                  </li>
                  <li class="can-disabled" @click="addNote('md')">
                    <img
                      class="dropdown-img"
                      src="/src/common/images/md.png"
                    /><span>Markdown</span>
                  </li>
                </ul>
              </el-popover>
            </div>
          </el-tooltip>
          <el-tooltip
            class="box-item"
            effect="dark"
            content="搜索"
            placement="right-start"
          >
            <router-link to="/Search">
              <div class="searchNote show">
                <el-icon>
                  <search />
                </el-icon>
              </div>
            </router-link>
          </el-tooltip>
          <el-tooltip
            class="box-item"
            effect="dark"
            content="笔记"
            placement="right-start"
          >
            <router-link to="/Note">
              <el-icon class="note el-icon-s-order">
                <document />
              </el-icon>
            </router-link>
          </el-tooltip>
          <el-tooltip
            class="box-item"
            effect="dark"
            content="笔记本"
            placement="right-start"
          >
            <div title="笔记本" @click.stop="doClickOpenNotebook">
              <el-icon
                class="book el-icon-s-platform"
                :class="{ 'note-book-open': notebookOpen }"
              >
                <notebook />
              </el-icon>
            </div>
          </el-tooltip>
          <el-tooltip
            class="box-item"
            effect="dark"
            content="回收站"
            placement="right-start"
          >
            <router-link to="/NoteTrash">
              <el-icon class="note el-icon-s-order">
                <delete />
              </el-icon>
            </router-link>
          </el-tooltip>
          <div class="userMenu" @click.stop="open = !open">
            <img :src="headerImg" alt="" class="userMenus-image" />
          </div>
        </div>
      </div>
      <div class="logout" @click="onLogout">
        <i class="el-icon-switch-button"></i>
      </div>
    </div>
    <div v-if="open">
      <user-pop v-clickoutside="handleClickOutside"></user-pop>
    </div>
    <div v-if="notebookOpen">
      <notebook-pop
        v-clickoutside="notebookClickOutside"
        v-on:PopOpenValue="PopOpenValue"
      ></notebook-pop>
    </div>
  </div>
</template>

<script>
import {
  Plus,
  Document,
  Notebook,
  Search,
  Delete,
} from "@element-plus/icons-vue";
import UserPop from "./UserPop.vue";
import NotebookPop from "./NotebookPop.vue";
import clickoutside from "../utils/click-outside";
import { mapActions, mapState } from "vuex";
import router from "@/router/index";

export default {
  directives: { clickoutside },
  name: "SiderBar",
  emits: ["addNoteValue"],
  components: {
    UserPop,
    NotebookPop,
    Plus,
    Document,
    Notebook,
    Search,
    Delete,
  },
  created() {
    console.log(this.$parent)
    this.GetNotebooksData();
  },
  data() {
    return {
      isCollapse: false,
      open: false,
      notebookOpen: false,
      showAdd: false,
    };
  },
  computed: {
    ...mapState({
      headerImg: (state) => state.user.userInfo.headerImg,
    }),
  },
  methods: {
    ...mapActions("notebook", ["GetNotebooksData"]),
    PopOpenValue: function () {
      this.notebookOpen = false;
    },
    handleClickOutside() {
      this.open = false;
      console.log("click outside");
    },
    notebookClickOutside() {
      this.notebookOpen = false;
    },
    doClickOpenNotebook() {
      this.notebookOpen = !this.notebookOpen;
    },
    addNote(value) {
      if (this.$route.path.split("/")[1] != "Note") {
        router.push({ name: "NoteDetail", params: { id: "add" } });
      } else {
        this.$emit("addNoteValue", value);
      }
    },
  },
};
</script>

<style scoped lang="less">
#sidebar {
  width: 73px;
  background-color: #f8f8f8;
  z-index: 99;
  height: 100%;
  position: absolute;
  .sidebar-con {
    background-color: #f8f8f8;
    border-right: 1px solid #ececec;
    min-height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    .login {
      display: flex;
      flex-direction: column;
      align-items: center;
      .user-image {
        display: inline-block;
        width: 36px;
        height: 36px;
        margin-top: 12px;
        margin-bottom: 30px;
        //position: relative;
        text-align: center;
        outline: none;
      }
      .icons {
        margin-top: 10px;

        // display: flex;
        // flex-direction: column;
        .searchNote,
        .book,
        .note,
        .trash,
        .addNote {
          position: relative;
          cursor: pointer;
          font-size: 20px;
          display: flex;
          justify-content: center;
          align-items: center;
          border-radius: 50%;
          width: 35px;
          height: 35px;
          // padding: 6px 0;
          margin: 15px;
          color: #2dbe60;
          border: 1px solid #f8f8f8;
          transition: all 250ms;
          &:hover {
            border: 1px solid #2dbe60;
            background-color: #2dbe60;
            color: #fff;
          }
        }
        .addNote .searchNote {
          background-color: #fff;
          border: 1px solid #ccc;
          margin-bottom: 40px;
        }
        .searchNote {
          margin-bottom: 56px;
        }
        .userMenu {
          cursor: pointer;
          border-top: 1px solid #e1e1e1;
          position: absolute;
          bottom: 0;
          left: 0;
          right: 0;
          padding: 17px;
          .userMenus-image {
            width: 36px;
            height: 36px;
            border-radius: 18px;
          }
        }
        .note-book-open {
          border: 1px solid #2dbe60;
          background-color: #2dbe60;
          color: #fff;
        }
      }
      .icons .show {
        // color:#fff;
        border: 1px solid #2dbe60;
        background-color: #fff;
      }
    }
    .logout {
      cursor: pointer;
      margin: 15px auto;
      font-size: 20px;
      width: 35px;
      height: 35px;
      color: #2dbe60;
      // background: yellow;
      display: flex;
      i {
        margin: auto;
      }
    }
  }
}

.el-popper .file-type-list {
  li {
    cursor: pointer;
    list-style: none;
    padding: 10px;
    border-radius: 4px;
    &:hover {
      background-color: #ecf5ff;
      color: #409eff;
    }

    span {
      color: #4e5a70;
      font-size: 14px;
      vertical-align: middle;
    }
    img {
      width: 20px;
      margin-right: 5px;
      vertical-align: middle;
    }
  }
}
</style>