<script>

  export let order = {}

  const close = (e) => {
    e.target.closest('.content').style.display = 'none'
    document.body.style.pointerEvents = "all"
  }

</script>

<div class="container">
  <div class="name">
    <button class="close" on:click={close}>закрыть</button>
  </div>
  <div>
    <strong>
      <span>{#if order.customer.surname != ""}{order.customer.surname}{/if} </span>
      <span>{order.customer.name} </span>
      <span>{#if order.customer.surname != ""}{order.customer.secondName}{/if}</span>
    </strong>
  </div>
  {#if order.content != ""}
    <div class="content">
      {@html (order.content + "<style>table:first-of-type td{border-bottom: 1px solid black;}</style>")}
    </div>
  {/if}
  {#if order.comment != ""}
    <div class="comment">
      <strong>Комментарий: </strong>{order.comment}
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
    padding: 5px;
    transform: translate(-50%, -50%);
    border: 1px solid black;
    border-radius: 5px;
    background-color: #FFF5EE;
    overflow:auto;
    box-shadow: 0px 0px 10px 0px black;
    z-index: 5;
  }
  .name{
    margin-left: 5px;
    margin-bottom: 5px;
  }
  .content{
    border: 1px solid black;
    border-radius: 10px;
    padding: 5px;
  }
  .comment{
    margin-top: 5px;
  }
  .close{
    float: right;
  }
  
</style>