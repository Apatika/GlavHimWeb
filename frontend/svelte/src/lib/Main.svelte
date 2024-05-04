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
  const timer = ms => new Promise(res => setTimeout(res, ms))
  const refreshInWork = async () => {
    while(true){
      getInWork()
      await timer(5000)
    }
  }
  //refreshInWork()

  const changeStatus = (id, status) => {
    fetch(`${uri}/orders/status`, {
      method: "PUT",
      body: JSON.stringify({id: id, status: status}),
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

  const toggleFullOrder = (event) => {
    let target = event.target.parentElement.nextElementSibling
    if (target.style.display != "flex") {
      target.style.display = "flex"
      event.target.innerHTML = 'ᐃ'
    }
    else {
      target.style.display = "none"
      event.target.innerHTML = 'ᐁ'
    }
  }

</script>

<div id="container">
  <div id="content">
    <button id="refresh" on:click={getInWork}>Обновить</button>
    <div id="table">
      {#each orders as order}
        <div>
          <div class="order">
            <div class="order-item cargo">{order.order.cargo}</div>
            <div class="order-item name">{order.client.surname} {order.client.name} {order.client.secondName}</div>
            <div class="order-item adress">{order.order.adress.city}</div>
            {#if order.order.payment} <div class="achtung"><strong>ЗА НАШ СЧЕТ</strong></div> {/if}
            {#if order.order.probes.length > 0} <div class="achtung"><strong>ПРОБНИКИ</strong></div> {/if}
            <div class="order-item status">
              <select bind:value={order.order.status} on:change={(e) => changeStatus(order.order.id, e.target.value)}>
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
            <button class="expander" on:click={toggleFullOrder}>ᐁ</button>
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
  .order{
    display: flex;
    height: 20px;
    color: black;
    font-size: 14px;
    border: 1px solid black;
  }
  .order-item{
    padding: 0px 3px;
    vertical-align: center;
    border-right: 1px solid black;
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
    display: none;
    position: relative;
    padding: 15px;
    box-shadow: 0px -5px 10px 0px black inset;
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
    color: red;
  }
  .full-item{
    display: flex;
  }
  .full-label{
    flex-basis: 25%;
  }
  .expander{
    text-align: center;
    vertical-align: auto;
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
    height: 90%;
  }
  #new-order{
    flex-basis: 25%;
  }
  #refresh{
    margin-left: 10px;
  }
</style>
