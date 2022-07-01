<template>
  <div class="mb-4">
    <div
      class="text-xl text-bold mb-1 px-1 cursor-pointer hover:bg-secondary hover:text-secondary-content"
      @click="toggleDomainVisible"
    >
      {{ domain['DomainName'] }}
      <span class="text-base italic">({{ linkCount() }})</span>
    </div>

    <template v-if="domainVisible">
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
    </template>
  </div>
</template>

<script setup>
import { ref } from 'vue'

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
