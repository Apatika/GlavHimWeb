<script>
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()

  export let order = {}
  export let isSearch = false

  const getColor = (status) => {
    let style = "background-color: "
    switch (status) {
      case "Принят В Работу":
        return style + "#FFFF00"
      case "СТОП":
        return style + "#FF00FF"
      case "Забор ПЭК":
        return style + "#B0C4DE"
      case "Заказан Забор":
        return style + "#00FFFF"
      case "Развозка":
        return style + "#FFA500"
      case "Нет Товара":
        return style + "#FF0000"
      case "Изменен!":
        return style + "#00FF00"
      case "В Маршрут":
        return style + "#008B8B"
      case "Передан":
        return style + "green"
      default:
        return style + "white"
    }
  }

  const changeStatus = (e, order, status) => {
    e.target.closest('.order').style.backgroundColor = getColor(e.target.value)
    order.order.status = status
    fetch(`${window.location.origin}/orders`, {
      method: "PUT",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('status', '')
    }).catch((err) => {
      alert(err)
    })
  }

</script>

<button class="toggle-button" on:click={e => dispatch('message', e)}></button>
<div class="order" style={getColor(order.order.status)}>
  <div class="item cargo">
    <div><strong>{order.order.cargo}</strong></div>
    <div class="additional">счет: 
      {#each order.order.invoice as invoice}
        <span>{invoice} </span>
      {/each}
    </div>
  </div>
  <div class="item name">{#if order.customer.surname != "" && order.customer.inn != ""}ИП {/if}<strong>{order.customer.surname} {order.customer.name}</strong> {order.customer.secondName}</div>
  <div class="item adress">
    <div><strong>{order.order.adress.city}</strong></div>
    {#if order.order.adress.adress != ""}
      <div>адрес: {order.order.adress.adress}</div>
    {:else}
      <div>терминал: {order.order.adress.terminal}</div>
    {/if}
  </div>
  {#if order.order.lastDate != ""} <div class="last-date"><strong>{order.order.lastDate}</strong></div> {/if}
  {#if Object.keys(order.order.probes).length > 0} <div class="achtung"><strong>🧴</strong></div> {/if}
  {#if order.order.payment} <div class="achtung" id="payment-span"><strong>💰</strong></div> {/if}
  {#if order.order.comment != ""} <div class="achtung">💬</div> {/if}
  {#if !isSearch}
    <div class="item status">
      <select bind:value={order.order.status}
              on:change={(e) => {changeStatus(e, order, e.target.value)}}>
        <option value=""></option>
        <option value="Принят В Работу">Принят В Работу</option>
        <option value="Развозка">Развозка</option>
        <option value="Забор ПЭК">Забор ПЭК</option>
        <option value="Заказан Забор">Заказан Забор</option>
        <option value="Нет Товара">Нет Товара</option>
        <option value="СТОП">СТОП</option>
        <option value="Отгружен">Отгружен</option>
        <option value="Изменен!" disabled>Изменен!</option>
        <option value="В Маршрут" disabled>В Маршрут</option>
        <option value="Передан" disabled>Передан</option>
      </select>
    </div>
  {/if}
</div>

<style>

  .toggle-button{
    position: absolute;
    width: 90%;
    height: 100%;
    border: none;
    background-color: transparent;
    z-index: 5;
  }
  .cargo{
    flex-basis: 15%;
  }
  .name{
    flex-basis: 25%;
  }
  .order{
    display: flex;
    height: 40px;
    color: black;
    font-size: 14px;
    border: 1px solid black;
    border-radius: 5px;
    background-color: white;
    transition: all .1s linear;
  }
  .item{
    padding: 5px 5px;
    border-right: 1px solid black;
    cursor: pointer;
    line-height: 1;
  }
  .adress{
    flex-basis: 15%;
    flex-grow: 1;
    border-right: none;
  }
  .status{
    position: relative;
    border-left: 1px solid black;
    cursor: pointer;
    z-index: 10;
  }
  .status select {
    height: 34px;
    text-align: center;
    background-color: transparent;
    border: none;
    color: black;
    font-weight: bold;
  }
  .status option:disabled{
    display: none;
  }
  .achtung{
    padding: 5px 0px;
    width: 30px;
    font-size: 18px;
    text-shadow:
    -1px -1px 0 black,
    1px -1px 0 black,
    -1px 1px 0 black,
    1px 1px 0 black;
  }
  .last-date{
    padding: 10px 0px;
    font-size: 12px;
    width: 75px;
    color: #0000CD;
    font-weight: bold;
  }
  #payment-span{
    margin-right: 5px;
  }
  @media (min-width:1365px) and (max-width:1600px){
    .order{
      font-size: 12px;
    }
  }
  @media (max-width:1364px){
    .adress, .achtung, .last-date, .additional{
      display: none;
    }
    .item, .status{
      font-size: 9px;
    }
    .status select{
      font-size: 9px;
    }
    .name{
      flex-grow: 1;
    }
  }

</style>