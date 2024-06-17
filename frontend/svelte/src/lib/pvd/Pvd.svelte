<script>

  let weight = null
  let count = null
  let sum = 0
  let reservePvd = []
  export let pvd = []

  const add = () => {
    weight = Number(weight)
    count = Number(count)
    if (isNaN(weight) || isNaN(count)){
      alert("ERROR: неверные значения")
      return
    }
    if (weight <= 0 || count <= 0) {
      alert("ERROR: нулевые значения")
      return
    }
    sum = 0
    pvd.forEach(p => {
      p.count += p.reserved
      p.reserved = 0
    })
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
      weight = null
      count = null
    }).catch((err) => {
      alert(err)
    })
  }

</script>

<div class="container">
  <div class="stored">
    <div class="add">
      <input class="add-input" type="tel" bind:value={weight} placeholder="кг">
      <input class="add-input" type="tel" bind:value={count} placeholder="кол-во">
      <button class="push" on:click={add}>Добавить</button>
    </div>
    <div class="in-store">
      <div>Сумма: <span class="reserved-sum">{sum}кг.</span></div>
      {#each pvd as p}
        <div>
          <span class="reserved-info"><span>{p.weight}кг.</span> - <span>{p.count}шт.</span></span>
          <button on:click={() => {if (p.reserved > 0) {p.reserved--; p.count++; sum -= p.weight}}}>-</button>
          <span class="reserved-span">{p.reserved}шт.</span>
          <button on:click={() => {if (p.count > 0) {p.reserved++; p.count--; sum += p.weight}}}>+</button>
        </div>
      {/each}
    </div>
  </div>
  <div class="reserved">

  </div>
</div>

<style>

  .container{
    margin: 10px;
    display: flex;
  }
  .stored{
    width: 300px;
    text-align: center;
  }
  .reserved{
    flex-grow: 1;
  }
  .add-input{
    width: 40px;
    text-align: center;
  }
  .reserved-info span, .reserved-span, .reserved-sum{
    display: inline-block;
    width: 45px;
  }

</style>