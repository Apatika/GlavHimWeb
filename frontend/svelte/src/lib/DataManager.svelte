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

  const get = () => {
    fetch(`${uri}/db/users`).then(function(response) {
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
  
  const add = () => {
    if (user.name == null || user.tel == null || user.email == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateEmail()){
      alert("Неправильный формат почты")
      return
    }
    fetch(`${uri}/db/users`, {
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

  const update = () => {
    if (user.name == null || user.tel == null || user.email == null){
      alert("Заполните пустые поля")
      return
    }
    if (!validateEmail()){
      alert("Неправильный формат почты")
      return
    }
    fetch(`${uri}/db/users`, {
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

  const del = () => {
    if (user.id == null){
      alert("Выберите пользователя")
      return
    }
    fetch(`${uri}/db/users`, {
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
  
  const select = () => {
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

  const telInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const telFormat = (event) => {
    event.preventDefault()
    let paste = (event.clipboardData || window.clipboardData).getData("text")
    paste = paste.replace(/[^0-9]/g, "")
    event.target.value = paste
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
    <input type="text" bind:value={user.tel} on:paste={telFormat} on:keypress={telInput} placeholder="Телефон">
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