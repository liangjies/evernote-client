<template>
  <sider-bar></sider-bar>
</template>

<script>
import SiderBar from "@/components/SiderBar.vue";
export default {
  name: "HomeHeader",
  components: {
    SiderBar,
  },
  created() {
    router.push({ name: "NoteDetail", params: { id: 0 } });
  },
  data() {
    return {};
  },
  computed: {},
  methods: {},
};
</script>

<style lang="less" scoped>
.home {
  margin-left: 240px;
  // .info-table
  // margin 30px
}
</style>
