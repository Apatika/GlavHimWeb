<script>
  import { createEventDispatcher } from 'svelte'
  const dispatch = createEventDispatcher()

  let chem = {
    id: null,
    name: null,
    sellValue: null,
    probeValue: null
  }

  export let chems = {}

  const add = () => {
    if (chem.name == null || chem.sellValue == null || chem.probeValue == null){
      alert("Заполните пустые поля")
      return
    }
    chem.sellValue = Number(chem.sellValue)
    chem.probeValue = Number(chem.probeValue)
    fetch(`${window.location.origin}/db/chems`, {
      method: "POST",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('message', {text: 'chems'})
      Object.keys(chem).forEach(c => chem[c] = null)
      alert("Запись Добавлена")
    }).catch((err) => {
      alert(err)
    })
  }

  const update = () => {
    if (chem.name == null || chem.sellValue == null || chem.probeValue == null){
      alert("Заполните пустые поля")
      return
    }
    chem.sellValue = Number(chem.sellValue)
    chem.probeValue = Number(chem.probeValue)
    fetch(`${window.location.origin}/db/chems`, {
      method: "PUT",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('message', {text: 'chems'})
      Object.keys(chem).forEach(c => chem[c] = null)
      alert("Запись Обновлена")
    }).catch((err) => {
      alert(err)
    })
  }

  const del = () => {
    if (chem.id == null){
      alert("Выберите транспортную")
      return
    }
    fetch(`${window.location.origin}/db/chems`, {
      method: "DELETE",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      dispatch('message', {text: 'chems'})
      Object.keys(chem).forEach(c => chem[c] = null)
      alert("Запись удалена")
    }).catch((err) => {
      alert(err)
      alert(JSON.stringify(chems))
    })
  }

  const select = (id) => {
    chem.id = null
    chem.name = null
    chem.sellValue = null
    chem.probeValue = null
    if (id == ""){
      return
    }
    chem = chems[id]
  }

  const chemInput = (event) => {if (!(event.key).match(/[0-9]/i)) event.preventDefault()}
  const chemFormat = () => {
    chem.sellValue = chem.sellValue.replace(/[^0-9]/g, "")
    chem.probeValue = chem.probeValue.replace(/[^0-9]/g, "")
  }

</script>

<div>
  <div class="container">
    <div>
      <span>Химия</span>
    </div>
    <div>
      <select value={chem.id} on:change={(e) => select(e.target.value)}>
        <option value="" selected>Новая</option>
        {#each Object.values(chems).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as chem}
          <option value={chem.id}>{chem.name}</option>
        {/each}
      </select>
    </div>
    <div>
      <span>ID: {chem.id}</span>
    </div>
    <div>
      <input type="text" bind:value={chem.name} placeholder="Название">
    </div>
    <div>
      <input type="text" bind:value={chem.sellValue} on:keypress={chemInput} on:input={chemFormat} placeholder="Объем на продажу, л">
    </div>
    <div>
      <input type="text" bind:value={chem.probeValue} on:keypress={chemInput} on:input={chemFormat} placeholder="Объем пробника, мл">
    </div>
    <div>
      {#if chem.id == null}
        <button on:click={add}>Добавить</button>
      {:else}
        <button on:click={update}>Изменить</button>
        <button on:click={del}>Удалить</button>
      {/if}
    </div>
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