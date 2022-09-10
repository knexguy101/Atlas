<script>
    export default {
        data() {
            return {
                rawProxiesText: ""
            }
        },
        methods: {
            saveProxies() {
                console.log(this.rawProxiesText.split("\n"))
                this.$store.state.proxies = this.rawProxiesText.split("\n")
                window.saveProxies(this.$store.state.proxies)
            }
        },
        mounted() {
            this.rawProxiesText = this.$store.state.proxies.join("\n")
        }
    }
</script>

<template>
    <div class="container" style="min-width: 70rem;">
        <div class="card shadow just" style="min-height: 40rem;">
            <div class="card-body">
                <div class="form-floating">
                    <textarea class="form-control" placeholder="Leave a comment here" id="floatingTextarea2" style="height: 30rem; resize: none;" v-model="rawProxiesText"></textarea>
                    <label for="floatingTextarea2">Proxies (IP:PORT:USER:PASS)</label>
                </div>
                <button class="btn btn-success mt-4" style="width: 10rem;" @click="saveProxies()">Save</button>
            </div>
        </div>
    </div>
</template>