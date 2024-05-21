<script>

  export let orders = {}

  const add = (e, order) => {
    changeStatus(order, "В Маршрут")
  }

  const del = (order) => {
    changeStatus(order, "Развозка")

  }

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
    }).catch((err) => {
      alert(err)
    })
  }

  const update = (o) => {
    let char = o.order.routeNum[0]
    if (char >= '0' && char <= '9'){
      o.order.routeNum = char
    }
    else {
      o.order.routeNum = ""
      return
    }
    fetch(`${window.location.origin}/orders`, {
      method: "PUT",
      body: JSON.stringify(o),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
    }).catch((err) => {
      alert(err)
    })
  }

  const delivered = (id) => {
    changeStatus(orders[id], 'Передан')
  }

</script>

<div class="container">
  <div id="all-routes">
    {#each Object.values(orders).filter(val => val.order.status == "Развозка") as order}
      <div class="route">
        <div>
          {order.order.cargo}
        </div>
        <div>
          {#if order.client.surname != ""}{order.client.surname} {/if}
          {order.client.name}
          {#if order.client.surname != ""} {order.client.secondName}{/if}          
        </div>
        <div>
          <span>{order.order.adress.city}, </span>
          {#if order.order.adress.adress != ""}
            {order.order.adress.adress}
          {:else}
            {order.order.adress.terminal}
          {/if}
        </div>
        <button class="toggle" on:click={(e) => add(e, order)}>></button>
      </div>
    {/each}
  </div>
  <div id="selected-routes">
    {#each Object.values(orders).filter(val => val.order.status == "В Маршрут") as order}
      <div class="route selected">
        <div>
          {order.order.cargo}
        </div>
        <div>
          {#if order.client.surname != ""}{order.client.surname} {/if}
          {order.client.name}
          {#if order.client.surname != ""} {order.client.secondName}{/if}          
        </div>
        <div>
          <span>{order.order.adress.city}, </span>
          {#if order.order.adress.adress != ""}
            {order.order.adress.adress}
          {:else}
            {order.order.adress.terminal}
          {/if}
        </div>
        <div class="move">
          <button on:click={() => del(order)}>X</button>
        </div>
      </div>
    {/each}
  </div>
  <div id="info">
    {#each Object.values(orders).filter(val => val.order.status == "В Маршрут").sort((a, b) => {return a.order.routeNum > b.order.routeNum ? 1 : (a.order.routeNum == b.order.routeNum ? 0 : -1)}) as order}
      <div class="info">
        <div class="item">
          <div class="label">Способ доставки:</div>
          <div class="value">{order.order.cargo}</div>
        </div>
        <div class="item">
          <div class="label">Наименование:</div>
          <div class="value">
            {#if order.client.surname != ""}{order.client.surname} {/if}
            {order.client.name}
            {#if order.client.surname != ""} {order.client.secondName}{/if}
          </div>
        </div>
        <div class="item">
          <div class="label">Менеджер:</div>
          <div class="value">{order.client.manager}</div>
        </div>
        <div class="item">
          <div class="label">Город:</div>
          <div class="value">{order.order.adress.city}</div>
        </div>
        {#if order.order.adress.adress != ""}
          <div class="item">
            <div class="label">Адрес:</div>
            <div class="value">{order.order.adress.adress}</div>
          </div>
        {:else}
          <div class="item">
            <div class="label">Терминал:</div>
            <div class="value">{order.order.adress.terminal}</div>
          </div>
        {/if}
        <div class="item">
          <div class="label">Контактное лицо:</div>
          <div class="value">
            {#each order.client.contact as contact}
              <div>{contact.name} - {contact.tel}</div>
            {/each}
          </div>
        </div>
        {#if order.order.comment != ""}
          <div class="item">
            <div class="label">Комментарий:</div>
            <div class="value">{order.order.comment}</div>
          </div>
        {/if}
        <div class="item">
          {#if order.order.payment}
            <div class="item">
              <div class="label"></div>
              <div class="value"><strong>ЗА НАШ СЧЕТ</strong></div>
            </div>
          {/if}
        </div>
        <button on:click={() => delivered(order.order.id)}>передан</button>
        <input class="route-num" type="text" bind:value={order.order.routeNum} on:input={() => update(order)}>
      </div>
    {/each}
  </div>
</div>

<style>

  .container{
    display: flex;
    padding: 30px;
  }
  .route{
    position: relative;
    padding: 5px;
    margin: 5px auto;
    border: 1px solid black;
    border-radius: 15px;
    background-color: #FFA500;
    box-shadow: 0px 0px 5px 0px grey;
    font-size: small;
    width: 300px;
  }
  .selected{
    background-color: #008B8B;
  }
  .toggle, .move{
    position: absolute;
    right: 20px;
    top: 50%;
    transform: translate(0, -50%);
  }
  .info{
    border-bottom: 1px solid black;
    margin-bottom: 3px;
  }
  .item{
    display: flex;
  }
  .label{
    width: 120px;
    font-weight: bold;
  }
  .value{
    width: 370px;
    word-wrap: break-word;
  }
  .route-num{
    width: 20px;
    background-color: #FFA500;
    border: none;
    border-radius: 5px;
    text-align: center;
  }
  #all-routes, #selected-routes{
    width: 350px;
    border-right: 1px solid black;
    height: 100vh;
  }
  #info{
    padding: 5px;
    width: 500px;
    font-size: 12px;
  }
  @media (max-width:1364px){
    .container{
      display: block;
      padding: 5px;
    }
    #all-routes, #selected-routes{
      display: none;
    }
    .route-num{
      pointer-events: none;
    }
  }

</style>