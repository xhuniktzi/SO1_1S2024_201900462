<template>
  <div id="app">
    <h1>Logs de MongoDB</h1>
    <button @click="fetchLogs">Actualizar Logs</button>
    <table>
      <thead>
        <tr>
          <th>Timestamp</th>
          <th>Mensaje</th>
          <th>Error</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="log in logs" :key="log._id">
          <td>{{ log.timestamp }}</td>
          <td>{{ log.message }}</td>
          <td>{{ log.error }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Log {
  _id: string
  timestamp: string
  message: string
  error: string
}

export default defineComponent({
  name: 'App',
  data() {
    return {
      logs: [] as Log[]
    }
  },
  methods: {
    async fetchLogs(): Promise<void> {
      try {
        const response = await fetch('http://localhost:3000/logs')
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        this.logs = (await response.json()) as Log[]
      } catch (error) {
        console.error('There was an error fetching the logs:', error)
      }
    }
  },
  mounted() {
    this.fetchLogs()
  }
})
</script>

<style>
table {
  width: 100%;
  border-collapse: collapse;
}
th,
td {
  border: 1px solid #ccc;
  padding: 8px;
  text-align: left;
}
th {
  color: black;
  background-color: #eee;
}
</style>
