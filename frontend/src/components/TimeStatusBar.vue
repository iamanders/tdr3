<!-- eslint-disable max-len -->
<template>
  <div v-if="status != null && status.hasRunningTime" class="time-status-bar container">
    <div class="panel p-4 mt-6 d-flex align-center">

      <svg xmlns="http://www.w3.org/2000/svg" width="26" height="26" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"></circle><polyline points="12 6 12 12 16 14"></polyline></svg>

      <div class="d-flex align-start flex-direction-column ml-2">
        <h3>{{ status.time?.client?.title }}, {{ status.time?.project?.title }}</h3>
        <p v-if="status.time?.code || status.time?.comment">
          {{ status.time?.code }}
          {{ status.time?.code && status.time?.comment ? ', ' : '' }}
          {{ status.time?.comment }}
        </p>
      </div>

      <div class="ml-auto d-flex align-center">
        <form @submit.prevent="submitStopForm()">
          <button type="submit" class="button button-danger">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect fill="currentColor" x="3" y="3" width="18" height="18" rx="2" ry="2"></rect></svg>
            <span class="ml-1">
              Stop ({{ durationDiff(status.time?.startedAt, dayjs().format('YYYY-MM-DD HH:mm:ss') ) }})
            </span>
          </button>
        </form>
      </div>

    </div>
  </div>
</template>

<script setup>
import dayjs from 'dayjs';
import { onMounted, ref, inject } from 'vue';
import dataStatus from '../data/status';
import dataTime from '../data/time';

const emitter = inject('emitter');

const status = ref(null);

async function loadView() {
  status.value = await dataStatus.getStatus();
}

function durationDiff(startDate, endDate) {
  const start = dayjs.utc(startDate);
  const end = dayjs.utc(endDate);
  return dayjs.duration(end.diff(start)).format('HH:mm');
}

onMounted(() => {
  setInterval(() => {
    loadView();
  }, 60 * 1000);
  loadView();
});

emitter.on('time-added-edited', () => {
  loadView();
});

async function submitStopForm() {
  try {
    await dataTime.stopTimes(dayjs().format('YYYY-MM-DD HH:mm:ss'));
  } catch (error) {
    // TODO
    console.log('Error:', error);
    return;
  }

  loadView();
  emitter.emit('time-status-bar-stopped', null);
}
</script>

<style scoped>
  .time-status-bar .panel {
    background-color: var(--color-light-green);
  }

  h3, p {
    margin: 0;
  }
</style>
