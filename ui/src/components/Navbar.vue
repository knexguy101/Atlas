<script>
    import { BIconWindowDash, BIconPerson, BIconServer, BIconDpad, BIconPlusCircle } from 'bootstrap-icons-vue'
    
    export default {
        components: {
            BIconWindowDash,
            BIconPerson,
            BIconServer,
            BIconDpad,
            BIconPlusCircle
        },
        methods: {
            redirect(endpoint) {
                this.$router.push("/" + endpoint)
            },
            github() {
                window.open("https://github.com/knexguy101/Atlas", "_blank", "resizable=yes, scrollbars=yes, titlebar=yes, width=800, height=900, top=10, left=10");
            }
        }
    }
</script>
    
<template>
    <div class="fixed-top">
        <nav class="navbar navbar-light bg-light shadow">
            <div class="container-fluid">
                <div class="justify-content-start">
                    <button class="btn btn-outline-primary me-2" type="button" style="min-width: 8rem;" @click="redirect('')"><BIconPlusCircle /> Create Task</button>
                    <button class="btn btn-outline-primary me-2" type="button" style="min-width: 8rem;" @click="redirect('tasks')"><BIconWindowDash /> Tasks</button>
                    <button class="btn btn-outline-primary me-2" type="button" style="min-width: 8rem;" @click="redirect('profiles')"><BIconPerson /> Profiles</button>
                    <button class="btn btn-outline-primary me-2" type="button" style="min-width: 8rem;" @click="redirect('proxies')"><BIconServer /> Proxies</button>
                    <button class="btn btn-outline-primary me-2" type="button" style="min-width: 8rem;" @click="redirect('accounts')"><BIconDpad /> Accounts</button>
                </div>
                <div class="justify-content-end mx-2">
                    <img class="icon" src="../assets/svg.svg" @click="github()"/>
                </div>
            </div>
        </nav>
    </div>
</template>

<style>
    .icon {
        max-width: 40px;
        max-height: 40px;
        cursor: pointer;
    }
</style>