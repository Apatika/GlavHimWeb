<script>
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  let user = {
    id: null,
    name: null,
    tel: null,
    email: null
  }
  
  export let managers = {}
  
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
      dispatch('message', {text: 'managers'})
      Object.keys(user).forEach(u => user[u] = null)
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
      dispatch('message', {text: 'managers'})
      Object.keys(user).forEach(u => user[u] = null)
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
      dispatch('message', {text: 'managers'})
      Object.keys(user).forEach(u => user[u] = null)
      alert("Менедежер Удален")
    }).catch((err) => {
      alert(err)
    })
  }
  
  const select = (id) => {
    user = structuredClone(userDefault)
    if (id == ""){
      return
    }
    user = managers[id]
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
    <select value={user.id} on:change={(e) => select(e.target.value)}>
      <option value="">Новый</option>
      {#each Object.values(managers).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as manager}
        <option value={manager.id}>{manager.name}</option>
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