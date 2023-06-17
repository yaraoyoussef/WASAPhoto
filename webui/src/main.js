import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import NotFound from './components/NotFound.vue'
import Photo from './components/Photo.vue'
import CommentOnPost from './CommentOnPost.vue'
import CommentModal from './CommentModal.vue'
import User from './User.vue'
import SideBar from './SideBar.vue'

import './assets/dashboard.css'
import './assets/main.css'
import '@fortawesome/fontawesome-free/css/all.css';


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("NotFound", NotFound);
app.component("Photo", Photo)
app.component("CommentOnPost", CommentOnPost)
app.component("CommentModal", CommentModal)
app.component("User", User)
app.component("SideBar", SideBar)
app.use(router)
app.mount('#app')
