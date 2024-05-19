<script>

  export let orders = {}
  let ids = []

</script>

<div class="container">
  <div id="all-routes">
    {#each Object.values(orders).filter(val => val.order.status == "Развозка") as order}
      <div class="route">
        <div>
          {order.order.cargo}
        </div>
        <div>
          {#if order.client.surname != ""}{order.client.surname}{/if}
          {order.client.name}
          {#if order.client.surname != ""}{order.client.secondName}{/if}          
        </div>
        <button class="toggle">></button>
      </div>
    {/each}
  </div>
  <div id="selected-routes"></div>
</div>

<style>

  .container{
    display: flex;
    padding: 30px;
  }
  .route{
    position: relative;
    margin-bottom: 5px;
    padding: 5px;
    border: 1px solid black;
    border-radius: 15px;
    background-color: #FFA500;
    box-shadow: 0px 0px 5px 0px grey;
    font-size: small;
    width: 300px;
  }
  .toggle{
    position: absolute;
    right: 20px;
    top: 50%;
    transform: translate(0, -50%);
  }
  #all-routes, #selected-routes{
    width: 450px;
    border-right: 1px solid black;
    height: 100vh;
  }

</style>