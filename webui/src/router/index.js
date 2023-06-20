import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Login.vue'
import NotFound from '../views/NotFound.vue'
import Profile from '../views/Profile.vue'
import Search from '../views/Search.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/', 
			redirect: '/login'
		},
		{
			path: '/login', 
			component: Login
		},
		{
			path: '/home', 
			component: HomeView
		},
		{
			path: '/users/:id', 
			component: Profile
		},
		{
			path: '/search',
			component: Search
		},
		{
			path: '/:catchAll(.*)',
			component: NotFound
		}
	]
})

export default router
