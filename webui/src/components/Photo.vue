<script>
export default {
    data() {
        return {
          photoSrc: "",
          liked: false,
          commentsList: [],
          likesList: []
        }
    },

    props: ['owner', 'likes', 'comments', 'dateAndTime', 'photoId'],

    methods: {
      getPhoto() {
        this.photoSrc = __API_URL__+"/users/"+this.owner+"/photos/"+this.photoId
      },

      async deletePhoto() {
        try {
          await this.$axios.delete("/users/"+this.owner+"/photos/"+this.photoId)
          this.$emit("removePhoto", this.photoId)
        } catch (e) {
        }
      }, 
// missing functions
      async changeLike() {
        if(this.cUserIsOwner) {return }
        const token = localStorage.getItem('token')
        try {
          if(!this.liked) {
            await this.$axios.put("/users/"+this.owner+"/photos/"+this.photoId+"/likes/"+token)
            this.likesList.push({
              userId: token,
              username: token
            })
          } else {
            await this.$axios.delete("/users/"+this.owner+"/photos/"+this.photoId+"/likes/"+token)
            this.likesList.pop()
          }
          this.liked = !this.liked;
        } catch (e) {}
      },
      
      deleteComment(val) {
        this.commentsList = this.commentsList.filter(item => item.commentId !== val)
      },

      addComment(comment) {
        this.commentsList.push(comment)
      }
    },

    async mounted() {
      await this.getPhoto()
      if(this.likes != null) {
        this.likesList = this.likes
      }
      // set liked to true if current user already liked the loaded picture
      if(this.likes != null) {
        this.liked = this.likesList.some(o => o.userId === localStorage.getItem('token'))
      }
      if(this.comments != null) {
        this.commentsList = this.comments
      }
    }
}
</script>

<template>
  <div class="container-fluid">
        <div class="photo-container">
          <hr>
          <div class="photo-head">
            <h1 class="user">{{ owner }}</h1>
            <button v-if="cUserIsOwner" class="delete-photo" @click="deletePhoto">
              <i class="delete-icon fas fa-trash"></i>
            </button>
          </div>
          <hr>
          <div class="picture-space">
            <img :src="photoSrc">
          </div>
          <hr>
          <div class="photo-bottom">
            <button class="like-btn">
              <i @click="changeLike" :class="'like '+(liked? 'fas fa-heart': 'far fa-heart')" ></i>
            </button>
            <button class="comment-btn">
              <i @click="addComment" class="comment far fa-comment"></i>
            </button>
            <h2 class="upload-info">Uploaded on {{ dateAndTime }}</h2>
          </div>
        </div>
  </div>
</template>

<style>
.container-fluid {
  background-color: white;
}
.photo-container {
  justify-content: center;
  width: 400px;
}
.photo-head {
  display: flex;
  justify-content: space-between;
}
.user {
  font-size: 20px;
}
.delete-icon {
  font-size: 15px;
}
.picture-space {
  height: 400px;
  justify-content: center;
}
.photo-bottom {
  display: flex;
}
.like {
  font-size: 20px;
}
.comment-btn {
  margin-left: 10px;
}
.comment {
  font-size: 20px;
}
.upload-info {
  margin-left: 100px;
  font-size: 15px;
}
</style>