<script>
  import { getContext } from 'svelte'

  const uri = getContext('uri')

  let user = {
    id: null,
    name: null,
    tel: null,
    email: null
  }
  
  let managers = []

  let get = () => {
    fetch(`${uri}/users`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then(function(d) {
      d.forEach(e => {
        e.id = e["_id"]
        delete e["_id"]
      });
      managers = d
    }).catch(function(err) {
      alert(err);
    })
  }

  get()
  
  let add = () => {
    if (user.name == null || user.tel == null || user.email == null){
      alert("Заполните пустые поля")
      return
    }
    fetch(`${uri}/users`, {
      method: "POST",
      body: JSON.stringify(user),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      user.id = null
      user.name = null
      user.tel = null
      user.email = null
      alert("Менедежер Добавлен")
    }).catch((err) => {
      alert(err)
    })
  }

  let update = () => {
    if (user.name == null || user.tel == null || user.email == null){
      alert("Заполните пустые поля")
      return
    }
    fetch(`${uri}/users`, {
      method: "PUT",
      body: JSON.stringify(user),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      alert("Менедежер Обновлен")
    }).catch((err) => {
      alert(err)
    })
  }

  let del = () => {
    if (user.id == null){
      alert("Выберите пользователя")
      return
    }
    fetch(`${uri}/users`, {
      method: "DELETE",
      body: JSON.stringify(user),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      user.id = null
      user.name = null
      user.tel = null
      user.email = null
      alert("Менедежер Удален")
    }).catch((err) => {
      alert(err)
    })
  }
  
  let select = () => {
    if (user.name == ""){
      user.id = null
      user.name = null
      user.tel = null
      user.email = null
      return
    }
    for (let i of managers){
      if (i.name == user.name){
        user.id = i.id
        user.name = i.name
        user.tel = i.tel
        user.email = i.email
        return
      }
    }
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
    <input type="text" bind:value={user.tel} placeholder="Телефон">
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
  }
  
</style>