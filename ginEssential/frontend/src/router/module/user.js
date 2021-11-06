/* Layout */
import Layout from "@/views/Layout/Layout.vue";
const userRoutes = {
  path: "/user",
  name: "user",
  component: Layout,
  redirect: "login",
  children: [
    {
      path: "/register",
      name: "Register",
      component: () => import("@/views/Register/Register.vue"),
    },
    {
      path: "/login",
      name: "Login",
      component: () => import("@/views/Login/Login.vue"),
    },
    {
      path: "profile",
      name: "profile",
      component: () => import("@/views/profile/Profile.vue"),
      meta: {
        auth: true,
      },
    },
  ],
};

export default userRoutes;
