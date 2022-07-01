<template>
  <div class="mb-4">
    <div
      class="text-xl text-bold mb-1 px-1 hover:bg-secondary hover:text-secondary-content"
      @click="toggleDomainVisible"
    >
      {{ domain['DomainName'] }}
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
import { useStore } from '../store'

const store = useStore()

const props = defineProps({
  domain: Object,
})

const domainVisible = ref(false)

const toggleDomainVisible = () => {
  domainVisible.value = !domainVisible.value
}
</script>
