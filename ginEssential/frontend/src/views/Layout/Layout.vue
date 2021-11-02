<template>
  <div>
    <b-container fluid style="padding: 0 0">
      <nav-bar></nav-bar>
    </b-container>

    <b-container class="main">
      <router-view></router-view>
    </b-container>
  </div>
</template>

<script>
import NavBar from "./Navbar.vue";
export default {
  components: {
    NavBar,
  },
};
</script>

<style></style>
