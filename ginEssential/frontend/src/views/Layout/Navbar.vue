<template>
  <div>
    <b-navbar
      toggleable="lg"
      type="dark"
      variant="info"
    >
      <b-container>
        <b-navbar-brand
          tag="a"
          @click="$router.push({ name: 'Home' })"
        >Blaunicorn
        </b-navbar-brand>
        <b-collapse
          id="nav-collapse"
          is-nav
        >
          <!-- Right aligned nav items -->
          <b-navbar-nav class="ml-auto">
            <!-- <b-nav-form>
              <b-form-input
                size="sm"
                class="mr-sm-2"
                placeholder="Search"
              ></b-form-input>
              <b-button
                size="sm"
                class="my-2 my-sm-0"
                type="submit"
              >Search</b-button>
            </b-nav-form>

            <b-nav-item-dropdown
              text="Lang"
              right
            >
              <b-dropdown-item href="#">EN</b-dropdown-item>
              <b-dropdown-item href="#">ES</b-dropdown-item>
              <b-dropdown-item href="#">RU</b-dropdown-item>
              <b-dropdown-item href="#">FA</b-dropdown-item>
            </b-nav-item-dropdown> -->

            <b-nav-item-dropdown
              right
              v-if="userInfo"
            >
              <!-- Using 'button-content' slot -->
              <template #button-content>
                <em>{{ userInfo.name }}-{{ userInfo1.name }}</em>
              </template>
              <b-dropdown-item @click="$router.replace({ name: 'profile' })">
                Profile</b-dropdown-item>
              <b-dropdown-item @click="logout">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown>
            <div v-else>
              <b-nav-item
                v-if="$route.name != 'Login'"
                @click="$router.replace({ name: 'Login' })"
              >登录
              </b-nav-item>
              <b-nav-item
                v-if="$route.name != 'Register'"
                @click="$router.replace({ name: 'Register' })"
              >注册</b-nav-item>
            </div>
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
// import storageService from "@/utils/storageService";
// Mode2  mapState
import { mapState, mapActions } from "vuex";

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
  methods: {
    ...mapActions("userModule", ["logout"]),
  },
};
</script>

<style></style>
