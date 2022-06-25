<template>
  <main>
    <div id="input" class="input-box">
      <input id="name" v-model="data.query" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="query">Query</button>
    </div>

    <ul id="result" class="result">
      <li v-for="elem in data.results" :key="elem.Uid">
        <div>Title: {{elem['Title']}}</div>
        <div>Subtitle: {{elem['Subtitle']}}</div>
        <div>Url: {{elem['Url']}}</div>
      </li>
    </ul>
  </main>
</template>

<script setup>
import {reactive} from 'vue'
import {Query} from '../../wailsjs/go/main/App'

const data = reactive({
  query: "",
  results: [],
})

function query() {
  Query(data.query).then(result => {
    console.log("Query returns result", result)
    if (result.length === 0) {
      data.results = []
    } else {
      data.results = result
    }
  })
}
</script>
