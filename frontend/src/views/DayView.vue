<!-- eslint-disable max-len -->
<template>

<!-- Header -->
<div class="container mt-7">
  <div  v-if="currentDate">
    <h1>{{ currentDate.format('dddd D MMM, YYYY') }}</h1>

    <div class="d-flex">
      <div class="d-flex grouped-buttons">
        <RouterLink class="button" :to="{ name: 'DayView', params: { dayDate: currentDate.subtract(1, 'day').format('YYYY-MM-DD') } }">
          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 12H6M12 5l-7 7 7 7"/></svg>
          <span class="ml-1">Previous day</span>
        </RouterLink>
        <RouterLink class="button" :to="{ name: 'DayView', params: { dayDate: currentDate.add(1, 'day').format('YYYY-MM-DD') } }">
          <span class="mr-1">Next day</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h13M12 5l7 7-7 7"/></svg>
        </RouterLink>
      </div>
      <RouterLink class="button button-outline ml-1" v-if="currentDate.format('YYYY-MM-DD') !== dayjs.utc().format('YYYY-MM-DD')" :to="{ name: 'DayView', params: {} }">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
        <span class="ml-1">Go to today</span>
      </RouterLink>

      <button class="ml-auto button button-secondary" @click="showAddTimeForm = !showAddTimeForm">
        <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="5" x2="12" y2="19"></line><line x1="5" y1="12" x2="19" y2="12"></line></svg>
        <span class="ml-1">Add time</span>
      </button>
    </div>
  </div>
</div>

<!-- Add time -->
<div class="container mt-7 mb-5">
  <div class="add-time-form panel p-5" v-if="showAddTimeForm">
      <TimeAddForm @created="loadView(); showAddTimeForm = false;" @cancel="showAddTimeForm = false"></TimeAddForm>
  </div>
</div>

<!-- Table -->
<div class="container mt-7">
  <div v-if="currentDate">
    <div v-if="times && times.length > 0">
      <table>
        <thead>
          <tr>
            <th>Client</th>
            <th>Project</th>
            <th>Code</th>
            <th>Comment</th>
            <th>Length</th>
            <th>Time</th>
            <th class="text-right"></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="time in times" :key="time.id">
            <td>{{ time.client.title }}</td>
            <td>{{ time.project.title }}</td>
            <td>{{ time.code }}</td>
            <td>{{ time.comment }}</td>
            <td>{{ timeHelpers.durationDiff(time.startedAt, time.stoppedAt) }}</td>
            <td>{{ dayjs.utc(time.startedAt).format('HH:mm') }} - {{ time.stoppedAt !== null ? dayjs.utc(time.stoppedAt).format('HH:mm') : '' }}</td>
            <td>
              <div class="d-flex align-center justify-end">
                <a href="#" @click.prevent="restartTime(time.client.title, time.project.title, time.code, time.comment)">
                  <svg class="vertical-align-middle" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><polygon points="5 3 19 12 5 21 5 3"></polygon></svg>
                </a>

                <RouterLink :to="{ name: 'TimeEdit', params: { id: time.id }}" class="ml-2 d-block">
                  <svg class="vertical-align-middle" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M20 14.66V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h5.34"></path><polygon points="18 2 22 6 12 16 8 16 8 12 18 2"></polygon></svg>
                </RouterLink>
              </div>
            </td>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <td></td>
            <td></td>
            <td></td>
            <td></td>
            <td>{{ timeHelpers.durationDiffMultiple(times) }}</td>
            <td>
              {{ dayjs.utc(times[0].startedAt).format('HH:mm') }} -
              {{ times[times.length - 1].stoppedAt !== null ? dayjs.utc(times[times.length - 1].stoppedAt).format('HH:mm') : `(${dayjs().format('HH:mm')})` }}
            </td>
            <td></td>
          </tr>
        </tfoot>
      </table>
    </div>

    <div v-else>
      <div class="panel p-4 d-flex align-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z"></path><line x1="4" y1="22" x2="4" y2="15"></line></svg>
        <span class="ml-2">Empty day! <a href="" @click.prevent="showAddTimeForm = !showAddTimeForm">Add time?</a></span>
      </div>
    </div>

  </div>
</div>

</template>

<script setup>
import { useRoute } from 'vue-router';
import dayjs from 'dayjs';
import {
  onMounted, ref, inject, watch,
} from 'vue';
import { setLoading } from '../stores/loader';
import dataTime from '../data/time';
import timeHelpers from '../lib/time-helpers';
import TimeAddForm from '../components/TimeAddForm.vue';

const emitter = inject('emitter');

const route = useRoute();

const currentDate = ref(null);
const times = ref(null);
const showAddTimeForm = ref(false);

async function loadView() {
  setLoading(true);
  currentDate.value = dayjs(route.params?.dayDate ? route.params.dayDate : new Date());
  const startDate = `${currentDate.value.format(currentDate.value.format('YYYY-MM-DD'))} 00:00:00`;
  const endDate = `${currentDate.value.format(currentDate.value.format('YYYY-MM-DD'))} 23:59:59`;
  times.value = await dataTime.getTimesBetween(startDate, endDate);
  setLoading(false);
}

async function restartTime(client, project, code, comment) {
  const startDate = dayjs().format('YYYY-MM-DD HH:mm:ss');

  try {
    setLoading(true);
    await dataTime.saveTime(client, project, code, comment, startDate, null);
    setLoading(false);
  } catch (error) {
    console.log('Error:', error);
  }

  emitter.emit('time-added-edited', null);
  loadView();
}

watch(() => route.path, () => {
  loadView();
});

onMounted(() => {
  document.title = `Day - ${process.env.VUE_APP_TITLE_POST}`;
  loadView();
});

emitter.on('time-status-bar-stopped', () => {
  loadView();
});

</script>
