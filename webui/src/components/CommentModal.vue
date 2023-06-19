<script>
export default {
    data() {
        return {
            commentTxt: "",
        };
    },
    props: ["modalId", "commentsList", "owner", "photoId"],
    methods: {
        async addComment() {
            try {
                let response = await this.$axios.post("/users/" + this.owner + "/photos/" + this.photoId + "/comments", {
                    commentId: -1,
                    userId: localStorage.getItem("token"),
                    comment: this.commentTxt
                }, { headers: {
                        "Content-Type": "application/json"
                    }
                });
                this.$emit("addComment", {
                    commentId: response.data.commentId,
                    userId: localStorage.getItem("token"),
                    comment: this.commentTxt
                });
                this.commentTxt = "";
            }
            catch (e) {
                console.log(e.toString());
            }
        },
        removeCommentParent(val) {
            this.$emit("removeComment", val);
        },
        addCommentParent() {
            this.$emit("addComment", newFormattedComment);
        }
    },
}
</script>

<template>
    <div class="modal fade my-modal" :id="modalId" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-cntnt">
                <div class="modal-header">
                    <h1 class="title" :id="modalId">Comments</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <CommentOnPost v-for="(comment, index) in commentsList"
                    :key="index"
                    :author="comment.userId"
                    :commentId="comment.commentId"
                    :photoId= "photoId"
                    :content="comment.comment"
                    :owner="owner"

                    @removeComment="removeCommentParent"
                    />
                </div>
                <div class="modal-footer">
                    <div class="row">
                        <div class="col">
                            <div class="container">
                                <textarea class="form-control" id="FormControlTextAera"
                                placeholder="Add a comment" rows="1" maxlength="500" v-model="commentTxt">
                                </textarea>
                            </div>
                        </div>
                        <div class="submit-btn">
                            <button type="button" class="btn btn-primary" 
                            @click.prevent="addComment"
                            :disabled="commentTxt.length < 1 || commentTxt.length > 500">
                            Post
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.my-modal {
    display: none;
}
.title {
    font-size: 25px;
}
.btn-close {
    position: absolute;
    top: 0;
    right: 0;
}
.modal-dialog {
    display: flex;
    align-items: center;
    justify-content: center;
    overflow-y: auto;
}
.modal-footer {
    display: flex;
    justify-content: center;
    width: 100%;
}
.row {
    width: 100%;
}
.container {
    margin-bottom: 5px;
    margin-block-end: auto;
}
.submit-btn {
    display: flex;
    align-items: center;
}
</style>