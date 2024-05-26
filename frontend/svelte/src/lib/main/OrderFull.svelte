<script>
  import { createEventDispatcher } from "svelte"
  const dispatch = createEventDispatcher()

  export let order = {}

</script>

<div class="container">
  <div class="customer">
    <div class="item">
      <div class="label"><strong>Менеджер:</strong></div>
      <div>{order.customer.manager}</div>
    </div>
    {#if order.customer.surname != ""}
      <div class="item">
        <div class="label"><strong>Фамилия:</strong></div>
        <div>{order.customer.surname}</div>
      </div>
    {/if}
      <div class="item">
        <div class="label"><strong>Имя:</strong></div>
        <div>{order.customer.name}</div>
      </div>
    {#if order.customer.secondName != ""}
      <div class="item">
        <div class="label"><strong>Отчество</strong></div>
        <div>{order.customer.secondName}</div>
      </div>
    {/if}
    {#if order.customer.inn != ""}
      <div class="item">
        <div class="label"><strong>ИНН:</strong></div>
        <div>{order.customer.inn}</div>
      </div>
    {/if}
    {#if order.customer.passportSerial != ""}
      <div class="item">
        <div class="label"><strong>Паспорт:</strong></div>
        <div>{order.customer.passportSerial} {order.customer.passportNum}</div>
      </div>
    {/if}
      <div class="item">
        <div class="label"><strong>Контакты:</strong></div>
        <div>
          {#each order.customer.contact as contact}
            <div>{contact.name} {contact.tel}</div>
          {/each}
        </div>
      </div>
    {#if order.customer.email != ""}
      <div class="item">
        <div class="label"><strong>Email:</strong></div>
        <div>{order.customer.email}</div>
      </div>
    {/if}
  </div>
  <div class="order">
    <div class="item">
      <div class="label"><strong>ТК:</strong></div>
      <div>{order.order.cargo}</div>
    </div>
    <div class="item">
      <div class="label"><strong>Город:</strong></div>
      <div>{order.order.adress.city}</div>
    </div>
    {#if order.order.adress.adress != ""} 
      <div class="item">
        <div class="label"><strong>Адрес:</strong></div>
        <div>{order.order.adress.adress}</div>
      </div>
    {:else}
      <div class="item">
        <div class="label"><strong>Терминал:</strong></div>
        <div>{order.order.adress.terminal}</div>
      </div>
    {/if}
    <div>
      <div class="item">
        <div class="label"><strong>Счет:</strong></div>
        <div>
          {#each order.order.invoice as invoice}
            <span>{invoice} </span>
          {/each}
        </div>
      </div>  
    </div>
    <div class="item">
      <div class="label"><strong>Крайняя дата:</strong></div>
      <div>{order.order.lastDate}</div>
    </div>
    <div class="item">
      <div class="label"><strong>Комментарий:</strong></div>
      <div>{order.order.comment}</div>
    </div>
  </div>
  <div class="probes">
    {#if Object.keys(order.order.probes).length > 0}<span><strong>ПРОБНИКИ</strong></span>{/if}
    {#each Object.values(order.order.probes).filter(val => val.probeCount > 0).sort((a, b) => {return a.name == b.name ? 0 : (a.name > b.name ? 1 : -1)}) as chem}
      <div>{chem.name} - {(chem.probeValue * chem.probeCount) / 1000} л.</div>
    {/each}
  </div>
  <div>
    <div>
      <button class="edit-button" on:click={() => dispatch('message', {order: order.order, customer: order.customer})}>Редактировать</button>
    </div>
    <div>
      {#if order.order.payment} <span style="color: red; font-size: 14px;"><strong>ЗА НАШ СЧЕТ!!!</strong></span> {/if}
    </div>
  </div>
</div>

<style>

  .container{
    display: flex;
    font-size: small;
  }
  .order{
    width: 350px;
  }
  .customer{
    width: 300px;
  }
  .item{
    display: flex;
    word-break: break-word;
    word-wrap: break-word;
    overflow-wrap: break-word;
  }
  .label{
    width: 100px;
    min-width: 100px;
  }
  .edit-button{
    float: right;
    height: 30px;
  }
  .probes{
    flex-grow: 1;
    text-align: center;
  }
  @media (min-width:1365px) and (max-width:1600px){}
  @media (max-width:1364px){
    .container{
      display: block;
      overflow: auto;
    }
  }

</style>