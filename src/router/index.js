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
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },

  // Vow of the disciple
  {
    path: '/vod',
    name: 'vod',
    component: () => import(/* webpackChunkName: "vod" */ '../views/vod/EncounterListView.vue')
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
  },
  {
    path: '/vod-exi',
    name: 'vod-exi',
    component: () => import(/* webpackChunkName: "vod-exi" */ '../views/vod/ExhibitionView.vue')
  },
  {
    path: '/xur',
    name: 'xur',
    component: () => import(/* webpackChunkName: "xur" */ '../views/XurView.vue')
  },
  {
    path: '/kf',
    name: 'kf',
    component: () => import(/* webpackChunkName: "kf" */ '../views/kf/EncounterListView.vue')
  },
  {
    path: '/kf-kfc',
    name: 'kf-kfc',
    component: () => import(/* webpackChunkName: "kf-kfc" */ '../views/kf/EntranceView.vue')
  },
  {
    path: '/riv',
    name: 'riv',
    component: () => import(/* webpackChunkName: "riv" */ '../views/riv/EncounterListView.vue')
  },
  {
    path: '/riv-rww',
    name: 'riv-rww',
    component: () => import(/* webpackChunkName: "riv-rww" */ '../views/riv/WishingWallView.vue')
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
