<script>
  import { fade } from 'svelte/transition';
  import { goto } from '@sapper/app';
  import { username } from '@/store/index.js';
  import api from '@/api/index.js';
  import Metadata from '@/components/Metadata.svelte';
  import Play from '@/assets/img/play.svg';
  import Button from '@/components/Button.svelte';

  let code = '';
  let visibleInput = false;

  function toggleInput() {
    visibleInput = !visibleInput;
  }

  function createRoom() {
    const params = {
      source_lang: 'jp-hiragana',
      destined_lang: 'en'
    };
    api.createRoom(createRoomSuccess, params, createRoomFail);
  }

  function createRoomSuccess(response) {
    const { room_id } = response.body;
    goto(`/room/${room_id}`);
  }

  function createRoomFail() {
    console.log('create room failed');
  }

  function handleJoin() {
    goto(`/room/${code}`);
  }
</script>

<style>
  .home {
    @apply mx-auto;
    @apply my-auto;
    @apply h-screen;
    @apply bg-blue-300;
  }
  .home-wrapper {
    @apply mx-auto;
    @apply my-auto;
    width: 200px;
  }
  .title {
    @apply text-center;
    text-shadow: -4px 0 white, 0 1px white, 1px 0 white, 0 -1px white;
  }
</style>

<Metadata />

<div class="home">
  <h1 class="title text-6xl md:text-8xl py-16">Overtype</h1>
  <div class="home-wrapper">
    {#if !visibleInput}
      <div in:fade out:fade>
        <div class="pb-4 md:pb-0 flex flex-col">
          <label for="name" class="input-label text-base mb-2">Your Username</label>
          <div>
            <div class="input-field shadow-md bg-white" for="username">
              <input
                id="username"
                type="text"
                class="placeholder-blue w-full p-2 no-outline text-dusty-blue-darker"
                name="username"
                placeholder="Your name"
                bind:value={$username} />
            </div>
          </div>
        </div>
      </div>

      <Button on:clicked={() => createRoom()} text="Create Room" class="mt-4">
        <div class="w-8" slot="icon">
          <Play />
        </div>
      </Button>

      <Button on:clicked={() => toggleInput()} text="Join Room" class="mt-4">
        <div class="w-8" slot="icon">
          <Play />
        </div>
      </Button>
    {:else}
      <div in:fade out:fade>
        <div class="pb-4 md:pb-0 flex flex-col">
          <label for="code" class="input-label text-base mb-2">Input Code</label>
          <div>
            <div class="input-field inline-flex items-baseline border-none shadow-md bg-white" for="username">
              <input
                id="code"
                type="text"
                class="placeholder-blue w-48 p-2 no-outline text-dusty-blue-darker"
                name="code"
                placeholder="Code"
                bind:value={code} />
            </div>
          </div>
        </div>

        <div class="button-primary menu" on:click={() => handleJoin()}>
          <Play />
        </div>
      </div>
    {/if}
  </div>
</div>
