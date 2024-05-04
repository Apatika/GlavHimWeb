<script>
  import NewOrder from "./NewOrder.svelte"
  import InWorkOrders from "./InWorkOrders.svelte"
</script>

<div id="container">
  <div id="content">
    <InWorkOrders></InWorkOrders>
  </div>
  <div id="new-order">
    <NewOrder></NewOrder>
  </div>
</div>

<style>
  #container{
    display: flex;
    height: 100%;
  }
  #content{
    flex-grow: 1;
    border-right: 1px solid black;
    padding: 10px 5px;
  }
  #new-order{
    flex-basis: 25%;
  }
</style>
