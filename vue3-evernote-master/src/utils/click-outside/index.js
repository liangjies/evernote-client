
const clickoutside = {
    mounted(el, binding) {
      function documentHandler(e) {
        if (el.contains(e.target)) {
          return false
        }
        if (binding.value) {
          binding.value(e)
        }
      }
      el.__vueClickOutside__ = documentHandler
      document.addEventListener('click', documentHandler)
    },
    beforeUnmount(el) {
        console.log("beforeUnmount")
      // 取消事件监听
      document.removeEventListener('click', el.__vueClickOutside__)
      delete el.__vueClickOutside__
    },
}
export default clickoutside;