<script>
  import Main from "./lib/Main.svelte";
  import Data from "./lib/Data.svelte";
  import { setContext } from 'svelte'

  let settings = {
    refreshRate: 5000,
  }
  fetch(`/settings`).then(function(response) {
    if (response.status != 200) throw new Error(response.statusText)
    return response.json();
  }).then(function(d) {
    settings = d
  }).catch(function(err) {
    alert(err);
  })

  setContext('settings', settings)

  let currentPage = Main

</script>

<div class="nav">
  <div>
    <a href="#top" on:click={() => currentPage = Main}>Главная</a>
  </div>
  <div>
    <a href="#top" on:click={() => currentPage = Data}>База Данных</a>
  </div>
</div>
<div class="content">
  <svelte:component this={currentPage} />
</div>

<style>
  .nav{
    flex-basis: 10%;
    background-color: darkblue;
    border-right: 2px solid black;
  }
  .nav div{
    display: flex;
    height: 40px;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    border-bottom: 1px solid white;
  }
  .content{
    background-color: white;
    flex-grow: 1;
  }
</style>
