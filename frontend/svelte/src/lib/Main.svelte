<script>
  import NewOrder from "./NewOrder.svelte";
  import { getContext } from 'svelte'

  const uri = getContext('uri')

  
  let orders = []

  const getInWork = () => {
    fetch(`${uri}/inwork`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then(function(d) {
      orders = d
    }).catch(function(err) {
      alert(err);
    })
  }
  getInWork()

</script>

<div id="container">
  <div id="content">
    <button id="refresh" on:click={getInWork}>Обновить</button>
    <div id="table">
      {#each orders as order}
      <div class="order">
        <div class="order-item cargo">{order.order.cargo}</div>
        <div class="order-item name">{order.client.name}</div>
        <div class="order-item adress">{order.order.adress.city}</div>
        <div class="order-item status">
          <select value={order.order.status}>
            <option value=""></option>
            <option value="Принят В Работу">Принят В Работу</option>
            <option value="СТОП">СТОП</option>
            <option value="Отгружен">Отгружен</option>
            <option value="Забор ПЭК">Забор ПЭК</option>
            <option value="Заказан Забор">Заказан Забор</option>
            <option value="Развозка">Развозка</option>
            <option value="Нет Товара">Нет Товара</option>
          </select>
        </div>
      </div>
      {:else}
      <span>Нет заказов</span>
      {/each}
    </div>
  </div>
  <div id="new-order">
    <NewOrder></NewOrder>
  </div>
</div>

<style>
  span{
    color: black;
  }
  select{
    height: 18px;
    margin: 1px;
    padding: 0px;
    text-align: center;
    background-color: transparent;
    border: none;
    color: black;
  }
  .cargo{
    flex-basis: 15%;
  }
  .name{
    flex-basis: 25%;
  }
  .adress{
    flex-basis: 30%;
    flex-grow: 1;
  }
  .order{
    display: flex;
    height: 20px;
    color: black;
    font-size: 14px;
    border-bottom: 1px solid black;
  }
  .order-item{
    padding: 0px 3px;
    vertical-align: center;
    border-right: 1px solid black;
  }
  .status{
    border-right: none;
  }
  #container{
    display: flex;
    height: 100%;
  }
  #content{
    flex-grow: 1;
    border-right: 1px solid black;
    padding: 10px 5px;
  }
  #table{
    margin: 10px 30px 10px 0px;
    border: 1px solid black;
    height: 90%;
  }
  #new-order{
    flex-basis: 25%;
  }
  #refresh{
    margin-left: 10px;
  }
</style>
