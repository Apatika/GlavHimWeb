<script>

  import DataCargo from "./DataCargo.svelte";
  import DataManager from "./DataManager.svelte";
  import DataChem from "./DataChem.svelte"

</script>

<div class="container">
  <DataManager></DataManager>
  <DataCargo></DataCargo>
  <DataChem></DataChem>
</div>

<style>
  .container{
    display: flex;
  }
</style>