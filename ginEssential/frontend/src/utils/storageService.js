// localstroage service

const PREFIX = "ginessential_";

// user module
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// save function
const set = (key, data) => {
  localStorage.setItem(key, data);
};

// read function
const get = (key) => {
  return localStorage.getItem(key);
};

export default {
  set,
  get,
  USER_INFO,
  USER_TOKEN,
};
