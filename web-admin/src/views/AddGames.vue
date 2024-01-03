<template>
    <div class="flex justify-center items-center">
        <div>
            <select v-model="homeID">
                <option value="" disabled selected>Select home</option>
                <option v-for="ht in teams" :key="ht.id" :value="ht.id">
                    {{ ht.team_name }}
                </option>
            </select>
            <select v-model="awayID">
                <option value="" disabled selected>Select away</option>
                <option v-for="aw in teams" :key="aw.id" :value="aw.id">
                    {{ aw.team_name }}
                </option>
            </select>
            <input v-model="date" type="date" placeholder="Enter teame name" />
            <button class="bg-blue text-white p-2" @click="addTeams">
                Add
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from '../axios'
import dayjs from 'dayjs'
import utc from 'dayjs/plugin/utc'
import advancedFormat from 'dayjs/plugin/advancedFormat'

dayjs.extend(advancedFormat)
dayjs.extend(utc)

const router = useRouter()

const teams = ref([] as any)
const date = ref('')
const homeID = ref('')
const awayID = ref('')

const getTeams = async () => {
    try {
        const response = await axios.get('/teams')
        teams.value = response.data
        console.log(teams.value)
    } catch (err: any) {
        console.error(err.message)
    }
}

const addTeams = async () => {
    await axios.post('/games', {
        home_team_id: homeID.value,
        away_team_id: awayID.value,
        game_date: dayjs.utc(date.value).format()
    })
    router.push('/')
}

onMounted(() => {
    getTeams()
})
</script>
