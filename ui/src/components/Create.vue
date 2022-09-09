<script>
    export default {
        data() {
            return {
                size: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]
            }
        }
    }
</script>

<template>
    <div class="container">
        <div class="card shadow">
            <h6 class="card-title">Create Tasks</h6>
            <div class="card-body text-center">
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">PID</span>
                    <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3">
                </div>
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Size</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3">
                        <option v-for="s in size" :key="s" :value="s">{{ (3.5 + (s * .5)) }}</option>
                    </select>
                    <div class="input-group-text">
                        <input class="form-check-input mt-0 mx-1" type="checkbox" value="" aria-label="Checkbox for following text input">
                        Kids
                    </div>
                </div>
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Account</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3">
                        <option v-for="(value, key) in this.$store.state.accounts" :key="key" :value="value">{{ value.email }}</option>
                    </select>
                </div>
                <div class="input-group mb-3">
                    <span class="input-group-text" id="basic-addon3">Proxy</span>
                    <select class="form-select" aria-label="Default select example" aria-describedby="basic-addon3">
                        <option v-for="p in this.$store.state.proxies" :key="p" :value="p">{{ p.split(":")[0] || p }}</option>
                    </select>
                </div>
                <button class="btn btn-primary mt-4">Create Task</button>
            </div>
        </div>
    </div>
</template>