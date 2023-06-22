<script>
export default {
	data: function() {
		return {
			errormsg: null,
			photos: []
		}
	},
	methods: {
		async load() {
			try {
				this.errormsg = null;
				let response = await this.$axios.get('/users/' + localStorage.getItem('token') + '/home')
				if (response.data != null) {
					this.photos = response.data
				}
			} catch (e) {
				this.errormsg = e.toString()
			}

		},

		async upload() {
			const input = document.getElementById('photo-input');
			try {
				const file = input.files[0];
				const reader = new FileReader();
				reader.readAsArrayBuffer(file);

				reader.onload = async () => {
					let response = await this.$axios.post("/users/"+localStorage.getItem('token')+"/photos/", reader.result, {
						headers: {
							'Content-Type': file.type
						},
					}) 
				}

				input.value = null;

			} catch(e) {
				this.errormsg = e.toString()
			}
		}
	},
	mounted() {
		this.load()
	}
}
</script>

<template>
	<div class="container">
		<div class="home-section screen">
			<div class="add-post-section">
				<input type="file" id="photo-input" class="photo-input" accept=".png, .jpeg">
				<button @click="upload" class="uploader">Upload</button>
			</div>
			<div class="photo-wrapper">
				<div class="photo-container">
					<Photo 
					v-for = "(photo, index) in photos"
					:key ="index"
					:owner="photo.owner"
					:photoId="photo.photoId"
					:comments="photo.comments != nil ? photo.comments : []"
					:likes="photo.likes != nil ? photo.likes : []"
					:dateAndTime="photo.dateAndTime"
					/>
				</div>
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
.add-post-section {
  margin-bottom: 20px;
  margin-left: auto;
}
.uploader {
  font-size: 17px;
}
.photo-input {
	font-size: 17px;
}
.photo-wrapper {
	display: flex;
	justify-content: center;
	margin-left: 200px;
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
	font-style: italic;
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
