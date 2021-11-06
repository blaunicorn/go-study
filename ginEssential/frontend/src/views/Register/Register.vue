<template>
  <div class="register">
    <b-row>
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card
          title="register"
          class="mt-5"
        >
          <b-form
            @submit="onSubmit"
            @reset="onReset"
          >
            <b-form-group
              description="We'll never share your telephone with anyone else."
            >
              <label
                for="feedback-user"
                class="mt-3"
              >Name</label>
              <b-form-input
                v-model="user.name"
                type="text"
                placeholder="Please enter your name (optional)"
              ></b-form-input>
              <label
                for="feedback-user"
                class="mt-3"
              >Telephone</label>
              <b-form-input
                v-model="user.telephone"
                type="number"
                placeholder="Please enter your telephone (required)"
                required
              ></b-form-input>
              <b-form-text
                text-variant="danger"
                v-if="showTelephoneValidate"
              >telephone must be 11 digits!
              </b-form-text>
              <label
                for="feedback-user"
                class="mt-3"
              >Password</label>
              <b-form-input
                class="mt-2"
                v-model="user.password"
                type="password"
                placeholder="Please enter your password (required)"
                required
              ></b-form-input>
              <b-form-invalid-feedback :state="validation">
                Your user password must be > 6 characters long.
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                Looks Good.
              </b-form-valid-feedback>
              <b-button
                class="mt-2"
                variant="outline-primary"
                block
                @click="register"
              >Register</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
// import storageService from "@/utils/storageService";
// import api from "@/api/user";
// import {
//   required,
//   minLength,
//   maxLength,
//   between,
// } from "vuelidate/lib/validators";
// const telephoneValidator = (value) => /^1[3|4|5|7]\d{9}$/.test(value);

// vuex mode 2 mapMutations
import { mapMutations, mapActions } from "vuex";
export default {
  data () {
    return {
      user: {
        name: "",
        telephone: null,
        password: "",
        checked: [],
      },
      foods: [
        { text: "Select One", value: null },
        "Carrots",
        "Beans",
        "Tomatoes",
        "Corn",
      ],
      show: true,
      showTelephoneValidate: false,
    };
  },
  computed: {
    validation () {
      return this.user.password.length > 6;
    },
  },
  // validations: {
  //   user: {
  //     telephone: {
  //       required,
  //       minLength: minLength(11),
  //       maxLength: maxLength(11),
  //     },
  //   },
  // },
  methods: {
    // vuex  mode1 1 mapMutations
    ...mapMutations("userModule", ["SET_TOKEN", "SET_USERINFO"]),
    ...mapActions("userModule", { userRegister: "register" }),
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
      // vuex  mode 1 mapActions
      // this.$store
      //   .dispatch("userModule/register", this.user)
      //   .then(() => {
      //     this.$router.replace({ name: "Home" });
      //   })
      //   .catch((err) => {
      //     console.log("err:", err.response.data.msg);
      //     this.$bvToast.toast(err.response.data.msg, {
      //       title: err.response.data.msg,
      //       variant: "danger",
      //       solid: true,
      //     });
      //   });

      // vuex  mode1 2 mapMutations
      // api
      //   .register({ ...this.user })
      //   .then((res) => {
      //     console.log(res.data);

      //     // save token
      //     // localStorage.setItem("token", res.data.data.token);
      //     // storageService.set(storageService.USER_TOKEN, res.data.data.token);
      //     console.log(this.$store);
      //     this.$store.commit("userModule/SET_TOKEN", res.data.data.token);

      //     this.SET_TOKEN(res.data.data.token);

      //     return api.info();
      //   })
      //   .then((response) => {
      //     console.log(response.data);
      //     // save user info
      //     // storageService.set(
      //     //   storageService.USER_INFO,
      //     //   JSON.stringify(response.data.data.user)
      //     // );
      //     this.$store.commit(
      //       "userModule/SET_USERINFO",
      //       response.data.data.user
      //     );

      //     this.SET_USERINFO(response.data.data.user);
      //     // Jump to home page
      //     this.$router.replace({ name: "Home" });
      //   })
      //   .catch((err) => {
      //     console.log("err:", err.response.data.msg);
      //     this.$bvToast.toast(err.response.data.msg, {
      //       title: err.response.data.msg,
      //       variant: "danger",
      //       solid: true,
      //     });
      //   });
      // alert(JSON.stringify(this.user));
    },
    onSubmit (event) {
      event.preventDefault();
      alert(JSON.stringify(this.form));
    },
    onReset (event) {
      event.preventDefault();
      // Reset our form values
      this.form.telephone = null;
      this.form.name = "";
      this.form.password = null;
      this.form.checked = [];
      // Trick to reset/clear native browser form validation state
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    },
  },
};
</script>

<style></style>
