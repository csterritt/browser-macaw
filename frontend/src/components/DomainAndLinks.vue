<template>
  <div class="mb-4">
    <div
      class="text-xl text-bold mb-1 px-1 cursor-pointer"
      @click="toggleDomainVisible"
    >
      <div
        class="px-2 hover:bg-primary hover:text-primary-content hover:rounded-md"
      >
        {{ domain['DomainName'] }}

        <span class="text-base italic">({{ linkCount() }})</span>
      </div>
    </div>

    <template v-if="domainVisible">
      <template v-for="elem in domain['Links']" :key="elem.Uid">
        <result-link :elem="elem"></result-link>

        <div class="divider my-1 mx-3 last-of-type:hidden"></div>
      </template>
    </template>
  </div>
</template>

<script setup>
import { ref } from 'vue'

import ResultLink from './ResultLink.vue'

const props = defineProps({
  domain: Object,
})

const domainVisible = ref(false)

const toggleDomainVisible = () => {
  domainVisible.value = !domainVisible.value
}

const linkCount = () => {
  const count = props.domain['Links'].length
  const links = count === 1 ? 'link' : 'links'
  return `${count} ${links}`
}
</script>
