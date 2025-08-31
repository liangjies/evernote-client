<template>
  <el-dialog
    v-model="dialogVisible"
    title="更换邮箱"
    :before-close="closeDialog"
  >
    <el-form-item class="login-animation3">
      <el-col :span="3">
        <span>密码</span>
      </el-col>
      <el-col :span="21">
        <el-input
          type="password"
          placeholder="请输入密码"
          clearable
          autocomplete="off"
          v-model="pass"
        >
        </el-input>
      </el-col>
    </el-form-item>
    <el-form-item class="login-animation3">
      <el-col :span="3">
        <span>邮箱</span>
      </el-col>
      <el-col :span="13">
        <el-input
          type="text"
          placeholder="请输入邮箱"
          clearable
          autocomplete="off"
          v-model="email"
        >
        </el-input>
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="5">
        <el-button
          type="warning"
          :loading="loading"
          :disabled="disabled"
          @click="getCheckCode"
        >
          {{ text }}
        </el-button>
      </el-col>
    </el-form-item>

    <el-form-item class="login-animation3">
      <el-col :span="3">
        <span>验证码</span>
      </el-col>
      <el-col :span="20">
        <el-input
          type="text"
          placeholder="请输入验证码"
          clearable
          autocomplete="off"
          v-model="verify"
        >
        </el-input>
      </el-col>
    </el-form-item>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="closeDialog">取消</el-button>
        <el-button type="primary" @click="changeEmail">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script>
import { UpdateEmail, verifyCode } from "@/api/user";
import { ElMessage } from "element-plus";
export default {
  name: "EmailSet",
  components: {},
  data() {
    return {
      dialogVisible: false,
      pass: "",
      email: "",
      verify: "",
      text: "获取验证码",
      loading: false,
      disabled: false,
      duration: 60,
      timer: null,
    };
  },
  methods: {
    async changeEmail() {
      // 校验密码
      if (!/^[\w]{6,25}$/.test(this.pass)) {
        ElMessage({
          showClose: true,
          message: "密码长度请在6-25字节",
          type: "error",
        });
        return;
      }
      // 校验邮箱
      if (
        /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/.test(
          this.pass
        ) ||
        this.email.length == 0
      ) {
        ElMessage({
          showClose: true,
          message: "请输入正确的邮箱",
          type: "error",
        });
        return;
      }
      // 校验验证码
      if (this.verify.length == 0) {
        ElMessage({
          showClose: true,
          message: "请输验证码",
          type: "error",
        });
        return;
      }
      const res = await UpdateEmail({
        password: this.pass,
        email: this.email,
        verifyCode: this.verify,
      });
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
        this.$store.commit("user/setEmail", this.email);
        this.dialogVisible = false;
      }
    },
    show() {
      this.dialogVisible = true;
    },
    closeDialog() {
      this.pass = "";
      this.email = "";
      this.dialogVisible = false;
    },
    getCheckCode() {
      if (
        /\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*/.test(this.email) ==
          false ||
        this.email.length == 0
      ) {
        ElMessage({
          showClose: true,
          message: "请输入正确的邮箱",
          type: "error",
        });
        return;
      }
      var captcha = new TencentCaptcha("2046626881", this._getCheckCode);
      captcha.show();
    },

    async _getCheckCode(e) {
      if (e.ticket == "") {
        return;
      }
      const res = await verifyCode({
        email: this.email,
        ticket: e.ticket,
        randstr: e.randstr,
      });
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
      } else {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "error",
        });
      }
      // 倒计时期间按钮不能单击
      if (this.duration !== 10) {
        this.disabled = true;
      }
      // 清除掉定时器
      this.timer && clearInterval(this.timer);
      // 开启定时器
      this.timer = setInterval(() => {
        const tmp = this.duration--;
        this.text = `${tmp}秒`;
        if (tmp <= 0) {
          // 清除掉定时器
          clearInterval(this.timer);
          this.duration = 10;
          this.text = "重新获取";
          // 设置按钮可以单击
          this.disabled = false;
        }
        console.info(this.duration);
      }, 1000);
    },
  },
};
</script>
<style lang="less">
.content {
  .content-box {
    margin-top: 32px;
    .title {
      color: #b2bccd;
      font-size: 12px;
    }
    .email-split {
      margin-top: 16px;
      display: flex;
      min-height: 24px;
      align-items: center;
      &:hover {
        .email-button {
          opacity: 1;
        }
      }
      .email-content {
        color: #07142a;
        font-size: 14px;
        line-height: 1;
      }
      .email-button {
        border: none;
        cursor: pointer;
        opacity: 0;
        outline: none;
        padding: 0;
        transition: opacity 0.2s ease-in;
        margin-left: 16px;
        background-color: transparent;
        color: #448aff;
        font-size: 14px;
      }
    }
  }
  .content-box {
    .split {
      margin-top: 16px;
      display: flex;
      min-height: 24px;
      align-items: center;
      &:hover {
        .button {
          opacity: 1;
        }
      }
      .content {
        color: #07142a;
        font-size: 14px;
        line-height: 1;
      }
      .button {
        border: none;
        cursor: pointer;
        opacity: 0;
        outline: none;
        padding: 0;
        transition: opacity 0.2s ease-in;
        margin-left: 16px;
        background-color: transparent;
        color: #448aff;
        font-size: 14px;
      }
    }
  }
  .content-avatar {
    top: 50px;
    right: 10px;
    width: 50px;
    height: 50px;
    overflow: hidden;
    position: absolute;
    border-radius: 45px;
    .avatar {
      width: 50px;
      height: 50px;
      img {
        color: transparent;
        width: 100%;
        height: 100%;
        object-fit: cover;
        text-align: center;
        text-indent: 10000px;
      }
    }
  }
}
.custom-dialog {
  .el-dialog__header {
    padding: 0 0 0 0;
  }
  .el-dialog__body {
    padding: 0 0 0 0;
  }
  .el-dialog__footer {
    padding: 0 0 0 0;
  }
}
.el-tabs--right .el-tabs__content,
.el-tabs--left .el-tabs__content {
  height: 100%;
}
</style>
