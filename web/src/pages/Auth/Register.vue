<template>
  <div class="text-3xl font-bold text-blue-500">Title</div>
</template>
<script lang="ts">
import { defineComponent, computed, ref } from "vue";
import { useRouter } from "vue-router";
import { useAuth } from "@websanova/vue-auth/src/v3.js";
export default defineComponent({
  name: "name",
  props: {},
  components: {},
  setup() {
    const auth = useAuth();

    const state = reactive({
      form: {
        body: {
          email: "super@starter.com",
          password: "testtest",
        },
        remember: false,
        fetchUser: true,
        staySignedIn: false,
        errors: {},
      },
    });
    onMounted(() => {
      console.log(auth.redirect());
    });
    function errors(res) {
      state.form.errors = Object.fromEntries(
        res.data.errors.map((item) => [item.field, item.msg])
      );
    }
    function loginDefault() {
      auth
        .login({
          data: state.form.body,
          remember: state.form.remember ? '{"name": "Default"}' : null,
          fetchUser: state.form.fetchUser,
          staySignedIn: state.form.staySignedIn,
          redirect: "/",
        })
        .then(null, (res) => {
          errors(res.response);
        });
    }
    function loginRedirect() {
      auth
        .login({
          data: state.form.body,
          redirect: { name: "user-account" },
          remember: state.form.remember ? '{"name": "Redirect"}' : null,
          fetchUser: state.form.fetchUser,
          staySignedIn: state.form.staySignedIn,
        })
        .then(null, (res) => {
          errors(res.response);
        });
    }
    function loginThen() {
      auth
        .login({
          data: state.form.body,
          redirect: null,
          fetchUser: state.form.fetchUser,
          staySignedIn: state.form.staySignedIn,
        })
        .then(
          (res) => {
            if (state.form.remember) {
              auth.remember(
                JSON.stringify({
                  name: auth.user().first_name,
                })
              );
            }
            router.push({ name: "user-account" });
          },
          (res) => {
            errors(res.response);
          }
        );
    }
    function loginComp() {
      authComp
        .login({
          body: state.form.body,
          remember: state.form.remember,
          fetchUser: state.form.fetchUser,
          staySignedIn: state.form.staySignedIn,
        })
        .then(null, (res) => {
          errors(res.response);
        });
    }
    function loginManual() {
      auth.token(null, "manual", false);
      auth.user({
        id: 1,
        first_name: "Manual",
        email: "test@manual.com",
        type: "user",
      });
      if (state.form.remember) {
        auth.remember(
          JSON.stringify({
            name: auth.user().first_name,
          })
        );
      } else {
        auth.unremember();
      }
      router.push({ name: "user-landing" });
    }
    return {
      state,
      loginThen,
      loginComp,
      loginManual,
      loginDefault,
      loginRedirect,
    };
  },
});
</script>
