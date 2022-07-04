<template>
  <div class="cursor-pointer" @click="openUrl">
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
</template>

<script setup>
import { BrowserOpenURL } from '../../wailsjs/runtime'

const props = defineProps({
  elem: Object,
})

const openUrl = () => {
  console.log(`opening URL ${props.elem['Url']}`)
  BrowserOpenURL(props.elem['Url'])
}
</script>
