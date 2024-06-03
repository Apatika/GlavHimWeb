<script>
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  let cargo = {
    id: null,
    name: null,
    uri: null,
    mainTel: null,
    managerTel: null
  }

  export let cargos = {}

  const add = () => {
    if (cargo.name == null || cargo.uri == null || cargo.mainTel == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateUrl()){
      alert("Неправильный формат адреса сайта")
      return
    }
    fetch(`${window.location.origin}/db/cargos`, {
      method: "POST",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('message', {text: 'cargos'})
      Object.keys(cargo).forEach(c => cargo[c] = null)
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
    fetch(`${window.location.origin}/db/cargos`, {
      method: "PUT",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('message', {text: 'cargos'})
      Object.keys(cargo).forEach(c => cargo[c] = null)
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
    fetch(`${window.location.origin}/db/cargos`, {
      method: "DELETE",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('message', {text: 'cargos'})
      Object.keys(cargo).forEach(c => cargo[c] = null)
      alert("Запись удалена")
    }).catch((err) => {
      alert(err)
    })
  }

  const select = (id) => {
    cargo.id = null
    cargo.name = null
    cargo.uri = null
    cargo.mainTel = null
    cargo.managerTel = null
    if (id == ""){
      return
    }
    cargo = cargos[id]
  }

  const telInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const telFormat = () => {
    cargo.mainTel = cargo.mainTel.replace(/[^0-9]/g, "")
    cargo.managerTel = cargo.managerTel.replace(/[^0-9]/g, "")
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
    <select value={cargo.id} on:change={(e) => select(e.target.value)}>
      <option value="">Новая</option>
      {#each Object.values(cargos).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as cargo}
        <option value={cargo.id}>{cargo.name}</option>
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
    <input type="text" bind:value={cargo.mainTel} on:keypress={telInput} on:input={telFormat} placeholder="Основной телефон">
  </div>
  <div>
    <input type="text" bind:value={cargo.managerTel} on:keypress={telInput} on:input={telFormat} placeholder="Телефон менеджера">
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
    display: flex;
    flex-direction: column;
    justify-content: center;
    margin: 5px;
    border: 1px solid black;
    width: 250px;
    height: 250px;
    text-align: center;
    border-radius: 50%;
    box-shadow: 0px 0px 10px 0px grey;
  }

</style>