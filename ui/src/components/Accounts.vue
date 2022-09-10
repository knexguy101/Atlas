<script>
    import { BIconTrash } from 'bootstrap-icons-vue'

    export default {
        data() {
            return {
                importLines: ""
            }
        },
        components: {
            BIconTrash
        },
        methods: {
            async importAccounts() {
                var res = await window.addAccounts(this.importLines.split("\n"))
                if(!res) {
                    return
                }

                for(var i of res) {
                    this.$store.commit("addAccount", i)
                }
            },
            async deleteAccount(id) {
                var res = await window.deleteAccount(id)
                if(!res) {
                    return 
                }

                this.$store.commit("removeAccount", id)
            }
        }
    }
</script>

<template>
    <div class="container" style="min-width: 70rem;">
        <div class="mt-5"></div>
        <div class="card shadow" style="min-height: 20rem;">
            <div class="card-body">
                <div class="form-floating">
                    <textarea class="form-control" placeholder="Email:Password" id="floatingTextarea2" style="height: 20rem; resize: none;" v-model="importLines"></textarea>
                    <label for="floatingTextarea2">Accounts (EMAIL:PASSWORD)</label>
                </div>
                <button class="btn btn-success mt-4" style="width: 10rem;" @click="importAccounts()">Import</button>
            </div>
        </div>
        <div class="mt-5"></div>
        <div class="container">
            <div class="row">
                <div v-for="(value, key) of this.$store.state.accounts" :key="key" class="col-6">
                    <div class="card mt-2">
                        <div class="card-body">
                            <div class="container">
                                <div class="row">
                                    <div class="col-5 text-start my-auto">
                                        <h6>{{ value.email }}</h6>
                                    </div>
                                    <div class="col-5 text-start my-auto">
                                        <h6 type="password">{{ value.password }}</h6>
                                    </div>
                                    <div class="col-2 text-start">
                                        <button class="btn btn-primary" @click="deleteAccount(value.id)"><BIconTrash /></button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>