<template>
  <footer class="footer">
    <div class="container content">
      <div>
        <nuxt-link to="/about">关于</nuxt-link>
        <a href="https://studio.txyl66.cn" target="_blank">工作室</a>
        <nuxt-link to="/tags">标签</nuxt-link>
        <a href="https://www.biqugesk.cc" target="_blank">文学</a>
        <nuxt-link to="/links">友链</nuxt-link>
      </div>
      <div>
        Powered by
        <a href="https://www.txyl66.cn" target="_blank" class="light">八斗米</a> © 2021-2022
        <a href="https://beian.miit.gov.cn/" target="_blank">粤ICP备2021098749号-1</a>
      </div>
    </div>
  </footer>
</template>

<script>
export default {}
</script>

<style lang="scss" scoped>
.footer {
  font-size: 14px;
  color: var(--text-color3);
  background: none;
  text-align: left;
  margin: 0 10px;
  a {
    color: var(--text-color3);
    text-decoration: none;
  }

  .light {
    color: #eb5424; // TODO
    font-weight: bold;
  }
}
</style>
