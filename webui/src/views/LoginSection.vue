<template>
<login-section v-bind="loginSectionData" /> </template>
<script>
    export default  {
        name: "App",
        components : {
        LoginSection,
        },
        data() {
        return { loginSectionData };
        },
    };
    const username1Data = {
        usernameEmail: "Username/Email",
    };
    const username2Data = {
        usernameEmail: "Password",
        className: "username",
    };
    const loginSectionData = {
        login: "Login",
        getTheApp: "Get the app",
        spanText1: "Get it on<br />",
        spanText2: "Google Play",
        spanText3: "Download on the‹br />",
        spanText4: "'App Store",
        username1Props: username1Data,
        username2Props: username2Data, 
    };
</script>