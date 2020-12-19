<script>
  import { wordCount } from '@/store/index.js';
  export let data;

  function onProgressChange(count) {
    return Math.ceil((count / $wordCount) * 12);
  }
</script>

<style>
  .progress {
    @apply bg-gray-200;
    @apply rounded-full;
    height: 8px;
  }

  .progress-value {
    @apply m-0;
    @apply p-0;
    @apply bg-purple-500;
    @apply rounded-full;
    @apply transition;
    @apply duration-1000;
    @apply ease-in-out;
    height: 8px;
  }

  .progress-image {
    @apply mt-4;
    @apply flex;
    min-width: 64px;
  }

  .progress-image > img {
    @apply w-16;
  }
</style>

<div class="visualization">
  {#if data && data.leader_board}
    {#each Object.keys(data.leader_board) as d, i}
      <div
        class={`
          progress-image
          min-w-4 w-${onProgressChange(data.leader_board[d].word_type)}/12
          justify-${onProgressChange(data.leader_board[d].word_type) === 0 ? 'start' : 'end'}`}>
        <img src="/images/car1.png" alt="mobil" />
      </div>
      <div class="progress">
        {#if onProgressChange(data.leader_board[d].word_type) > 0}
          <div class={`progress-value w-${onProgressChange(data.leader_board[d].word_type)}/12`}>&nbsp;</div>
        {:else}
          <div>&nbsp;</div>
        {/if}
      </div>
    {/each}
  {/if}
</div>
