<template>
  <template v-if="store.results.length > 0">
    <div class="italic mb-2">Found {{ store.results.length }} sites:</div>
    <div id="result" class="result">
      <domain-and-links
        v-for="domain in store.results"
        :key="domain['DomainName']"
        :domain="domain"
      >
      </domain-and-links>
    </div>
  </template>

  <template v-else>
    <div v-if="store.oneQueryRun">Nothing found for that query.</div>
  </template>
</template>

<script setup>
import { useStore } from '../store'

import DomainAndLinks from './DomainAndLinks.vue'

const store = useStore()
</script>
