<template>
  <template v-if="store.results.length > 0">
    <ul id="result" class="result">
      <li v-for="elem in store.results" :key="elem.Uid">
        <div>Title: {{ elem['Title'] }}</div>
        <div>Subtitle: {{ elem['Subtitle'] }}</div>
        <div>Url: {{ elem['Url'] }}</div>
      </li>
    </ul>
  </template>

  <template v-else>
    <div v-if="store.oneQueryRun">Nothing found for that query.</div>
  </template>
</template>

<script setup>
import { useStore } from '../store'

const store = useStore()
</script>
