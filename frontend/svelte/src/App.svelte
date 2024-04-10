<script>
  import Main from "./lib/Main.svelte";
  import NewOrder from "./lib/NewOrder.svelte";
  import Add from "./lib/Add.svelte";

  let currentPage = Main
</script>

<div class="nav">
  <div>
    <a href="#top" on:click={() => currentPage = Main}>Главная</a>
  </div>
  <div>
    <a href="#top" on:click={() => currentPage = NewOrder}>Новый Заказ</a>
  </div>
  <div>
    <a href="#top" on:click={() => currentPage = Add}>Добавить</a>
  </div>
</div>
<div class="content">
  <svelte:component this={currentPage}/>
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
