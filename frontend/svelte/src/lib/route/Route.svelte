<script>

  export let orders = {}
  export let ids = []

  const add = (e, order) => {
    changeStatus(order, "В Маршрут")
    ids.push(order.order.id)
  }

  const del = (order) => {
    changeStatus(order, "Развозка")
    const index = ids.indexOf(order.order.id);
    if (index > -1) {
      ids.splice(index, 1)
    }
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
    {#each ids as id}
      <div class="info">
        <div class="item">
          <div class="label">Способ доставки:</div>
          <div class="value">{orders[id].order.cargo}</div>
        </div>
        <div class="item">
          <div class="label">Наименование:</div>
          <div class="value">
            {#if orders[id].client.surname != ""}{orders[id].client.surname} {/if}
            {orders[id].order.cargo}
            {#if orders[id].client.surname != ""} {orders[id].client.secondName}{/if}
          </div>
        </div>
        <div class="item">
          <div class="label">Менеджер:</div>
          <div class="value">{orders[id].client.manager}</div>
        </div>
        <div class="item">
          <div class="label">Город:</div>
          <div class="value">{orders[id].order.adress.city}</div>
        </div>
        {#if orders[id].order.adress.adress != ""}
          <div class="item">
            <div class="label">Адрес:</div>
            <div class="value">{orders[id].order.adress.adress}</div>
          </div>
        {:else}
          <div class="item">
            <div class="label">Терминал:</div>
            <div class="value">{orders[id].order.adress.terminal}</div>
          </div>
        {/if}
        <div class="item">
          <div class="label">Контактное лицо:</div>
          <div class="value">
            {#each orders[id].client.contact as contact}
              <div>{contact.name} - {contact.tel}</div>
            {/each}
          </div>
        </div>
        {#if orders[id].order.comment != ""}
          <div class="item">
            <div class="label">Комментарий:</div>
            <div class="value">{orders[id].order.comment}</div>
          </div>
        {/if}
        <div class="item">
          {#if orders[id].order.payment}
            <div class="item">
              <div class="label"></div>
              <div class="value"><strong>ЗА НАШ СЧЕТ</strong></div>
            </div>
          {/if}
        </div>
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

</style>