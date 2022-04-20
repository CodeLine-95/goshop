<template>
  <div class="home">
    <Content></Content>
  </div>

</template>

<script lang="ts" setup>
import Content from '../Home/Content.vue'
</script>

<style>
.home{
  position: absolute;
  top: 60px;
  left: 200px;
  bottom: 0;
  z-index: 9999;
  background-color: #545c64;
  width: calc(100% - 200px);
}
</style>