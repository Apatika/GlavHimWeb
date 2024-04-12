<script>

  let cargo = {
    id: null,
    name: null,
    uri: null,
    mainTel: null,
    managerTel: null
  }

  let cargos =[]

  let get = () => {
    fetch('http://localhost:8081/cargos').then(function(response) {
      return response.json();
    }).then(function(d) {
      cargos = d
    }).catch(function(err) {
      alert(err);
    })
  }

  get()

  let add = () => {
    if (cargo.name == null || cargo.tel == null || cargo.email == null){
      alert("Заполните пустые поля")
      return
    }
    fetch("http://localhost:8081/cargos", {
      method: "POST",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(() => {
      get()
      alert("Запись Добавлена")
    }).catch((err) => {
      alert(err)
    })
    cargo.id = null
    cargo.name = null
    cargo.tel = null
    cargo.email = null
  }

  let update = () => {
    if (cargo.name == null || cargo.tel == null || cargo.email == null){
      alert("Заполните пустые поля")
      return
    }
    fetch("http://localhost:8081/cargos", {
      method: "PUT",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(() => {
      get()
      alert("Запись Обновлена")
    }).catch((err) => {
      alert(err)
    })
  }

  let del = () => {
    if (cargo.id == null){
      alert("Выберите транспортную")
      return
    }
    fetch("http://localhost:8081/cargos", {
      method: "DELETE",
      body: JSON.stringify(cargo),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(() => {
      get()
      alert("Запись удалена")
    }).catch((err) => {
      alert(err)
    })
    cargo.id = null
    cargo.name = null
    cargo.tel = null
    cargo.email = null
  }

  let select = () => {
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
    <input type="text" bind:value={cargo.mainTel} placeholder="Основной телефон">
  </div>
  <div>
    <input type="text" bind:value={cargo.managerTel} placeholder="Телефон менеджера">
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
  }

</style>