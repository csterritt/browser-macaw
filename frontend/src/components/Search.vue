<template>
  <main>
    <div class="tabs">
      <a :class="store.queryTabClass" @click="store.makeQueryTabActive">
        Query
      </a>

      <a :class="store.resultsTabClass" @click="store.makeResultsTabActive">
        Results
      </a>

      <a :class="store.aboutTabClass" @click="store.makeAboutTabActive">
        About
      </a>
    </div>

    <query-tab v-if="store.queryTabActive"></query-tab>

    <results-tab v-if="store.resultsTabActive"></results-tab>

    <about-tab v-if="store.aboutTabActive"></about-tab>
  </main>
</template>

<script setup>
import { useStore } from '../store'
import AboutTab from './AboutTab.vue'
import QueryTab from './QueryTab.vue'
import ResultsTab from './ResultsTab.vue'

const store = useStore()
</script>
