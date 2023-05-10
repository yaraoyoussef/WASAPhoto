<template>
    <div :class="['username-1', className || '*']">
      <div class="username-email poppins-semi-bold-eerie-black-20px">{{ usernameEmail }}</div>
      <div class="placeholder">
        <input class="write-here poppins-normal-black-20px" name="writehere" placeholder="Write here" type="text" />
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "Username",
    props: ["usernameEmail", "className"],
  };
  </script>
  
  <style>
  .username-1 {
    align-items: flex-start;
    display: flex;
    flex-direction: column;
    gap: 10px;
    justify-content: center;
    position: relative;
    width: 420px;
  }
  
  .username-email {
    letter-spacing: 0;
    line-height: normal;
    margin-top: -1px;
    position: relative;
    width: fit-content;
  }
  .placeholder {
  align-items: center;
  align-self: stretch;
  background-color: var(--white);
  border: 2px solid;
  border-color: var(--sonic-silver);
  border-radius: 5px;
  display: flex;
  gap: 10px;
  height: 50px;
  overflow: hidden;
  padding: 10px;
  position: relative;
}

.write-here {
  background-color: transparent;
  border: 0;
  height: 30px;
  letter-spacing: 0;
  line-height: normal;
  margin-top: -2px;
  min-width: 102px;
  padding: 0;
  position: relative;
}

.write-here::placeholder {
  color: #00000099;
}

.username-1 .username .placeholder {
  align-items: flex-start;
}

</style>
  