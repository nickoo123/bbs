<template>
  <section class="main">
    <div class="container">
      <div class="widget">
        <div class="widget-header">关于</div>
        <div class="widget-content content">
          <h2>简介</h2>
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
