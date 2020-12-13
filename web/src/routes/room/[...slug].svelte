<script context="module">
  export async function preload({ params }) {
    let [roomId] = params.slug;

    return { roomId };
  }
</script>

<script>
  import { onMount } from 'svelte';
  import Cookies from 'js-cookie';

  import { querySerializer } from '@/util/query.js';

  import { username } from '@/store/index.js';

  import api from '@/api/index.js';
  import Metadata from '@/components/Metadata.svelte';

  export let roomId;
  let roomData, sourceText, destinedText, participantId;
  let currentIndex = 0;
  let currentText = '';
  let hasError = false;

  onMount(() => {
    getRoomByCode();

    const params = {
      code: roomId,
      participant_id: participantId,
      participant_name: $username
    };

    const ws = new WebSocket(`${ENV.WEB_SOCKET_BASE_URL}/api/join-room${querySerializer(params)}`);

    ws.onopen = function (event) {
      console.log('masuk');
      ws.send(
        JSON.stringify(
          Object.assign(params, {
            action: 'join',
            current_word_count: 10
          })
        )
      );
    };

    ws.onmessage = function (event) {
      console.log(event);
    };
  });

  participantId = Cookies.get('participantId');
  if (!participantId) {
    Cookies.set('participantId', `participant_${new Date().getTime()}`);
    participantId = Cookies.get('participantId');
  }

  function getRoomByCode() {
    const params = {
      code: roomId
    };
    api.getRoomByCode(getRoomByCodeSuccess, params, getRoomByCodeFail);
  }

  function getRoomByCodeSuccess(response) {
    roomData = response.body;
    sourceText = roomData.source_text.split(' ');
    destinedText = roomData.destined_text.split(' ');
  }

  function getRoomByCodeFail() {
    console.log('get room by code failed');
  }

  function handleInputChange(event) {
    if (event.data === ' ') {
      console.log('change', event, destinedText[currentIndex], event.target.value);
      if (destinedText[currentIndex] === event.target.value.replace(/\s/g, '')) {
        hasError = false;
        currentIndex += 1;
        currentText = '';
      } else {
        hasError = true;
      }
    }
  }
</script>

<style>
  .room {
    @apply mx-auto;
    @apply my-auto;
    @apply h-screen;
    @apply bg-blue-300;
  }
  .room-wrapper {
    @apply mx-auto;
    @apply my-auto;
    @apply text-center;
    max-width: 600px;
  }
  .text-wrapper {
    @apply bg-blue-200;
    @apply p-4;
    @apply my-2;
  }

  .text {
    @apply text-xl;
  }

  .input-field {
    @apply font-bold;
  }

  .success {
    @apply text-green-500;
    @apply font-bold;
  }

  .error {
    @apply text-red-500;
    @apply font-bold;
  }
</style>

<Metadata />

<div class="room">
  <div class="room-wrapper">
    <div>Hallo, {$username} - {participantId}</div>
    {#if roomData}
      <div>{roomData.source_lang} - {roomData.destined_lang}</div>
      <div class="text-wrapper flex flex-wrap justify-start">
        {#each sourceText as s, i}
          <div class="text" class:success={i < currentIndex} class:error={hasError && i === currentIndex}>
            {s}
            &nbsp; &nbsp;
          </div>
        {/each}
      </div>

      <div class="input-field inline-flex items-baseline border-none shadow-md bg-white w-full" for="username">
        <input
          id="inputText"
          type="text"
          class="placeholder-blue w-full p-2 no-outline text-dusty-blue-darker"
          name="inputText"
          placeholder="Type here"
          bind:value={currentText}
          on:input={handleInputChange} />
      </div>
    {/if}
  </div>
</div>
