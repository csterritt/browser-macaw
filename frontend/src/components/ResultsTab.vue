<template>
  <template v-if="store.results.length > 0">
    <div id="result" class="result">
      <template v-for="elem in store.results" :key="elem.Uid">
        <a :href="elem['Url']" rel="nofollow" target="_blank" class="ml-1">
          <div v-if="elem['Title']">{{ elem['Title'] }}</div>
          <div v-if="elem['Subtitle']">{{ elem['Subtitle'] }}</div>
        </a>

        <div class="divider my-1 last-of-type:hidden"></div>
      </template>
    </div>
  </template>

  <template v-else>
    <div v-if="store.oneQueryRun">Nothing found for that query.</div>
  </template>
</template>

<script setup>
import { useStore } from '../store'

const store = useStore()
</script>
