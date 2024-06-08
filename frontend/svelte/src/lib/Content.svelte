<script>

  export let order = {}

  const close = (e) => {
    e.target.closest('.content').style.display = 'none'
    document.body.style.pointerEvents = "all"
  }

</script>

<div class="container print">
  <div class="buttons">
    <button on:click={() => window.print()}>печать</button>
    <button class="close" on:click={close}>закрыть</button>
  </div>
  <div class="name">
    <strong>
      <span>{#if order.customer.surname != ""}{order.customer.surname}{/if} </span>
      <span>{order.customer.name} </span>
      <span>{#if order.customer.surname != ""}{order.customer.secondName}{/if}</span>
    </strong>
  </div>
  {#if order.content != ""}
    <div class="content">
      {@html (order.content + "<style>table:nth-child(odd) td{border-bottom: 1px solid black;}table:not(:first-child){border-top: 2px solid black}table:nth-child(even){border-top: none}</style>")}
    </div>
  {/if}
  {#if order.comment != ""}
    <div class="comment">
      <strong>Комментарий: </strong>{order.comment}
    </div>
  {/if}
  {#if Object.keys(order.probes).length > 0}
    <div class="probes">
      <span><strong>ПРОБНИКИ</strong></span>
      {#each Object.values(order.probes).filter(val => val.probeCount > 0).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as chem}
        <div>{chem.name} - {(chem.probeValue * chem.probeCount) / 1000} л.</div>
      {/each}
    </div>
  {/if}
</div>

<style>

  .container{
    position: absolute;
    width: 800px;
    height: 90%;
    left: 50%;
    top: 50%;
    padding: 15px;
    transform: translate(-50%, -50%);
    border: 1px solid black;
    border-radius: 5px;
    background-color: #FFF5EE;
    overflow:auto;
    box-shadow: 0px 0px 10px 0px black;
    z-index: 5;
  }
  .name, .comment, .probes{
    margin: 5px 15px;
  }
  .comment{
    border-bottom: 1px solid black;
  }
  .content{
    border: 1px solid black;
    border-radius: 10px;
    padding: 5px;
  }
  .close{
    float: right;
  }
  @media print{
    :global(body){
      visibility: hidden;
    }
    .print{
      visibility: visible;
      overflow:visible;
      box-shadow: none;
      border: none;
      transform: none;
      left: 0;
      top: 0;
    }
    .buttons{
      visibility: hidden;
    }
  }
  
</style>