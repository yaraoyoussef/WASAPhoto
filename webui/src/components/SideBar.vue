<script>
export default {
    data() {
        return {
            text: ""
        }
    },
    methods: {
        logout() {
            localStorage.removeItem('token'),
            this.$emit('logoutSidebar', false)
        },

        myProfile() {
            this.$emit('updateView', "/users/"+localStorage.getItem('token').trim())
        },

        feed() {
            this.$emit('updateView', "/home")
        },

        search() {
            this.$emit('searchUsers', this.text)
            this.text = ""
        }
    }
}
</script>

<template>
    <div class="sidebar">
        <div class="top">
            <h2 class="title">WasaPhoto</h2>
            <div class="search-section">
                <input class="search-input" type="search" v-model="text" placeholder="Search for users">
                <button class="search-button" type="submit" @click.prevent="search">Search</button>
            </div>
        </div>
        <div class="profile-section">
            <div class="profile-icon"> 
                <i class="uuser fas fa-user"></i>              
            </div>
            <button @click="myProfile" type="button" class="profile-header">
                My Profile
            </button>
        </div>
        <ul class="menu-list">
            <li>
                <button type="button" @click="feed">
                    <i class="fas fa-home"></i>
                    Feed
                </button>
            </li>
            <li>
                <button type="button" @click="logout">
                    <i class="fas fa-sign-out-alt"></i>
                    Logout
                </button>
            </li>
        </ul>
    </div>   
</template>

<style>
body {
    margin: 0;
    padding: 0;
}

.sidebar {
    position: fixed;
    top: 0;
    left: 0;
    width: 250px;
    bottom: 0;
    background-color: #8792f8;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.title {
    font-size: 30px;
    font-weight: bold;
    margin-left: 35px;
    margin-bottom: 20px;
}

.search-section {
    display: flex;
    align-items: center;
}

.search-input {
    width: 100%;
}

.search-button {
    margin-left: 10px;
    background: #b9bddf;
}

.profile-section {
    display: flex;
    align-items: center;
}

.profile-icon {
    width: 60px;
    height: 60px;
    background-color: rgb(212, 203, 203);
    border-radius: 50%;
    margin-right: 10px;
}

.uuser {
    font-size: 30px;
    margin-left: 16px;
    margin-top: 10px;
}

.profile-header {
    font-weight: bold;
    font-size: 25px;
    background: #b9bddf;
}
.profile-header:hover {
    transform: scale(1.1);
}

.menu-list {
    list-style: none;
    padding: 0;
    margin: 0;
}

.menu-list li {
    margin-bottom: 10px;
}

.menu-list li button {
    text-decoration: none;
    color: #333;
    background: #b9bddf;
    font-size: 20px;
}

.menu-list li button:hover {
    transform: scale(1.1);
}
    
.content {
    margin-left: 250px;
    padding: 20px;
}
</style>