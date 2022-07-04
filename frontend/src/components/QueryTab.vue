<template>
  <form class="input-box" @submit.prevent="store.runQuery()">
    <div class="form-control w-full max-w-xs">
      <label class="label">
        <span class="label-text">Find any of these words:</span>
      </label>

      <input
        v-model="store.queryWords"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />

      <label class="label mt-2">
        <span class="label-text">Find this exact phrase:</span>
      </label>

      <input
        v-model="store.exactPhrase"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />

      <label class="label mt-2">
        <span class="label-text">These words must occur in the results:</span>
      </label>

      <input
        v-model="store.mustWords"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />

      <label class="label mt-2">
        <span class="label-text"
          >These words must NOT occur in the results:</span
        >
      </label>

      <input
        v-model="store.mustNotWords"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />

      <label class="label mt-2">
        <span class="label-text">Words in the URL that must occur:</span>
      </label>

      <input
        v-model="store.onlyDomain"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />
    </div>

    <button class="btn btn-primary ml-2 mt-4" type="submit">Query</button>
  </form>
</template>

<script setup>
import { useStore } from '../store'

const store = useStore()
</script>
