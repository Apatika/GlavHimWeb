<script>

  let data = {
    client: {
      id: null,
      type: 0,
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
    },
    order: {
      id: null,
      clientId: null,
      payment: null,
      adress: {city: null, adress: null, terminal: "основной"},
      invoice: [null],
      cargo: null,
      lastDate: null,
      comment: null,
      probes: [],
      pvd: []
    }
  }

  let managers = []
  let toAdress = false

  fetch('http://localhost:8081/users').then(function(response) {
    return response.json();
  }).then(function(d) {
    managers = d
  }).catch(function(err) {
    alert(err);
  })

  const addContact = () => data.client.contact = data.client.contact.concat({name: "", tel: ""})
  const removeContact = () => {
    if (data.client.contact.length > 1){
      data.client.contact = data.client.contact.slice(0, -1)
    }
  }
  const addInvoice = () => data.order.invoice = data.order.invoice.concat(null)
  const removeInvoice = () => {
    if (data.order.invoice.length > 1){
      data.order.invoice = data.order.invoice.slice(0, -1)
    }
  }

</script>

<div id="search">
  <input type="text" placeholder="Поиск...">
</div>
<div id="container">
  <div id="main">
    <div id="client-name">
      <select bind:value={data.client.type}>
        <option value="0">ИП</option>
        <option value="1">Юр.Лицо</option>
        <option value="2">Физ.Лицо</option>
      </select>
      {#if data.client.type != 1}
        <input type="text" bind:value={data.client.surname} placeholder="Фамилия">
        <input type="text" bind:value={data.client.name} placeholder="Имя">
        <input type="text" bind:value={data.client.secondName} placeholder="Отчество">
      {:else}
        <input type="text" id="name" bind:value={data.client.name} placeholder="Наименование компании">
      {/if}
    </div>
    <div>
      <div class="underline">
        <div class="lables">
          <div class="flex">
            <div class="lables">ID:</div>
            <div>{data.client.id}</div>
          </div>
          <div class="flex">
            <div class="lables">Менеджер:</div>
            <div>
              <select bind:value={data.client.manager}>
                {#each managers as manager}
                  <option value={manager.name}>{manager.name}</option>
                {/each}
              </select>
            </div>
          </div>
          <div class="flex">
            <div class="lables">
              {#if data.client.type != 2}
                ИНН: 
              {:else}
                Паспорт: 
              {/if}
            </div>
            <div>
              {#if data.client.type != 2}
                <input type="text" bind:value={data.client.inn} placeholder="ИНН">
              {:else}
                <input type="text" bind:value={data.client.passportSerial} placeholder="Серия">
                <input type="text" bind:value={data.client.passportNum} placeholder="Номер">
              {/if}
            </div>
          </div>
          <div class="flex">
            <div class="lables">Контакт:</div>
            <div>
              {#each data.client.contact as contact}
                <div>
                  <input type="text" bind:value={contact.name} placeholder="Контакт">
                  <input type="text" bind:value={contact.tel} placeholder="Телефон">
                </div>
              {/each}
              <button on:click={addContact}>+</button>
              <button on:click={removeContact}>-</button>
              <span>добавить/удалить контакт</span>
            </div>
          </div>
          <div class="flex">
            <div class="lables">Email:</div>
            <div>
              <input type="text" bind:value={data.client.email} placeholder="Email">
            </div>
          </div>
        </div>
      </div>
      <div class="underline">
        <div class="lables">
          <div class="flex">
            <div class="lables">ТК:</div>
            <div>
              <select bind:value={data.order.cargo}></select>
            </div>
          </div>
          <div class="flex">
            <div class="lables">За Наш Счет:</div>
            <input type="checkbox" bind:checked={data.order.payment}>
          </div>
          <div class="flex">
            <div class="lables">Город:</div>
            <input type="text" bind:value={data.order.adress.city} placeholder="Город">
          </div>
          <div class="flex">
            {#if toAdress}
              <div class="lables">Адрес:</div>
              <input type="text" bind:value={data.order.adress.adress} placeholder="Адрес">
            {:else}
              <div class="lables">Терминал:</div>
              <input type="text" bind:value={data.order.adress.terminal} placeholder="Терминал">
            {/if}
            <label><input type="checkbox" bind:checked={toAdress}> До Адреса</label>
          </div>
          <div class="flex">
            <div class="lables">Счет:</div>
            <div>
              {#each data.order.invoice as invoice}
                <div>
                  <input type="text" bind:value={invoice} placeholder="Счет">
                </div>
              {/each}
              <button on:click={addInvoice}>+</button>
              <button on:click={removeInvoice}>-</button>
              <span>добавить/удалить счет</span>
            </div>
          </div>
          <div class="flex">
            <div class="lables">Крайняя дата</div>
            <div>
              <input type="date" bind:value={data.order.lastDate}>
            </div>
          </div>
          <div class="flex">
            <div class="lables">Комментарий</div>
            <div>
              <textarea  rows="4" bind:value={data.order.comment} placeholder="Комментарий"></textarea>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  #search{
    margin-left: 5px;
    margin-top: 10px;
  }
  #container{
    display:flex;
    margin: 30px 0px 5px 50px;
  }
  #main{
    padding: 5px;
    border-right: 1px solid blue;
  }
  .flex{
    display: flex;
  }
  .lables{
    flex-basis: 30%;
  }
  .underline{
    border-bottom: 1px solid blue;
  }
  .flex div{
    margin-bottom: 2px;
  }
</style>