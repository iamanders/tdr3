<template>
  <div class="main-navigation">
    <router-link to="/day" class="brand">
      <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>
      <span>TDR</span>
    </router-link>

    <ul class="nav">
      <li>
        <router-link to="/day" :class="{ 'active': route.path.startsWith('/day') }">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
          <span class="ml-1">Day</span>
        </router-link>
      </li>
      <li>
        <router-link to="/week" :class="{ 'active': route.path.startsWith('/week') }">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
          <span class="ml-1">Week</span>
        </router-link>
      </li>
      <li>
        <router-link to="/month" :class="{ 'active': route.path.startsWith('/month') }">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
          <span class="ml-1">Month</span>
        </router-link>
      </li>
      <li>
        <router-link to="/report" :class="{ 'active': route.path.startsWith('/report') }">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16c0 1.1.9 2 2 2h12a2 2 0 0 0 2-2V8l-6-6z"/><path d="M14 3v5h5M16 13H8M16 17H8M10 9H8"/></svg>
          <span class="ml-1">Report</span>
        </router-link>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router';

const route = useRoute();

</script>

<style lang="scss" scoped>

.main-navigation {
  padding: 2rem 1rem;
  height: 100%;
  background-color: #ffffff;
}

.brand {
  font-weight: 900;
  padding: 1rem;
  display: flex;
  align-items: center;
  text-decoration: none;

  span { margin-left: .5rem; }
}

ul.nav {
  margin: .5rem 0 0 0;
  padding: 0;
  list-style: none;

  li {
    margin-bottom: .25rem;

    a {
      display: flex;
      align-items: center;
      padding: .5rem 1rem;
      text-decoration: none;
      border-radius: 1rem;

      &:hover {
        background-color: #eee;
      }

      &.active {
        background-color: #eee;
      }
    }
  }
}
</style>
