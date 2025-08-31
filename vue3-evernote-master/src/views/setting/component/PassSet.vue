<template>
  <el-dialog v-model="dialogVisible" title="修改密码" :before-close="closeDialog">
    <p>原密码</p>
    <el-input size="large" v-model="oldPass" type="password"></el-input>
    <p>新密码</p>
    <el-input size="large" v-model="newPass1" type="password"></el-input>
    <p>确认新密码</p>
    <el-input size="large" v-model="newPass2" type="password"></el-input>
    <p class="tip-text">密码长度不少于6字节，由数字、字母及常用符号组成，字母区分大小写</p>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeDialog">取消</el-button>
        <el-button type="primary" @click="setPass">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script>
import { ChangePassword } from "@/api/user";
import { ElMessage } from "element-plus";
export default {
  name: "PassSet",
  components: {},
  data() {
    return {
      dialogVisible: false,
      oldPass: "",
      newPass1: "",
      newPass2: "",
    };
  },
  methods: {
    show() {
      this.dialogVisible = true;
    },
    async setPass() {
      // 校验
      if (this.newPass1 !== this.newPass2) {
        ElMessage({
          showClose: true,
          message: "确认新密码与新密码不相同",
          type: "error",
        });
        return;
      }
      // 校验
      if (!/^[\w]{6,25}$/.test(this.newPass1)) {
        ElMessage({
          showClose: true,
          message: "密码长度请在6-25字节",
          type: "error",
        });
        return;
      }

      const res = await ChangePassword({
        oldPass: this.oldPass,
        newPass: this.newPass1,
      });
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
        this.dialogVisible = false;
      }
    },
    closeDialog() {
      this.oldPass = "";
      this.newPass1 = "";
      this.newPass2 = "";
      this.dialogVisible = false;
    },
  },
};
</script>
<style lang="less">
.tip-text {
  color: #b2bccd;
  font-size: 12px;
  margin-top: 8px;
}
</style>
