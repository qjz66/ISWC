<template>
  <div id="app">
    <router-view />
    <!-- <Home/> -->
  </div>
</template>

<script>
import Home from './views/Home.vue'

export default {
  name: 'App',
  components: {
    Home
  },
  mounted() {
    // this.$router.push('/home')
  }
}
</script>

<style>
#app {
  width: 100vw;
  height: 100vh;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
