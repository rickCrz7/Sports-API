<template>
    <div class="flex justify-center items-center">
        <div class="input-group flex gap-1">
            <select v-model="sports_id">
                <option value="" disabled selected>Select sport</option>
                <option
                    v-for="sport in sports"
                    :key="sport.id"
                    :value="sport.id"
                >
                    {{ sport.sport_name }}
                </option>
            </select>
            <select
                v-model="league_id"
                :disabled="!sports_id"
                @click="filterLeague"
            >
                <option value="" disabled selected>Select league</option>
                <option
                    v-for="league in filteredLeagues"
                    :key="league.id"
                    :value="league.id"
                >
                    {{ league.league_name }}
                </option>
            </select>
            <input
                v-model="teamName"
                type="text"
                placeholder="Enter teame name"
            />
            <button class="secondary-action-buttons" @click="addTeams">
                <i class="i-mdi:plus"></i>
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from '../axios'

const router = useRouter()

const sports = ref([] as any)
const leagues = ref([] as any)
const filteredLeagues = ref([] as any)
const teamName = ref('')
const sports_id = ref('')
const league_id = ref('')

const getSports = async () => {
    try {
        const response = await axios.get('/sports')
        sports.value = response.data
        console.log(sports.value)
    } catch (err: any) {
        console.error(err.message)
    }
}

const getLeagues = async () => {
    try {
        const response = await axios.get(`/leagues`)
        leagues.value = response.data
        console.log(leagues.value)
    } catch (err: any) {
        console.error(err.message)
    }
}

const filterLeague = () => {
    console.log(sports_id.value)
    console.log(leagues.value)
    filteredLeagues.value = leagues.value.filter((league: any) => {
        return league.sport_id === sports_id.value
    })
}

const addTeams = async () => {
    console.log(sports_id.value)
    await axios.post('/teams', {
        team_name: teamName.value,
        sport_id: sports_id.value,
        league_id: league_id.value
    })
    router.push('/')
}

onMounted(() => {
    getSports()
    getLeagues()
})
</script>
