<script>
    import Navbar from './components/Navbar.vue'

    export default {
        components: {
            Navbar
        },
        mounted() {
            window.appLoaded()
        }
    }
</script>

<template>
    <Navbar></Navbar>
    <router-view></router-view>
</template>