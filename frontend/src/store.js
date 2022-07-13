import { defineStore } from 'pinia'
import { Query } from '../wailsjs/go/main/App'

import { ACTIVE_TAB_CLASS, INACTIVE_TAB_CLASS } from './constants'

const StandardizeString = /[\r\n\t ]+/g

const standardizeArg = (str) => {
  return str.replace(StandardizeString, ' ').trim()
}

// useStore could be anything like useUser, useCart
// the first argument is a unique id of the store across your application
export const useStore = defineStore('main', {
  state: () => ({
    queryTabClass: ACTIVE_TAB_CLASS,
    resultsTabClass: INACTIVE_TAB_CLASS,
    aboutTabClass: INACTIVE_TAB_CLASS,
    oneQueryRun: false,
    queryWords: '',
    queryAllWords: '',
    exactPhrase: '',
    mustNotWords: '',
    inUrl: '',
    results: [],
    errorFound: null,
  }),

  getters: {
    queryTabActive: (state) => state.queryTabClass === ACTIVE_TAB_CLASS,
    resultsTabActive: (state) => state.resultsTabClass === ACTIVE_TAB_CLASS,
    aboutTabActive: (state) => state.aboutTabClass === ACTIVE_TAB_CLASS,
    invalidInput: (state) => {
      const queryWords = standardizeArg(state.queryWords)
      const queryAllWords = standardizeArg(state.queryAllWords)
      const exactPhrase = standardizeArg(state.exactPhrase)
      const inUrl = standardizeArg(state.inUrl)

      return !(
        queryWords.length > 0 ||
        queryAllWords.length > 0 ||
        exactPhrase.length > 0 ||
        inUrl.length > 0
      )
    },
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

    copyUrlToClipboard(url) {
      const type = 'text/plain'
      const blob = new Blob([url], { type })
      const data = [new ClipboardItem({ [type]: blob })]

      navigator.clipboard.write(data).catch((err) => {
        // this.errorFound = `Unable to copy to the clipboard: ${err}`
        console.log(`Unable to copy to the clipboard: ${err}`)
      })
    },

    runQuery() {
      const queryWords = standardizeArg(this.queryWords)
      const queryAllWords = standardizeArg(this.queryAllWords)
      const exactPhrase = standardizeArg(this.exactPhrase)
      const inUrl = standardizeArg(this.inUrl)
      const mustNotWords = standardizeArg(this.mustNotWords)

      if (
        queryWords.length > 0 ||
        queryAllWords.length > 0 ||
        exactPhrase.length > 0 ||
        inUrl.length > 0 ||
        (queryWords.length > 0 && mustNotWords.length > 0)
      ) {
        return Query({
          Words: queryWords,
          AllWords: queryAllWords,
          ExactPhrase: exactPhrase,
          InUrl: inUrl,
          MustNotWords: mustNotWords,
        })
          .then((result) => {
            this.oneQueryRun = true
            this.makeResultsTabActive()
            if (result.length === 0) {
              this.results = []
            } else {
              this.results = result
            }
          })
          .catch((err) => {
            console.log(`Query returned error ${JSON.stringify(err, null, 2)}`)
            try {
              const json = JSON.parse(err)
              console.log(`Error as JSON is ${JSON.stringify(json, null, 2)}`)
              this.errorFound = json
            } catch (jsonErr) {
              console.log(
                `Oops, caught a JSON error trying to parse that error.`
              )
              this.errorFound = {
                message: 'Error found (but could not be parsed).',
              }
            }
          })
      }
    },
  },
})
