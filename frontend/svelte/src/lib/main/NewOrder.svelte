<script>
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()

  export let order = {
    id: null,
    customerId: null,
    payment: null,
    catalog: null,
    catalogCount: null,
    toadress: true,
    adress: {city: "Санкт-Петербург", adress: null, terminal: "основной"},
    invoice: [null],
    cargo: null,
    lastDate: null,
    comment: null,
    content: null,
    probes: {},
    customer: {
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
  }
  const setInitial = () => {
    Object.keys(order).forEach(o => {if (o != "customer") order[o] = null})
    Object.keys(order.customer).forEach(c => order.customer[c] = null)
    order.customer.contact = [{name: null, tel: null}]
    order.customer.type = '0'
    order.cargo = "город"
    order.toadress = false
    order.adress = {city: "Санкт-Петербург", adress: null, terminal: "основной"}
    order.invoice = [null]
    order.probes = {}
  }

  let cities = []
  let customerList = []
  let probeCountSum = 0
  export let chems = {}
  export let managers = {}
  export let cargos = {}
  $: {
    probeCountSum = 0
    Object.keys(chems).forEach(key => probeCountSum += chems[key].probeCount)
  }

  const check = () => {
    if (order.customer.manager == null || order.customer.manager == ""){
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
    if (order.customer.name == null || order.customer.name == ""){
      alert("Имя не заполнено")
      return false
    }
    if (order.customer.contact[0].tel == null || order.customer.contact[0].tel == ""){
      alert("Контакт не заполнен")
      return false
    }
    if ((order.customer.type == "0" || order.customer.type == "2") && 
        (order.customer.surname == null || order.customer.surname == "")){
      alert("ФИО заполнены не полностью")
      return false
    }
    if ((order.customer.type == "1" || order.customer.type == "0") && (order.customer.inn == null || order.customer.inn == "")){
      alert("ИНН не заполнен")
      return false
    }
    if (order.cargo == "город" || order.cargo == "забрать заказ" || order.cargo == "самовывоз") {
      order.payment = false
    }
    return true
  }

  const addContact = () => order.customer.contact = order.customer.contact.concat({name: "", tel: ""})
  const removeContact = () => {
    if (order.customer.contact.length > 1){
      order.customer.contact = order.customer.contact.slice(0, -1)
    }
  }

  const push = () => {
    if (!check()) return
    order.customer.adress = order.adress
    order.customerId = order.customer.id
    for (let [k, v] of Object.entries(chems)){
      if (chems[k].probeCount != 0){
        order.probes[k] = v
      }
    }
    fetch(`${window.location.origin}/orders`, {
      method: "POST",
      body: JSON.stringify(order),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      setInitial()
      dispatch('message', false)
    }).catch((err) => {
      alert(err)
    })
    for (let k of Object.keys(chems)){
      chems[k].probeCount = 0
    }
  }

  const update = () => {
    if (!check()) return
    order.customer.adress = order.adress
    Object.keys(order.probes).forEach(key => {
      delete order.probes[key]
    })
    for (let [k, v] of Object.entries(chems)){
      if (chems[k].probeCount != 0){
        order.probes[k] = v
      }
    }
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
      dispatch('message', {
        id: order.id,
        update: true
      })
    }).catch((err) => {
      alert(err)
    })
  }

  const telInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const telFormat = () => {
    for (let i = 0; i < (order.customer.contact).length; ++i){
      (order.customer.contact)[i].tel = ((order.customer.contact)[i].tel).replace(/[^0-9]/g, "")
    }
  }
  const innFormat = () => order.customer.inn = (order.customer.inn).replace(/[^0-9]/g, "")
  const passportFormat = () => {
    try{
      order.customer.passportNum = (order.customer.passportNum).replace(/[^0-9]/g, "")
      order.customer.passportSerial = (order.customer.passportSerial).replace(/[^0-9]/g, "")
    }
    catch(err){}
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

  const searchCustomer = (e) => {
    fetch(`${window.location.origin}/customers?customer=${e.target.value}`).then(function(response) {
      return response.json();
    }).then((data) => {
      if (data.length == 0) throw new Error("storage empty")
      customerList = data
    }).catch(() => {
      return
    })
  }

  const getCustomer = (e) => {
    for (let v of customerList){
      if (e.target.value == v.id){
        order.customer = v
        order.adress = v.adress
        return
      }
    }
  }

  const probesShow = (e) => {
    let target = e.target.previousElementSibling
    target.style.height = '500px'
    target.style.padding = '10px'
  }

  const probesClose = (e) => {
    let target = e.target.closest('.probes-container')
    target.style.height = '0px'
    target.style.padding = '0px'
  }

  const cargoChange = (e) => {
    let value = e.target.value
    if (value == "город" || value == "забрать заказ" || value == "самовывоз") {
      order.toadress = true
      order.adress.city = "Санкт-Петербург"
    }
    else {
      order.toadress = false
      order.adress.city = ""
    }
    if (value == "самовывоз") {
      order.adress.adress = "Склад"
    } else {
      order.adress.adress = ""
    }
  }

  const loadFile = (e) => {
    const files = e.target.files
    if (files.length != 0) {
      order.content = ""
      order.invoice = []
    }
    for (let file of files){
      const num = /\d+/.exec(file.name)[0]
      order.invoice.push(num)
      const reader = new FileReader()
      reader.readAsText(file)
      reader.onload = () => {
        let result = new DOMParser().parseFromString(reader.result, "text/html")
        let tables = result.querySelectorAll('table')
        order.content += tables[3].outerHTML + tables[4].outerHTML
      }
    }
  }

</script>

<div class="container">
  {#if order.id == null}
    <div class="search-field">
      <input list="customer-list" type="text" on:input={searchCustomer} on:change={getCustomer} placeholder="Поиск...">
      {#if order.id == null}
        <datalist id="customer-list">
          {#each customerList as cl}
            <option value={cl.id}>{cl.surname} {cl.name} {cl.secondName}</option>
          {/each}
        </datalist>
      {/if}
    </div>
  {:else}
    <div>
      <button class="close-button" on:click={() => dispatch('message', {id: order.id, update: false})}>Закрыть</button>
    </div>
  {/if}
  <div class="main">
    <div class="customer-name">
      <select bind:value={order.customer.type}>
        <option value="0" selected>ИП</option>
        <option value="1">Юр.Лицо</option>
        <option value="2">Физ.Лицо</option>
      </select>
      {#if order.customer.type != "1"}
        <div>
          <input type="text" bind:value={order.customer.surname} placeholder="Фамилия">
        </div>
        <div>
          <input type="text" bind:value={order.customer.name} placeholder="Имя">
        </div>
        <div>
          <input type="text" bind:value={order.customer.secondName} placeholder="Отчество">
        </div>
      {:else}
        <div>
          <input class="name" type="text" bind:value={order.customer.name} placeholder="Наименование компании">
        </div>    
      {/if}
    </div>
    <div>
      <div class="underline">
        <div class="lable">
          <div class="flex">
            <div class="lable">Менеджер:</div>
            <div>
              <select bind:value={order.customer.manager}>
                {#each Object.values(managers) as manager}
                  <option value={manager.name}>{manager.name}</option>
                {/each}
              </select>
            </div>
          </div>
          <div class="flex">
            <div class="lable">
              {#if order.customer.type != "2"}
                ИНН: 
              {:else}
                Паспорт: 
              {/if}
            </div>
            <div class="nums">
              {#if order.customer.type != "2"}
                <input class="inn" type="text" bind:value={order.customer.inn} on:keypress={telInput} on:input={innFormat} placeholder="ИНН">
              {:else}
                <div>
                  <input type="text" bind:value={order.customer.passportSerial} on:keypress={telInput} on:input={passportFormat} placeholder="Серия">
                  <input type="text" bind:value={order.customer.passportNum} on:keypress={telInput} on:input={passportFormat} placeholder="Номер">
                </div>
              {/if}
            </div>
          </div>
          <div class="flex">
            <div class="lable">Контакт:</div>
            <div class="contact">
              {#each order.customer.contact as contact}
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
            <div class="lable">Email:</div>
            <div>
              <input type="text" bind:value={order.customer.email} placeholder="Email">
            </div>
          </div>
        </div>
      </div>
      <div class="underline">
        <div class="lable">
          <div class="flex">
            <div class="lable">ТК:</div>
            <div>
              <select bind:value={order.cargo} on:change={cargoChange}>
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
            <div class="lable">За Наш Счет:</div>
            <input type="checkbox" bind:checked={order.payment}>
          </div>
          <div class="flex">
            <div class="lable">Город:</div>
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
            <div class="lable">До Адреса</div>
            <div>
              <input type="checkbox" bind:checked={order.toadress}>
            </div>
          </div>
          <div class="flex">
            {#if order.toadress}
              <div class="lable">Адрес:</div>
              <div>
                <input type="text" bind:value={order.adress.adress} placeholder="Адрес">
              </div>
            {:else}
              <div class="lable">Терминал:</div>
              <div>
                <input type="text" bind:value={order.adress.terminal} placeholder="Терминал">
              </div>
            {/if}
          </div>
          <div class="flex">
            <div class="lable">Cчет:</div>
            <div>
              <input type="file" on:change={loadFile} size="60" accept="text/html" multiple>
            </div>
          </div>
          <div class="flex">
            <div class="lable">Крайняя дата:</div>
            <div>
              <input type="date" bind:value={order.lastDate}>
            </div>
          </div>
          <div class="flex">
            <div class="lable">Каталог:</div>
            <div>
              <input type="checkbox" bind:checked={order.catalog}>
              {#if order.catalog}<input class="catalog-count" type="text" bind:value={order.catalogCount} placeholder="кол-во">{/if}
            </div>
          </div>
          <div class="flex">
            <div class="lable">Комментарий:</div>
            <div>
              <textarea  rows="4" bind:value={order.comment} placeholder="Комментарий"></textarea>
            </div>
          </div>
          <div>
            <div class="probes-container">
              <button on:click={probesClose}>закрыть</button>
              {#each Object.values(chems).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as chem}
                <div class="probes">
                  <div class="probes-name">
                    <span>{chem.name}</span>
                  </div>
                  <div class="probes-count">
                    <button on:click={() => {if (chem.probeCount > 0) {chem.probeCount--}}}>-</button>
                    <span class="probes-count-span">{(chem.probeValue * chem.probeCount) / 1000} л</span>
                    <button on:click={() => {chem.probeCount++}}>+</button>
                  </div>
                </div>
              {/each}
            </div>
            <button on:click={probesShow}>ПРОБНИКИ</button>
            {#if probeCountSum != 0} <strong style="color: red;">Добавлены пробники</strong> {probeCountSum}шт.{/if}
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
  input[type="text"], textarea{
    background-color: rgba(0, 0, 0, 0.1);
    border: none;
    border-radius: 3px;
    border-bottom: 1px solid red;
    outline: none;
    color: rgb(229, 245, 15);
    font-size: 14px;
  }
  input::placeholder, textarea::placeholder{
    color: lightgray;
  }
  input[type="file"]{
    width: 200px;
  }
  select{
    color: rgb(229, 245, 15);
    border-bottom: 1px solid red;
  }
  .container{
    margin: 30px 0px 5px 15px;
    color: white;
  }
  .search-field{
    margin-left: 5px;
    margin-top: 10px;
  }
  .search-field input{
    width: 250px;
  }
  .main{
    padding: 5px;
  }
  .flex{
    display: flex;
    border-bottom: 1px solid grey;
    border-image: linear-gradient(to right, transparent 0%, #778899 50%, transparent 85%) 1;
    padding: 2px 0px;
  }
  .flex div{
    margin-bottom: 2px;
  }
  .lable{
    flex-basis: 35%;
    font-size: 14px;
  }
  .underline{
    border-bottom: 1px solid #778899;
  }
  .contact input, .nums input{
    width: 100px;
  }
  .nums .inn {
    width: 110px;
  }
  .close-button{
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
  .probes{
    display: flex;
  }
  .probes-name{
    width: 150px;
  }
  .probes-count{
    width: 100px;
  }
  .probes-count-span{
    display: inline-block;
    width: 50px;
    margin: 0px 3px;
    text-align: center;
  }
  .probes-container{
    position: absolute;
    background-color: #FFF5EE;
    border: 1px solid black;
    border-right: none;
    border-radius: 5px;
    box-shadow: 0px 0px 10px 0px black;
    top: 0;
    right: 0;
    width: 300;
    height: 0px;
    overflow: auto;
    text-align: center;
    transition: all .2s linear;
  }
  .catalog-count{
    width: 50px;
    text-align: center;
  }
  @media (max-width:1364px){
    .probes{
      margin-bottom: 2px;
    }
    .probes-name{
      font-size: 9px;
      width: 100px;
    }
    .probes-count{
      font-size: 9px;
      width: 100px;
    }
    .probes-count-span{
      width: 30px;
    }
  }

</style>