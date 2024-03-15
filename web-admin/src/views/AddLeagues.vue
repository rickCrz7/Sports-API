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
            <input
                v-model="leagueName"
                type="text"
                placeholder="Enter league name name"
            />
            <button class="secondary-action-buttons" @click="addLeague">
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
const leagueName = ref('')
const sports_id = ref('')

const getSports = async () => {
    try {
        const response = await axios.get('/sports')
        sports.value = response.data
        console.log(sports.value)
    } catch (err: any) {
        console.error(err.message)
    }
}

const addLeague = async () => {
    console.log(sports_id.value)
    await axios.post('/leagues', {
        league_name: leagueName.value,
        sport_id: sports_id.value
    })
    router.push('/')
}

onMounted(() => {
    getSports()
})
</script>
