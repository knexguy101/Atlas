<script>
    export default {
        data() {
            return {
                size: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25],
                currentTask: {
                    sku: "",
                    rawSize: 0,
                    rawSizeKids: false,
                    rawProxy: "",
                    accountId: "",
                    profileId: "",
                    isProfileTask: false,
                    startTime: 0,
                    monitorRelease: true
                }
            }
        },
        methods: {
            async createTask() {
                var res = await window.createTask(this.currentTask)
                if(res === null) {
                    return
                }

                this.$store.commit("addTask", res)
            }
        }
    }
</script>

<template>
    <div class="container">
        <div class="card shadow semiWide">
            <h6 class="card-title">Create Tasks</h6>
            <div class="card-body text-center">
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">PID</span>
                    <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentTask.sku">
                </div>
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Size</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3" v-model="currentTask.rawSize">
                        <option v-for="s in size" :key="s" :value="s">{{ (3.5 + (s * .5)) }}</option>
                    </select>
                    <div class="input-group-text">
                        <input class="form-check-input mt-0 mx-1" type="checkbox" value="" aria-label="Checkbox for following text input" v-model="currentTask.rawSizeKids">
                        Kids
                    </div>
                </div>
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Account</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3" v-model="currentTask.accountId">
                        <option v-for="(value, key) in this.$store.state.accounts" :key="key" :value="value.id">{{ value.email }}</option>
                    </select>
                </div>
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Proxy</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3" v-model="currentTask.rawProxy">
                        <option v-for="p in this.$store.state.proxies" :key="p" :value="p">{{ p.split(":")[0] || p }}</option>
                    </select>
                </div>
                <div class="form-check text-start mb-1">
                    <input class="form-check-input" type="checkbox" value="" id="flexCheckChecked" v-model="currentTask.monitorRelease">
                    <label class="form-check-label" for="flexCheckChecked">
                        Monitor Release
                    </label>
                </div>
                <div class="form-check text-start mb-1">
                    <input class="form-check-input" type="checkbox" value="" id="flexCheckChecked" v-model="currentTask.isProfileTask">
                    <label class="form-check-label" for="flexCheckChecked">
                        Profile Task
                    </label>
                </div>
                <div v-if="currentTask.isProfileTask" class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Profile</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3" v-model="currentTask.profileId">
                        <option v-for="p in this.$store.state.profiles" :key="p" :value="p.id">{{ p.title }}</option>
                    </select>
                </div>
                <button class="btn btn-primary mt-4" @click="createTask()">Create Task</button>
            </div>
        </div>
    </div>
</template>

<style>
    .semiWide {
        min-width: 30rem;
    }
</style>