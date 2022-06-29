<template>
  <template v-if="store.results.length > 0">
    <div id="result" class="result">
      <template v-for="elem in store.results" :key="elem.Uid">
        <div class="hover:bg-primary">
          <a :href="elem['Url']" rel="nofollow" target="_blank">
            <div v-if="elem['Title']" class="ml-2">{{ elem['Title'] }}</div>
            <div v-if="elem['Subtitle']" class="ml-2">
              {{ elem['Subtitle'] }}
            </div>
          </a>
        </div>

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
