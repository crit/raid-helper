import { createRouter, createWebHashHistory } from 'vue-router'
import RaidListView from '../views/RaidListView.vue'

const routes = [
  {
    path: '/',
    name: 'raid-list',
    component: RaidListView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },

  // Vow of the disciple
  {
    path: '/vod',
    name: 'vod',
    component: () => import(/* webpackChunkName: "vod" */ '../views/vod/EncounterListView.vue'),
  },
  {
    path: '/vod-ent',
    name: 'vod-ent',
    component: () => import(/* webpackChunkName: "vod-ent" */ '../views/vod/EntranceView.vue')
  },
  {
    path: '/vod-acq',
    name: 'vod-acq',
    component: () => import(/* webpackChunkName: "vod-acq" */ '../views/vod/AcquisitionView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
