<template>
  <div id="input" class="input-box">
    <input
      id="name"
      v-model="store.query"
      autocomplete="off"
      class="input input-bordered w-full max-w-xs"
      type="text"
      @keyup.enter="store.runQuery()"
    />

    <button class="btn btn-primary ml-2" @click="store.runQuery()">Query</button>
  </div>
</template>

<script setup>
import { useStore } from '../store'

const store = useStore()
</script>
