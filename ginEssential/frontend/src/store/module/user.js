import storageService from "@/utils/storageService";
import api from "@/api/user";
console.log(storageService.get(storageService.USER_INFO));
console.log(storageService.get(storageService.USER_TOKEN));
const userModule = {
  namespaced: true,
  state: {
    token: storageService.get(storageService.USER_TOKEN),
    userInfo: storageService.get(storageService.USER_INFO)
      ? JSON.parse(storageService.get(storageService.USER_INFO))
      : null,
  },

  mutations: {
    SET_TOKEN (state, token) {
      // updata localstorage
      storageService.set(storageService.USER_TOKEN, token);
      // updata state
      state.token = token;
    },
    SET_USERINFO (state, userInfo) {
      // updata localstorage
      storageService.set(storageService.USER_INFO, JSON.stringify(userInfo));
      // updata state
      state.userInfo = userInfo;
    },
  },

  actions: {
    register (context, user) {
      return new Promise((resolve, reject) => {
        api
          .register({ ...user })
          .then((res) => {
            console.log(res.data);

            // save token
            context.commit("SET_TOKEN", res.data.data.token);

            return api.info();
          })
          .then((response) => {
            console.log(response.data);
            // save user info
            // storageService.set(
            //   storageService.USER_INFO,
            //   JSON.stringify(response.data.data.user)
            // );
            context.commit("SET_USERINFO", response.data.data.user);
            resolve(response);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    login (context, user) {
      return new Promise((resolve, reject) => {
        api
          .login({ ...user })
          .then((res) => {
            console.log(res.data);

            // save token
            context.commit("SET_TOKEN", res.data.data.token);

            return api.info();
          })
          .then((response) => {
            console.log(response.data);
            // save user info
            // storageService.set(
            //   storageService.USER_INFO,
            //   JSON.stringify(response.data.data.user)
            // );
            context.commit("SET_USERINFO", response.data.data.user);
            resolve(response);
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    logout ({ commit }) {
      // clear token
      commit("SET_TOKEN", "");
      // storageService.set(storageService.USER_TOKEN, "");
      // clear userinfo
      commit("SET_USERINFO", "");
      // storageService.set(storageService.USER_INFO, "");
    },
  },
};

export default userModule;
