<script>
  import DataCargo from "./DataCargo.svelte";
  import DataManager from "./DataManager.svelte";
  import DataChem from "./DataChem.svelte"
  import { createEventDispatcher } from 'svelte'

  const dispatch = createEventDispatcher()

  export let managers = {}
  export let cargos = {}
  export let chems = {}

</script>

<div class="container">
  <DataManager {managers} on:message={(e) => dispatch('message', {text: e.detail.text})}></DataManager>
  <DataCargo {cargos} on:message={(e) => dispatch('message', {text: e.detail.text})}></DataCargo>
  <DataChem {chems} on:message={(e) => dispatch('message', {text: e.detail.text})}></DataChem>
</div>

<style>
  .container{
    display: flex;
    margin-left: 30px;
  }
</style>