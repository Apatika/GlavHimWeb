<script>
  import NewOrder from "./NewOrder.svelte"
  
  $: orders = []

  let isEdit = false

  const updateOrders = (data) => {
    orders = []
    for (let v of Object.values(data)){
      orders.push(v)
    }
    orders = orders.sort((a, b) => {
      if (a.order.status > b.order.status) return -1
      if (a.order.status < b.order.status) return 1
      return 0
    })
  }

  const getInWork = () => {
    fetch(`${window.location.origin}/inwork`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then(function(d) {
      updateOrders(d)
    }).catch(function(err) {
      alert(err);
    })
  }
  getInWork()
  const timer = ms => new Promise(res => setTimeout(res, ms))
  const refreshInWork = async () => {
    let refresh = 10000
    await fetch(`/settings`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json()
    }).then((d) => {
      refresh = d.refreshRate
    }).catch((err) => {
      alert(err)
    })
    while(true){
      await timer(refresh)
      if (!isEdit) getInWork()
    }
  }
  //refreshInWork()

  const changeStatus = (order, status) => {
    order.order.status = status
    fetch(`${window.location.origin}/orders/status`, {
      method: "PUT",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      getInWork()
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
      default:
        return "white"
    }
  }

  const toggleFullOrder = (event) => {
    let target = event.target.parentElement.nextElementSibling
    if (target.style.height != '160px') {
      getInWork()
      document.querySelectorAll('.full-order').forEach((elem) => {
        elem.style.padding = null
        elem.style.height = '0px'
      })
      document.querySelectorAll('.order').forEach((elem) => {
        elem.style.transform = 'scale(1)'
        elem.style.boxShadow = 'none'
      })
      target.style.padding = '15px'
      target.style.height = '160px'
      event.target.parentElement.style.transform = 'scale(1.005)'
      event.target.parentElement.style.boxShadow = '0px 0px 10px black'
    }
    else {
      target.style.padding = null
      target.style.height = '0px'
      event.target.parentElement.style.transform = 'scale(1)'
      event.target.parentElement.style.boxShadow = 'none'
    }
  }

  const onEdit = (e) => {
    e.target.parentElement.parentElement.previousElementSibling.style.display = "block"
    document.body.style.pointerEvents = "none"
    e.target.parentElement.parentElement.previousElementSibling.style.pointerEvents = "all"
    isEdit = true
  }

</script>

<div class="container">
  <div id="table-container">
    <div>
      <button id="refresh" on:click={getInWork}>Обновить</button>
    </div>
    <div id="table">
      {#each orders as order}
        <div class="editor">
          <NewOrder order={order.order} client={order.client} on:message={() => {isEdit = false; getInWork()}}></NewOrder>
        </div>
        <div>
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <!-- svelte-ignore a11y-no-static-element-interactions -->
          <div class="order" style="background-color: {getColor(order.order.status)}">
            <div class="order-item cargo" on:click={toggleFullOrder}>{order.order.cargo}</div>
            <div class="order-item name" on:click={toggleFullOrder}>{order.client.surname} {order.client.name} {order.client.secondName}</div>
            <div class="order-item adress" on:click={toggleFullOrder}>{order.order.adress.city}</div>
            {#if order.order.payment} <div class="achtung"><strong>ЗА НАШ СЧЕТ</strong></div> {/if}
            {#if order.order.probes.length > 0} <div class="achtung"><strong>ПРОБНИКИ</strong></div> {/if}
            <div class="order-item status">
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
                      <span>{invoice}, </span>
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
            <div class="probes"></div>
            <button class="edit" on:click={onEdit}>Редактировать</button>
          </div>
        </div>
      {:else}
        <span>Нет заказов</span>
      {/each}
    </div>
  </div>
  <div id="new-order">
    <div id="new-order-title"><strong>Новый Заказ</strong></div>
    <NewOrder on:message={() => {isEdit = false; getInWork()}}></NewOrder>
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
    height: 20px;
    color: black;
    font-size: 14px;
    border: 1px solid black;
    border-radius: 5px;
    background-color: white;
    transition: all .1s linear;
  }
  .order-item{
    padding: 0px 3px;
    vertical-align: center;
    border-right: 1px solid black;
    cursor: pointer;
  }
  .adress{
    flex-basis: 15%;
    flex-grow: 1;
    border-right: none;
  }
  .status{
    border-left: 1px solid black;
    border-right: 1px solid black;
  }
  .full-order{
    display: flex;
    height: 0px;
    border-radius: 5px;
    background-color: #FFEBCD;
    box-shadow: 0px -3px 3px 0px black inset;
    transition: all .2s linear;
    overflow: hidden;
  }
  .full-order-order{
    flex-basis: 35%;
    font-size: small;
  }
  .full-order-client{
    flex-basis: 25%;
    font-size: small;
  }
  .achtung{
    padding: 0px, 5px;
    flex-basis: 10%;
    color: #0000CD;
    pointer-events: none;
  }
  .full-item{
    display: flex;
  }
  .full-label{
    flex-basis: 25%;
  }
  .edit{
    float: right;
    height: 30px;
  }
  .probes{
    flex-grow: 1;
  }
  .editor{
    display: none;
    position: absolute;
    width: 500px;
    height: 90%;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    border: 1px solid black;
    border-radius: 5px;
    background-color: #FFEBCD;
    overflow:auto;
    box-shadow: 0px 0px 10px 0px black;
    z-index: 5;
  }
  #new-order{
    flex-basis: 21%;
    background-color: #FFEBCD;
    border: 1px solid black;
    border-radius: 5px;
    box-shadow: 0px 0px 10px 0px black;
  }
  #new-order-title {
    text-align: center;
  }
  #table-container{
    margin: 10px 30px;
    flex-grow: 1;
  }
  #table{
    border-radius: 5px;
    box-shadow: 0px 0px 10px 0px black, 0px 0px 10px 0px black inset;
  }
  #refresh{
    margin-left: 10px;
    margin-bottom: 7px;
  }
</style>