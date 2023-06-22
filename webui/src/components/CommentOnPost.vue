<script>
export default {
    data() {
        return {
            cUser: ""
        }
    }, 
    props: ['content', 'author', 'owner', 'commentId', 'photoId'],
    methods: {
        async deleteComment() {
            try {
                await this.$axios.delete("/users/"+this.owner+"/photos/"+this.photoId+"/comments/"+this.commentId)
                this.$emit('removeComment', this.commentId)
            } catch(e) {
                console.log(e.toString())
            }
        }
    },
    mounted() {
        this.cUser = localStorage.getItem('token')
    }
}
</script>

<template>
    <div class="container-fluid">
        <div class="cmt-cnt">
            <hr>
            <div class="c-header">
                <h5>@{{author}}</h5>
                <button v-if="cUser === author || cUser === owner" class="del-btn" @click="deleteComment">
                    <i class="x-icon fas fa-times"></i>
                </button>
            </div>
            <div class="c-content">
                <h6 class="content-txt">{{content}}</h6>
            </div>
        </div>
    </div>
</template>

<style>
.container-fluid {
    background-color: white;
}
.cmt-cnt {
    width: 400px;
    justify-content: center;
}
.c-header {
    display: flex;
    justify-content: space-between; 
    font-size: 18px;
}
.content-txt {
    font-size: 15px;
}
.del-btn {
    font-size: 15px;
    border: none;
}
.del-btn:hover {
    border: none;
    color: red;
    transform: scale(1.1);
}
</style>