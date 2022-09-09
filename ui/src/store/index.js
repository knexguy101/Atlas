import { createStore } from 'vuex'

// Create a new store instance.
const store = createStore({
  state () {
    return {
        profiles: {},
        tasks: {},
        proxies: []
    }
  },
  mutations: {
    addProfile(state, profile) {
      state.profiles[profile.id] = profile
    },
    addTask(state, task) {
      state.tasks[task.id] = task
    },
    setTaskStatus(state, id, status) {
      state.tasks[id].status = status
    },
    proxies(state, proxy) {
      state.proxies.push(proxy)
    }
  }
})

export default store