<script>
  import DataCargo from "./DataCargo.svelte";
  import DataManager from "./DataManager.svelte";
  import DataChem from "./DataChem.svelte"

  export let managers = []
  export let cargos = []
  export let chems = []

</script>

<div class="container">
  <DataManager {managers}></DataManager>
  <DataCargo {cargos}></DataCargo>
  <DataChem {chems}></DataChem>
</div>

<style>
  .container{
    display: flex;
  }
</style>