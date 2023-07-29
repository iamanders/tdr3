<!-- eslint-disable max-len -->
<template>
  <div class="container mt-7">

    <!-- Header -->
    <h1>Report</h1>

    <!-- Form -->
    <div class="panel p-5 pt-3 pb-3">
      <form class="grid" @submit.prevent="searchSubmit()">

        <!-- Row 1 -->
        <div class="col col-md-3">
          <label class="label" for="client">Client</label>
          <input type="text" id="client" v-model="form.client">
        </div>
        <div class="col col-md-3">
          <label class="label" for="project">Project</label>
          <input type="text" id="project" v-model="form.project">
        </div>
        <div class="col col-md-3">
          <label class="label" for="code">Code</label>
          <input type="text" id="code" v-model="form.code">
        </div>
        <div class="col col-md-3">
          <label class="label" for="comment">Comment</label>
          <input type="text" id="comment" v-model="form.comment">
        </div>

        <!-- Row 2 -->
        <div class="col col-md-3">
          <label class="label" for="startDate">From</label>
          <input type="date" id="startDate" v-model="form.startDate" required>
        </div>
        <div class="col col-md-3">
          <label class="label" for="startDate">To</label>
          <input type="date" id="startDate" v-model="form.stopDate" required>
        </div>

        <div class="col col-md-3 ml-auto d-flex justify-end">
          <button v-if="Array.isArray(times)" @click.prevent="resetForm" type="button" class="button button-outline mt-4 mr-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg>
          </button>

          <button type="submit" class="button mt-4 submit-button">
            <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"></circle><line x1="21" y1="21" x2="16.65" y2="16.65"></line></svg>
            <span class="ml-1">Search</span>
          </button>
        </div>
      </form>

    </div>

    <!-- Empty result -->
    <div v-if="Array.isArray(times) && times.length === 0" class="mt-6">
      <div class="panel p-4 d-flex align-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path><line x1="12" y1="9" x2="12" y2="13"></line><line x1="12" y1="17" x2="12.01" y2="17"></line></svg>
        <span class="ml-2">No results found!</span>
      </div>
    </div>

    <!-- Results -->
    <div v-if="Array.isArray(times) && times.length > 0" class="mt-6">
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
            <td class="text-right">Total:</td>
            <td>{{ timeHelpers.durationDiffGroupedMultiple(times) }}</td>
            <td></td>
            <td></td>
          </tr>
        </tfoot>
      </table>
    </div>
  </div>

</template>

<script setup>
import dayjs from 'dayjs';
import { onMounted, ref } from 'vue';
import { setLoading } from '../stores/loader';
import dataTime from '../data/time';
import timeHelpers from '../lib/time-helpers';

const form = ref({});
const times = ref(null);

async function searchSubmit() {
  setLoading(true);
  times.value = null;

  const startDate = `${form.value.startDate} 00:00:00`;
  const stopDate = `${form.value.stopDate} 23:59:59`;

  times.value = timeHelpers.groupTimesByDay(await dataTime.searchTimes(
    form.value.client,
    form.value.project,
    form.value.code,
    form.value.comment,
    startDate,
    stopDate,
  ));
  setLoading(false);
}

function resetForm() {
  times.value = null;
  form.value = {
    client: '',
    project: '',
    code: '',
    comment: '',
    startDate: dayjs().utc().subtract('7', 'day').format('YYYY-MM-DD'),
    stopDate: dayjs().utc().format('YYYY-MM-DD'),
  };
}

onMounted(() => {
  document.title = `Report - ${process.env.VUE_APP_TITLE_POST}`;
  resetForm();
});

</script>

<style lang="scss" scoped>
.submit-button {
  padding-left: 2rem;
  padding-right: 2rem;
}
</style>
