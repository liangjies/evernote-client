<template>
  <el-form size="large" class="login-content-form">
    <el-form-item class="login-animation1">
      <el-input
        type="text"
        placeholder="用户名"
        clearable
        autocomplete="off"
        v-model="username"
      >
        <template #prefix>
          <el-icon class="el-input__icon"><user /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation3">
      <el-col :xs="14" :sm="14" :md="15" :lg="15" :xl="15">
        <el-input
          type="text"
          placeholder="请输入邮箱"
          clearable
          autocomplete="off"
          v-model="email"
          ><template #prefix>
            <el-icon class="el-input__icon"><user /></el-icon>
          </template>
        </el-input>
      </el-col>
      <el-col :xs="0" :sm="0" :md="1" :lg="1" :xl="1"></el-col>
      <el-col :xs="8" :sm="8" :md="7" :lg="7" :xl="7">
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
      <el-col :span="24">
        <el-input
          type="text"
          placeholder="请输入验证码"
          clearable
          autocomplete="off"
          v-model="verify"
          ><template #prefix>
            <el-icon class="el-input__icon"><position /></el-icon>
          </template>
        </el-input>
      </el-col>
    </el-form-item>
    <el-form-item class="login-animation2">
      <el-input
        placeholder="密码"
        autocomplete="off"
        v-model="password"
        show-password
      >
        <template #prefix>
          <el-icon class="el-input__icon"><unlock /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <!-- <el-form-item class="login-animation3">
      <el-col :span="15">
        <el-input
          type="text"
          maxlength="4"
          placeholder="验证码"
          clearable
          autocomplete="off"
          v-model="login.captcha"
        >
          <template #prefix>
            <el-icon class="el-input__icon"><position /></el-icon>
          </template>
        </el-input>
      </el-col>
      <el-col :span="1"></el-col>
      <el-col :span="8">
        <div class="login-content-code">
          <img
            v-if="picPath"
            :src="picPath"
            alt="请输入验证码"
            @click="loginVerify()"
          />
        </div>
      </el-col>
    </el-form-item> -->
    <el-form-item class="login-animation4">
      <el-button
        type="primary"
        class="login-content-submit"
        round
        @click="onRegister"
      >
        <span>注册</span>
      </el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import { User, Unlock, Position } from "@element-plus/icons-vue";
import { captcha, verifyCode, register } from "@/api/user";
import { mapActions } from "vuex";
import { ElMessage } from "element-plus";
export default {
  name: "Register",
  components: { User, Unlock, Position },
  created() {
    // this.loginVerify();
  },
  data() {
    return {
      picPath: "",
      login: {
        captcha: "",
        captchaId: "",
        notice: "",
        isError: false,
      },
      username: "",
      password: "",
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
    ...mapActions("user", ["LoginIn"]),
    onRegister() {
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
      if (!/^[\w\u4e00-\u9fa5]{3,15}$/.test(this.username)) {
        ElMessage({
          showClose: true,
          message: "用户名3~15个字符，仅限于字母数字下划线中文",
          type: "error",
        });
        return;
      }
      this._onRegister();
    },
    async _onRegister() {
      const res = await register({
        email: this.email,
        userName: this.username,
        passWord: this.password,
        verifyCode: this.verify,
      });
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
        if (location.href.indexOf("#reloaded") == -1) {
          location.href = location.href + "#reloaded";
          location.reload();
        }
      }
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
      }, 1000);
    },
    loginVerify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.picPath;
        this.login.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="less">
.loop(@i) when (@i > 0) {
  .loop((@i - 1)); // 递归调用自身
  .login-animation@{i} {
    opacity: 1;
    animation-name: error-num;
    animation-duration: 0.5s;
    animation-fill-mode: forwards;
    animation-delay: calc(@i / 10) + s;
  }
}
.login-content-form {
  .loop(4);
  .login-content-password {
    display: inline-block;
    width: 20px;
    cursor: pointer;
    &:hover {
      color: #909399;
    }
  }
  .login-content-code {
    width: 100%;
    padding: 0;
    font-weight: bold;
    letter-spacing: 5px;
  }
  .login-content-submit {
    width: 100%;
    letter-spacing: 2px;
    font-weight: 300;
  }
}
</style>
