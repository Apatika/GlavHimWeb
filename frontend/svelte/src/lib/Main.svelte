<script>
  import NewOrder from "./NewOrder.svelte"
  import InWorkOrders from "./InWorkOrders.svelte"
</script>

<div id="container">
  <div id="content">
    <InWorkOrders></InWorkOrders>
  </div>
</div>

<style>
  #container{
    display: flex;
    height: 100%;
  }
  #content{
    flex-grow: 1;
    padding-left: 8px;
  }
</style>
