<script context="module">
  export async function preload({ params }) {
    let [roomId] = params.slug;

    return { roomId };
  }
</script>

<script>
  import { onMount, onDestroy } from 'svelte';
  import Cookies from 'js-cookie';

  import { username, code, participantId, action, currentWordCount } from '@/store/index.js';
  import { initWS, closeWS } from '@/util/websocket.js';

  import api from '@/api/index.js';
  import Metadata from '@/components/Metadata.svelte';
  import Play from '@/assets/img/play.svg';

  export let roomId;
  let roomData, sourceText, destinedText;
  let currentIndex = 0;
  let currentText = '';
  let hasError = false;

  onMount(() => {
    console.log(roomId);
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
    <div>Hallo, {$username} - {$participantId}</div>
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

      <div class="button-primary w-24" on:click={() => handleChangeAction('ready')}>Ready</div>
    {/if}
  </div>
</div>
