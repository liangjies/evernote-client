<template>
  <el-dialog v-model="dialogVisible" title="修改昵称" :before-close="closeDialog">
    <el-input size="large" :value="nickName"></el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="changeNickName">确定</el-button>
      </span>
    </template>
  </el-dialog>
</template>
<script>
import { UpdateNickName } from "@/api/user";
import { ElMessage } from "element-plus";
export default {
  name: "NickNameSet",
  props: ["nickName"],
  components: {},
  data() {
    return {
      dialogVisible: false,
    };
  },
  methods: {
    async changeNickName() {
      // 校验
      if (!/^[\w\u4e00-\u9fa5]{1,15}$/.test(this.nickName)) {
        ElMessage({
          showClose: true,
          message: "昵称1-15个字符，仅限字母数字下划线中文",
          type: "error",
        });
        return;
      }
      const res = await UpdateNickName({ nickName: this.nickName });
      if (res.code === 200) {
        ElMessage({
          showClose: true,
          message: res.msg,
          type: "success",
        });
        this.$store.commit("user/setNickName", this.nickName);
        this.dialogVisible = false;
      }
    },
    show() {
      this.dialogVisible = true;
    },
    closeDialog(){
        this.nickName = "";
    }
  },
};
</script>
<style lang="less">
</style>
