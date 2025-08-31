<template>
  <el-form size="large" class="login-content-form">
    <el-form-item class="login-animation1">
      <el-input
        type="text"
        placeholder="用户名"
        clearable
        autocomplete="off"
        v-model="login.username"
      >
        <template #prefix>
          <el-icon class="el-input__icon"><user /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <el-form-item class="login-animation2">
      <el-input
        placeholder="密码"
        autocomplete="off"
        v-model="login.password"
        show-password
      >
        <template #prefix>
          <el-icon class="el-input__icon"><unlock /></el-icon>
        </template>
      </el-input>
    </el-form-item>
    <!-- <el-form-item class="login-animation3" @keyup.enter="onSignIn">
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
        @click="onSignIn"
        :disabled="btlStatus"
      >
        <span>登录</span>
      </el-button>
    </el-form-item>
  </el-form>
</template>
<script>
import { User, Unlock, Position } from "@element-plus/icons-vue";
import { captcha } from "@/api/user";
import { mapActions } from "vuex";
export default {
  name: "Account",
  components: { User, Unlock, Position },
  created() {
    // this.loginVerify();
  },
  data() {
    return {
      btlStatus: false,
      picPath: "",
      login: {
        username: "demo",
        password: "demo",
        // captcha: "",
        // captchaId: "",
        ticket: "",
        randstr: "",
        notice: "",
        isError: false,
      },
    };
  },
  methods: {
    ...mapActions("user", ["LoginIn"]),
    onSignIn() {
      var captcha = new TencentCaptcha("2046626881", this._onSignIn);
      captcha.show();
    },
    async _onSignIn(e) {
      console.log("腾讯验证码回调", e);
      
      // 检查验证码回调结果
      if (!e.ticket || !e.randstr) {
        console.error("验证码参数缺失", e);
        this.btlStatus = false;
        return;
      }
      
      this.login.ticket = e.ticket;
      this.login.randstr = e.randstr;
      this.btlStatus = true;
      
      console.log("发送登录请求", this.login);
      const data = await this.LoginIn(this.login);
      console.log("登录响应", data);
      
      // 验证码错误刷新验证码
      if (typeof data.code != "undefined" && data.code != 200) {
        this.btlStatus = false;
        console.error("登录失败", data);
        //   this.loginVerify();
      }
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
  margin-top: 20px;
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
    margin-top: 15px;
  }
}
</style>
