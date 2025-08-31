<template>
  <el-upload
    class="avatar-uploader"
    :headers="headers"
    :action="UploadAvatar"
    :show-file-list="false"
    :on-success="handleAvatarSuccess"
    :before-upload="beforeAvatarUpload"
  >
    <img :src="headerImg" class="avatar" />
    <el-icon style="z-index: 1000" class="plus"><Plus /></el-icon>
  </el-upload>
</template>

<script setup>
import { ref } from "vue";
import { Plus } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { useStore } from "vuex";
import { computed } from "vue";
import { UploadAvatar } from "@/api/upload";
const store = useStore();
const headerImg = computed(() => store.state.user.userInfo.headerImg);
const headers = ref({ "x-token": store.getters["user/token"] });
const handleAvatarSuccess = (res, file) => {
  if (res.code === 200) {
    ElMessage({
      showClose: true,
      message: res.msg,
      type: "success",
    });
    store.commit("user/setHeaderImg", res.data.file.url);
  } else {
    ElMessage({
      showClose: true,
      message: res.msg,
      type: "error",
    });
  }
};
const beforeAvatarUpload = (file) => {
  console.log("beforeAvatarUpload");
  // 校验
  if (file.type === "image/jpeg" || file.type === "image/png") {
    if (file.size / 1024 / 1024 > 2) {
      ElMessage.error("Avatar picture size can not exceed 2MB!");
      return false;
    }
  } else {
    ElMessage.error("Avatar picture must be JPG format!");
    return false;
  }
  // 上传
};
</script>

<style lang="less" scoped>
.avatar-uploader .avatar {
  width: 50px;
  height: 50px;
  border-radius: 45px;
  overflow: hidden;
}
.plus {
  position: absolute;
  top: 0;
  right: 0;
  z-index: 1000;
  width: 50px;
  height: 50px;
  line-height: 50px;
  text-align: center;
  color: #fff;
  background-color: #409eff;
  border-radius: 50%;
  font-size: 20px;
  opacity: 0;
  &:hover {
    opacity: 1;
  }
}
</style>

<style>
.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>