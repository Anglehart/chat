<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue';
import { format } from 'date-fns'

interface IncomeMessage {
  CreatedDate: string,
  GUID: string,
  Text: string
}

interface IMessage {
  guid: string,
  text: string,
  createdDate: string,
}

const socket = ref<WebSocket | null>(null);
const messages = ref<IMessage[]>([]);
const status = ref('Отключено')
const isConnected = ref(false)

const WS_URL = 'wss://echo.websocket.org'

const message = ref<string>("")
const result = ref<{result: string}>({result: ""})


const getMessages = async () => {
  const resp = await fetch("/get-messages", {
    headers: {
      'Content-Type': 'application/json', 
    }
  })
  const response: IncomeMessage[] = await resp.json();

  messages.value = response.map(el => {
    return {
      guid: el.GUID,
      text: el.Text,
      createdDate: format(new Date(el.CreatedDate), 'HH:mm dd.MM.yy'),
    }
  });
}

const connect = () => {
  socket.value = new WebSocket(WS_URL);

  socket.value.onopen = () => {
    isConnected.value = true;
    status.value = 'Подключено';
    console.log('WebSocket connected');
  };

  socket.value.onmessage = (event) => {
    messages.value.push(`Ответ сервера: ${event.data}`);
  };

  socket.value.onerror = (error) => {
    console.error('WebSocket error:', error);
    status.value = 'Ошибка';
  };

  socket.value.onclose = () => {
    isConnected.value = false;
    status.value = 'Отключено';
    console.log('WebSocket disconnected');
    setTimeout(() => connect(), 3000);
  };
};

const filteredMessages = computed(() => {
  return messages.value.filter(el => !!el.text)
})

onMounted(async () => {
  await getMessages()
  connect()
});

onUnmounted(() => {
  if (socket.value) {
    socket.value.close();
  }
});

const onSendMessage = () => {
  if (isConnected.value && socket.value) {
    socket.value.send(message.value);
    message.value = ""
  }
}
</script>

<template>
  <div class="messager">
    <p class="messager__status">Статус: {{ status }}</p>
    <ul class="message__list">
      <li v-for="(msg, index) in filteredMessages" :key="index">
        <div class="message__baloon">
          <span class="message__baloon-text">{{ msg.text }}</span> 
          <span class="message__baloon-time">{{ msg.createdDate }}</span>
        </div>
      </li>
    </ul>
    <textarea
      class="message__textarea" 
      v-model="message"
      placeholder="Введите сообщение..."
    />
    <button 
      @click="onSendMessage"
      :disabled="!isConnected"
    >
      Отправить сообщение
    </button>
  </div>
</template>
<style>
  .messager {
    max-width: 600px;
    margin: 20px auto;
    padding: 0 15px;
  }

  .messager__status {
    margin-bottom: 10px;
    text-align: right;
  }

  .message__wrap {
    display: flex;
    flex-direction: column;
  }

  .message__textarea {
    width: 100%;
    height: 100px;
    margin: 20px 0;
    resize: none;
  }

  .message__list {
    list-style: none;
    padding-left: 0;
  }

  .message__baloon {
    background-color: aliceblue;
    width: fit-content;
    padding: 5px 10px;
    border-radius: 6px;
    margin-bottom: 10px;
  }

  .message__baloon-time {
    font-size: 12px;
    margin-left: 10px;
  }

  .message__baloon-text {
    font-size: 16px;
  }

</style>
