<template>
    <div class="flex justify-center items-center">
        <div class="input-group flex gap-1">
            <input v-model="element" type="text" />
            <button class="secondary-action-buttons" @click="updateElement">
                Update
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from '../axios'

const router = useRouter()

const id = router.currentRoute.value.params.id

const from = router.currentRoute.value.query.from

const element = ref()

const sportID = ref()

const leagueID = ref()

const getElement = async () => {
    try {
        console.log(from)
        switch (from) {
            case 'Sports':
                var response = await axios.get(`/sports/${id}`)
                element.value = response.data.sport_name
                break
            case 'Leagues':
                response = await axios.get(`/leagues/${id}`)
                element.value = response.data.league_name
                sportID.value = response.data.sport_id
                break
            case 'Teams':
                response = await axios.get(`/teams/${id}`)
                element.value = response.data.team_name
                sportID.value = response.data.sport_id
                leagueID.value = response.data.league_id
                break
            default:
                console.log('Invalid route')
                break
        }
    } catch (err) {
        console.log(err)
    }
}

const updateElement = async () => {
    try {
        switch (from) {
            case 'Sports':
                await axios.put(`/sports/${id}`, { sport_name: element.value })
                break
            case 'Leagues':
                await axios.put(`/leagues/${id}`, {
                    league_name: element.value,
                    sport_id: sportID.value
                })
                break
            case 'Teams':
                await axios.put(`/teams/${id}`, {
                    team_name: element.value,
                    league_id: leagueID.value,
                    sport_id: sportID.value
                })
                break
            default:
                console.log('Invalid route')
                break
        }
        router.push('/')
    } catch (err) {
        console.log(err)
    }
}

onMounted(() => {
    getElement()
})
</script>
