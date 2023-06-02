<script>
export default {
    data: function() {
        return {

        }
    },
}
</script>

<template>
    <div class="container-fluid" v-if="!cUserBanned && exists">
        <div class="pre-page">
            <button class="exitbtn">
                <i class="xbtn fas fa-xmark"></i>
            </button>
        </div>
        <div class="page-head">
            <h1 class="username">{{ username }}</h1>
            <h2 class="nposts">Posts: {{ nPosts }}</h2>
            <h3 class="nfollowers">Followers: {{ nFollowers }}</h3>
            <h4 class="nfollowing">Followings: {{ nFollowings }}</h4>
            <button v-if="!cUser && !banned" @click="follow" class="followbtn">
                {{followState ? "Unfollow" : "Follow" }}
            </button>
            <button v-if="!cUser" @click="ban" class="banbtn">
                {{banSate ? "Unban" : "Ban" }}
            </button>
            <button v-else @click="openSettings" class="settingsbtn">
               <i class="settings-icon fas fa-gear"></i> 
            </button>
        </div>
        <hr class="hr">
        <div class="posts-title-section">
            <h5 class="posts-title">Posts</h5>     
        </div>
        <hr class="hr">
        <div class="posts-section">
            <div v-if="!cUserBanned && nPosts > 0" class="photo-container">
                <Photo v-for="(photo, index) in photos"
                :key="index"
                :owner="this.$router.params.id"
                :photo_id="photo.photoId"
                >
            </Photo>
            </div>
        </div>
    </div>
</template>

<style>
.container-fluid {

}
.page-head {

}
</style>