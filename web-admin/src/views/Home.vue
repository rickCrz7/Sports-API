<template>
    <main>
        <div class="flex justify-center pt-8 h-2">
            <span class="title">Dashboard</span>
        </div>
        <div class="flex justify-end items-baseline pr-24">
            <button class="button-color text-3xl" @click="darkMode">
                <i class="i-mdi:theme-light-dark"></i>
            </button>
        </div>
        <div class="admin-page h-3/4 mt-24">
            <div class="dashboard">
                <h2>Sports</h2>
                <div class="h-11/12">
                    <ul v-for="sport in sports" :key="sport.id">
                        <li class="select">
                            <span
                                class="text-red hover:text-black cursor-pointer"
                                @click="editName(sport.id, 'Sports')"
                                >{{ sport.sport_name }}</span
                            >
                            <button
                                class="button-color"
                                @click="deleteSport(sport.id)"
                            >
                                <i class="i-mdi:trashcan"></i>
                            </button>
                        </li>
                    </ul>
                </div>
                <div class="flex justify-center">
                    <button class="action-buttons" @click="addSport">
                        Add
                    </button>
                </div>
            </div>
            <div class="dashboard">
                <h2>Leagues</h2>
                <div class="h-11/12">
                    <ul v-for="league in leagues" :key="league.id">
                        <li class="select">
                            <span
                                class="text-red hover:text-black cursor-pointer"
                                @click="editName(league.id, 'Leagues')"
                                >{{ league.league_name }}</span
                            >
                            <button
                                class="button-color"
                                @click="deleteLeague(league.id)"
                            >
                                <i class="i-mdi:trashcan"></i>
                            </button>
                        </li>
                    </ul>
                </div>
                <div class="flex justify-center">
                    <button class="action-buttons" @click="addLeague">
                        Add
                    </button>
                </div>
            </div>
            <div class="dashboard">
                <h2>Teams</h2>
                <div class="overflow-y-auto h-11/12">
                    <ul v-for="team in teams" :key="team.id">
                        <li class="select">
                            <span
                                class="text-red hover:text-black cursor-pointer"
                                @click="editName(team.id, 'Teams')"
                                >{{ team.team_name }}</span
                            >
                            <button
                                class="button-color"
                                @click="deleteTeam(team.id)"
                            >
                                <i class="i-mdi:trashcan"></i>
                            </button>
                        </li>
                    </ul>
                </div>
                <div class="flex justify-center">
                    <button class="action-buttons" @click="addTeams">
                        Add
                    </button>
                </div>
            </div>
            <div class="dashboard">
                <h2>Games</h2>
                <div class="overflow-y-auto h-11/12">
                    <ul v-for="g in games" :key="g.id">
                        <li class="select" @click="updateScore(g.id)">
                            <span
                                class="text-red hover:text-black cursor-pointer"
                                >{{
                                    `${g.home_team} ${g.home_score}-${g.away_score} ${g.away_team}`
                                }}</span
                            >
                            <br />
                            <span class="flex justify-between">
                                <span class="text-gray-400">{{
                                    g.game_date
                                }}</span>
                                <button
                                    class="button-color ml-2"
                                    @click="deleteGame(g.id)"
                                >
                                    <i class="i-mdi:trashcan"></i>
                                </button>
                            </span>
                        </li>
                    </ul>
                </div>
                <div class="flex justify-center">
                    <button class="action-buttons" @click="addGames">
                        Add
                    </button>
                </div>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from '../axios'
import { Sports, League, Teams } from '../models/structs'
import { useRouter } from 'vue-router'

const router = useRouter()

const leagues = ref([] as League[])
const sports = ref([] as Sports[])
const teams = ref([] as Teams[])
const games = ref([] as any[])

const getSports = async () => {
    try {
        const response = await axios.get('/sports')
        sports.value = response.data
        console.log(sports.value)
    } catch (err: any) {
        console.error(err.message)
    }
}

const addSport = async () => {
    router.push('/add/sports')
}

const deleteSport = async (id: number) => {
    try {
        const confirmed = confirm('Are you sure you want to delete this sport?')
        if (confirmed) {
            await axios.delete(`/sports/${id}`)
            getSports()
        }
    } catch (err: any) {
        console.error(err.message)
    }
}

const getLeagues = async () => {
    try {
        const response = await axios.get(`/leagues`)
        leagues.value = response.data
        console.log(leagues.value)
        // if (leagues.value.length === 0) {
        //     leagues.value = [
        //         {
        //             id: 0,
        //             sport_id: 0,
        //             league_name: 'No leagues found'
        //         }
        //     ]
        // }
    } catch (err: any) {
        console.error(err.message)
    }
}

const addLeague = async () => {
    router.push('/add/leagues')
}

const deleteLeague = async (id: number) => {
    try {
        const confirmed = confirm(
            'Are you sure you want to delete this league?'
        )
        if (confirmed) {
            await axios.delete(`/leagues/${id}`)
            getLeagues()
        }
    } catch (err: any) {
        console.error(err.message)
    }
}

const getTeams = async () => {
    try {
        const response = await axios.get('/teams')
        teams.value = response.data
    } catch (err: any) {
        console.error(err.message)
    }
}

const addTeams = async () => {
    router.push('/add/teams')
}

const deleteTeam = async (id: number) => {
    try {
        const confirmed = confirm('Are you sure you want to delete this team?')
        if (confirmed) {
            await axios.delete(`/teams/${id}`)
            getTeams()
        }
    } catch (err: any) {
        console.error(err.message)
    }
}

const getGames = async () => {
    try {
        const response = await axios.get('/games')
        console.log(response.data)
        games.value = response.data
    } catch (err: any) {
        console.error(err.message)
    }
}

const addGames = async () => {
    router.push('/add/games')
}

const updateScore = async (id: string) => {
    router.push('/update/score/' + id)
}

const deleteGame = async (id: number) => {
    try {
        const confirmed = confirm('Are you sure you want to delete this game?')
        if (confirmed) {
            await axios.delete(`/games/${id}`)
            getGames()
        }
    } catch (err: any) {
        console.error(err.message)
    }
}

const darkMode = () => {
    var element = document.body
    element.classList.toggle('dark-mode')
}

const editName = (id: number, where: string) => {
    router.push({
        name: 'Edit',
        params: { id },
        query: {
            from: where
        }
    })
}

onMounted(() => {
    getSports()
    getLeagues()
    getTeams()
    getGames()
})
</script>
