<script setup lang="ts">
const props = defineProps({
  src: {
    type: String,
    required: true
  }
})
</script>

<template>
  <img :src="src" style="width: 250px; position: absolute; bottom: 120px; right: 50px; height: auto; transform: rotate(-25deg);">
</template>
