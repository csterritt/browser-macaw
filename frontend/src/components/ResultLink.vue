<template>
  <div class="flex flex-row justify-between ml-2">
    <div
      class="grow max-w-full overflow-hidden mr-4 cursor-pointer hover:bg-secondary hover:text-secondary-content hover:rounded-md"
      @click="openUrl"
    >
      <div v-if="elem['Title']" class="ml-3">{{ elem['Title'] }}</div>
      <div v-if="elem['Subtitle']" class="ml-3">
        {{ elem['Subtitle'] }}
      </div>
      <div
        v-if="elem['BodyPart']"
        class="ml-6 text-ellipsis overflow-hidden whitespace-nowrap"
      >
        {{ elem['BodyPart'] }}
      </div>
    </div>

    <button
      class="btn btn-primary"
      @click="store.copyUrlToClipboard(elem['Url'])"
    >
      Copy
    </button>
  </div>
</template>

<script setup>
import { useStore } from '../store'

import { BrowserOpenURL } from '../../wailsjs/runtime'

const props = defineProps({
  elem: Object,
})

const store = useStore()

const openUrl = () => {
  console.log(`opening URL ${props.elem['Url']}`)
  BrowserOpenURL(props.elem['Url'])
}
</script>
