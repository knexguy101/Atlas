import { createStore } from 'vuex'

// Create a new store instance.
const store = createStore({
  state () {
    return {
        profiles: {},
        tasks: {},
        accounts: {},
        proxies: []
    }
  },
  mutations: {
    addProfile(state, profile) {
      state.profiles[profile.id] = profile
    },
    removeProfile(state, id) {
      delete state.profiles[id]
    },
    addTask(state, task) {
      state.tasks[task.id] = task
    },
    removeTask(state, id){
      delete state.tasks[id]
    },
    setTaskStatus(state, st) {
      console.log(st)
      state.tasks[st.id].status = st.msg
    },
    addProxy(state, p) {
      state.proxies.push(p)
    },
    addAccount(state, account) {
      state.accounts[account.id] = account
    },
    removeAccount(state, id) {
      delete state.accounts[id]
    }
  }
})

export default store