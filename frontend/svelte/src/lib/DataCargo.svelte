<script>
  import { getContext } from 'svelte'

  const uri = getContext('uri')

  let cargo = {
    id: null,
    name: null,
    uri: null,
    mainTel: null,
    managerTel: null
  }

  let cargos =[]

  const get = () => {
    fetch(`${uri}/db/cargos`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then(function(d) {
      d.forEach(e => {
        e.id = e["_id"]
        delete e["_id"]
      });
      cargos = d
    }).catch(function(err) {
      alert(err);
    })
  }

  get()

  const add = () => {
    if (cargo.name == null || cargo.uri == null || cargo.mainTel == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateUrl()){
      alert("Неправильный формат адреса сайта")
      return
    }
    fetch(`${uri}/db/cargos`, {
      method: "POST",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      cargo.id = null
      cargo.name = null
      cargo.mainTel = null
      cargo.managerTel = null
      cargo.uri = null
      alert("Запись Добавлена")
    }).catch((err) => {
      alert(err)
    })
  }

  const update = () => {
    if (cargo.name == null || cargo.uri == null || cargo.mainTel == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateUrl()){
      alert("Неправильный формат адреса сайта")
      return
    }
    fetch(`${uri}/db/cargos`, {
      method: "PUT",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      alert("Запись Обновлена")
    }).catch((err) => {
      alert(err)
    })
  }

  const del = () => {
    if (cargo.id == null){
      alert("Выберите транспортную")
      return
    }
    fetch(`${uri}/db/cargos`, {
      method: "DELETE",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      cargo.id = null
      cargo.name = null
      cargo.mainTel = null
      cargo.managerTel = null
      cargo.uri = null
      alert("Запись удалена")
    }).catch((err) => {
      alert(err)
    })
  }

  const select = () => {
    if (cargo.name == ""){
      cargo.id = null
      cargo.name = null
      cargo.uri = null
      cargo.mainTel = null
      cargo.managerTel = null
      return
    }
    for (let i of cargos){
      if (i.name == cargo.name){
        cargo.id = i.id
        cargo.name = i.name
        cargo.uri = i.uri
        cargo.mainTel = i.mainTel
        cargo.managerTel = i.managerTel
        return
      }
    }
  }

  const telInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const telFormat = (event) => {
    event.preventDefault()
    let paste = (event.clipboardData || window.clipboardData).getData("text")
    paste = paste.replace(/[^0-9]/g, "")
    event.target.value = paste
  }

  const validateUrl = () => {
    return cargo.uri.match(/^(ftp|http|https):\/\/[^ "]+$/)
  }

</script>

<div class="container">
  <div>
    <span>Транспортные</span>
  </div>
  <div>
    <select bind:value={cargo.name} on:change={select}>
      <option value="">Новая</option>
      {#each cargos as cargo}
        <option value={cargo.name}>{cargo.name}</option>
      {/each}
    </select>
  </div>
  <div>
    <span>ID: {cargo.id}</span>
  </div>
  <div>
    <input type="text" bind:value={cargo.name} placeholder="Имя">
  </div>
  <div>
    <input type="text" bind:value={cargo.uri} placeholder="Сайт">
  </div>
  <div>
    <input type="text" bind:value={cargo.mainTel} on:keypress={telInput} on:paste={telFormat} placeholder="Основной телефон">
  </div>
  <div>
    <input type="text" bind:value={cargo.managerTel} on:keypress={telInput} on:paste={telFormat} placeholder="Телефон менеджера">
  </div>
  <div>
    {#if cargo.id == null}
      <button on:click={add}>Добавить</button>
    {:else}
      <button on:click={update}>Изменить</button>
      <button on:click={del}>Удалить</button>
    {/if}
  </div>
</div>

<style>

  .container{
    margin: 5px;
    padding: 5px;
    border: 1px solid black;
    width: 250px;
    height: 250px;
    text-align: center;
  }

</style>