import { defineStore } from 'pinia'
import { Query } from '../wailsjs/go/main/App'

// useStore could be anything like useUser, useCart
// the first argument is a unique id of the store across your application
export const useStore = defineStore('main', {
  state: () => ({
    oneQueryRun: false,
    query: '',
    results: [],
  }),

  actions: {
    runQuery: function () {
      if (this.query.trim().length > 0) {
        Query(this.query.trim()).then((result) => {
          this.oneQueryRun = true
          if (result.length === 0) {
            this.results = []
          } else {
            this.results = result
          }
        })
      }
    },
  },
})
