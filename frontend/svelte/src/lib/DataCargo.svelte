<script>

  let cargoDefault = {
    id: null,
    name: null,
    uri: null,
    mainTel: null,
    managerTel: null
  }

  let cargo = structuredClone(cargoDefault)

  $: cargos = []

  const get = () => {
    fetch(`${window.location.origin}/db/cargos`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then((d) => {
      cargos = []
      for (let v of Object.values(d)){
        cargos.push(v)
      }
      cargos = cargos.sort((a, b) => {
        let fa = a.name.toLowerCase()
        let fb = b.name.toLowerCase()
        if (fa > fb) return 1
        if (fa < fb) return -1
        return 0
      })
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
    fetch(`${window.location.origin}/db/cargos`, {
      method: "POST",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      get()
      cargo = structuredClone(cargoDefault)
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
      get()
      cargo = structuredClone(cargoDefault)
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
      get()
      cargo = structuredClone(cargoDefault)
      alert("Запись удалена")
    }).catch((err) => {
      alert(err)
    })
  }

  const select = () => {
    if (cargo.name == ""){
      cargo = structuredClone(cargoDefault)
      return
    }
    for (let i of cargos){
      if (i.name == cargo.name){
        cargo = structuredClone(i)
        return
      }
    }
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
    margin: 5px;
    padding: 5px;
    border: 1px solid black;
    width: 250px;
    height: 250px;
    text-align: center;
  }

</style>