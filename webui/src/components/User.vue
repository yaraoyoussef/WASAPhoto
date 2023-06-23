<script>
export default {
    props: ['id'],
    methods: {
        chosen() {
            this.$emit('chosen', this.id)
        }
    },
}
</script>

<template>
    <div class="container">
        <div class="section">
            <div class="card-container">
                <div class="user-card border-start">
                    <div class="card-body">
                        <h5 @click="chosen" class="user-card-title">@{{id }}</h5>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
.container {
    margin-bottom: 2rem;
    margin-top: 2rem;
}

.section {
    display: flex;
    justify-content: center;
}

.user-card {
    background-color: #3f729b;
    color: white;
    border-style: solid;
    border-width: 1px;
    border-color: #000;
    padding: 10px;
}

.user-card-title {
    display: flex;
    justify-content: center;
} 
</style>