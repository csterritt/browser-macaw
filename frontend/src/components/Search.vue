<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.query" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="query">Query</button>
    </div>
  </main>
</template>

<script setup>
import {reactive} from 'vue'
import {Query} from '../../wailsjs/go/main/App'

const data = reactive({
  query: "",
  resultText: "Please enter your query below",
})

function query() {
  Query(data.query).then(result => {
    data.resultText = result
  })
}
</script>
