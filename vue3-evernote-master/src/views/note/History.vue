<template>
  <div class="note-history">
    <el-dialog
      v-model="dialogFormVisible"
      width="98%"
      top="2vh"
      :show-close="false"
    >
      <div class="modal-body">
        <div class="history-list-wrap pull-left">
          <div class="history-list-header">
            历史记录 (<span class="history-num">{{ historyTotal }}</span
            >)
          </div>
          <div class="history-list list-group">
            <a
              class="list-group-item"
              v-for="(data, index) in historyData"
              :key="data.id"
              @click="clickSwitch(index)"
              :class="{ active: activeIndex == index }"
              ><span class="badge" v-if="index == 0">当前</span
              ><span class="badge" v-if="index != 0">#{{ data.version }}</span
              >{{ _formateDate(data.createdAt) }}</a
            >
          </div>
        </div>
        <div class="history-content-wrap pull-left">
          <div class="history-content-header">
            <el-button class="button-type" type="primary" @click="doRecover"
              >从该版本还原</el-button
            >
            <div class="dialog-close" @click="dialogFormVisible = false">
              <el-icon class="dialog-close-icon"><Close /></el-icon>
            </div>
          </div>
          <div class="history-content">
            <div v-html="activeContent"></div>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { GetHistories, RecoverHistory } from "@/api/history";
import { getFullDate } from "@/utils/util";
import { ElMessage } from "element-plus";
export default {
  name: "History",
  data() {
    return {
      dialogFormVisible: false,
      historyData: [],
      historyTotal: 0,
      activeIndex: 0,
      activeContent: "",
    };
  },
  created() {},
  methods: {
    // 根据ID获取历史记录
    async getHistories(id) {
      const res = await GetHistories({ id: id });
      if (res.code === 200 && res.data.total > 0) {
        this.historyData = res.data.list;
        this.historyTotal = res.data.total;
        this.activeContent = this.historyData[0].content;
      }
    },
    // 格式化日期
    _formateDate(dateStr) {
      if (dateStr == "") {
        return "";
      } else {
        return getFullDate(dateStr);
      }
    },
    // 切换历史内容
    clickSwitch(index) {
      this.activeContent = this.historyData[index].content;
      this.activeIndex = index;
    },
    // 还原笔记历史版本
    async doRecover() {
      if (this.activeIndex == 0) {
        ElMessage({
          showClose: true,
          message: "已是当前版本无需还原",
          type: "error",
        });
        return;
      }
      const res = await RecoverHistory(this.historyData[this.activeIndex]);
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
      }
    },
    // 清空显示框
    _cleanForm() {
      this.historyData = [];
      this.historyTotal = 0;
      this.activeIndex = 0;
      this.activeContent = "";
    },
    // 点击历史记录
    openHistory(data) {
      this._cleanForm();
      this.dialogFormVisible = true;
      this.getHistories(data.id);
    },
  },
};
</script>

<style lang="less" scoped>
.modal-body {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 0;
  .history-list-wrap {
    position: absolute;
    top: 0;
    bottom: 0;
    width: 250px;
    border-right: 1px solid #ccc;
    box-shadow: 1px 1px 10px #ddd;
    .history-list-header {
      line-height: 50px;
      font-size: 16px;
      font-weight: bold;
      padding-left: 15px;
      border-bottom: 1px solid #eee;
      background-color: #ccc;
    }
    .history-list {
      position: absolute;
      top: 51px;
      bottom: 0;
      left: 0;
      right: 0;
      overflow-y: auto;
      margin-bottom: 0;
      .list-group-item {
        cursor: pointer;
        position: relative;
        display: block;
        padding: 10px 15px;
        margin-bottom: -1px;
        background-color: #fff;
        border: 1px solid #ddd;
        .badge {
          min-width: 10px;
          padding: 3px 7px;
          font-size: 12px;
          color: #fff;
          background-color: #999;
          border-radius: 10px;
          float: right;
        }
        &:hover {
          color: #000;
          background-color: #eee;
          border-color: #eee;
        }
      }
      .active {
        color: #000;
        background-color: #eee;
        border-color: #eee;
        .badge {
          color: #428bca;
          background-color: #fff;
        }
      }
    }
  }
  .history-content-wrap {
    position: absolute;
    top: 0;
    bottom: 0;
    right: 0;
    left: 250px;
    .history-content-header {
      height: 51px;
      border-bottom: 1px solid #eee;
      box-shadow: 5px 0px 5px #ccc;
      .button-type {
        margin-left: 10px;
        margin-top: 5px;
      }
      .dialog-close {
        position: absolute;
        top: 0px;
        right: 0;
        padding: 0;
        width: 54px;
        height: 54px;
        background: 0 0;
        border: none;
        outline: 0;
        cursor: pointer;
        .dialog-close-icon {
          font-size: 16px;
          position: absolute;
          top: 50%;
          left: 50%;
          transform: translate(-50%, -50%);
        }
        &:hover {
          .dialog-close-icon {
            color: #409eff;
          }
        }
      }
    }
    .history-content {
      position: absolute;
      top: 51px;
      bottom: 0;
      right: 0;
      left: 0;
      padding-top: 5px;
      padding-right: 5px;
      padding-left: 10px;
      overflow-y: auto;
    }
  }
}
</style>
<style lang="less">
ol {
  padding-left: 20px;
}
.note-history {
  .el-dialog {
    height: 95%;
  }
}
</style>