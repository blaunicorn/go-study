import axios from "axios";
import storageService from "./storageService";
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 5000,
});

// Add a request interceptor
service.interceptors.request.use(
  function (config) {
    // do something before request is sent
    // Object.assign(config.headers, {
    //   Authorization: "Bearer " + storageService.get(storageService.USER_TOKEN),
    // });
    config.headers.Authorization =
      "Bearer " + storageService.get(storageService.USER_TOKEN);
    return config;
  },
  function (error) {
    // do something with request error
    return Promise.reject(error);
  }
);

// add a response interceptor
service.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);

export default service;
