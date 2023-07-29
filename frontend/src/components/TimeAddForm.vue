<template>
  <h1>Add time</h1>

  <ValidationErrors :errors="validationErrors"></ValidationErrors>

  <form v-on:submit.prevent="submitForm()">

    <div class="grid">
      <!-- Row 1 -->
      <div class="col col-md-6 col-lg-3">
        <label class="label" for="client">Client</label>
        <input type="text" id="client" v-model="form.client" ref="focusInput">
      </div>
      <div class="col col-md-6 col-lg-3">
        <label class="label" for="project">Project</label>
        <input type="text" id="project" v-model="form.project">
      </div>
      <div class="col col-md-6 col-lg-3">
        <label class="label" for="code">Code</label>
        <input type="text" id="code" v-model="form.code">
      </div>
      <div class="col col-md-6 col-lg-3">
          <label class="label" for="comment">Comment</label>
        <input type="text" id="comment" v-model="form.comment">
      </div>

      <!-- Row 2 -->
      <div class="col col-md-6 col-lg-3">
        <label class="label" for="startDate">Date</label>
        <input type="date" id="startDate" v-model="form.startDate">
      </div>
      <div class="col col-md-6 col-lg-3">
        <label class="label" for="startTime">Start/stop time</label>
        <div class="mt-1 d-flex align-center">
          <input type="time" id="startTime" v-model="form.startTime">
          <span class="d-block ml-2 mr-2">-</span>
          <input type="time" id="stopTime" v-model="form.stopTime">
        </div>
      </div>
    </div>

    <!-- Action -->
    <div class="mt-3 d-flex">
      <button type="submit" class="button mr-2">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"></polyline></svg>
        <span class="ml-2 pr-1">Add time</span>
      </button>
      <button type="button" class="button button-outline" @click.prevent="emit('cancel');">
        Cancel
      </button>
    </div>

  </form>

</template>

<script setup>
import dayjs from 'dayjs';
import { toast } from 'vue3-toastify';
import {
  onMounted, ref, defineEmits, inject,
} from 'vue';
import { setLoading } from '../stores/loader';
import dataTime from '../data/time';
import ValidationErrors from './ValidationErrors.vue';

const emitter = inject('emitter');

const emit = defineEmits(['created', 'cancel']);

const form = ref({
  client: '',
  project: '',
  code: '',
  comment: '',
  startDate: dayjs().format('YYYY-MM-DD'),
  startTime: dayjs().format('HH:mm'),
  stopTime: '',
});
const validationErrors = ref(null);
const focusInput = ref(null);

async function submitForm() {
  setLoading(true);

  const start = dayjs.utc(`${form.value.startDate} ${form.value.startTime}:00`)
    .format('YYYY-MM-DD HH:mm:ss');
  const stop = form.value.stopTime ? dayjs.utc(`${form.value.startDate} ${form.value.stopTime}:00`)
    .format('YYYY-MM-DD HH:mm:ss') : null;

  try {
    await dataTime.saveTime(
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
  emit('created');
  setLoading(false);

  toast('Time added!', { type: 'success' });
}

onMounted(() => {
  focusInput.value.focus();
});
</script>
