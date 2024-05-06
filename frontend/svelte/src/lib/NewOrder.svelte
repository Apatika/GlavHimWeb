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

  let managers = []
  let cargos = []

  fetch(`${window.location.origin}/db/users`).then(function(response) {
    return response.json();
  }).then(function(data) {
    if (data.length == 0) throw new Error("storage empty") 
    managers = data
  }).catch(function(err) {
    alert(err);
  })

  fetch(`${window.location.origin}/db/cargos`).then(function(response) {
    return response.json();
  }).then(function(data) {
    if (data.length == 0) throw new Error("storage empty")
    cargos = data
  }).catch(function(err) {
    alert(err);
  })

  const check = () => {
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

  const pushOrder = (id) => {
    order.clientId = id
    fetch(`${window.location.origin}/orders`, {
      method: "POST",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      order = structuredClone(orderDefault)
    }).catch((err) => {
      alert(err)
    })
  }

  const push = () => {
    // @ts-ignore
    if (!check()) return
    client.adress = order.adress
    fetch(`${window.location.origin}/clients`, {
      method: "POST",
      body: JSON.stringify(client),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      else return response.json()
    }).then(data => {
      client = structuredClone(clientDefault)
      pushOrder(data.id)
    }).catch((err) => {
      alert(err)
    })
  }

  const update = () => {
    client.adress = order.adress
    fetch(`${window.location.origin}/clients`, {
      method: "PUT",
      body: JSON.stringify(client),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
    }).catch((err) => {
      alert(err)
    })
    order.status = "Изменен!"
    fetch(`${window.location.origin}/orders`, {
      method: "PUT",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      alert("Запись обновлена")
      let closer = document.querySelectorAll('.close').forEach((elem) => {
        elem.parentElement.parentElement.style.display = "none"
      })
      document.body.style.pointerEvents = "all"
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

  const onClose = (e) => {
    e.target.parentElement.parentElement.style.display = "none"
    document.body.style.pointerEvents = "all"
    dispatch('message', false)
  }

</script>

{#if order.id == null}
  <div id="search">
    <input type="text" placeholder="Поиск...">
  </div>
{:else}
  <div>
    <button class="close" on:click={onClose}>закрыть</button>
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
                {#each managers as manager}
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
                <input type="text" bind:value={client.inn} placeholder="ИНН">
              {:else}
                <div>
                  <input type="text" bind:value={client.passportSerial} placeholder="Серия">
                  <input type="text" bind:value={client.passportNum} placeholder="Номер">
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
              <button on:click={addContact}>+</button>
              <button on:click={removeContact}>-</button>
              <span>добавить/удалить</span>
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
                {#each cargos as cargo}
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
              <input type="text" bind:value={order.adress.city} placeholder="Город">
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
              <button on:click={addInvoice}>+</button>
              <button on:click={removeInvoice}>-</button>
              <span>добавить/удалить</span>
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
    <button on:click={push}>Отправить</button>
  {:else}
    <button on:click={update}>Изменить</button>
  {/if}
</div>

<style>
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
    border-image: linear-gradient(to right, white 0%, blue 50%, white 100%) 1;
    padding: 2px 0px;
  }
  .lables{
    flex-basis: 35%;
  }
  .underline{
    border-bottom: 1px solid blue;
  }
  .flex div{
    margin-bottom: 2px;
  }
  #contact input, #nums input{
    width: 90px;
  }
  .close{
    position: fixed;
    right: 0;
    top: 0;
    margin: 5px;
  }
</style>