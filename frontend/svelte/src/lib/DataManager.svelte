<script>

  let userDefault = {
    id: null,
    name: null,
    tel: null,
    email: null
  }

  let user = structuredClone(userDefault)
  
  export let managers = []

  const get = () => {
    fetch(`${window.location.origin}/db/users`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then((d) => {
      managers = []
      for (let v of Object.values(d)){
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
  }
  
  const add = () => {
    if (user.name == null || user.tel == null || user.email == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateEmail()){
      alert("Неправильный формат почты")
      return
    }
    fetch(`${window.location.origin}/db/users`, {
      method: "POST",
      body: JSON.stringify(user),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      get()
      user = structuredClone(userDefault)
      alert("Менедежер Добавлен")
    }).catch((err) => {
      alert(err)
    })
  }

  const update = () => {
    if (user.name == null || user.tel == null || user.email == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateEmail()){
      alert("Неправильный формат почты")
      return
    }
    fetch(`${window.location.origin}/db/users`, {
      method: "PUT",
      body: JSON.stringify(user),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      get()
      user = structuredClone(userDefault)
      alert("Менедежер Обновлен")
    }).catch((err) => {
      alert(err)
    })
  }

  const del = () => {
    if (user.id == null){
      alert("Выберите пользователя")
      return
    }
    fetch(`${window.location.origin}/db/users`, {
      method: "DELETE",
      body: JSON.stringify(user),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      get()
      user = structuredClone(userDefault)
      alert("Менедежер Удален")
    }).catch((err) => {
      alert(err)
    })
  }
  
  const select = () => {
    if (user.name == ""){
      user = structuredClone(userDefault)
      return
    }
    for (let i of managers){
      if (i.name == user.name){
        user = structuredClone(i)
        return
      }
    }
  }

  const telInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const telFormat = () => {
    user.tel = user.tel.replace(/[^0-9]/g, "")
  }

  const validateEmail = () => {
    return user.email.match(
      /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    )
  }
  
</script>
  

<div class="container">
  <div>
    <span>Менеджеры</span>
  </div>
  <div>
    <select bind:value={user.name} on:change={select}>
      <option value="">Новый</option>
      {#each managers as manager}
        <option value={manager.name}>{manager.name}</option>
      {/each}
    </select>
  </div>
  <div>
    <span>ID: {user.id}</span>
  </div>
  <div>
    <input type="text" bind:value={user.name} placeholder="Имя">
  </div>
  <div>
    <input type="text" bind:value={user.tel} on:keypress={telInput} on:input={telFormat} placeholder="Телефон">
  </div>
  <div>
    <input type="text" bind:value={user.email} placeholder="Email">
  </div>
  <div>
    {#if user.id == null}
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