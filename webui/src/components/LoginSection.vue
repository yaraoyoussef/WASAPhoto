<template>
    <div class="container-center-horizontal">
      <div class="login-section screen">
        <div class="padding">
          <div class="login-form">
            <img class="logo-frame" src="https://anima-uploads.s3.amazonaws.com/releases/645a175ad9b92dc7038c3bb/img/logo-frame.svg" alt="Logo_Frame" />
            <div class="login-wrapper">
              <div class="input-fields">
                <username :usernameEmail="username1Props.usernameEmail" />
                <username :username&mail="username2Props.usernameEmail" :className="username2Props.className" />
              </div>
              <div class="login-button">
                <div class="login">{{ login }}</div>
              </div>
            </div>
          </div>
          <div class="get-the-app">{{ getTheApp }}</div>
          <div class="download-app">
            <div class="download-app-item">
              <img class="vector" src="https://anima-uploads.s3.amazonaws.com/releases/645a175ad9b92dc7038c3bb/img/vector.svg" alt="Vector" />
              <div class="x-on poppins-normal-white-12px">
                <span class="poppins-normal-white-12px">{{ spanText1 }}</span>
                <span class="poppins-bold-white-12px">{{ spanText2 }}</span>
              </div>
            </div>
            <div class="download-app-item">
              <img class="vector-1" src="https://anima-uploads.s3.amazonaws.com/releases/645a175ad9b92dc7038c3bb/img/vector-1.svg" alt="Vector-1" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="container-center-horizontal">
      <div class="login-section screen">
        <div class="padding">
          <div class="login-form">
            <img class="logo-frame" src="https://anima-uploads.s3.amazonaws.com/releases/645a175ad9b92dc7038c3bb/img/logo-frame.svg" alt="Logo_Frame">
            <div class="login-wrapper">
              <div class="input-fields">
                <username :usernameEmail="username1Props.usernameEmail" />
                <username :username&mail="username2Props.usernameEmail" :className="username2Props.className" />
              </div>
              <div class="login-button">
                <div class="login">{{ login }}</div>
              </div>
            </div>
          </div>
          <div class="get-the-app">{{ getTheApp }}</div>
          <div class="download-app">
            <div class="download-app-item">
              <img class="vector" src="https://anima-uploads.s3.amazonaws.com/releases/645a175ad9b92dc7038c3bb/img/vector.svg" alt="Vector" />
              <div class="-on poppins-normal-white-12px">
                <span class="poppins-normal-white-12px">{{ spanText3 }}</span>
                <span class="poppins-bold-white-12px">{{ spanText4 }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import Username from './Username'
  
  export default {
    name: "LoginSection",
    components: {
      Username,
    },
    props: [
      "login",
      "getTheApp",
      "spanText1",
      "spanText2",
      "spanText3",
      "spanText4",
      "username1Props",
      "username2Props"
    ]
  }
  </script>
  
  <style>
  .login-section {
    align-items: flex-start;
    background-color: var(--white);
    border: 1px none;
    display: flex;
    width: 1280px;
  }
.padding {
      align-items: center;
      display: flex;
      flex-direction: column;
      gap: 20px;
      height: 832px;
      justify-content: center;
      padding: 50px;
      position: relative;
      width: 1280px;
    }
  .login-form {
    align-items: center;
    border: 2px solid;
    border-color: var(--blueberry);
    border-radius: 20px;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    position: relative;
    width: fit-content;
  }
  
  .logo-frame {
    align-self: stretch;
    height: 75.95px;
    min-width: 500px;
    position: relative;
  }
  
  .login-wrapper {
    align-items: center;
    background-color: #bfd7ff;
    display: flex;
    flex-direction: column;
    gap: 40px;
    padding: 40px;
    position: relative;
    width: 500px;
  }
  
  .input-fields {
    align-items: flex-start;
    align-self: stretch;
    display: flex;
    flex-direction: column;
    gap: 10px;
    position: relative;
  }
  
  .login-button {
    align-items: center;
    align-self: stretch;
    background-color: var(--blueberry);
    border-radius: 20px;
    cursor: pointer;
    display: flex;
    gap: 10px;
    justify-content: center;
    overflow: hidden;
    padding: 10px 20px;
    position: relative;
    transition: all 0.2s ease;
  }
  .login-button:hover {
  background-color: #2836b1;
  color: var(--white);
  font-family: var(--font-family-inter);
  font-size: var(--font-size-1);
  font-weight: 700;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -1px;
  position: relative;
  white-space: nowrap;
  width: fit-content;
  }
.get-the-app {
  color: var(--black);
  font-family: var(--font-family-poppins);
  font-size: var(--font-size-m);
  font-weight: 700;
  letter-spacing: 0;
  line-height: 12px;
  position: relative;
  white-space: nowrap;
  width: fit-content;
  }

.download-app {
  align-items: flex-start;
  display: flex;
  gap: 10px;
  position: relative;
  width: fit-content;
  }

.download-app-item {
  align-items: center;
  background-color: var(--black);
  border-radius: 5px;
  display: flex;
  gap: 10px;
  justify-content: center;
  overflow: hidden;
  padding: 10px;
  position: relative;
  width: fit-content;
  }
  .vector {
  height: 19.98px;
  margin-left: -1px;
  min-width: 17.95px;
  position: relative;
}

.x-on {
  letter-spacing: 0;
  line-height: 12px;
  margin-top: -1px;
  position: relative;
  width: fit-content;
}

.vector-1 {
  height: 18px;
  min-width: 14px;
  position: relative;
}
</style>  