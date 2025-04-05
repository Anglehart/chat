<script setup lang="ts">
import { ref } from 'vue';


const query = ref<string>("")
const result = ref<{result: string}>({result: ""})
const onDouble = async () => {
  const req = {
    number: Number(query.value),
  };

  const resp = await fetch("http://localhost:8080/double", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json', 
    },
    body: JSON.stringify(req)
  })

  result.value = await resp.json();
}

const onHalfed = async () => {
  const req = {
    number: Number(query.value),
  };

  const resp = await fetch("http://localhost:8080/half", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json', 
    },
    body: JSON.stringify(req)
  })

  result.value = await resp.json();
}
</script>

<template>
  <div>
    <input 
      v-model="query"
      @keydown.enter="onDouble"
    />
    <button @click="onDouble">Удвоить</button>
    <button @click="onHalfed">Разделить на два</button>
    {{ result.result }}
  </div>

</template>
