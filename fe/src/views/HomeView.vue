<script setup lang="ts">
import { ref } from 'vue';


const query = ref<string>("")
const result = ref<{result: string}>({result: ""})
const onSendMessage = async () => {
  const req = {
    message: query.value,
  };

  const resp = await fetch("http://localhost:8080/save-message", {
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
  <div class="message__wrap">
    <textarea
      class="message__textarea" 
      v-model="query"
      @keydown.enter="onSendMessage"
    />
    <button @click="onSendMessage">Отправить</button>
    {{ result.result }}
  </div>

</template>
<style>
  .message__wrap {
    display: flex;
    flex-direction: column;
  }

  .message__textarea {
    width: 500px;
    height: 300px;
    margin-bottom: 12px;
  }

</style>
