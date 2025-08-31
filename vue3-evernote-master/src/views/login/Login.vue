<template>
  <div class="login-container">
    <div class="login-icon-group">
      <div class="login-icon-group-title">
        <img src="@/assets/logo-mini.1646368269034.svg" />
        <div class="login-icon-group-title-text font25">笔记</div>
      </div>
      <img
        src="@/assets/login-icon-two.1646368269034.svg"
        class="login-icon-group-icon"
      />
    </div>
    <div class="login-content">
      <div class="login-content-main">
        <h4 class="login-content-title ml15">登录</h4>
        <div>
          <el-tabs v-model="tabsActiveName">
            <el-tab-pane label="登录" name="account">
              <account></account>
            </el-tab-pane>
            <el-tab-pane label="注册" name="register">
              <Register />
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { register } from "@/api/user";
import { ElMessage } from "element-plus";
import Account from "@/views/login/component/Account.vue";
import Register from "@/views/login/component/Register.vue";
export default {
  name: "Login",
  components: {
    Account,
    Register,
  },
  data() {
    return {
      tabsActiveName: "account",
      isShowLogin: true,
      isShowRegister: false,
      register: {
        username: "",
        password: "",
        notice: "",
        isError: false,
      },
      picPath: "",
    };
  },
  created() {},
  methods: {
    ...mapActions("user", ["LoginIn"]),
    showRegister() {
      this.isShowLogin = false;
      this.isShowRegister = true;
    },
    showLogin() {
      this.isShowLogin = true;
      this.isShowRegister = false;
    },
    async onLogin() {
      if (!/^[\w\u4e00-\u9fa5]{3,15}$/.test(this.login.username)) {
        this.login.isError = true;
        this.login.notice = "用户名3-15个字符，仅限字母数字下划线中文";
        return;
      }
      // if (!/^.{6,16}$/.test(this.login.password)) {
      //   this.login.isError = true;
      //   this.login.notice = "密码长度为6~16个字符";
      //   return;
      // }
      return await this.LoginIn(this.login);
    },
    async onRegister() {
      if (!/^[\w\u4e00-\u9fa5]{3,15}$/.test(this.register.username)) {
        this.register.isError = true;
        this.register.notice = "用户名3~15个字符，仅限于字母数字下划线中文";
        return;
      }
      // if (!/^.{6,16}$/.test(this.register.password)) {
      //   this.register.isError = true;
      //   this.register.notice = "密码长度为6~16个字符";
      //   return;
      // }
      const res = await register({
        username: this.register.username,
        password: this.register.password,
      });
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
        this.$store.commit("setUsername", this.login.username);
        this.$store.commit("setToken", res.data.token);
        this.$store.commit("changeIsLogin", true);
        this.$router.push({
          path: "/Note",
        });
      }
    },
  },
};
</script>

<style lang="less" scoped>
#login {
  position: absolute;
  width: 100%;
  min-height: 100%;
  background-color: #f3f3f3;
  z-index: 100;
  display: flex;
  // flex-direction: column;
  .all {
    margin: 120px auto;
    .icons {
      display: flex;
      justify-content: center;
      font-size: 60px;
      margin: 20px auto 20px auto;
    }
    .con {
      background-color: #fff;
      padding-bottom: 20px;
      margin: auto;
      max-width: 400px;
      // height: 400px;
      border: 1px solid #dedede;
      border-radius: 2px;
      p {
        font-size: 12px;
        margin-left: 25px;
      }

      h3 {
        margin-top: 20px;
        font-size: large;
        text-align: center;
        cursor: pointer;
      }
      .login,
      .register {
        height: 0;
        overflow: hidden;
        transition: height 0.4s;
        margin-top: 20px;
        .error {
          color: red;
        }
        .input,
        .button {
          width: 350px;
          margin: 10px 25px;
        }
        .button {
          margin-top: 20px;
          background-color: #2dbe60;
          color: #fff;
          border-radius: 5px;
          border: none;
          font-size: medium;
          &:hover {
            background-color: #23a04f;
            border-color: #23a04f;
          }
        }
        &.show {
          height: 210px;
        }
      }
    }
  }
}
</style>
<style scoped lang="less">
.login-container {
  width: 100%;
  height: 100%;
  position: relative;
  background: var(--el-color-white);
  .login-icon-group {
    width: 100%;
    height: 100%;
    position: relative;
    .login-icon-group-title {
      position: absolute;
      top: 50px;
      left: 80px;
      display: flex;
      align-items: center;
      img {
        width: 30px;
        height: 30px;
      }
      &-text {
        padding-left: 15px;
        color: var(--el-color-primary);
      }
    }
    &::before {
      content: "";
      position: absolute;
      bottom: 0;
      left: 0;
      width: 60%;
      overflow: hidden;
      height: 80%;
      -webkit-mask-box-image: url("data:image/svg+xml,%3Csvg width='1200' height='770' xmlns='http://www.w3.org/2000/svg' fill='none'%3E%3Cg%3E%3Cpath id='svg_1' d='M58.4 47.77C104.6 59.51 135.26 67.37 162.11 78.04C188.97 88.72 226.33 102.69 265.92 123.55C305.51 144.4 366.96 167.09 441.43 121.52C515.9 75.95 546.48 61.01 577.69 46.27C608.9 31.53 625.86 23.69 680.26 12.28C734.65 0.87 837.29 10.7 867.29 21.8C897.29 32.9 935.51 51.9 962.21 95.45C988.9 139.01 972.91 177.36 951.37 221.39C929.83 265.43 883.49 306 890.44 337.33C897.4 368.66 974.73 412.18 974.73 411.47C974.73 412.18 1066.36 457.62 1106.36 491.06C1146.36 524.5 1178.8 563.36 1184.03 579.63C1189.26 595.9 1200.4 622.49 1181.55 676.88C1162.71 731.26 1127.16 764.32 1115.31 778.64C1103.45 792.96 5.34 783.61 4.32 784.63C3.3 785.65 -172.34 2.38 1.13 35.04L58.4 47.77L58.4 47.77Z' fill='%23409eff'%2F%3E%3C%2Fg%3E%3C%2Fsvg%3E");
      background: var(--el-color-primary-light-5);
      transition: all 0.3s ease;
    }
    &::after {
      content: "";
      width: 150px;
      height: 300px;
      position: absolute;
      right: 0;
      top: 0;
      -webkit-mask-box-image: url("data:image/svg+xml,%3Csvg width='150' height='300' xmlns='http://www.w3.org/2000/svg' fill='none'%3E%3Cg%3E%3Cpath id='svg_1' d='M-0.56 -0.28C41.94 36.17 67.73 18.94 93.33 33.96C118.93 48.98 107.58 73.56 101.94 89.76C96.29 105.96 50.09 217.83 47.87 231.18C45.64 244.52 46.02 255.2 64.4 270.05C82.79 284.91 121.99 292.31 111.98 289.81C101.97 287.32 153.96 301.48 151.83 299.9C149.69 298.32 149.98 -1.36 149.71 -1.18C149.98 -1.36 -43.06 -36.74 -0.56 -0.28L-0.56 -0.28Z' fill='%23409eff'%2F%3E%3C%2Fg%3E%3C%2Fsvg%3E");
      background: var(--el-color-primary-light-5);
      transition: all 0.3s ease;
    }
    &-icon {
      width: 60%;
      height: 70%;
      position: absolute;
      left: 0;
      bottom: 0;
    }
  }
  @media screen and (min-width: 720px) {
    .login-content {
      position: absolute;
      width: 500px;
      padding: 20px;
      right: 200px;
      top: 50%;
    }
  }
  @media screen and (max-width: 720px) {
    .login-content {
      position: absolute;
      width: 95%;
      left: 0; 
      right: 0;
      margin:0 auto;
      top: 50%;
    }
  }
  .login-content {
    transform: translateY(-50%) translate3d(0, 0, 0);
    background-color: var(--el-color-white);
    border: 5px solid var(--el-color-primary-light-8);
    border-radius: 5px;
    overflow: hidden;
    z-index: 1;
    height: 460px;
    .login-content-main {
      margin: 0 auto;
      width: 80%;
      .login-content-title {
        color: var(--el-text-color-primary);
        font-weight: 500;
        font-size: 22px;
        text-align: center;
        letter-spacing: 4px;
        margin: 10px 0 10px;
        white-space: nowrap;
        z-index: 5;
        position: relative;
        transition: all 0.3s ease;
      }
    }
    .login-content-main-sacn {
      position: absolute;
      top: 0;
      right: 0;
      width: 50px;
      height: 50px;
      overflow: hidden;
      cursor: pointer;
      transition: all ease 0.3s;
      color: var(--el-text-color-primary);
      &-delta {
        position: absolute;
        width: 35px;
        height: 70px;
        z-index: 2;
        top: 2px;
        right: 21px;
        background: var(--el-color-white);
        transform: rotate(-45deg);
      }
      &:hover {
        opacity: 1;
        transition: all ease 0.3s;
        color: var(--el-color-primary) !important;
      }
      i {
        width: 47px;
        height: 50px;
        display: inline-block;
        font-size: 48px;
        position: absolute;
        right: 2px;
        top: -1px;
      }
    }
  }
}
</style>