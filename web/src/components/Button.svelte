<script>
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();
  let clazz;

  export { clazz as class };
  export let text;
  export let iconButton;

  function handleClick() {
    dispatch('clicked');
  }
</script>

<style>
</style>

<div class={`button-primary flex flex-wrap justify-space-between ${clazz || ''}`} on:click={() => handleClick()}>
  <slot name="icon" />
  {#if !iconButton}
    <div class="self-center ml-2">{text}</div>
  {/if}
</div>
