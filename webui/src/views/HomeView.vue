<script>
export default {
	data: function() {
		return {
			errormsg: null,
			photos: []
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		async load() {
			try {
				this.errormsg = null;
				let response = await this.$axios.get('/users/' + localStorage.getItem('id') + '/home')
				if (response.data != null) {
					this.photos = response.data
				}
			} catch (e) {
				this.errormsg = e.toString()
			}

		}
	},
	mounted() {
		this.load()
		this.refresh()
	}
}
</script>

<template>
	<div class="container">
		<div class="home-section screen">
            <h1 class="title">WASAPhoto</h1>
            <div class="photo-container">
                <Photo 
                v-for = "(photo, index) in photos"
                :key ="index"
                :owner="photo.owner"
                :photo_id="photo.photoId"
                :comments="photo.comments != nil ? photo.comments : []"
                :likes="photo.likes != nil ? photo.likes : []"
                :date="photo.date"
                />
            </div>
            <div v-if="photos.length==0" class="empty-container">
                <h2 class="empty">
                    No content, follow your friends to be able to view their newest posts
                </h2>
            </div>
            <div class="err-container">
                <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
            </div>
		</div>
	</div>
</template>

<style>
.home-section {
	background-color:white;
}
.title {
    text-align: start;
    flex-direction: column;
    align-items: center;
    height: 100%;
}
.photo-container {
    align-items: center;
    text-align: center;
    flex-direction: column;
	gap: 20px;
}
.empty {
    align-items: center;
    text-align: center;
    flex-direction: column;
	font-size: 30px;
	color: black;
}
.err-container {
    align-items: center;
    text-align: center;
    flex-direction: column;
    font-size: 20px;
    color: red;
}
</style>
