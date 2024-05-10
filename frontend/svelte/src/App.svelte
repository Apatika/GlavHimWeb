<script>
  import Main from "./lib/Main.svelte";
  import Data from "./lib/Data.svelte";

  let currentPage = Main

</script>

<div class="nav">
  <div>
    <button on:click={() => currentPage = Main}>Главная</button>
  </div>
  <div>
    <button on:click={() => currentPage = Data}>База Данных<button>
  </div>
  <div id="placeholder"></div>
</div>
<div class="content">
  <svelte:component this={currentPage} />
</div>

<style>
  button{
    display: block;
    width: 100%;
    background-color: transparent;
    border: none;
  }
  .nav{
    flex-basis: 10%;
    background-color: darkblue;
  }
  .nav div{
    display: flex;
    height: 40px;
    justify-content: center;
    align-items: center;
    box-shadow: -5px 0px 5px 0px white inset;
  }
  .content{
    background-color: white;
    flex-grow: 1;
  }
  #placeholder{
    box-shadow: -5px 0px 5px 0px white inset;
    height: 100%;
    overflow: hidden;
  }
</style>
