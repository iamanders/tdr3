<!-- eslint-disable max-len -->
<template>
  <div class="container mt-5 grid">
    <div class="col col-md-6 col-lg-4">

      <h1>Edit time</h1>

      <div v-if="dataLoaded">
        <ValidationErrors :errors="validationErrors"></ValidationErrors>

        <form @submit.prevent="submitForm()">

          <div>
            <label class="label" for="client">Client</label>
            <input type="text" id="client" v-model="form.client">
          </div>

          <div class="mt-2">
            <label class="label" for="project">Project</label>
            <input type="text" id="project" v-model="form.project">
          </div>

          <div class="mt-2">
            <label class="label" for="code">Code</label>
            <input type="text" id="code" v-model="form.code">
          </div>

          <div class="mt-2">
            <label class="label" for="comment">Comment</label>
            <input type="text" id="comment" v-model="form.comment">
          </div>

          <div class="mt-2">
            <label class="label" for="startDate">Date</label>
            <input type="date" id="startDate" v-model="form.startDate">
          </div>

          <div class="mt-2">
            <label class="label" for="startTime">Start/stop time</label>
            <div class="d-flex align-center">
              <input type="time" id="startTime" v-model="form.startTime">
              <span class="d-block ml-2 mr-2">-</span>
              <input type="time" id="stopTime" v-model="form.stopTime">
            </div>
          </div>

          <div class="mt-6 d-flex">
            <input type="submit" class="button" value="Save changes">
            <RouterLink
              :to="{ name: 'DayView', params: { dayDate: form.startDate } }"
              class="button button-outline ml-2">
              Back
            </RouterLink>

            <button type="button" @click.prevent="deleteTime()" class="button button-danger ml-auto">
              Delete
            </button>
          </div>

        </form>

      </div>
    </div>
  </div>
</template>

<script setup>
import dayjs from 'dayjs';
import { toast } from 'vue3-toastify';
import { onMounted, ref, inject } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import ValidationErrors from '../components/ValidationErrors.vue';
import dataTime from '../data/time';

const route = useRoute();
const router = useRouter();

const emitter = inject('emitter');

const dataLoaded = ref(false);
const validationErrors = ref(null);
const time = ref(null);
const form = ref({
  client: '',
  project: '',
  code: '',
  comment: '',
  startDate: '',
  startTime: '',
  stopTime: '',
});

async function loadData() {
  try {
    time.value = await dataTime.getTime(route.params?.id);
    form.value.client = time.value.client?.title;
    form.value.project = time.value.project?.title;
    form.value.code = time.value.code;
    form.value.comment = time.value.comment;
    form.value.startDate = dayjs.utc(time.value.startedAt).format('YYYY-MM-DD');
    form.value.startTime = dayjs.utc(time.value.startedAt).format('HH:mm');
    form.value.stopTime = time.value.stoppedAt !== null ? dayjs.utc(time.value.stoppedAt).format('HH:mm') : '';
  } catch (error) {
    // if (error.message === '404Exception') {
    //   // TODO
    // } else {
    //   // TODO
    // }
  }
  dataLoaded.value = true;
}

async function submitForm() {
  const start = dayjs.utc(`${form.value.startDate} ${form.value.startTime}:00`)
    .format('YYYY-MM-DD HH:mm:ss');
  const stop = form.value.stopTime ? dayjs.utc(`${form.value.startDate} ${form.value.stopTime}:00`)
    .format('YYYY-MM-DD HH:mm:ss') : null;

  try {
    await dataTime.updateTime(
      time.value.id,
      form.value.client,
      form.value.project,
      form.value.code !== '' ? form.value.code : null,
      form.value.comment !== '' ? form.value.comment : null,
      start,
      stop,
    );
  } catch (error) {
    if (error.message === 'ValidationErrorException') {
      validationErrors.value = error.validationErrors;
    } else {
      console.log('Error:', error);
    }
    return;
  }

  emitter.emit('time-added-edited', null);

  router.push({ name: 'DayView', params: { dayDate: form.value.startDate } }).then(() => {
    toast('Changes saved!', { type: 'success' });
  });
}

async function deleteTime() {
  // eslint-disable-next-line no-restricted-globals, no-alert
  if (!confirm('Are you sure?')) {
    return;
  }

  try {
    await dataTime.deleteTime(time.value.id);
  } catch (error) {
    if (error.message === '404Exception') {
      // TODO
    } else {
      console.log('Error:', error.message);
    }
    return;
  }

  emitter.emit('time-added-edited', null);

  router.push({ name: 'DayView', params: { dayDate: form.value.startDate } }).then(() => {
    toast('Time deleted!', { type: 'success' });
  });
}

onMounted(() => {
  document.title = `Edit time - ${process.env.VUE_APP_TITLE_POST}`;
  loadData();
});

</script>
