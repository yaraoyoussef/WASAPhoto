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
			this.$router.replace(otherView)
		},
		search(val) {
			this.searchValue = val
			this.$router.replace("/search")
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
	<div class="container-fluid">
		<div class="row">
		<div class="col p-0">
			<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<router-link class="navbar-brand" to="/">WasaPhoto</router-link>
			<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
				aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarNav">
				<ul class="navbar-nav mr-auto">
				<!-- Other navbar items go here -->
				</ul>
				<form class="form-inline my-2 my-lg-0">
				<input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search"
					v-model="searchValue">
				<button class="btn btn-outline-primary my-2 my-sm-0" type="button" @click="search(searchValue)">
					Search
				</button>
				</form>
				<ul class="navbar-nav ml-auto">
				<li class="nav-item">
					<button class="btn btn-outline-danger" type="button" @click="logout(false)">
					Logout
					</button>
				</li>
				</ul>
			</div>
			</nav>
			<router-view></router-view>
		</div>
		</div>
	</div>
</template>
  
<style>
.container-fluid {
padding: 0;
}

.navbar {
background-color: #f8f9fa;
}

.navbar-brand {
font-weight: bold;
}

.form-inline {
margin-right: 10px;
}

.btn-outline-primary {
margin-left: 10px;
}

.btn-outline-danger {
margin-left: 10px;
}
</style>
  
  
