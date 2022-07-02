import { defineStore } from 'pinia'
import { Query } from '../wailsjs/go/main/App'

import { ACTIVE_TAB_CLASS, INACTIVE_TAB_CLASS, SHOW_RESULTS } from './constants'

// useStore could be anything like useUser, useCart
// the first argument is a unique id of the store across your application
export const useStore = defineStore('main', {
  state: () => ({
    queryTabClass: ACTIVE_TAB_CLASS,
    resultsTabClass: INACTIVE_TAB_CLASS,
    aboutTabClass: INACTIVE_TAB_CLASS,
    oneQueryRun: false,
    queryWords: '',
    exactPhrase: '',
    mustWords: '',
    mustNotWords: '',
    onlyDomain: '',
    results: [],
  }),

  getters: {
    queryTabActive: (state) => state.queryTabClass === ACTIVE_TAB_CLASS,
    resultsTabActive: (state) => state.resultsTabClass === ACTIVE_TAB_CLASS,
    aboutTabActive: (state) => state.aboutTabClass === ACTIVE_TAB_CLASS,
  },

  actions: {
    makeQueryTabActive() {
      this.queryTabClass = ACTIVE_TAB_CLASS
      this.resultsTabClass = INACTIVE_TAB_CLASS
      this.aboutTabClass = INACTIVE_TAB_CLASS
    },

    makeResultsTabActive() {
      this.queryTabClass = INACTIVE_TAB_CLASS
      this.resultsTabClass = ACTIVE_TAB_CLASS
      this.aboutTabClass = INACTIVE_TAB_CLASS
    },

    makeAboutTabActive() {
      this.queryTabClass = INACTIVE_TAB_CLASS
      this.resultsTabClass = INACTIVE_TAB_CLASS
      this.aboutTabClass = ACTIVE_TAB_CLASS
    },

    runQuery() {
      if (this.queryWords.trim().length > 0) {
        Query({ Words: this.queryWords.trim() }).then((result) => {
          this.oneQueryRun = true
          this.makeResultsTabActive()
          if (result.length === 0) {
            this.results = []
          } else {
            this.results = result
          }
        })
      } else if (this.exactPhrase.trim().length > 0) {
        Query({ ExactPhrase: this.exactPhrase.trim() }).then((result) => {
          this.oneQueryRun = true
          this.makeResultsTabActive()
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
