<script>
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();

  let clientDefault = {
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
  let orderDefault = {
    id: null,
    clientId: null,
    payment: null,
    toadress: false,
    adress: {city: null, adress: null, terminal: "основной"},
    invoice: [null],
    cargo: null,
    lastDate: null,
    comment: null,
    probes: [],
  }

  export let client = structuredClone(clientDefault)
  export let order = structuredClone(orderDefault)

  let cities = []
  export let managers = {}
  export let cargos = {}

  const check = () => {
    if (client.manager == null || client.manager == ""){
      alert("Менеджер не выбран")
      return false
    }
    if (order.cargo == null || order.cargo == ""){
      alert("Способ доставки не выбран")
      return false
    }
    if (order.adress.city == null || order.adress.city == ""){
      alert("Город не заполнен")
      return false
    }
    if (order.toadress && (order.adress.adress == null || order.adress.adress == "")){
      alert("Адрес не заполнен")
      return false
    }
    if (client.name == null || client.name == ""){
      alert("Имя не заполнено")
      return false
    }
    if (client.contact[0].tel == null || client.contact[0].tel == ""){
      alert("Контакт не заполнен")
      return false
    }
    if ((client.type == "0" || client.type == "2") && 
        (client.surname == null || client.surname == "")){
      alert("ФИО заполнены не полностью")
      return false
    }
    if ((client.type == "1" || client.type == "0") && (client.inn == null || client.inn == "")){
      alert("ИНН не заполнен")
      return false
    }
    return true
  }

  const addContact = () => client.contact = client.contact.concat({name: "", tel: ""})
  const removeContact = () => {
    if (client.contact.length > 1){
      client.contact = client.contact.slice(0, -1)
    }
  }
  const addInvoice = () => order.invoice = order.invoice.concat(null)
  const removeInvoice = () => {
    if (order.invoice.length > 1){
      order.invoice = order.invoice.slice(0, -1)
    }
  }

  const push = () => {
    if (!check()) return
    client.adress = order.adress
    order.clientId = client.id
    fetch(`${window.location.origin}/orders`, {
      method: "POST",
      body: JSON.stringify({client: client, order: order}),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      order = structuredClone(orderDefault)
      client = structuredClone(clientDefault)
      dispatch('message', false)
    }).catch((err) => {
      alert(err)
    })
  }

  const update = () => {
    client.adress = order.adress
    order.status = "Изменен!"
    fetch(`${window.location.origin}/orders`, {
      method: "PUT",
      body: JSON.stringify({client: client, order: order}),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      alert("Запись обновлена")
      dispatch('message', false)
    }).catch((err) => {
      alert(err)
    })
  }

  const telInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const telFormat = () => {
    for (let i = 0; i < (client.contact).length; ++i){
      (client.contact)[i].tel = ((client.contact)[i].tel).replace(/[^0-9]/g, "")
    }
  }
  const innFormat = () => client.inn = (client.inn).replace(/[^0-9]/g, "")
  const passportFormat = () => {
    client.passportNum = (client.passportNum).replace(/[^0-9]/g, "")
    client.passportSerial = (client.passportSerial).replace(/[^0-9]/g, "")
  }

  const onClose = (e) => {
    dispatch('message', false)
  }

  const searchCity = (e) => {
    fetch(`${window.location.origin}/cities?city=${e.target.value}`).then(function(response) {
    return response.json();
  }).then((data) => {
    if (data.length == 0) throw new Error("storage empty")
    cities = data
  }).catch(() => {
    return
  })
  }

</script>

{#if order.id == null}
  <div id="search">
    <input type="text" placeholder="Поиск...">
  </div>
{:else}
  <div>
    <button class="close" on:click={onClose}>Закрыть</button>
  </div>
{/if}
<div id="container">
  <div id="main">
    <div id="client-name">
      <select bind:value={client.type}>
        <option value="0" selected>ИП</option>
        <option value="1">Юр.Лицо</option>
        <option value="2">Физ.Лицо</option>
      </select>
      {#if client.type != "1"}
        <div>
          <input type="text" bind:value={client.surname} placeholder="Фамилия">
        </div>
        <div>
          <input type="text" bind:value={client.name} placeholder="Имя">
        </div>
        <div>
          <input type="text" bind:value={client.secondName} placeholder="Отчество">
        </div>
      {:else}
        <div>
          <input type="text" id="name" bind:value={client.name} placeholder="Наименование компании">
        </div>    
      {/if}
    </div>
    <div>
      <div class="underline">
        <div class="lables">
          <!--
          <div class="flex">
            <div class="lables">ClientID:</div>
            <div>{client.id}</div>
          </div>
          <div class="flex">
            <div class="lables">OrderID:</div>
            <div>{order.id}</div>
          </div>
          <div class="flex">
            <div class="lables">OrderClientID:</div>
            <div>{order.clientId}</div>
          </div>
          -->
          <div class="flex">
            <div class="lables">Менеджер:</div>
            <div>
              <select bind:value={client.manager}>
                {#each Object.values(managers) as manager}
                  <option value={manager.name}>{manager.name}</option>
                {/each}
              </select>
            </div>
          </div>
          <div class="flex">
            <div class="lables">
              {#if client.type != "2"}
                ИНН: 
              {:else}
                Паспорт: 
              {/if}
            </div>
            <div id="nums">
              {#if client.type != "2"}
                <input type="text" bind:value={client.inn} on:keypress={telInput} on:input={innFormat} placeholder="ИНН">
              {:else}
                <div>
                  <input type="text" bind:value={client.passportSerial} on:keypress={telInput} on:input={passportFormat} placeholder="Серия">
                  <input type="text" bind:value={client.passportNum} on:keypress={telInput} on:input={passportFormat} placeholder="Номер">
                </div>
              {/if}
            </div>
          </div>
          <div class="flex">
            <div class="lables">Контакт:</div>
            <div id="contact">
              {#each client.contact as contact}
                <div>
                  <input type="text" bind:value={contact.name} placeholder="Контакт">
                  <input class="tel" type="text" bind:value={contact.tel} on:keypress={telInput} on:input={telFormat} placeholder="Телефон">
                </div>
              {/each}
              <button class="add-del-button" on:click={addContact}>+</button>
              <button class="add-del-button" on:click={removeContact}>-</button>
              <span class="add-del">добавить/удалить</span>
            </div>
          </div>
          <div class="flex">
            <div class="lables">Email:</div>
            <div>
              <input type="text" bind:value={client.email} placeholder="Email">
            </div>
          </div>
        </div>
      </div>
      <div class="underline">
        <div class="lables">
          <div class="flex">
            <div class="lables">ТК:</div>
            <div>
              <select bind:value={order.cargo}>
                <option value="город">город</option>
                <option value="самовывоз">самовывоз</option>
                <option value="забрать заказ">забрать заказ</option>
                {#each Object.values(cargos) as cargo}
                  <option value={cargo.name}>{cargo.name}</option>
                {/each}
              </select>
            </div>
          </div>
          <div class="flex">
            <div class="lables">За Наш Счет:</div>
            <input type="checkbox" bind:checked={order.payment}>
          </div>
          <div class="flex">
            <div class="lables">Город:</div>
            <div>
              <input list="city-list" type="text" bind:value={order.adress.city} on:input={searchCity} placeholder="Город">
              {#if order.id == null}
                <datalist id="city-list">
                  {#each cities as city}
                    <option value={city.city}>{city.city}</option>
                  {/each}
                </datalist>
              {/if}
            </div>
          </div>
          <div class="flex">
            <div class="lables">До Адреса</div>
            <div>
              <input type="checkbox" bind:checked={order.toadress}>
            </div>
          </div>
          <div class="flex">
            {#if order.toadress}
              <div class="lables">Адрес:</div>
              <div>
                <input type="text" bind:value={order.adress.adress} placeholder="Адрес">
              </div>
            {:else}
              <div class="lables">Терминал:</div>
              <div>
                <input type="text" bind:value={order.adress.terminal} placeholder="Терминал">
              </div>
            {/if}
          </div>
          <div class="flex">
            <div class="lables">Счет:</div>
            <div>
              {#each order.invoice as invoice}
                <div>
                  <input type="text" bind:value={invoice} placeholder="Счет">
                </div>
              {/each}
              <button class="add-del-button" on:click={addInvoice}>+</button>
              <button class="add-del-button" on:click={removeInvoice}>-</button>
              <span class="add-del">добавить/удалить</span>
            </div>
          </div>
          <div class="flex">
            <div class="lables">Крайняя дата</div>
            <div>
              <input type="date" bind:value={order.lastDate}>
            </div>
          </div>
          <div class="flex">
            <div class="lables">Комментарий</div>
            <div>
              <textarea  rows="4" bind:value={order.comment} placeholder="Комментарий"></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  {#if order.id == null}
    <button class="new-order-button" on:click={push}>Отправить</button>
  {:else}
    <button on:click={update}>Изменить</button>
  {/if}
</div>

<style>
  button{
    border: none;
    border-radius: 5px;
    box-shadow: 0px 0px 5px 0px grey;
    background-color: #191970;
    outline: none;
    color: #fdf5df;
  }
  button:active{
    box-shadow: none;
  }
  #search{
    margin-left: 5px;
    margin-top: 10px;
  }
  #container{
    margin: 30px 0px 5px 15px;
  }
  #main{
    padding: 5px;
  }
  .flex{
    display: flex;
    border-bottom: 1px solid grey;
    border-image: linear-gradient(to right, transparent 0%, #778899 50%, transparent 85%) 1;
    padding: 2px 0px;
  }
  .lables{
    flex-basis: 35%;
    font-size: 14px;
  }
  .underline{
    border-bottom: 1px solid #778899;
  }
  .flex div{
    margin-bottom: 2px;
  }
  #contact input, #nums input{
    width: 110px;
  }
  .close{
    position: fixed;
    right: 0;
    top: 0;
    margin: 5px;
  }
  .new-order-button{
    margin: 10px 0px;
  }
  .add-del{
    font-size: 14px;
  }
  .add-del-button{
    width: 20px;
    height: 20px;
  }
</style>