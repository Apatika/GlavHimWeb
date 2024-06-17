<script>
  import Main from "./lib/main/Main.svelte";
  import Data from "./lib/data/Data.svelte";
  import Route from "./lib/route/Route.svelte";
  import Search from "./lib/search/Search.svelte";
  import Pvd from "./lib/pvd/Pvd.svelte";

  let managers = {}
  let cargos = {}
  let chems = {}
  let orders = {}
  let pvd = []

  window.onload = () => {
    document.querySelector('#main').style.display = 'block'
    getManagers()
    getCargos()
    getChems()
    getPvd()
  }

  let socket = new WebSocket(`ws://${window.location.host}/inwork`);
  socket.onmessage = (event) => {
    orders = JSON.parse(event.data)
  }
  socket.onclose = (event) => {
    if (event.wasClean) {
      alert(`[close] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`);
    } else {
      alert('[close] Соединение прервано');
    }
  }
  socket.onerror = (error) => {
  }

  const getPvd = () => {
    fetch(`${window.location.origin}/pvd`).then(function(response) {
      return response.json();
    }).then(function(data) {
      if (data.length == 0) throw new Error("storage empty")
      pvd = data
    }).catch(function(err) {
      alert(err)
    })
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
    if (window.screen.width < 1366) document.querySelector('#menu-button').style.display = 'inline-block'
    document.querySelector('.mobile-nav').style.display = "none"
    document.querySelector('.content').style.display = "block"
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

  const mobileMenu = () => {
    if (window.screen.width < 1366) document.querySelector('#menu-button').style.display = 'none'
    document.querySelector('.content').style.display = 'none'
    document.querySelector('.mobile-nav').style.display = 'block'
  }

</script>

<div class="nav">
  <button name='main' on:click={change}>Главная</button>
  <button name='data' on:click={change}>База Данных</button>
  <button name='route' on:click={change}>Маршрут</button>
  <button name='search' on:click={change}>Поиск Заказов</button>
  <button name='pvd' on:click={change}>Пленка</button>
</div>
<div class="mobile-nav">
  <button name='main' on:click={change}>Главная</button>
  <button name='route' on:click={change}>Маршрут</button>
</div>
<div class="content">
  <div class="page" id="main">
    <Main {orders} {managers} {cargos} {chems}></Main>
  </div>
  <div class="page" id="data">
    <Data {managers} {cargos} {chems} on:message={(e) => update(e.detail.text)}></Data>
  </div>
  <div class="page" id="route">
    <Route {orders}></Route>
  </div>
  <div class="page" id="search">
    <Search></Search>
  </div>
  <div class="page" id="pvd">
    <Pvd {pvd}></Pvd>
  </div>
</div>
<button id="menu-button" on:click={mobileMenu}>меню</button>

<style>
  button{
    display: block;
    position: relative;
    border: none;
    margin-left: 2px;
    margin-top: 1px;
    font-size: 18px;
    height: 40px;
    background-color: rgba(0, 0, 255, 0.2);
  }
  button:hover{
    background-color: rgba(0, 0, 255, 0.3);
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
    width: 150px;
    background-color: rgba(0, 0, 0, 0.3);
  }
  .mobile-nav{
    display: none;
    width: 100%;
    height: 100vh;
    background-color: var(--border-main);
  }
  .mobile-nav button{
    display: block;
    width: 100%;
    border-top: none;
    border-bottom: 1px solid white;
    box-shadow: none;
    border-radius: 0;
  }
  .content{
    flex-grow: 1;
  }
  #menu-button{
    position: absolute;
    display: none;
    top: 5px;
    right: 5px;
    border-radius: 50%;
    font-size: 7px;
    border: none;
    height: 30px;
    width: 30px;
    text-align: center;
    padding: 0;
  }
  @media (min-width:960px){
    button {
      font-size: 14px;
    }
  }
  @media (max-width:1364px){
    .nav{
      display: none;
    }
    #menu-button{
      display: inline-block;
    }
    .content{
      flex-grow: 0;
      width: 100%;
    }
  }
</style>
