import dayjs from 'dayjs';
import duration from 'dayjs/plugin/duration';
import utc from 'dayjs/plugin/utc';
import isoWeek from 'dayjs/plugin/isoWeek';
import mitt from 'mitt';
import Vue3Toasity from 'vue3-toastify';
import 'vue3-toastify/dist/index.css';

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

// Mitt
const emitter = mitt();

// Day.js
dayjs.extend(utc);
dayjs.extend(duration);
dayjs.extend(isoWeek);

// Vue3 toastify
const toastifyOptions = {
  autoClose: 3000,
  hideProgressBar: true,
  theme: 'colored',
};

// Vue
createApp(App)
  .use(router)
  .use(Vue3Toasity, toastifyOptions)
  .provide('emitter', emitter)
  .mount('#app');
