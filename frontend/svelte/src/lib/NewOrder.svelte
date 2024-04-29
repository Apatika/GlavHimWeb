<script>
  import { getContext } from 'svelte'

  const uri = getContext('uri')

  let client = {
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
    tel: null,
    email: null,
  }
  let order = {
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
    pvd: []
  }

  let managers = []
  let cargos = []

  fetch(`${uri}/db/users`).then(function(response) {
    return response.json();
  }).then(function(data) {
    managers = data
  }).catch(function(err) {
    alert(err);
  })

  fetch(`${uri}/db/cargos`).then(function(response) {
    return response.json();
  }).then(function(data) {
    cargos = data
  }).catch(function(err) {
    alert(err);
  })

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
    fetch(`${uri}/orders`, {
      method: "POST",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then((data) => {
      if (data.code != 200) throw new Error(data.message)
      order.id = null
      order.clientId = null
      order.payment = null
      order.adress = {city: null, adress: null, terminal: "основной"}
      order.invoice = [null]
      order.cargo = null
      order.lastDate = null
      order.comment = null
      order.probes = []
      order.pvd = []
      order.toadress = false
    }).catch((err) => {
      alert(err)
    })
  }

  const push = () => {
    // @ts-ignore
    client.type = Number(client.type)
    client.adress = order.adress
    fetch(`${uri}/clients`, {
      method: "POST",
      body: JSON.stringify(client),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then((data) => {
      if (data.code != 200) throw new Error(data.message)
      client.id = null
      client.type = '0'
      client.adress = null
      client.name = null
      client.surname = null
      client.secondName = null
      client.manager = null
      client.inn = null
      client.passportSerial = null
      client.passportNum = null
      client.contact = [{name: null, tel: null}]
      client.tel = null
      client.email = null
      pushOrder(data.id)
    }).catch((err) => {
      alert(err)
    })
  }

</script>

<div id="search">
  <input type="text" placeholder="Поиск...">
</div>
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
          <div class="flex">
            <div class="lables">ID:</div>
            <div>{client.id}</div>
          </div>
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
                  <input type="text" bind:value={contact.tel} placeholder="Телефон">
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
  <button on:click={push}>Отправить</button>
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
</style>