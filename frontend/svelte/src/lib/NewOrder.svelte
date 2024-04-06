<script>

  let data = {
    id: "123123",
    type: 0,
    name: "",
    surname: "",
    secondName: "",
    manager: "",
    inn: "",
    passportSerial: null,
    passportNum: null,
    contact: "",
    tel: "",
    email: "",
  }

</script>

<div id="search">
  <input type="text" placeholder="Поиск...">
</div>
<div id="main">
  <div id="client-name">
    <select bind:value={data.type}>
      <option value="0">ИП</option>
      <option value="1">Юр.Лицо</option>
      <option value="2">Физ.Лицо</option>
    </select>
    {#if data.type != 1}
      <input type="text" bind:value={data.surname} placeholder="Фамилия...">
      <input type="text" bind:value={data.name} placeholder="Имя...">
      <input type="text" bind:value={data.secondName} placeholder="Отчество...">
    {:else}
      <input type="text" id="name" bind:value={data.name} placeholder="Наименование компании...">
    {/if}
  </div>
  <div class="flex">
    <div class="lables">
      <div>ID: </div>
      <div>Менеджер: </div>
      <div>
        {#if data.type != 2}
          ИНН: 
        {:else}
          Паспорт: 
        {/if}
      </div>
      <div>Контакт: </div>
      <div>Email: </div>
    </div>
    <div>
      <div>{data.id}</div>
      <div>
        <select bind:value={data.manager}>
        
        </select>
      </div>
      <div>
        {#if data.type != 2}
          <input type="text" bind:value={data.inn} placeholder="ИНН...">
        {:else}
          <input type="text" bind:value={data.passportSerial} placeholder="Серия...">
          <input type="text" bind:value={data.passportNum} placeholder="Номер...">
        {/if}
      </div>
      <div>
        <input type="text" bind:value={data.contact} placeholder="Контакт">
        <input type="text" bind:value={data.tel} placeholder="Телефон">
      </div>
      <div>
        <input type="text" bind:value={data.email} placeholder="Email">
      </div>
    </div>
  </div>
  <div class="flex">
    <div>
      <div>ТК: </div>
      <div>За Наш Счет: </div>
      <div>Город: </div>
      <div>Адрес: </div>
      <div>Счет: </div>
    </div>
    <div>

    </div>
  </div>
</div>

<style>
  #search{
    margin-left: 5px;
    margin-top: 10px;
  }
  #main{
    margin: 30px 0px 5px 50px;
  }
  .flex{
    display: flex;
  }
  .lables{
    flex-basis: 5%;
  }
  .flex{
    border-bottom: 1px solid blue;
  }
  .flex div{
    margin-bottom: 2px;
  }
</style>