<script>
  import NewOrder from "./NewOrder.svelte"
  import Order from "../Order.svelte"
  import OrderFull from "../OrderFull.svelte"
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()
  
  let editCustomer = {
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
    customerId: null,
    payment: null,
    toadress: false,
    adress: {city: null, adress: null, terminal: "основной"},
    invoice: [null],
    cargo: null,
    lastDate: null,
    comment: null,
    probes: {},
  }

  let probeCountSum = 0
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

  const toggleFullOrder = (event) => {
    let parent = event.target.closest('.order')
    let target = parent.nextElementSibling
    let height = "160px"
    if (window.screen.width < 1366) height = "300px"
    if (target.style.height != height) {
      closeFullOrder()
      target.style.padding = '15px'
      target.style.height = height
      parent.style.transform = 'scale(1.005)'
      parent.style.boxShadow = '0px 0px 10px black'
    }
    else {
      target.style.padding = null
      target.style.height = '0px'
      parent.style.transform = 'scale(1)'
      parent.style.boxShadow = 'none'
    }
  }

  const dispatchEdit = (detail) => {
    closeFullOrder()
    Object.keys(chems).forEach(key => chems[key].probeCount = 0)
    let target = document.querySelector('#editor')
    target.style.display = "none"
    document.body.style.pointerEvents = "all"
    if (detail.update){
      dispatch('message', {routeId: detail.id})
    }
  }

  const onEdit = (order, customer) => {
    for (let [k, v] of Object.entries(order.probes)){
      chems[k] = v
      probeCountSum += v.probeCount
    }
    editCustomer = customer
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
      <NewOrder order={editOrder} customer={editCustomer} {managers} {cargos} {chems} {probeCountSum} on:message={(e) => dispatchEdit(e.detail)}></NewOrder>
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
          <div class="order">
            <Order {order} on:message={(e) => toggleFullOrder(e.detail)} on:status={closeFullOrder}></Order>
          </div>
          <div class="full-order">
            <OrderFull {order} on:message={(e) => {onEdit(e.detail.order, e.detail.customer)}}></OrderFull>
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
  .container {
    display: flex;
  }
  .container div:not(:last-child) {
    align-items: stretch;
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
  .order{
    position: relative;
  }
  .full-order{
    height: 0px;
    border-radius: 5px;
    background-color: #FFF5EE;
    box-shadow: 0px -3px 3px 0px black inset;
    transition: all .2s linear;
    overflow: hidden;
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
    #new-order, .legend{
      display: none;
    }
    #new-order{
      position: absolute;
      width: 100%;
      height: 100vh;
      z-index: 15;
    }
    #table-container{
      width: 200px;
      margin: 10px 40px 10px 5px;
    }
    #editor{
      height: 100vh;
      width: 100%;
      font-size: 7px;
    }
    .full-order{
      box-shadow: none;
      overflow: auto;
    }
    #new-order-expand{
      position: fixed;
      display: block;
      bottom: 0;
      width: 100%;
      z-index: 20;
    }
  }
</style>