<script>
export default {
    data: function() {
        return {
            users: [],
            errMsg: null
        }
    },
    props: ['searchValue'],
    watch: {
        searchValue: function() {
            this.loadUsers()
        }
    },
    methods: {
        async loadUsers() {
            this.errMsg = null;
            if(this.searchValue === undefined || this.searchValue === "") {
                this.users = []
                return
            }
            try {
                console.log("we have this search value: ", this.searchValue)
                let response = await this.$axios.get("/users", {params: {id: this.searchValue}});
                console.log(response.data)
                this.users = response.data
            } catch (e) {
                this.errMsg = e.toString();
                console.log("error no users")
            }
        },
        getProfile(val) {
            this.$router.replace("/users/"+val)
        }
    },
    async mounted() {
        if(!localStorage.getItem('token')) {
            this.$router.replace("/login")
        }
        console.log("about to load users in search.vue")
        await this.loadUsers()
    }
}
</script>

<template>
    <div class="container-fluid">
        <User v-for="(user, index) in users"
        :key="index"
        :id="user.ID"
        @chosen="getProfile"
        />
        <p v-if="users.length==0" class="empty">No results</p>
        <ErrorMsg class="error" v-if="errMsg" :msg="errMsg"></ErrorMsg> 
    </div>
</template>

<style>
.empty {
    font-size: 20px;
    font-style: italic;
    color: black;
}
.error {
    font-size: 20px;
    color: red;
}
</style>