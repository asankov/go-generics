<script setup lang="ts">
const props = defineProps({
  hrefs: {
    type: Array as () => string[],
    required: true
  }
})
</script>

<template>
  <div style="position: fixed; left: 10px; bottom: 5px; padding: 10px; margin: 0;">
    <div v-for="(href, index) in props.hrefs" :key="index">
      <a :href="href" target="_blank">{{ href }}</a>
    </div>
  </div>
</template>