<script>
  // import Nav from '../components/Nav.svelte';

  // export let segment;
</script>

<style>
  .wrapper {
    width: 100%;
    height: 100%;
    position: relative;
    z-index: 0;
    @apply bg-blue-300;
    @apply h-screen;
    overflow-x: hidden;
  }

  .wrapper img {
    position: absolute;
  }

  .cloud1 {
    width: 200px;
    animation: move 40s infinite alternate;
    margin-top: 5%;
  }

  .cloud2 {
    width: 200px;
    animation: move 30s infinite alternate-reverse;
    margin-top: 20%;
  }

  .cloud3 {
    width: 200px;
    animation: move 65s infinite alternate;
    margin-top: 40%;
  }

  @media screen and (max-width: 768px) {
    .cloud2 {
      margin-top: 40vh;
    }
    .cloud3 {
      margin-top: 70vh;
    }
  }
  main {
    @apply mx-auto;
    @apply my-auto;
    z-index: 1;
    margin: 0 auto;
    position: relative;
  }

  @keyframes move {
    0% {
      left: -200px;
    }
    100% {
      left: 100%;
    }
  }
</style>

<!-- <Nav {segment} /> -->
<div class="wrapper">
  <img class="cloud1" src="/images/cloud2.png" alt="cloud" />
  <img class="cloud2" src="/images/cloud2.png" alt="cloud" />
  <img class="cloud3" src="/images/cloud2.png" alt="cloud" />

  <main>
    <slot />
  </main>
</div>
