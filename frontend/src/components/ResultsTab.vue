<template>
  <template v-if="store.results.length > 0">
    <div id="result" class="result">
      <section v-for="domain in store.results" :key="domain" class="mb-4">
        <div
          class="text-xl text-bold mb-1 px-1 hover:bg-secondary hover:text-secondary-content"
        >
          {{ domain['DomainName'] }}
        </div>

        <template v-for="elem in domain['Links']" :key="elem.Uid">
          <div class="hover:bg-primary hover:text-primary-content">
            <a :href="elem['Url']" rel="nofollow" target="_blank">
              <div v-if="elem['Title']" class="ml-3">{{ elem['Title'] }}</div>
              <div v-if="elem['Subtitle']" class="ml-3">
                {{ elem['Subtitle'] }}
              </div>
            </a>
          </div>

          <div class="divider my-1 mx-3 last-of-type:hidden"></div>
        </template>
      </section>
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
