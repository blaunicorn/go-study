<template>
  <div>
    <h3>
      您好<span v-if="userInfo">{{ userInfo.name }}</span>
    </h3>
  </div>
</template>

<script>
import { mapState } from "vuex";
export default {
  computed: {
    userInfo () {
      // return JSON.parse(storageService.get(storageService.USER_INFO));

      // mode 1 this.$store
      // console.log("userInfo:", this.$store.state.userModule.userInfo);
      return this.$store.state.userModule.userInfo;
    },
    // mode mapState
    ...mapState({ userInfo1: (state) => state.userModule.userInfo }),
  },
};
</script>

<style></style>
