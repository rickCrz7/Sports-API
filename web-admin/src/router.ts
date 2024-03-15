import { createRouter, createWebHistory } from 'vue-router'
import Home from './views/Home.vue'

const routes = [
    { path: '/', name: 'Home', component: Home },
    {
        path: '/add/sports',
        name: 'AddSports',
        component: () => import('./views/AddSports.vue')
    },
    {
        path: '/add/leagues',
        name: 'AddLeagues',
        component: () => import('./views/AddLeagues.vue')
    },
    {
        path: '/add/teams',
        name: 'AddTeams',
        component: () => import('./views/AddTeams.vue')
    },
    {
        path: '/add/games',
        name: 'AddGames',
        component: () => import('./views/AddGames.vue')
    },
    {
        path: '/update/score/:id',
        name: 'UpdateScore',
        component: () => import('./views/UpdateScore.vue')
    },
    {
        path: '/edit/:id',
        name: 'Edit',
        component: () => import('./views/EditName.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})
export default router
