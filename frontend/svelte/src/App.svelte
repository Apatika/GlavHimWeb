<script>
  import Main from "./lib/Main.svelte";
  import Data from "./lib/Data.svelte";

  let currentPage = Main

</script>

<div class="nav">
  <button on:click={() => currentPage = Main}>Главная</button>
  <button on:click={() => currentPage = Data}>База Данных</button>
</div>
<div class="content">
  <svelte:component this={currentPage} />
</div>

<style>
  button{
    display: block;
    position: relative;
    border: none;
    margin-left: 2px;
    margin-top: 1px;
    font-size: 18px;
    height: 40px;
    border-top: 1px solid #fdf5df;
    border-radius: 5px;
    background-color: #191970;
    box-shadow: 0px 0px 10px 2px black;
  }
  button:hover{
    box-shadow: 0px 0px 10px 2px black, 0px 0px 2px 0px white inset;
  }
  button:active{
    box-shadow: 0px 0px 10px 0px black inset;
  }
  .nav{
    display: flex;
    flex-direction: column;
    background-color: #fdf5df;
    width: 150px;
  }
  .content{
    background-color: #fdf5df;
    flex-grow: 1;
  }
</style>
