<template>
    <div class="flex justify-center items-center">
        <div class="input-group flex gap-1">
            <input
                v-model="sportName"
                type="text"
                placeholder="Enter sport name"
            />
            <button class="secondary-action-buttons" @click="addSport">
                <i class="i-mdi:plus"></i>
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
