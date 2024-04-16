<script>

  import { getContext } from 'svelte'
    import Data from './Data.svelte';

  const uri = getContext('uri')

  let chem = {
    id: null,
    name: null,
    sellValue: null,
    probeValue: null
  }

  let chems = []

  let get = () => {
    fetch(`${uri}/chems`).then(function(response) {
      if (response.status != 200) throw new Error(response.statusText)
      return response.json();
    }).then(function(d) {
      d.forEach(e => {
        e.id = e["_id"]
        delete e["_id"]
      });
      chems = d
    }).catch(function(err) {
      alert(err);
    })
  }

  get()

  let add = () => {
    if (chem.name == null || chem.sellValue == null || chem.probeValue == null){
      alert("Заполните пустые поля")
      return
    }
    chem.sellValue = Number(chem.sellValue)
    chem.probeValue = Number(chem.probeValue)
    fetch(`${uri}/chems`, {
      method: "POST",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      chem.id = null
      chem.name = null
      chem.sellValue = null
      chem.probeValue = null
      alert("Запись Добавлена")
    }).catch((err) => {
      alert(err)
    })
  }

  let update = () => {
    if (chem.name == null || chem.sellValue == null || chem.probeValue == null){
      alert("Заполните пустые поля")
      return
    }
    chem.sellValue = Number(chem.sellValue)
    chem.probeValue = Number(chem.probeValue)
    fetch(`${uri}/chems`, {
      method: "PUT",
      body: JSON.stringify(chem),
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

  let del = () => {
    if (chem.id == null){
      alert("Выберите транспортную")
      return
    }
    fetch(`${uri}/chems`, {
      method: "DELETE",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then((response) => {
      return response.json()
    }).then(data => {
      if (data.code != 200) throw new Error(data.message)
      get()
      chem.id = null
      chem.name = null
      chem.sellValue = null
      chem.probeValue = null
      alert("Запись удалена")
    }).catch((err) => {
      alert(err)
    })
  }

let select = () => {
    if (chem.name == ""){
      chem.id = null
      chem.name = null
      chem.sellValue = null
      chem.probeValue = null
      return
    }
    for (let i of chems){
      if (i.name == chem.name){
        chem.id = i.id
        chem.name = i.name
        chem.sellValue = i.sellValue
        chem.probeValue = i.probeValue
        return
      }
    }
  }

</script>

<div>
  <div class="container">
    <div>
      <span>Химия</span>
    </div>
    <div>
      <select bind:value={chem.name} on:change={select}>
        <option value="">Новая</option>
        {#each chems as chem}
          <option value={chem.name}>{chem.name}</option>
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
      <input type="text" bind:value={chem.sellValue} placeholder="Объем на продажу, л">
    </div>
    <div>
      <input type="text" bind:value={chem.probeValue} placeholder="Объем пробника, мл">
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
    margin: 5px;
    padding: 5px;
    border: 1px solid black;
    width: 250px;
    height: 250px;
    text-align: center;
  }

</style>