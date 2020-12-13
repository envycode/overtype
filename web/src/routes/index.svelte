<script>
  import { fade } from 'svelte/transition';
  import api from '@/api/index.js';
  import Metadata from '@/components/Metadata.svelte';
  import Play from '@/assets/img/play.svg';

  let code = '';
  let username = 'Username';
  let visibleInput = false;

  function toggleInput() {
    visibleInput = !visibleInput;
  }

  function createRoomFail() {
    console.log('create room failed');
  }

  function createRoom() {
    const params = {
      source_lang: 'jp-hiragana',
      destined_lang: 'en'
    };
    api.getContent(
      response => {
        const { result } = response.body;
        console.log(result, response);
        success && success(response);
      },
      params,
      createRoomFail()
    );
  }
</script>

<style>
  .home {
    @apply mx-auto;
    @apply my-auto;
    @apply h-screen;
    @apply bg-blue-300;
  }

  .title {
    @apply text-center;
    text-shadow: -4px 0 white, 0 1px white, 1px 0 white, 0 -1px white;
  }

  .menu {
    @apply w-48;
    @apply mt-8;
  }

  .home-wrapper {
    @apply mx-auto;
    @apply my-auto;
    width: 200px;
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
            <div class="input-field inline-flex items-baseline border-none shadow-md bg-white" for="username">
              <input
                id="username"
                type="text"
                class="placeholder-blue w-48 p-2 no-outline text-dusty-blue-darker"
                name="username"
                placeholder="Your name"
                bind:value={username} />
            </div>
          </div>
        </div>
      </div>
      <div class="button-primary menu flex flex-wrap justify-space-between" on:click={createRoom}>
        <div class="w-8">
          <Play />
        </div>
        <div class="self-center ml-2">Create Room</div>
      </div>

      <div class="button-primary menu flex flex-wrap justify-space-between" on:click={toggleInput}>
        <div class="w-8">
          <Play />
        </div>
        <div class="self-center ml-2">Join Room</div>
      </div>
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

        <div class="button-primary menu">
          <Play />
        </div>
      </div>
    {/if}
  </div>
</div>
