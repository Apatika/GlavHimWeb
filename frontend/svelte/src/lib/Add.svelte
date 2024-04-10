<script>

let user = {
  name: null,
  tel: null,
  email: null
}

let addUser = () => {
  if (user.name == null || user.tel == null || user.email == null){
    alert("Заполните пустые поля")
    return
  }
  fetch("http://localhost:8081/users", {
    method: "POST",
    body: JSON.stringify(user),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  })
  user.name = null
  user.tel = null
  user.email = null
}

</script>

<div class="container">
  <div id="user">
    <div>
      <span>Добавить менеджера</span>
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
      <button on:click={addUser}>Добавить</button>
    </div>
  </div>
</div>

<style>
  .container{
    display: flex;
  }

  #user{
    margin: 5px;
    padding: 5px;
    border: 1px solid black;
  }
</style>