<template>
    <div class="flex justify-center items-center">
        <div>
            <input
                v-model="sportName"
                type="text"
                placeholder="Enter sport name"
            />
            <button class="bg-blue text-white p-2" @click="addSport">
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

const sportName = ref('')

const addSport = async () => {
    await axios.post('/sports', { sport_name: sportName.value })
    router.push('/')
}
</script>
