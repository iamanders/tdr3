/* eslint-disable no-return-assign */
import { ref, computed } from 'vue';

const loading = ref(false);

export const isLoading = computed(() => loading.value);

export const setLoading = (value) => loading.value = value;
