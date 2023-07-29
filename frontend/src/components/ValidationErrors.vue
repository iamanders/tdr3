<template>
    <div v-if="Object.keys(errors ?? []).length > 0" class="alert alert-danger mb-5">
        <ul>
            <li v-for="error in errors" :key="error">{{ error }}</li>
        </ul>
    </div>
</template>

<script setup>
import { defineProps } from 'vue';

defineProps({
  errors: Object,
});

</script>
