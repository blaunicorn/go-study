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
          title="Login"
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
                @click="login"
              >Login</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
// import storageService from "@/utils/storageService";
import { mapActions } from "vuex";
// import api from "@/api/user";
// import {
//   required,
//   minLength,
//   maxLength,
//   between,
// } from "vuelidate/lib/validators";
// const telephoneValidator = (value) => /^1[3|4|5|7]\d{9}$/.test(value);
export default {
  data () {
    return {
      user: {
        telephone: null,
        password: "",
        checked: [],
      },
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
    ...mapActions("userModule", { userLogin: "login" }),
    login (event) {
      event.preventDefault();
      if (this.user.telephone && this.user.telephone.length !== 11) {
        this.showTelephoneValidate = false;
        return;
      }
      console.log(this.user);

      this.userLogin(this.user)
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

      // api
      //   .login({ ...this.user })
      //   .then((res) => {
      //     console.log(res.data);

      //     // save token
      //     // localStorage.setItem("token", res.data.data.token);
      //     // storageService.set(storageService.USER_TOKEN, res.data.data.token);
      //     this.$store.commit("userModule/SET_TOKEN", res.data.data.token);
      //     api.info().then((response) => {
      //       console.log(response.data);
      //       // save user info
      //       // storageService.set(
      //       //   storageService.USER_INFO,
      //       //   JSON.stringify(res.data.data.user)
      //       // );

      //       this.$store.commit(
      //         "userModule/SET_USERINFO",
      //         response.data.data.user
      //       );
      //       // Jump to home page
      //       this.$router.replace({ name: "Home" });
      //     });
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
