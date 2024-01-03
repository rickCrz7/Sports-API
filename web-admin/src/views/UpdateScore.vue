<template>
    <div class="flex justify-center items-center">
        <div>
            <input
                v-model="homeScore"
                type="number"
                placeholder="Enter home score"
            />
            <input
                v-model="awayScore"
                type="number"
                placeholder="Enter away score"
            />
            <button class="bg-blue text-white p-2" @click="addScore">
                Add
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from '../axios'

const router = useRouter()

const id = router.currentRoute.value.params.id

const homeScore = ref()
const awayScore = ref()

const addScore = async () => {
    try {
        await axios.post('/scores', {
            game_id: id,
            home_score: homeScore.value,
            away_score: awayScore.value
        })
        router.push('/')
    } catch (err) {
        console.log(err)
    }
}
</script>
