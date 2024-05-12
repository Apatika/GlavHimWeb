<script>
  import Main from "./lib/Main.svelte";
  import Data from "./lib/Data.svelte";

  let managers = []
  let cargos = []
  let chems = []

  window.onload = () => {
    document.querySelector('#main').style.display = 'block'
    fetch(`${window.location.origin}/db/users`).then(function(response) {
      return response.json();
    }).then(function(data) {
      if (data.length == 0) throw new Error("storage empty") 
      managers = []
      for (let v of Object.values(data)){
        managers.push(v)
      }
      managers = managers.sort((a, b) => {
        if (a.name > b.name) return 1
        if (a.name < b.name) return -1
        return 0
      })
    }).catch(function(err) {
      alert(err);
    })

    fetch(`${window.location.origin}/db/cargos`).then(function(response) {
      return response.json();
    }).then(function(data) {
      if (data.length == 0) throw new Error("storage empty")
      cargos = []
      for (let v of Object.values(data)){
        cargos.push(v)
      }
      cargos = cargos.sort((a, b) => {
        if (a.name > b.name) return 1
        if (a.name < b.name) return -1
        return 0
      })
    }).catch(function(err) {
      alert(err);
    })

    fetch(`${window.location.origin}/db/chems`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then((d) => {
      chems = []
      for (let v of Object.values(d)){
        chems.push(v)
      }
      chems = chems.sort((a, b) => {
        if (a.name > b.name) return 1
        if (a.name < b.name) return -1
        return 0
      })
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
 

</script>

<div class="nav">
  <button name='main' on:click={change}>Главная</button>
  <button name='data' on:click={change}>База Данных</button>
</div>
<div class="content">
  <div class="page" id="main">
    <Main {managers} {cargos}></Main>
  </div>
  <div class="page" id="data">
    <Data {managers} {cargos} {chems}></Data>
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
    border-top: 1px solid #fdf5df;
    border-radius: 5px;
    background-color: #191970;
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
    background-color: #fdf5df;
    width: 150px;
  }
  .content{
    background-color: #fdf5df;
    flex-grow: 1;
  }
</style>
