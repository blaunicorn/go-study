import Vue from "vue";
import VueRouter from "vue-router";
import store from "@/store";
// import Home from "../views/Home.vue";

Vue.use(VueRouter);
/* Layout */
import Layout from "@/views/Layout/Layout.vue";
import userRoutes from "./module/user";
const routes = [
  {
    path: "/",
    name: "Home",
    component: Layout,
    // redirect: '/dashboard',
    // children: [
    //   {
    //     path: "/register",
    //     name: "Register",
    //     component: () => import("../views/Register/Register.vue"),
    //   },
    //   {
    //     path: "/login",
    //     name: "Login",
    //     component: () => import("../views/Login/Login.vue"),
    //   },
    // ],
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  // {
  //   path: "/register",
  //   name: "Register",
  //   component: () => import("../views/Register/Register.vue"),
  // },
  // {
  //   path: "/login",
  //   name: "Login",
  //   component: () => import("../views/Login/Login.vue"),
  // },
  userRoutes,
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.auth) {
    // 判断是否需要登陆权限
    // 判断用户是否登陆
    if (store.state.userModule.token) {
      // 这里还需要判断token的有效性，比如有没有过期，需要后台发放token的时候，带上token的有效期
      // 如果token无效，需要请求token
      next();
    } else {
      next({ name: "Login" });
    }
  } else {
    next();
  }
});

export default router;
