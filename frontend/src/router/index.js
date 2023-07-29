import { createRouter, createWebHashHistory } from 'vue-router';

const routes = [
  // Home
  { path: '/', redirect: '/day' },

  // Day/week/month/report
  { path: '/day/:dayDate?', name: 'DayView', component: () => import(/* webpackChunkName: "day-view" */ '../views/DayView.vue') },
  { path: '/week/:year?/:week?', name: 'WeekView', component: () => import(/* webpackChunkName: "week-view" */ '../views/WeekView.vue') },
  { path: '/month/:year?/:month?', name: 'MonthView', component: () => import(/* webpackChunkName: "month-view" */ '../views/MonthView.vue') },
  { path: '/report', name: 'ReportView', component: () => import(/* webpackChunkName: "report-view" */ '../views/ReportView.vue') },

  // Time
  { path: '/time/:id', name: 'TimeEdit', component: () => import(/* webpackChunkName: "time-edit" */ '../views/TimeEdit.vue') },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
