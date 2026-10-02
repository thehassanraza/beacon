<template>
  <img
    src="/brand-mark.png"
    alt=""
    :width="size"
    :height="size"
    class="brand-mark"
    aria-hidden="true"
    draggable="false"
  />
</template>

<script setup>
defineProps({
  size: { type: [Number, String], default: 28 },
});
</script>

<style scoped>
.brand-mark {
  flex-shrink: 0;
  display: block;
  object-fit: contain;
  user-select: none;
}
</style>
