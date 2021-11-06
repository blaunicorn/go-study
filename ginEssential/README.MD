# learning gin videos

# ginEssential

```
|-- ginEssential
  |--server  后端代码
    |-- common // 公用包
    |  |-- database.go //数据库
    |  |-- jwt   // 认证
    |-- util // 工具包
    |-- dto //数据传输对象
    |-- response // 统一返回格式
    |-- main.go  //入口文件
    |-- routes.go  // 路由
  |--frontend  // 前端代码
```

## 一、项目内容

基于 gin 框架的 demo 项目,本项目基于 B 站视频 Go 语言 Gin+Vue 前后端分离实战 - OceanLearn
[https://www.bilibili.com/video/BV1CE411H7bQ?t=7](https://www.bilibili.com/video/BV1CE411H7bQ?t=7)

**frontend**为前端 vue 代码

## 二、怎样运行该项目

### 1. 运行后端程序

> 先确保你电脑上正确安装了 golang 环境

从git@github.com:blaunicorn/go-study.git master 分支拉取后端 golang 代码

```shell
# 拉取代码
git clone -b master https://github.com/kotlindev/gin-demo.git backend
# 进入项目目录
cd  backend
# 安装项目依赖
go get
```

打开 `config/application.yaml` 文件，修改数据库链接配置，修改项目运行端口，确保端口不被占用，参考如下

```yaml
server:
  port: 8081
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: api
  username: pan
  password: www.phy.xyz
  charset: utf8
  loc: Asia/Shanghai
```

启动项目

```
go run routes.go main.go
```

如果看到命令行终端输出以下路由信息，代表项目运行正常。如果不正常，检查一下数据库地址还有账号密码是否正确，同时确保运行的端口没有被占用

```
[GIN-debug] POST   /api/auth/register        --> jkdev.cn/api/controller.Register (5 handlers)
[GIN-debug] POST   /api/auth/login           --> jkdev.cn/api/controller.Login (5 handlers)
[GIN-debug] GET    /api/auth/info            --> jkdev.cn/api/controller.Info (6 handlers)
[GIN-debug] POST   /categories               --> jkdev.cn/api/controller.ICategoryController.Create-fm (5 handlers)
[GIN-debug] PUT    /categories/:id           --> jkdev.cn/api/controller.ICategoryController.Update-fm (5 handlers)
[GIN-debug] GET    /categories/:id           --> jkdev.cn/api/controller.ICategoryController.Show-fm (5 handlers)
[GIN-debug] DELETE /categories/:id           --> jkdev.cn/api/controller.ICategoryController.Delete-fm (5 handlers)
[GIN-debug] POST   /posts                    --> jkdev.cn/api/controller.IPostController.Create-fm (6 handlers)
[GIN-debug] PUT    /posts/:id                --> jkdev.cn/api/controller.IPostController.Update-fm (6 handlers)
[GIN-debug] GET    /posts/:id                --> jkdev.cn/api/controller.IPostController.Show-fm (6 handlers)
[GIN-debug] DELETE /posts/:id                --> jkdev.cn/api/controller.IPostController.Delete-fm (6 handlers)
[GIN-debug] POST   /posts/page/list          --> jkdev.cn/api/controller.IPostController.PageList-fm (6 handlers)
[GIN-debug] Listening and serving HTTP on :8081
```

### 2. 运行前端程序

> 先确保你电脑上正确安装了 npm 环境，并安装了 vue、yarn

从 vue 分支拉取前端 vue 代码

```shell
# 拉取代码
git clone -b vue https://github.com/kotlindev/gin-demo.git vue
# 进入项目目录
cd  vue
# 安装项目依赖
yarn install
```

根据 1 中的 后端代码的运行端口，修改 `.env.development.local` 和 `.env.development` 两个配置文件，修改配置如下为

```
VUE_APP_BASE_URL = http://localhost:8081/api/
```

在运行项目

```shell
yarn serve
```

[video address]https://www.bilibili.com/video/BV1CE411H7bQ

1. first section

```
mkdir ginessential
cd ginessential
cd . > main.go
```

edit helllowrd

```
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

```

```
go run main.go
```

```
go mod init github.com/oceanlearn.teach/ginessential
```

download gin dependency

```
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

```

2. second section

3. 封装统一的请求返回格式

```
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// {
// 	code: 20001,
// 	data: XXX,
// 	msg: XXX
// }

func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

// 业务成功
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

//业务失败
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

```

7. 统一配置管理

```
// viper module

```

8. 前端环境搭建

```
//  采用nvm 确保多个项目开发环境统一
 nvm install v10.14.2
 nvm ls
 node -v
 nvm user v12.16.1
 node -v
 node use v10.14.2
node -v
// install yarn
 npm install -g yarn

# install vue-cli
 npm install -g @vue/cli
# OR
yarn global add @vue/cli
```

9. 创建 vue 项目代码规范

```
 vue create frontend
eslintrc.js

```

10. 编写注册登录页面

```
yarn add vue bootstrap-vue
yarn add bootstrap@4.5.3
```

11. 表单验证

```
// 三种方式
// one
              <b-form-text text-variant="danger" v-if="showTelephoneValidate"
                >telephone must be 11 digits!
              </b-form-text>

                  register (event) {
      event.preventDefault();
      if (this.user.telephone.length !== 11) {
        this.showTelephoneValidate = false;
        return;
      }
      alert(JSON.stringify(this.user));
    },
// twor
              <b-form-invalid-feedback :state="validation">
                Your user password must be > 6 characters long.
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                Looks Good.
              </b-form-valid-feedback>

                  validation () {
      return this.user.password.length > 6;
    },
// three
yarn add vuelidate --save

```

12. axios 请求和 cros 设置

```
yarn add axios
// 在入口main.js中导入axios 并将axios写入vue的原型，这样就能更简单的使用。

import axios from 'axios'
import Qs from 'qs'
//QS是axios库中带的，不需要我们再npm安装一个

Vue.prototype.axios = axios;
Vue.prototype.qs = Qs;

```

```
    register () {
      // event.preventDefault();
      if (!this.user.telephone || this.user.telephone.length !== 11) {
        this.showTelephoneValidate = true;
        return;
      } else {
        this.showTelephoneValidate = false;
      }
      if (!this.user.password || this.user.password.length < 6) {
        return;
      }
      const api = "http://localhost:8081/api/auth/register";
      this.axios
        .post(api, { ...this.user })
        .then((res) => {
          console.log(res.data);

          // save token
          localStorage.setItem("token", res.data.data.token);
          // Jump to home page
          this.$router.replace({ name: "Home" });
        })
        .catch((err) => {
          console.log("err:", err.response.data.msg);
          this.$bvToast.toast(err.response.data.msg, {
            title: err.response.data.msg,
            variant: "danger",
            solid: true,
          });
        });
      // alert(JSON.stringify(this.user));
    },
```

13. axios 封装

```
// frontend\src\utils\request.js
import axios from "axios";
import storageService from "./storageService";
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 5000,
  headers: {
    "X-Custom-Header":
      "Bearer " + storageService.get(storageService.USER_TOKEN),
  },

});
service.interceptors.request.use(
  function (config) {
    // do something before request is sent
    Object.assign(config.headers, {
      Authorization: "Bearer " + storageService.get(storageService.USER_TOKEN),
    });
    return config;
  },
  function (error) {
    // do something with request error
    return Promise.reject(error);
  }
);
export default service;

```

```
// frontend\src\utils\storageService.js
// localstroage service

const PREFIX = "ginessential_";

// user module
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// save function
const set = (key, data) => {
  localStorage.setItem(key, JSON.stringify(data));
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

```

```
// frontend\src\views\Register\Register.vue
      api
        .register({ ...this.user })
        .then((res) => {
          console.log(res.data);

          // save token
          // localStorage.setItem("token", res.data.data.token);
          storageService.set(storageService.USER_TOKEN, res.data.data.token);
          // Jump to home page
          this.$router.replace({ name: "Home" });
        })
```

14. 使用 vuex 重构 login 页面数据

```
// frontend\src\views\Layout\Navbar.vue
<script>
// import storageService from "@/utils/storageService";
// Mode2  mapState
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

```

```
// frontend\src\views\Register\Register.vue
    // vuex  mode1 1 mapMutations
    ...mapMutations("userModule", ["SET_TOKEN", "SET_USERINFO"]),

    register () {
      // event.preventDefault();
      if (!this.user.telephone || this.user.telephone.length !== 11) {
        this.showTelephoneValidate = true;
        return;
      } else {
        this.showTelephoneValidate = false;
      }
      if (!this.user.password || this.user.password.length < 6) {
        return;
      }
      // const api = "http://localhost:8081/api/auth/register";
      // this.axios.post(api, { ...this.user })
      api
        .register({ ...this.user })
        .then((res) => {
          console.log(res.data);

          // save token
          // localStorage.setItem("token", res.data.data.token);
          // storageService.set(storageService.USER_TOKEN, res.data.data.token);
          console.log(this.$store);
          this.$store.commit("userModule/SET_TOKEN", res.data.data.token);

          this.SET_TOKEN(res.data.data.token);

          return api.info();
        })
        .then((response) => {
          console.log(response.data);
          // save user info
          // storageService.set(
          //   storageService.USER_INFO,
          //   JSON.stringify(response.data.data.user)
          // );
          this.$store.commit(
            "userModule/SET_USERINFO",
            response.data.data.user
          );

          this.SET_USERINFO(response.data.data.user);
          // Jump to home page
          this.$router.replace({ name: "Home" });
        })
        .catch((err) => {
          console.log("err:", err.response.data.msg);
          this.$bvToast.toast(err.response.data.msg, {
            title: err.response.data.msg,
            variant: "danger",
            solid: true,
          });
        });
      // alert(JSON.stringify(this.user));
    },
```

```
// frontend\src\store\module\user.js
    ...mapActions("userModule", { userRegister: "register" }),
    ...
// vuex  mode1 2 mapActions
      // vuex  mode 2 mapActions
      this.userRegister(this.user)
        .then(() => {
          this.$router.replace({ name: "Home" });
        })
        .catch((err) => {
          console.log("err:", err.response.data.msg);
          this.$bvToast.toast(err.response.data.msg, {
            title: err.response.data.msg,
            variant: "danger",
            solid: true,
          });
        });
```

16. 前端登陆和 router navigation guards

```
// frontend\src\store\module\user.js
    logout ({ commit }) {
      // clear token
      commit("SET_TOKEN", "");
      // storageService.set(storageService.USER_TOKEN, "");
      // clear userinfo
      commit("SET_USERINFO", "");
      // storageService.set(storageService.USER_INFO, "");
    },
```

```
// frontend\src\router\index.js
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
```

17. 文章分类

```
// server\model\category.go
// 根据约定规范，一般文件名应该采取小写字母+“_”方式。
package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

```

```
// server\controller\category_controller.go
package controller

import (
	"strconv"

	"github.com/blaunicorn/oceanlearn.teach/ginessential/common"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/model"
	"github.com/blaunicorn/oceanlearn.teach/ginessential/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() CategoryController {
	db := common.GetDB()

	db.AutoMigrate(model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "Data validation error, category name is required")
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "success")

}
func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "Data validation error, category name is required")
	}

	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RowsAffected < 1 {
		response.Fail(ctx, nil, "category is not esint")

	}

	// 更新分类
	// 方式1 map
	// 方式2 结构体
	// 方式3 name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "success")
}
func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category, categoryId).RowsAffected < 1 {
		response.Fail(ctx, nil, "category is not esint")
	}
	response.Success(ctx, gin.H{"category": category}, "success")
}
func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category, categoryId).RowsAffected < 1 {
		response.Fail(ctx, nil, "category is not esint")
		return
	}
	if err := c.DB.Delete(&model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, gin.H{"err": err}, "删除失败")
		return
	}
	response.Success(ctx, nil, "delete success")
}

```

18. 文章分类重构

```
// 修改为自动数据验证

```

19. 文章 api

```

```
