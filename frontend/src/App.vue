<template>

  <LoaderBar></LoaderBar>

  <div class="d-flex">
    <div class="side">
      <MainNavigation></MainNavigation>
    </div>
    <main class="main">
      <TimeStatusBar></TimeStatusBar>
      <router-view/>
    </main>
  </div>

</template>

<script setup>
import LoaderBar from './components/LoaderBar.vue';
import MainNavigation from './components/MainNavigation.vue';
import TimeStatusBar from './components/TimeStatusBar.vue';
import './scss/framework.scss';
import './scss/app.scss';

</script>

<style lang="scss" scoped>

.side {
  flex: 0 0 250px;
  min-height: 100vh;
}
.main {
  flex: 1 1 100%;
}

</style>
