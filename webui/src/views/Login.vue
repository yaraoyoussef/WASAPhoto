<script>
export default {
    data() {
      return {
        id: "",
        errormsg: null,
        disabled: true
      }
    },
    methods: {
        async login() {
            this.errormsg = null
            try {
              let response = await this.$axios.post("/session", {ID: this.id.trim()});
              localStorage.setItem('token', response.data.ID)
              this.$router.replace('/home')
              this.$emit('updateLoggedChild', true)
            } catch (e) {
                this.errormsg = e.toString()
            }
        },
    },
    mounted() {
      if(localStorage.getItem('item')) {
        this.$router.replace('/home')
      }
    }
}
</script>

<template>
  <div class="container-center-horizontal">
    <div class="login-section screen">
      <h1>WASAPhoto</h1>
      <form @submit.prevent="login" class="d-flex flex-column align-items-center justify-content-center">
        <div class="padding">
          <div class="login-form">
            <h2>Register / Login</h2>
            <div class="login-wrapper">
              <div class="input-field">
                <input
                  type="text"
                  class="form-control"
                  v-model="id"
                  minlength="5"
                  maxlength="15"
                  placeholder="enter an identifier"
                />
              </div>
              <div class="login-button">
                <button class="login-btn" :disabled="id == null || id.trim().length < 5">
                  Register / Login
              </button>
              </div>
            </div>
          </div>
          <div class="err-container">
                <ErrorMsg class="err-txt" v-if="errormsg" :msg="errormsg"></ErrorMsg>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>
  
<style>
.login-section {
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: flex-start;
    height: 100%;
}
.padding {
    align-items: center;
    display: flex;
    flex-direction: column;
    gap: 50px;
    justify-content: center;
    padding: 10px;
    position: relative;
}
.login-form {
    align-items: center;
    border: 2px solid;
    background-color: #7696e8fe;
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    position: relative;
    width: fit-content;
    margin-top: 70px;
}

.login-wrapper {
    align-items: center;
    background-color: #cbddfc;
    display: flex;
    flex-direction: column;
    gap: 40px;
    padding: 40px;
    position: relative;
    width: 500px;
    }

.input-field {
    align-items: flex;
    align-self: stretch;
    display: flex;
    flex-direction: column;
}

.login-button {
    align-items: center;
    align-self: flex;
    justify-content: center;
    padding: 10px 20px;
}
.err-container {
    text-align: center;
}
.err-txt {
    color: red;
    font-size: 20px;
    font-style: italic;
}
</style>  