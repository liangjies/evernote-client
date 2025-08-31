<template>
  <div id="main">
    <div id="header">
      <div class="main" style="width: 1000px">
        <a href="/">
          <span id="header-logo" style="display: flex; justify-content: left; align-items: center">
            <img style="height: 35px; margin-right: 10px" src="@/common/images/logo.png" alt="" />
            <span style="font-size: 28px; width: 100px">爱笔记</span>
          </span>
        </a>
        <ul>
          <li id="navbar-homepage">
            <a href="/"><span class="is-active">首页</span></a>
          </li>
          <li id="navbar-opensource">
            <a href="https://github.com/wnasui/evernote-client" target="_blank"><span class="">开源地址 </span></a>
          </li>
        </ul>
        <a style="margin-left: auto" id="header-developer" href="/login"><span style="font-size: 16px">登录/注册</span></a>
      </div>
    </div>
    <div id="content">
      <div id="first_screen">
        <div id="main-content">
          <div class="text-group">
            <div class="weui-flex logo-box">
              <div class="weui-flex__item">
                <img class="logo" src="@/common/images/logo.png" alt="酷安" />
              </div>
            </div>
            <div class="weui-flex">
              <div class="weui-flex__item">
                <p class="title">开源、免费</p>
              </div>
              <div class="login-show">
                <a style="margin-left: auto" id="header-developer" href="/login"><span
                    style="font-size: 16px">登录/注册</span></a>
              </div>
              <div class="home-show">
                <img src="@/common/images/home_show.png" alt="Website Screenshot">
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style>
@import "@/common/css/index_v12.css";
</style>
<style>
body {
  background-color: #f5f5f5;
}

#footer {
  margin-top: 20px;
  z-index: 10;
  padding: 40px;
  display: block;
}

#footer .content {
  width: 1000px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
}

#footer .content .footer-left .footer-text-box {
  margin-left: 20px;
}

#footer .content .footer-left .footer-navbar {
  margin-top: 40px;
}

#footer .content .footer-left .footer-navbar a {
  font-size: 16px;
}

#footer .content .footer-left .icon-group {
  display: flex;
  justify-content: left;
}

#footer .content .footer-left .icon-group .footer-logo-box {
  width: 82px;
  height: 82px;
  background-color: #fff;
  border-radius: 20px;
}

#footer .content .footer-left .icon-group .footer-logo-box img {
  width: 100%;
}

#footer .content .footer-left .icon-group .footer-text-box {
  display: flex;
  justify-content: center;
  flex-direction: column;
  margin-left: 10px;
}

#footer .content .footer-left .icon-group .footer-text-box .title {
  font-size: 30px;
  color: #212121;
}

#footer .content .footer-left .icon-group .footer-text-box .sub-title {
  font-size: 20px;
  color: #757575;
}

#footer .content .footer-left .footer-info {
  font-size: 14px;
}

#footer .content .footer-right {
  display: flex;
  align-items: center;
  text-align: center;
}

#footer .content .footer-right img {
  border-radius: 5px;
}

.footer-navbar ul li {
  display: inline;
}

.footer-navbar ul li a {
  line-height: 2em;
  color: #212121;
  border-right: #212121 solid 1px;
  padding: 0 5px;
}

.footer-navbar .first {
  padding-left: 0;
}

.footer-navbar .last {
  border-right: 0;
}

.footer-info p {
  color: #212121;
  line-height: 2em;
}

.footer-info p img {
  height: 1.2em;
  margin-bottom: -2px;
}

.footer-qrcode-box {
  margin: 20px;
}

.footer-qrcode-box img {
  width: 120px;
  border: 1px solid #e6e6e6;
}

.footer-qrcode-box p {
  font-size: 15px;
  color: #212121;
}

/*动态列表样式*/
#feed-page {
  display: flex;
  width: 100%;
  justify-content: center;
}

#feed-content {
  width: 960px;
  display: flex;
  justify-content: space-between;
}

#feed-part {
  width: 632px;
}

#rank-part {
  width: 306px;
}

.feed-emoji {
  height: 18px;
  margin-bottom: -3px;
}

#rank-part .rank-list ul li {
  list-style-type: none;
  font-size: 14px;
  font-weight: 400;
  line-height: 20px;
  margin-top: 20px;
  display: flex;
  justify-content: space-between;
}

#rank-part .rank-list ul li .rank-group {
  border-radius: 4px;
  height: 20px;
  width: 20px;
  display: inline-block;
  vertical-align: top;
  margin-right: 8px;
  background-color: transparent;
  color: #212121;
}

#rank-part .rank-list ul li .rank-num {
  display: block;
  height: 20px;
  line-height: 20px;
  text-align: center;
}

#rank-part .rank-list ul li:nth-of-type(1) .rank-group {
  background: #f44336;
  color: #ffffff;
}

#rank-part .rank-list ul li:nth-of-type(2) .rank-group {
  background: #f57c00;
  color: #ffffff;
}

#rank-part .rank-list ul li:nth-of-type(3) .rank-group {
  background: #ffc107;
  color: #ffffff;
}

#rank-part .rank-list ul li .hot-group {
  min-width: 70px;
}

.home-show {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

</style>