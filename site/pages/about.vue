<template>
  <section class="main">
    <div class="container">
      <div class="widget">
        <div class="widget-header">关于</div>
        <div class="widget-content content">
          <h2>简介</h2>
          <p>
            BBS-GO&nbsp;是一款基于Go语言研发的精美小巧的开源社区论坛系统。初期该项目仅用过学习和交流，开源之后越来越多的小伙伴儿开始喜欢和关注他，这也是我长期升级和维护的动力。
          </p>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  head() {
    return {
      title: this.$siteTitle('关于'),
    }
  },
}
</script>

<style lang="scss" scoped></style>
