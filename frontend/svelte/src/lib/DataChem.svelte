<script>

  import { getContext } from 'svelte'

  const uri = getContext('uri')

  let chem = {
    id: null,
    name: null,
    sellValue: null,
    probeValue: null
  }

  let chems = []

  const get = () => {
    fetch(`${uri}/db/chems`).then(function(response) {
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

  const add = () => {
    if (chem.name == null || chem.sellValue == null || chem.probeValue == null){
      alert("Заполните пустые поля")
      return
    }
    chem.sellValue = Number(chem.sellValue)
    chem.probeValue = Number(chem.probeValue)
    fetch(`${uri}/db/chems`, {
      method: "POST",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
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

  const update = () => {
    if (chem.name == null || chem.sellValue == null || chem.probeValue == null){
      alert("Заполните пустые поля")
      return
    }
    chem.sellValue = Number(chem.sellValue)
    chem.probeValue = Number(chem.probeValue)
    fetch(`${uri}/db/chems`, {
      method: "PUT",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      get()
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
    fetch(`${uri}/db/chems`, {
      method: "DELETE",
      body: JSON.stringify(chem),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
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

  const select = () => {
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
    margin: 5px;
    padding: 5px;
    border: 1px solid black;
    width: 250px;
    height: 250px;
    text-align: center;
  }

</style>