<script setup>
import Menu from './components/Base/Menu.vue'
import Head from './components/Base/Head.vue'
import Home from './components/Base/Home.vue'
</script>

<template>
  <Menu></Menu>
  <Head></Head>
  <Home></Home>
</template>

<style>
*{
  margin: 0;
  padding: 0;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
