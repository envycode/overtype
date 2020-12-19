<script context="module">
  export async function preload({ params }) {
    let [roomId] = params.slug;

    return { roomId };
  }
</script>

<script>
  import { onMount, onDestroy } from 'svelte';
  import Cookies from 'js-cookie';

  import { username, code, participantId, action, wordCount, currentWordCount, wsStore } from '@/store/index.js';
  import { initWS, closeWS } from '@/util/websocket.js';

  import api from '@/api/index.js';
  import Metadata from '@/components/Metadata.svelte';
  import Button from '@/components/Button.svelte';
  import DataVisualization from '@/components/DataVisualization.svelte';
  import Play from '@/assets/img/play.svg';

  export let roomId;
  let roomData, sourceText, destinedText;
  let currentIndex = 0;
  let currentText = '';
  let hasError = false;

  onMount(() => {
    code.set(roomId);
    action.set('join');
    currentWordCount.set(0);
    participantId.set(Cookies.get('participantId'));
    if (!$participantId) {
      Cookies.set('participantId', `participant_${new Date().getTime()}`);
      participantId.set(Cookies.get('participantId'));
    }

    getRoomByCode();
    initWS();
  });

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
    wordCount.set(sourceText.length);
  }

  function getRoomByCodeFail() {
    console.log('get room by code failed');
  }

  function handleInputChange(event) {
    if (event.data === ' ') {
      if (destinedText[currentIndex] === event.target.value.replace(/\s/g, '')) {
        action.set('sync');
        hasError = false;
        currentIndex += 1;
        currentText = '';
        currentWordCount.set(currentIndex);
      } else {
        hasError = true;
      }
    }
  }

  function handleChangeAction(value) {
    action.set(value);
  }

  onDestroy(() => {
    closeWS();
  });
</script>

<style>
  .room {
    @apply mx-auto;
    @apply my-auto;
    @apply h-screen;
    @apply bg-blue-300;
    @apply p-4;
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
    @apply bg-white;
    @apply w-full;
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
    <div>Hallo, {$username} - {$participantId}</div>

    {#if roomData}
      <DataVisualization data={$wsStore} />

      <div>{roomData.source_lang} - {roomData.destined_lang}</div>

      <div class="text-wrapper flex flex-wrap justify-start">
        {#each sourceText as s, i}
          <div class="text" class:success={i < currentIndex} class:error={hasError && i === currentIndex}>
            {s}
            &nbsp; &nbsp;
          </div>
        {/each}
      </div>

      {#if $wsStore && $wsStore.my_state === 0}
        <div class="flex justify-center">
          <Button on:clicked={() => handleChangeAction('ready')} text="Ready" class="mt-4">
            <div class="w-8" slot="icon">
              <Play />
            </div>
          </Button>
        </div>
      {:else if $wsStore && $wsStore.my_state === 1 && $wsStore.room_state === 0}
        <div class="flex justify-center mt-4">Waiting . . .</div>
      {:else if $wsStore && $wsStore.room_state === 1 && $wsStore.leader_board[$participantId].state !== 2}
        <div class="input-field border-none shadow-md" for="username">
          <input
            id="inputText"
            type="text"
            class="placeholder-blue w-full p-2 no-outline text-dusty-blue-darker"
            name="inputText"
            placeholder="Type here"
            bind:value={currentText}
            on:input={handleInputChange} />
        </div>
      {:else if $wsStore && $wsStore.leader_board[$participantId].state === 2}
        <div class="flex justify-center mt-4">Finish :)</div>
      {/if}
    {/if}
  </div>
</div>
