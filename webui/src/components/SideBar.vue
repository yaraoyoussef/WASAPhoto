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
            console.log("text in sidebar: ", this.text)
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
                <i class="fas fa-user"></i>              
            </div>
            <button @click="myProfile" type="button" class="profile-header">
                My Profile
            </button>
        </div>
        <ul class="menu-list">
            <li>
                <button type="button" @click="feed">
                    <i class="fas fa-house"></i>
                    Feed
                </button>
            </li>
            <li>
                <button type="button" @click="logout">
                    <i class="fas fa-right-from-bracket"></i>
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
    background-color: #f1f1f1;
    padding: 20px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.title {
    font-size: 30px;
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
}

.profile-section {
    display: flex;
    align-items: center;
}

.profile-icon {
    width: 60px;
    height: 60px;
    background-color: gray;
    border-radius: 50%;
    margin-right: 10px;
}

.profile-header {
    font-weight: bold;
    font-size: 25px;
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