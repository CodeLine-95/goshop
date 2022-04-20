<template>
<!-- 侧边栏 -->
  <div class="menu-side">
    <el-menu active-text-color="#ffd04b" background-color="#545c64" class="el-menu-vertical-demo"
             default-active="1" text-color="#fff" :collapse="isCollapse" @open="handleOpen" @close="handleClose">
      <el-menu-item index="1">
        <el-icon><icon-menu /></el-icon>
        <template #title>控制台</template>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'

const isCollapse = ref(false)
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}
</script>

<style>
.menu-side{
  position: fixed;
  background-color: #545c64;
  top: 60px;
  left: 0;
  bottom: 0;
  z-index: 999;
  width: 200px;
  overflow-x: hidden;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  /*width: 200px;*/
  min-height: 400px;
  border-right: none;
}
</style>
