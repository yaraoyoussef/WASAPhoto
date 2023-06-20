<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data() {
		return {
			loggedIn: false,
			searchValue: ""
		}
	},
	methods: {
		logout(val) {
			this.loggedIn = val
			this.$router.replace("/login");
		},
		updateLogged(val) {
			this.loggedIn = val
		},
		updateView(otherView) {
			console.log("updateView func in app.vue")
			this.$router.replace(otherView)
		},
		search(val) {
			this.searchValue = val
			this.$router.replace("/search")
		}
	},

	computed: {
		showSideBar() {
			return this.$route.path!=='/login';
		}
	},

	created() {
		if(!localStorage.getItem('notFirstStart')) {
			localStorage.clear()
			localStorage.setItem('notFirstStart', true)
		}
	},

	mounted() {
		if(!localStorage.getItem('token')) {
			this.$router.replace("/login")
		} else {
			this.loggedIn = true
		}
	}
}
</script>

<template>
	<div class="container">
		<SideBar v-if="loggedIn && showSideBar"
		@logoutSidebar="logout"
		@updateView="updateView"
		@searchUsers="search"
		/>
		<div :class=" { 'main-content': showSideBar, 'main-content-full': !showSideBar }">
			<RouterView 
			@updateLoggedChild="updateLogged"
			@updateView="updateView"
			:searchValue="searchValue"
			/>
		</div>
	</div>
</template>
  
<style>
.container {
	display: flex;
}
.main-content {
	flex-grow: 1;
	margin-left: 250px;
}
.main-content-full {
	width: 100%;
}
</style>
  
  
