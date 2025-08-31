<template>
  <md-editor
    v-model="text"
    @onChange="onChange"
    @onHtmlChanged="onHtmlChanged"
    :toolbarsExclude="toolbarsExclude"
    :autoDetectCode="true"
    :noPrettier="true"
    :showCodeRowNumber="true"
  />
</template>

<script>
export default {
  name: "NoteMd",
};
</script>
<script setup>
import { ref } from "vue";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import TurndownService from "turndown";
import { gfm, tables, strikethrough } from "turndown-plugin-gfm";
import {
  getCurrentInstance,
  onMounted,
  onUnmounted,
  computed,
  watch,
} from "vue";
const text = ref("");
const toolbarsExclude = ["github"];
var turndownService = new TurndownService({
  codeBlockStyle: "fenced",
});
turndownService.use(gfm);
turndownService.use([tables, strikethrough]);
turndownService.addRule("pre2Code", {
  filter: ["pre"],
  replacement(content) {
    const len = content.length;
    // 除了pre标签，里面是否还有code标签包裹，有的话去掉首尾的`（针对微信文章）
    const isCode = content[0] === "`" && content[len - 1] === "`";
    const result = isCode ? content.substr(1, len - 2) : content;
    //   let newresult = result.replace(/,/g,",\n")
    return "```js\n" + result + "\n```\n";
  },
});
const onPaste = (event) => {
  let data = event.clipboardData.getData("text/html");
  let htmlValue = turndownService.turndown(data);
  let value = text.value;
  if (htmlValue) {
    setTimeout(() => {
      if (text.value.indexOf(value) == -1) {
        return;
      }
      text.value = value + "\n" + htmlValue;
    }, 100);
  }
};
onMounted(() => {
  window.addEventListener("paste", onPaste);
});

onUnmounted(() => {
  window.removeEventListener("paste", onPaste);
});
const onChange = (v) => {
    // console.log(v);
};
const onHtmlChanged = (v) => {
    console.log(v);
};
</script>
<style lang="less" scoped>
#editor #md-editor-v3.md-editor {
  height: 85vh;
}
// .md-editor-footer {
//   height: 40px;
// }
</style>