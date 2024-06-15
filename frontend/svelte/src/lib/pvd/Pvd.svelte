<script>

  let weight = 0
  let count = 0
  export let pvd = []

  const add = () => {
    if (weight == 0 || count == 0) {
      alert("ERROR: нулевые значения")
      return
    }
    fetch(`${window.location.origin}/pvd`, {
      method: "POST",
      body: JSON.stringify({weight: weight, count: count}),
      headers: {
        "Content-type": "application/json; charset=UTF-8"
      }
    }).then(response => {
      if (!response.ok) return response.text().then(text => {throw new Error(text)})
      return response.json()
    }).then((data) => {
      if (data.length == 0) throw new Error("storage empty")
      pvd = data
    }).catch((err) => {
      alert(err)
    })
  }

</script>

<div class="container">
  <div class="stored">
    <div class="add">
      <input type="number" bind:value={weight}>
      <input type="number" bind:value={count}>
      <button class="push" on:click={add}>Добавить</button>
    </div>
    <div class="in-store">
      {#each pvd as p}
        {p.weight} - {p.count}
      {/each}
    </div>
  </div>
  <div class="reserved">

  </div>
</div>

<style>

  .container{
    display: flex;
  }
  .stored{
    width: 300px;
  }
  .reserved{
    flex-grow: 1;
  }

</style>