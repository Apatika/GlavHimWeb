<script>
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()
  import NewOrder from "./NewOrder.svelte"
  
  let editClient = {
    id: null,
    type: '0',
    adress: null,
    name: null,
    surname: null,
    secondName: null,
    manager: null,
    inn: null,
    passportSerial: null,
    passportNum: null,
    contact: [{name: null, tel: null}],
    email: null,
  }
  let editOrder = {
    id: "",
    clientId: null,
    payment: null,
    toadress: false,
    adress: {city: null, adress: null, terminal: "основной"},
    invoice: [null],
    cargo: null,
    lastDate: null,
    comment: null,
    probes: {},
  }

  export let managers = {}
  export let cargos = {}
  export let chems = {}
  export let orders = {}

  const closeFullOrder = () => {
    document.querySelectorAll('.full-order').forEach((elem) => {
      elem.style.padding = null
      elem.style.height = '0px'
    })
    document.querySelectorAll('.order').forEach((elem) => {
      elem.style.transform = 'scale(1)'
      elem.style.boxShadow = 'none'
    })
  }

  const changeStatus = (order, status) => {
    closeFullOrder()
    order.order.status = status
    fetch(`${window.location.origin}/orders/status`, {
      method: "PUT",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
    }).catch((err) => {
      alert(err)
    })
  }

  const getColor = (status) => {
    switch (status) {
      case "Принят В Работу":
        return "#FFFF00"
      case "СТОП":
        return "#FF00FF"
      case "Забор ПЭК":
        return "#B0C4DE"
      case "Заказан Забор":
        return "#00FFFF"
      case "Развозка":
        return "#FFA500"
      case "Нет Товара":
        return "#FF0000"
      case "Изменен!":
        return "#00FF00"
      case "В Маршрут":
        return "#008B8B"
      case "Передан":
        return "green"
      default:
        return "white"
    }
  }

  const toggleFullOrder = (event) => {
    let target = event.target.closest('.order').nextElementSibling
    let height = "160px"
    if (window.screen.width < 1366) height = "300px"
    if (target.style.height != height) {
      closeFullOrder()
      target.style.padding = '15px'
      target.style.height = height
      event.target.closest('.order').style.transform = 'scale(1.005)'
      event.target.closest('.order').style.boxShadow = '0px 0px 10px black'
    }
    else {
      target.style.padding = null
      target.style.height = '0px'
      event.target.closest('.order').style.transform = 'scale(1)'
      event.target.closest('.order').style.boxShadow = 'none'
    }
  }

  const dispatchEdit = (detail) => {
    closeFullOrder()
    let target = document.querySelector('#editor')
    target.style.display = "none"
    document.body.style.pointerEvents = "all"
    for (let k of Object.keys(chems)){
      chems[k].probeCount = 0
    }
    if (detail.update){
      dispatch('message', {routeId: detail.id})
    }
  }

  const onEdit = (order, client) => {
    editClient = client
    editOrder = order
    let target = document.querySelector('#editor')
    target.style.display = "block"
    document.body.style.pointerEvents = "none"
    target.style.pointerEvents = "all"
  }

  const sortOrders = (a, b) => {
    return a.order.status > b.order.status ? -1 : (a.order.status < b.order.status ? 1 : (a.order.cargo < b.order.cargo ? -1 : 1))
  }

  const toggleNewOrder = (e) => {
    let target = document.querySelector('#new-order')
    if (target.style.display == "block"){
      target.style.display = "none"
      e.target.innerHTML = "НОВЫЙ ЗАКАЗ"
    } else {
      target.style.display = "block"
      e.target.innerHTML = "ЗАКРЫТЬ"
    }
  }

</script>

<div class="container">
  <div id="table-container">
    <div id="editor">
      <NewOrder order={editOrder} client={editClient} {managers} {cargos} {chems} on:message={(e) => dispatchEdit(e.detail)}></NewOrder>
    </div>
    <div class="legend">
      <span>🧴 - пробники</span>
      <span>💰 - за наш счет</span>
      <span>💬 - комментарий</span>
      <span><strong>2024-05-17</strong> - крайняя дата</span>
    </div>
    <div id="table">
      {#each Object.values(orders).sort(sortOrders) as order}
        <div>
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <!-- svelte-ignore a11y-no-static-element-interactions -->
          <div class="order" style="background-color: {getColor(order.order.status)}">
            <div class="order-item cargo" on:click={toggleFullOrder}>
              <div><strong>{order.order.cargo}</strong></div>
              <div class="additional">счет: 
                {#each order.order.invoice as invoice}
                  <span>{invoice} </span>
                {/each}
              </div>
            </div>
            <div class="order-item name" on:click={toggleFullOrder}><strong>{order.client.surname} {order.client.name}</strong> {order.client.secondName}</div>
            <div class="order-item adress" on:click={toggleFullOrder}>
              <div><strong>{order.order.adress.city}</strong></div>
              {#if order.order.adress.adress != ""}
                <div>адрес: {order.order.adress.adress}</div>
              {:else}
                <div>терминал: {order.order.adress.terminal}</div>
              {/if}
            </div>
            {#if order.order.lastDate != ""} <div class="last-date" on:click={toggleFullOrder}><strong>{order.order.lastDate}</strong></div> {/if}
            {#if Object.keys(order.order.probes).length > 0} <div class="achtung" on:click={toggleFullOrder}><strong>🧴</strong></div> {/if}
            {#if order.order.payment} <div class="achtung" id="payment-span" on:click={toggleFullOrder}><strong>💰</strong></div> {/if}
            {#if order.order.comment != ""} <div class="achtung">💬</div> {/if}
            <div class="order-item status" on:click={() => {return}}>
              <select bind:value={order.order.status}
                      on:change={(e) => {changeStatus(order, e.target.value); e.target.parentElement.parentElement.style.backgroundColor = getColor(e.target.value)}}>
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
          </div>
          <div class="full-order">
            <div class="full-order-client">
              <div class="full-item">
                <div class="full-label"><strong>Менеджер:</strong></div>
                <div>{order.client.manager}</div>
              </div>
              {#if order.client.surname != ""}
                <div class="full-item">
                  <div class="full-label"><strong>Фамилия:</strong></div>
                  <div>{order.client.surname}</div>
                </div>
              {/if}
                <div class="full-item">
                  <div class="full-label"><strong>Имя:</strong></div>
                  <div>{order.client.name}</div>
                </div>
              {#if order.client.secondName != ""}
                <div class="full-item">
                  <div class="full-label"><strong>Отчество</strong></div>
                  <div>{order.client.secondName}</div>
                </div>
              {/if}
              {#if order.client.inn != ""}
                <div class="full-item">
                  <div class="full-label"><strong>ИНН:</strong></div>
                  <div>{order.client.inn}</div>
                </div>
              {/if}
              {#if order.client.passportSerial != ""}
                <div class="full-item">
                  <div class="full-label"><strong>Паспорт:</strong></div>
                  <div>{order.client.passportSerial} {order.client.passportNum}</div>
                </div>
              {/if}
                <div class="full-item">
                  <div class="full-label"><strong>Контакты:</strong></div>
                  <div>
                    {#each order.client.contact as contact}
                      <div>{contact.name} {contact.tel}</div>
                    {/each}
                  </div>
                </div>
              {#if order.client.email != ""}
                <div class="full-item">
                  <div class="full-label"><strong>Email:</strong></div>
                  <div>{order.client.email}</div>
                </div>
              {/if}
            </div>
            <div class="full-order-order">
              <div class="full-item">
                <div class="full-label"><strong>ТК:</strong></div>
                <div>{order.order.cargo}</div>
              </div>
              <div class="full-item">
                <div class="full-label"><strong>Город:</strong></div>
                <div>{order.order.adress.city}</div>
              </div>
              {#if order.order.adress.adress != ""} 
                <div class="full-item">
                  <div class="full-label"><strong>Адрес:</strong></div>
                  <div>{order.order.adress.adress}</div>
                </div>
              {:else}
                <div class="full-item">
                  <div class="full-label"><strong>Терминал:</strong></div>
                  <div>{order.order.adress.terminal}</div>
                </div>
              {/if}
              <div>
                <div class="full-item">
                  <div class="full-label"><strong>Счет:</strong></div>
                  <div>
                    {#each order.order.invoice as invoice}
                      <span>{invoice} </span>
                    {/each}
                  </div>
                </div>  
              </div>
              <div class="full-item">
                <div class="full-label"><strong>Крайняя дата:</strong></div>
                <div>{order.order.lastDate}</div>
              </div>
              <div class="full-item">
                <div class="full-label"><strong>Комментарий:</strong></div>
                <div>{order.order.comment}</div>
              </div>
            </div>
            <div class="probes">
              {#if Object.keys(order.order.probes).length > 0}<span><strong>ПРОБНИКИ</strong></span>{/if}
              {#each Object.values(order.order.probes).filter(val => val.probeCount > 0).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as chem}
                <div>{chem.name} - {(chem.probeValue * chem.probeCount) / 1000} л.</div>
              {/each}
            </div>
            <div>
              <div>
                <button class="edit" on:click={() => onEdit(order.order, order.client)}>Редактировать</button>
              </div>
              <div>
                {#if order.order.payment} <span style="color: red; font-size: 14px;"><strong>ЗА НАШ СЧЕТ!!!</strong></span> {/if}
              </div>
            </div>
          </div>
        </div>
      {:else}
        <span>Нет заказов</span>
      {/each}
    </div>
  </div>
  <div id="new-order">
    <div id="new-order-title"><strong>Новый Заказ</strong></div>
    <NewOrder {managers} {cargos} {chems}></NewOrder>
  </div>
  <button id="new-order-expand" on:click={toggleNewOrder}>НОВЫЙ ЗАКАЗ</button>
</div>

<style>
  span{
    color: black;
  }
  select{
    height: 34px;
    text-align: center;
    background-color: transparent;
    border: none;
    color: black;
    font-weight: bold;
  }
  .container {
    display: flex;
  }
  .container div:not(:last-child) {
    align-items: stretch;
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
  .order-item{
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
    border-left: 1px solid black;
    border-right: 1px solid black;
    cursor: pointer;
  }
  .status option:disabled{
    display: none;
  }
  .full-order{
    display: flex;
    height: 0px;
    border-radius: 5px;
    background-color: #FFF5EE;
    box-shadow: 0px -3px 3px 0px black inset;
    transition: all .2s linear;
    overflow: hidden;
  }
  .full-order-order{
    width: 350px;
    font-size: small;
  }
  .full-order-client{
    width: 300px;
    font-size: small;
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
  .full-item{
    display: flex;
    word-break: break-word;
    word-wrap: break-word;
    overflow-wrap: break-word;
  }
  .full-label{
    width: 100px;
    min-width: 100px;
  }
  .edit{
    float: right;
    height: 30px;
  }
  .probes{
    flex-grow: 1;
    text-align: center;
    font-size: small;
  }
  .legend{
    margin: 5px;
  }
  .legend span{
    margin-right: 10px;
    color: white;
  }
  .legend strong{
    color: #0000CD;
  }
  #payment-span{
    margin-right: 5px;
  }
  #editor{
    display: none;
    position: absolute;
    width: 500px;
    height: 90%;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    border: 1px solid black;
    border-radius: 5px;
    background-color: #FFF5EE;
    overflow:auto;
    box-shadow: 0px 0px 10px 0px black;
    z-index: 5;
  }
  #new-order{
    width: 370px;
    background-color: #FFF5EE;
    border: 1px solid black;
    border-radius: 5px;
    box-shadow: 0px 0px 10px 0px black;
  }
  #new-order-title {
    text-align: center;
    color: #778899;
  }
  #table-container{
    margin: 10px 30px;
    flex-grow: 1;
  }
  #table{
    border-radius: 5px;
    box-shadow: 0px 0px 10px 0px black, 0px 0px 10px 0px black inset;
  }
  #new-order-expand{
    display: none;
    position: absolute;
    bottom: 0;
    right: 0;
    width: 370px;
    height: 40px;
    border-radius: none;
  }
  @media (min-width:1365px) and (max-width:1600px){
    .order{
      font-size: 12px;
    }
    .full-label{
      flex-basis: 30%;
      min-width: 30%;
    }
    #editor{
      width: 400px;
    }
    #new-order-expand{
      display: inline-block;
    }
    #new-order{
      display: none;
      position: absolute;
      right: 0;
      bottom: 40px;
    }
  }
  @media (max-width:1364px){
    .adress, .achtung, #new-order, .last-date, .legend, .additional{
      display: none;
    }
    .full-order{
      display: block;
      overflow: auto;
    }
    .edit{
      margin-bottom: 5px;
    }
    .name{
      flex-grow: 1;
    }
    #table-container{
      width: 200px;
    }
    .order-item, .status{
      font-size: 9px;
    }
    #editor{
      width: 300px;
      font-size: 7px;
    }
  }
</style>