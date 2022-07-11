<template>
  <form class="input-box" @submit.prevent="runTheQuery">
    <div class="form-control w-full max-w-xs">
      <label class="label">
        <span class="label-text">
          Find all of these words (they all must occur):
        </span>
      </label>

      <input
        v-model="store.queryAllWords"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />

      <label class="label">
        <span class="label-text">
          Find any of these words (any of them may occur):
        </span>
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
        <span class="label-text">
          These words must NOT occur in the results:
        </span>
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
        v-model="store.inUrl"
        autocomplete="off"
        class="input input-bordered w-full max-w-xs"
        type="text"
      />
    </div>

    <button
      class="btn btn-primary ml-2 mt-4"
      type="submit"
      :disabled="store.invalidInput"
    >
      Query
    </button>
  </form>

  <input
    type="checkbox"
    id="query-tab-error-dialog"
    class="modal-toggle"
    :checked="showErrorDialog"
  />
  <div class="modal">
    <div class="modal-box relative">
      <label
        for="query-tab-error-dialog"
        class="btn btn-sm btn-circle absolute right-2 top-2"
      >
        ✕
      </label>
      <h3 class="text-lg font-bold">Something bad happened:</h3>
      <p class="py-4">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

import { useStore } from '../store'

const store = useStore()

const showErrorDialog = ref(false)
const errorMessage = ref('')

const runTheQuery = () => {
  console.log('Running the query...')
  store.runQuery().then(() => {
    if (store.errorFound != null) {
      errorMessage.value = store.errorFound['message']
      showErrorDialog.value = true
    }
  })
}
</script>
