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
    box-shadow: 0px 0px 10px 0px grey;
  }
  .nav div{
    display: flex;
    height: 40px;
    justify-content: center;
    align-items: center;
    box-shadow: 0px 0px 10px 0px grey;
  }
  .content{
    background-color: #fdf5df;
    flex-grow: 1;
  }
  #placeholder{
    box-shadow: 0px 0px 10px 0px grey;
    height: 100%;
    overflow: hidden;
  }
</style>
