<template>
  <main>
    <div id="input" class="input-box">
      <input
        id="name"
        v-model="data.query"
        autocomplete="off"
        class="input"
        type="text"
      />
      <button class="btn" @click="query">Query</button>
    </div>

    <template v-if="runOneQuery">
      <template v-if="data.results.length > 0">
        <ul id="result" class="result">
          <li v-for="elem in data.results" :key="elem.Uid">
            <div>Title: {{ elem['Title'] }}</div>
            <div>Subtitle: {{ elem['Subtitle'] }}</div>
            <div>Url: {{ elem['Url'] }}</div>
          </li>
        </ul>
      </template>

      <template v-else>
        <div>Nothing found for that query.</div>
      </template>
    </template>
  </main>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { Query } from '../../wailsjs/go/main/App'

const data = reactive({
  query: '',
  results: [],
})
const runOneQuery = ref(false)

function query() {
  if (data.query.trim().length > 0) {
    runOneQuery.value = true
    Query(data.query.trim()).then((result) => {
      if (result.length === 0) {
        data.results = []
      } else {
        data.results = result
      }
    })
  }
}
</script>
