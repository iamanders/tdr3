<!-- eslint-disable max-len -->
<template>

<!-- Header -->
<div class="container mt-7">
  <div v-if="currentDate">
    <h1 class="m-0">Week {{ currentDate.isoWeek() }}, {{  currentDate.year() }}</h1>
    <h3 class="mb-4">{{ currentDate.format('D MMM') }} - {{ nextDate.subtract(1, 'day').format('D MMM') }}</h3>

    <div class="d-flex">
      <div class="d-flex grouped-buttons">
        <RouterLink class="button" :to="{ name: 'WeekView', params: { year: previousDate.year(), week: previousDate.isoWeek() } }">
          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 12H6M12 5l-7 7 7 7"/></svg>
          <span class="ml-1">Previous week</span>
        </RouterLink>
        <RouterLink class="button" :to="{ name: 'WeekView', params: { year: nextDate.year(), week: nextDate.isoWeek() } }">
          <span class="mr-1">Next week</span>
          <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h13M12 5l7 7-7 7"/></svg>
        </RouterLink>
      </div>

      <RouterLink class="button button-outline ml-1" v-if="showTodayButton()" :to="{ name: 'WeekView', params: {} }">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
        <span class="ml-1">Go to today</span>
      </RouterLink>

    </div>
  </div>
</div>

<!-- Table -->
<div class="container mt-7 mb-6">
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

          <template v-for="day in times" :key="day.formattedDate">
            <tr class="sub-header"><td colspan="7">{{ day.formattedDate }}</td></tr>

            <tr v-for="time in day.times" :key="time.id">
              <td>{{ time.client.title }}</td>
              <td>{{ time.project.title }}</td>
              <td>{{ time.code }}</td>
              <td>{{ time.comment }}</td>
              <td>{{ timeHelpers.durationDiff(time.startedAt, time.stoppedAt) }}</td>
              <td>{{ dayjs.utc(time.startedAt).format('HH:mm') }} - {{ time.stoppedAt !== null ? dayjs.utc(time.stoppedAt).format('HH:mm') : '' }}</td>
              <td>
                <div class="d-flex align-center justify-end">
                  <RouterLink :to="{ name: 'TimeEdit', params: { id: time.id }}" class="ml-2 d-block">
                    <svg class="vertical-align-middle" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><path d="M20 14.66V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h5.34"></path><polygon points="18 2 22 6 12 16 8 16 8 12 18 2"></polygon></svg>
                  </RouterLink>
                </div>
              </td>
            </tr>

            <tr class="sub-footer">
              <td colspan="4"></td>
              <td>{{ timeHelpers.durationDiffMultiple(day.times) }}</td>
              <td colspan="2"></td>
            </tr>
          </template>

        </tbody>
        <tfoot>
          <tr>
            <td></td>
            <td></td>
            <td></td>
            <td class="text-right">Total in week:</td>
            <td>{{ timeHelpers.durationDiffGroupedMultiple(times) }}</td>
            <td></td>
            <td></td>
          </tr>
        </tfoot>
      </table>
    </div>

    <div v-else>
      <div class="panel p-4 d-flex align-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z"></path><line x1="4" y1="22" x2="4" y2="15"></line></svg>
        <span class="ml-2">Empty week!</span>
      </div>
    </div>

  </div>
</div>

</template>

<script setup>
import dayjs from 'dayjs';
import { onMounted, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import { setLoading } from '../stores/loader';
import dataTime from '../data/time';
import timeHelpers from '../lib/time-helpers';

const route = useRoute();

const currentDate = ref(null);
const nextDate = ref(null);
const previousDate = ref(null);
const times = ref(null);

async function loadView() {
  setLoading(true);

  if (route.params.year && route.params.week) {
    currentDate.value = dayjs().year(route.params.year).isoWeek(route.params.week).isoWeekday(1);
  } else {
    currentDate.value = dayjs().isoWeekday(1);
  }
  currentDate.value = currentDate.value.utc().hour(0).minute(0).second(0);

  nextDate.value = currentDate.value.utc().add(7, 'day');
  previousDate.value = currentDate.value.utc().subtract(7, 'day');

  const startDate = `${currentDate.value.format(currentDate.value.format('YYYY-MM-DD'))} 00:00:00`;
  const endDate = `${currentDate.value.format(nextDate.value.subtract('1', 'second').format('YYYY-MM-DD'))} 23:59:59`;
  times.value = timeHelpers.groupTimesByDay(await dataTime.getTimesBetween(startDate, endDate));

  setLoading(false);
}

function showTodayButton() {
  if (currentDate.value.year() !== dayjs().year()) {
    return true;
  }

  // eslint-disable-next-line max-len
  if (currentDate.value.isoWeek() === dayjs().isoWeek() && currentDate.value.year() === dayjs().year()) {
    return false;
  }

  return true;
}

watch(() => route.path, () => {
  loadView();
});

onMounted(() => {
  document.title = `Week - ${process.env.VUE_APP_TITLE_POST}`;
  loadView();
});

</script>
