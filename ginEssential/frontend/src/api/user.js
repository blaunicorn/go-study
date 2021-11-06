import request from "@/utils/request";

// user register
function register ({ name, telephone, password }) {
  // export function register ({ name, telephone, password }) {
  return request.post("auth/register", { name, telephone, password });
}

// get user info
function info () {
  // export function register ({ name, telephone, password }) {
  return request.get("auth/info");
}
// user login
function login ({ telephone, password }) {
  // export function register ({ name, telephone, password }) {
  return request.post("auth/login", { telephone, password });
}

export default {
  register,
  info,
  login,
};
