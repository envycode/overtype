<script context="module">
  import api from '@/api/index.js';
  import Metadata from '@/components/Metadata.svelte';
  export async function preload({ params }) {
    const data = {
      source_lang: 'jp-hiragana',
      destined_lang: 'en'
    };
    api.getContent(
      response => {
        const { result } = response.body;
        console.log(result, response);
        success && success(response);
      },
      data,
      createRoomFail()
    );
  }
</script>

<script>
  export let post;
</script>

<style>
</style>

<Metadata />
