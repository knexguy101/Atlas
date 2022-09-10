<script>
    import { UNREF } from '@vue/compiler-core'
import { BIconPlay, BIconStop, BIconTrash } from 'bootstrap-icons-vue'

    export default {
        components: {
            BIconPlay,
            BIconStop,
            BIconTrash
        },
        methods: {
            async startTask(id) {
                let res = await window.startTask(id)
                if(!res) {
                    return
                }
            },
            async stopTask(id) {
                let res = await window.stopTask(id)
                if(!res) {
                    return
                }
            },
            async deleteTask(id) {
                let res = await window.deleteTask(id)
                if(!res) {
                    return
                }

                this.$store.commit("removeTask", id)
            }
        }
    }
</script>

<template>
    <div class="container wide">
        <div class="mt-5"></div>
        <div class="card mt-2" v-for="(value, key) of this.$store.state.tasks" :key="key">
            <div class="card-body">
                <div class="container">
                    <div class="row">
                        <div class="col-1 text-truncate text-start my-auto">
                            <h5 class="text-primary">{{ (value.isProfileTask ? "Profile" : "SNKRS") }}</h5>
                        </div>
                        <div class="col-2 text-truncate text-start my-auto">
                            <h5 class="text-secondary">{{ value.sku }}</h5>
                        </div>
                        <div class="col-3 text-truncate my-auto">
                            <h6>{{value.accountEmail}}</h6>
                        </div>
                        <div class="col-3 text-truncate my-auto">
                            <h6>{{value.status}}</h6>
                        </div>
                        <div class="col-3 text-end">
                            <div class="btn-group" role="group" aria-label="Basic example">
                                <button class="badge bg-primary imageButton" @click="startTask(value.id)"><BIconPlay /> Start</button>
                                <button class="badge bg-warning text-black mx-1 imageButton" @click="stopTask(value.id)"><BIconStop /> Stop</button>
                                <button class="badge bg-danger imageButton" @click="deleteTask(value.id)"><BIconTrash /> Delete</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style>
    .wide {
        min-width: 70rem;
    }

    .imageButton {
        height: 24px;
    }
</style>