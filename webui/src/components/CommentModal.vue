<script>
export default {
    data() {
        return {
            commentTxt: "",
        }
    },
    props: ['modalId', 'commentsList', 'owner', 'photoId'],

    methods: {
        async addComment() {
            try {
                let response = await this.$axios.post("/users/"+this.owner+"/photos/"+this.photoId+"/comments", {
                    userId: localStorage.getItem('token'),
                    comment: this.commentTxt
                }, { headers: {
                    'Content-Type': 'application/json'
                    }
                })
                this.$emit('addComment', {
                    commentId: response.data.commentId,
                    photoId: this.photoId,
                    userId: localStorage.getItem('token'),
                    comment: this.commentTxt
                })
                this.commentTxt = ""
            } catch (e) {
                console.log(e.toString())
            }
        },
        deleteCommentParent(val) {
            this.$emit('deleteComment', val)
        },
        addCommentParent() {
            this.$emit('addComment', newFormattedComment)
        }
    }
}
</script>

<template>
    
</template>

<style>
</style>