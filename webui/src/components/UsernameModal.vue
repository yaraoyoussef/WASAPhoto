<script>
export default {
    data() {
        return {
            username: "",
        }
    },
    props:['modalId'],

    methods: {
        editUsername() {
            if(this.username) {
                this.$emit("editUsername", this.username);
                this.username = "";
            }
        },
    }
}
</script>

<template>
    <div class="modal fade my-modal" :id="modalId" tabindex="-1" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="title" :id="modalId">Edit Username</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <div class="row">
                        <div class="col">
                            <div class="container">
                                <input class="form-control" type="text" id="FormControlTextAera"
                                placeholder="new username" maxlength="15" v-model="username">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                    <div class="submit-btn">
                        <button type="button" class="btn btn-primary" 
                        @click.prevent="editUsername"
                        :disabled="username===null || username.trim().length>15 || username.trim().length<3">
                        Save
                        </button>
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
    top: 10px;
    right: 5px;
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
