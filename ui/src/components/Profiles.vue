<script>
    export default {
        data() {
            return {
                states: [ 'AL', 'AK', 'AS', 'AZ', 'AR', 'CA', 'CO', 'CT', 'DE', 'DC', 'FM', 'FL', 'GA', 'GU', 'HI', 'ID', 'IL', 'IN', 'IA', 'KS', 'KY', 'LA', 'ME', 'MH', 'MD', 'MA', 'MI', 'MN', 'MS', 'MO', 'MT', 'NE', 'NV', 'NH', 'NJ', 'NM', 'NY', 'NC', 'ND', 'MP', 'OH', 'OK', 'OR', 'PW', 'PA', 'PR', 'RI', 'SC', 'SD', 'TN', 'TX', 'UT', 'VT', 'VI', 'VA', 'WA', 'WV', 'WI', 'WY' ],
                countries: [ 'US' ],
                months: ['01', '02', '03', '04', '05', '06', '07', '08', '09', '10', '11', '12'],
                years: ['22', '23', '24', '25', '26', '27', '28', '29', '30'],
                currentPage: 0,
                billingSame: true,
                selectedProfile: "",
                currentProfile: {
                    id: "",
                    title: "",
                    shipping: {
                        fname: "",
                        lname: "",
                        a1: "",
                        a2: "",
                        city: "",
                        zip: "",
                        province: "",
                        country: "",
                        phone: ""
                    },
                    billing: {
                        fname: "",
                        lname: "",
                        a1: "",
                        a2: "",
                        city: "",
                        zip: "",
                        province: "",
                        country: "",
                        phone: ""
                    },
                    payment: {
                        cc: "",
                        month: "01",
                        year: "22",
                        cvv: ""
                    }
                }
            }
        },
        methods: {
            changePage(num) {
                this.currentPage = num
            },
            async loadProfile() {
                var p = await window.getProfile(this.selectedProfile)
                if(p === null) {
                    return
                }

                this.currentProfile = p
            },
            async deleteProfile() {
                var res = await window.deleteProfile(this.selectedProfile)
                if(!res) {
                    return
                }

                this.$store.commit("removeProfile", this.selectedProfile)
            },
            async saveProfile() {
                var copyProfile = JSON.parse(JSON.stringify(this.currentProfile))
                if(this.billingSame) {
                    copyProfile.billing = copyProfile.shipping
                }
                var res = await window.saveProfile(copyProfile)
                if(!res) {
                    return
                }
                
                this.currentProfile = res
                this.$store.commit("addProfile", res)
            }
        }   
    }
</script>

<template>
    <div class="container profileCard">
        <div class="container">
            <div class="row">
                <div class="col-6 text-start">
                    <div class="input-group mb-3">
                        <select class="form-select" v-model="selectedProfile">
                            <option v-for="(value, key) of this.$store.state.profiles" :key="key" :value="value.id">{{value.title}}</option>   
                        </select>
                        <button class="btn btn-outline-primary" type="button" @click="loadProfile()">Load</button>
                        <button class="btn btn-outline-danger" type="button" @click="deleteProfile()">Delete</button>
                    </div>
                </div>
                <div class="col-6 text-end">
                    <div class="btn-group" role="group" aria-label="Basic radio toggle button group">
                        <input type="radio" class="btn-check" name="btnradio" id="btnradio1" autocomplete="off" checked @click="changePage(0)">
                        <label class="btn btn-outline-primary" for="btnradio1">Shipping</label>

                        <input type="radio" class="btn-check" name="btnradio" id="btnradio2" autocomplete="off" @click="changePage(1)">
                        <label class="btn btn-outline-primary" for="btnradio2">Billing</label>

                        <input type="radio" class="btn-check" name="btnradio" id="btnradio3" autocomplete="off" @click="changePage(2)">
                        <label class="btn btn-outline-primary" for="btnradio3">Payment</label>
                    </div>
                </div>
            </div>
        </div>
        <div v-if="currentPage == 0" class="card shadow mt-4">
            <div class="card-body">
                <div class="container">
                    <div class="row justify-content-center">
                        <div class="col-4">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Title</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.title">
                            </div>
                        </div>
                    </div>
                    <div class="row mt-2">
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">First Name</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.fname">
                            </div>
                        </div>
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Last Name</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.lname">
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Street</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.a1">
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Apt</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.a2">
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Zip</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.zip">
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-4">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">City</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.city">
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Phone</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.shipping.phone">
                            </div>
                        </div>
                        <div class="col-2">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">State</span>
                                <select class="form-select" aria-label="Default select example" v-model="currentProfile.shipping.province">
                                    <option v-for="s of states" :key="s" :value="s" >{{s}}</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Country</span>
                                <select class="form-select" aria-label="Default select example" v-model="currentProfile.shipping.country">
                                    <option v-for="c of countries" :key="c" :value="c" >{{c}}</option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-else-if="currentPage == 1" class="card shadow mt-4">
            <div class="card-body">
                <div class="container">
                    <div class="row justify-content-start">
                        <div class="col-4">
                            <div class="form-check text-start">
                                <input class="form-check-input" type="checkbox" value="" id="flexCheckChecked" :checked="billingSame">
                                <label class="form-check-label" for="flexCheckChecked">
                                    Billing same as shipping
                                </label>
                            </div>
                        </div>
                    </div>
                    <div class="row mt-2">
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">First Name</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.fname">
                            </div>
                        </div>
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Last Name</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.lname">
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Street</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.a1">
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Apt</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.a2">
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Zip</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.zip">
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-4">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">City</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.city">
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Phone</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.billing.phone">
                            </div>
                        </div>
                        <div class="col-2">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">State</span>
                                <select class="form-select" aria-label="Default select example" v-model="currentProfile.billing.province">
                                    <option v-for="s of states" :key="s" :value="s" >{{s}}</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-3">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Country</span>
                                <select class="form-select" aria-label="Default select example" v-model="currentProfile.billing.country">
                                    <option v-for="c of countries" :key="c" :value="c" >{{c}}</option>
                                </select>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div v-else class="card shadow mt-4">
            <div class="card-body">
                <div class="container">
                    <div class="row mt-2">
                        <div class="col-6">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Card</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.payment.cc">
                            </div>
                        </div>
                        <div class="col-2">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Month</span>
                                <select class="form-select" aria-label="Default select example" v-model="currentProfile.payment.month">
                                    <option v-for="s of months" :key="s" :value="s" >{{s}}</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-2">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">Year</span>
                                <select class="form-select" aria-label="Default select example" v-model="currentProfile.payment.year">
                                    <option v-for="s of years" :key="s" :value="s" >{{s}}</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-2">
                            <div class="input-group mb-3">
                                <span class="input-group-text" id="basic-addon3">CVV</span>
                                <input type="text" class="form-control" id="basic-url" aria-describedby="basic-addon3" v-model="currentProfile.payment.cvv">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <button class="btn btn-success mt-4" style="width: 10rem;" @click="saveProfile()">Save</button>
    </div>
</template>

<style>
    .profileCard {
        min-width: 70rem;
    }
</style>