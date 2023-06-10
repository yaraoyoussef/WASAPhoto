<script>
export default {
    data: function() {
      return {
        errMsg: null,
        cUserBanned: false,
        exists: false,

        username: "",

        followState: false,
        banSate: false,

        nPosts: 0,
        nFollowers: 0,
        nFollowings: 0,

        photos: [],
        followers: [],
        followings: []
      }
    },
    watch: {
      currentPath(newId, oldId) {
        if(newId !== oldId) {
          this.loadInfo()
        }
      }
    },
    computed: {
      currentPath() {
        return this.$route.params.id
      },

      cUser() {
        return this.$route.params.id === localStorage.getItem('token')
      }
    },
    methods: {
      async follow(){
        try {
          if(this.followState) {
            await this.$axios.delete("/users/"+localStorage.getItem('token')+"follow/"+this.$route.params.otherUserId);
            this.nFollowings -=1
          } else {
            await this.$axios.put("/users/"+localStorage.getItem('token')+"follow/"+this.$route.params.otherUserId);
            this.nFollowings +=1
          }
          this.followState = !this.followState
        } catch(e) {
          this.errMsg = e.toString();
        }
      },

      async ban() {
        try {
          if(this.banSate) {
            await this.$axios.delete("/users/"+localStorage.getItem('token')+"ban/"+this.$route.params.otherUserId);
          } else {
            await this.$axios.put("/users/"+localStorage.getItem('token')+"ban/"+this.$route.params.otherUserId);
            this.followState = false
          }
          this.banSate = !this.banSate
        } catch(e) {
          this.errMsg = e.toString();
        }
      },

      async loadInfo() {
        if(this.$route.params.id === undefined){
          return
        }
        try {
          let response = await this.$axios.get("/users/"+this.$route.params.id);
          this.banSate = false
          this.exists = true
          this.cUserBanned = false

          if(response.status === 206) {
            this.banSate = true
            return
          }

          if(response.status === 204) {
            this.exists = false
          }
          this.username = response.data.username
          this.nFollowers = response.data.followers != null ? response.data.followers.length : 0
          this.nFollowings = response.data.followings != null ? response.data.followings.length : 0
          this.nPosts = response.data.photos != null ? response.data.photos.length : 0
          this.photos = response.data.photos != null ? response.data.photos : []
          this.followers = response.data.followers != null ? response.data.followers : []
          this.followings = response.data.followings != null ? response.data.followings : []
          this.followState = response.data.followers != null ? response.data.followers.find(obj => obj.id === localStorage.getItem('token')) : false

        } catch(e) {
          this.cUserBanned = true
        }
      },

      deletePhoto(photoId) {
        this.photos = this.photos.filter(item => item.photoId !== photoId)
      },

      editUsername() {
        // idk what to write here
      },

      logout() {
        this.$router.push("/session")
        // clear the session 
      },

      async mounted() {
        await this.loadInfo()
      }
    }
};
</script>

<template>
    <div class="container-fluid" v-if="!cUserBanned && exists">
        <div class="pre-page">
          <h1 class="username">{{ username }}</h1>
          <button class="exit">
            <i class="exit-icon fas fa-xmark"></i>
          </button>
        </div>
        <div class="page-head">
          <h2 class="element">Posts: {{ nPosts }}</h2>
          <h3 class="element">Followers: {{ nFollowers }}</h3>
          <h4 class="element">Followings: {{ nFollowings }}</h4>
          <button v-if="!cUser && !banned" @click="follow" class="element">
              {{followState ? "Unfollow" : "Follow" }}
          </button>
          <button v-if="!cUser" @click="ban" class="element">
              {{banSate ? "Unban" : "Ban" }}
          </button>
          <button v-else @click="logout" class="element">
            <i class="logout-icon fas fa-right-from-bracket"></i>
          </button>
          <button v-else @click="editUsername" class="element">
              <i class="edit-icon fas fa-pen-to-square"></i> 
          </button>
        </div>
        <hr class="hr">
        <h5 class="post-title">Posts</h5>     
        <hr class="hr">
        <div class="posts-section">
            <div v-if="!cUserBanned && nPosts > 0" class="photo-container">
                <Photo v-for="(photo, index) in photos"
                :key="index"
                :owner="this.$router.params.id"
                :photoId="photo.photoId"
                :comments="photo.comments"
                :likes="photo.likes"
                :date="photo.dateAndTime"
                >
            </Photo>
            </div>
            <div v-else class="empty-container">
                <h6 class="empty">No Posts</h6>
            </div>
        </div>
        <div class="err-container">
          <ErrMsg v-if="errMsg" :msg="errMsg"></ErrMsg>
        </div>
    </div>
    <div v-else class="otherPage">
        <NotFound />
    </div>
</template>

<style>
.container-fluid {
  background-color: white;
}
.pre-page {
  display: flex;
  justify-content: space-between;
  margin-bottom: 30px;
}
.username {
  font-size: 35px;
  color: black;
  margin-bottom: 15px;
  margin-top: 10px;
}
.exit-icon {
  color: black;
  margin-bottom: 15px;
  margin-top: 10px;
}
.page-head {
  display: flex;
}
.element {
  margin-right: 20px;
  font-size: 20px;
  margin-bottom: 10px;
  margin-top: 10px;
}
.element:last-child {
  margin-right: 10px;
}
.post-title {
  text-align: start;
  align-items: center;
  flex-direction: column;
  font-size: 20px;
  color: black;
}
.posts-section {
  align-items: center;
  text-align: center;
  flex-direction: column;
}
.empty {
  align-items: center;
  text-align: center;
  flex-direction: column;
	font-size: 30px;
	color: black;
}
.err-container {
    align-items: center;
    text-align: center;
    flex-direction: column;
    font-size: 20px;
    color: red;
}
</style>