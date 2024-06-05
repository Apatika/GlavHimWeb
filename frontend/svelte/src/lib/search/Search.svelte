<script>
  import Order from "../Order.svelte"
  import OrderFull from "../OrderFull.svelte"

  let orders = {}
  let option = "id"
  let id = ""
  let payment = false
  let limit = 20
  let month = "January"
  let customerList = []

  const find = () => {
    if (option == "id" && id == "") {
      alert("выбери клиента")
      return
    }
    if (option == "last" && (limit <= 0 || limit > 100)) {
      alert("значение должно быть от 0 до 100")
      return
    }
    fetch(`${window.location.origin}/search?id=${id}&payment=${payment}&month=${month}&limit=${limit}`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json()
    }).then((data) => {
      if (Object.keys(data).length == 0) throw new Error()
      orders = data
    }).catch(() => {
      orders = {}
      alert('записей не найдено')
    })
  }

  const sortOrders = (a, b) => {
    return a.status > b.status ? -1 : (a.status < b.status ? 1 : (a.cargo < b.cargo ? -1 : 1))
  }

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

  const change = (e) => {
    if (e.target.value == "payment"){
      payment = true
    } else {
      payment = false
    }
    if (e.target.value == "last"){
      id = ""
    }
  }

</script>

<div class="container">
  <div id="options">
    <select id="option" bind:value={option} on:change={change}>
      <option value="id">По Клиенту</option>
      <option value="payment">За Наш Счет</option>
      <option value="last">Последние</option>
    </select>
    {#if option == "payment"}
      <select id="month" bind:value={month}>
        <option value="01">Январь</option>
        <option value="02">Февраль</option>
        <option value="03">Март</option>
        <option value="04">Апрель</option>
        <option value="05">Май</option>
        <option value="06">Июнь</option>
        <option value="07">Июль</option>
        <option value="08">Август</option>
        <option value="09">Сентябрь</option>
        <option value="10">Октябрь</option>
        <option value="11">Ноябрь</option>
        <option value="12">Декабрь</option>
      </select>
    {/if}
    {#if option == "id"}
      <input id="customer-input" list="list" type="text" bind:value={id} on:input={searchCustomer} placeholder="Поиск...">
      <datalist id="list">
        {#each customerList as cl}
          <option value={cl.id}>{cl.surname} {cl.name} {cl.secondName}</option>
        {/each}
      </datalist>
    {/if}
    {#if option == "last"}
      <input id="limit-input" type="text" bind:value={limit}>
    {/if}
    <button on:click={find}>Показать</button>
  </div>
  <div id="table">
    {#each Object.values(orders).sort(sortOrders) as order}
      <div>
        <div class="order">
          <Order {order} isSearch={true} on:message={(e) => toggleFullOrder(e.detail)}></Order>
        </div>
        <div class="full-order">
          <OrderFull {order} isSearch={true}></OrderFull>
        </div>
      </div>
    {/each}
  </div>
</div>

<style>

.container {
  margin: 10px 30px;
  flex-grow: 1;
}
.container div:not(:last-child) {
  align-items: stretch;
}
#table{
  border-radius: 5px;
  box-shadow: 0px 0px 10px 0px black, 0px 0px 10px 0px black inset;
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
#options{
  margin-bottom: 5px;
}
#customer-input{
  width: 250px;
}
#limit-input{
  width: 25px;
}
select, input{
  border: 1px solid black;
  background-color: #FFF5EE;
  border-radius: 3px;
  color: black;
}
@media (max-width:1364px){
  .full-order{
    box-shadow: none;
    overflow: auto;
  }
}

</style>