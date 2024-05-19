<script>
  import Main from "./lib/main/Main.svelte";
  import Data from "./lib/data/Data.svelte";
  import Route from "./lib/route/Route.svelte";

  let managers = {}
  let cargos = {}
  let chems = {}
  let orders = {}
  let routeIds = []
  let onOpen = true

  window.onload = () => {
    document.querySelector('#main').style.display = 'block'
    getManagers()
    getCargos()
    getChems()
  }

  const getRoute = () => {
    let arr = Object.values(orders).filter(elem => elem.order.status == "В Маршрут")
    arr.forEach(elem => routeIds.push(elem.order.id))
  }

  let socket = new WebSocket("ws://localhost:8081/inwork");
  socket.onmessage = (event) => {
    orders = JSON.parse(event.data)
    if (onOpen){
      getRoute()
      onOpen = false
    }
  }
  socket.onclose = (event) => {
    if (event.wasClean) {
      alert(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
    } else {
      alert('[close] Соединение прервано');
    }
  }
  socket.onerror = (error) => {
    alert(`[error]`)
  }

  const getManagers = () => {
    fetch(`${window.location.origin}/db/users`).then(function(response) {
      return response.json();
    }).then(function(data) {
      if (data.length == 0) throw new Error("storage empty") 
      managers = data
    }).catch(function(err) {
      alert(err);
    })
  }
  
  const getCargos = () => {
    fetch(`${window.location.origin}/db/cargos`).then(function(response) {
      return response.json();
    }).then(function(data) {
      if (data.length == 0) throw new Error("storage empty")
      cargos = data
    }).catch(function(err) {
      alert(err);
    })
  }

  const getChems = () => {
    fetch(`${window.location.origin}/db/chems`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then((data) => {
      chems = data
    }).catch(function(err) {
      alert(err);
    })
  }

  const change = (e) => {
    document.querySelectorAll('.page').forEach((elem) => {
      elem.style.display = 'none'
    })
    let name = e.target.name
    document.querySelector(`#${name}`).style.display = 'block'
  }

  const update = (text) => {
    switch (text) {
      case 'managers':
        getManagers()
        break
      case 'cargos':
        getCargos()
        break
      case 'chems':
        getChems()
        break
      default:
        alert('update case not alowed')
    }
  }

  const removeRoute = (id) => {
    const index = routeIds.indexOf(id);
    if (index > -1) {
      routeIds = routeIds.filter(elem => elem != id)
    }
  }

</script>

<div class="nav">
  <button name='main' on:click={change}>Главная</button>
  <button name='data' on:click={change}>База Данных</button>
  <button name='route' on:click={change}>Маршрут</button>
</div>
<div class="content">
  <div class="page" id="main">
    <Main {orders} {managers} {cargos} {chems} on:message={(e) => removeRoute(e.detail.routeId)}></Main>
  </div>
  <div class="page" id="data">
    <Data {managers} {cargos} {chems} on:message={(e) => update(e.detail.text)}></Data>
  </div>
  <div class="page" id="route">
    <Route {orders} ids={routeIds}></Route>
  </div>
</div>

<style>
  button{
    display: block;
    position: relative;
    border: none;
    margin-left: 2px;
    margin-top: 1px;
    font-size: 18px;
    height: 40px;
    border-top: 1px solid var(--background-main);
    border-radius: 5px;
    background-color: var(--border-main);
    box-shadow: 0px 0px 10px 2px black;
  }
  button:hover{
    box-shadow: 0px 0px 10px 2px black, 0px 0px 2px 0px white inset;
  }
  button:active{
    box-shadow: 0px 0px 10px 0px black inset;
  }
  .page{
    display: none;
  }
  .nav{
    display: flex;
    flex-direction: column;
    background-color: var(--background-main);
    width: 150px;
  }
  .content{
    background-color: var(--background-main);
    flex-grow: 1;
  }
  @media (min-width:960px){
    button {
      font-size: 14px;
    }
  }
</style>
