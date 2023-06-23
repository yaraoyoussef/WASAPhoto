<script>
export default {
    data: function() {
      return {
        errMsg: null,
        cUserBanned: false,
        exists: false,

        username: "",

        followState: false,
        banState: false,

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
            await this.$axios.delete("/users/"+localStorage.getItem('token')+"/follow/"+this.$route.params.id);
            this.nFollowers -=1
          } else {
            await this.$axios.put("/users/"+localStorage.getItem('token')+"/follow/"+this.$route.params.id);
            this.nFollowers +=1
          }
          this.followState = !this.followState
        } catch(e) {
          this.errMsg = e.toString();
        }
      },

      async ban() {
        try {
          if(this.banState) {
            await this.$axios.delete("/users/"+localStorage.getItem('token')+"/ban/"+this.$route.params.id);
            this.loadInfo()
          } else {
            await this.$axios.put("/users/"+localStorage.getItem('token')+"/ban/"+this.$route.params.id);
            if(this.followState) {
              this.nFollowers-=1;
            }
            if (this.followings.find(obj => obj.ID === localStorage.getItem('token'))) {
              this.nFollowings -= 1;
            }
            this.followState = false  
          }
          this.banState = !this.banState
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
          this.banState = false
          this.exists = true
          this.cUserBanned = false

          if(response.status === 206) {
            this.banState = true
            return
          }

          if(response.status === 204) {
            this.exists = false
          }
          this.username = response.data.username
          this.nFollowers = response.data.followers != null ? response.data.followers.length : 0
          this.nFollowings = response.data.following != null ? response.data.following.length : 0
          this.nPosts = response.data.posts
          this.photos = response.data.photos != null ? response.data.photos : []
          this.followers = response.data.followers != null ? response.data.followers : []
          this.followings = response.data.following != null ? response.data.following : []
          this.followState = response.data.followers != null ? response.data.followers.find(obj => obj.ID === localStorage.getItem('token')) : false
        } catch(e) {
          this.cUserBanned = true
        }
      },

      deletePhoto(photoId) {
        this.photos = this.photos.filter(item => item.photoId !== photoId)
        this.nPosts-=1
      },

      async editUsername(newUsername) {
        try {
          let response = this.$axios.put("/users/"+this.$route.params.id, {username: newUsername})
          await this.loadInfo()
        } catch (e) {
          this.errMsg = e.toString()
        }
      },

      goHome() {
        this.$router.replace("/home")
      },
    },

    async mounted() {
        //console.log("mounting profile")
        await this.loadInfo()
      }
};
</script>

<template>
    <div class="container-fluid" v-if="!cUserBanned && exists">
      <UsernameModal :modalId="'usernameMod'"
      @editUsername="editUsername"
      />
        <div class="pre-page">
          <h1 class="username">{{username}}</h1>
          <button class="exit" @click="goHome">
            <i class="exit-icon fas fa-times"></i>
          </button>
        </div>
        <div class="page-head">
          <h2 class="element">Posts: {{nPosts}}</h2>
          <h3 class="element">Followers: {{nFollowers}}</h3>
          <h4 class="element">Followings: {{nFollowings}}</h4>
          <button v-if="!cUser && !banState" @click="follow" class="element cbtn fbtn">
              {{followState ? "Unfollow" : "Follow"}}
          </button>
          <button v-if="!cUser" @click="ban" class="element cbtn bbtn">
              {{banState ? "Unban" : "Ban" }}
          </button>
          <button v-else data-bs-toggle="modal" :data-bs-target="'#usernameMod'" class="element lastof">
              <i class="edit-icon fas fa-pen-square"></i> 
          </button>
        </div>
        <hr class="hr">
        <h5 class="post-title">Posts</h5>     
        <hr class="hr">
        <div class="posts-section">
            <div v-if="!banState && nPosts > 0" class="photo-container">
              <div class="photo-wrapper">
                <Photo v-for="(photo, index) in photos"
                :key="index"
                :owner="this.$route.params.id"
                :photoId="photo.photoId"
                :comments="photo.comments"
                :likes="photo.likes"
                :dateAndTime="photo.dateAndTime"
                :cUserIsOwner="cUser"
                @removePhoto="deletePhoto"
                />
              </div>
            </div>
            <div v-else class="empty-container">
                <h6 class="empty">No Posts</h6>
            </div>
        </div>
        <div class="err-container">
          <ErrorMsg v-if="errMsg" :msg="errMsg"></ErrorMsg>
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
.exit {
  border: none;
  background: transparent;
}
.exit:hover {
  transform: scale(1.5);
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
.cbtn {
  border: none;
  color: white;
}
.cbtn:hover {
  transform: scale(1.1)
}
.fbtn {
  background: #3f729b;
}
.bbtn {
  background: #C01111;
}
.edit-icon {
  color: rgb(34, 34, 154);
  font-size: 25px;
}
.element:last-child {
  margin-right: 10px;
}
.lastof {
  border: none;
  background: transparent;
}
.lastof:hover {
  transform: scale(1.2);
}
.post-title {
  text-align: start;
  align-items: center;
  flex-direction: column;
  font-size: 20px;
  color: black;
}
.posts-section {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  grid-gap: 70px;
  margin-left: 120px;
}
.photo-wrapper {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  grid-gap: 70px;
}
.empty-container {
  display: flex;
  align-items: start;
  justify-content: start;
}
.empty {
  align-items: center;
  text-align: start;
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