<template>
    <div class="bar" :class="{ 'active': isLoading }"></div>
</template>

<script setup>
import { isLoading } from '../stores/loader';
</script>

<style scoped>
    .bar {
        background-color: transparent;
        position: absolute;
        top: 3px;
        left: 3px;
        right: 3px;
        height: 6px;
        border-radius: 3px;
    }
    .bar.active {
        background-color: var(--color-primary);
        animation: pulse .66s infinite;
    }

    @keyframes pulse {
        0% {
            background-color: transparent;
        }
        70% {
            background-color: var(--color-primary);
        }
        100% {
            background-color: transparent;
        }
    }
</style>
