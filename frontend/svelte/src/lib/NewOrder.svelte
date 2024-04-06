<script>

  let data = {
    id: "",
    type: 0,
    name: "",
    surname: "",
    secondName: "",
    manager: "",
    inn: "",
    passportSerial: null,
    passportNum: null,
    contact: [{name: "", tel: ""}],
    tel: "",
    email: "",
  }

  const addContact = () => data.contact = data.contact.concat({name: "", tel: ""})
  const removeContact = () => data.contact = data.contact.slice(0, -1)

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
      <input type="text" bind:value={data.surname} placeholder="Фамилия">
      <input type="text" bind:value={data.name} placeholder="Имя">
      <input type="text" bind:value={data.secondName} placeholder="Отчество">
    {:else}
      <input type="text" id="name" bind:value={data.name} placeholder="Наименование компании">
    {/if}
  </div>
  <div class="flex underline">
    <div class="lables">
      <div class="flex">
        <div class="lables">ID:</div>
        <div>{data.id}</div>
      </div>
      <div class="flex">
        <div class="lables">Менеджер:</div>
        <div>
          <select bind:value={data.manager}>
          
          </select>
        </div>
      </div>
      <div class="flex">
        <div class="lables">
          {#if data.type != 2}
            ИНН: 
          {:else}
            Паспорт: 
          {/if}
        </div>
        <div>
          {#if data.type != 2}
            <input type="text" bind:value={data.inn} placeholder="ИНН">
          {:else}
            <input type="text" bind:value={data.passportSerial} placeholder="Серия">
            <input type="text" bind:value={data.passportNum} placeholder="Номер">
          {/if}
        </div>
      </div>
      <div class="flex">
        <div class="lables">Контакт:</div>
        <div>
          {#each data.contact as contact}
            <div>
              <input type="text" bind:value={contact.name} placeholder="Контакт">
              <input type="text" bind:value={contact.tel} placeholder="Телефон">
            </div>
          {/each}
          <button id="add-contact" on:click={addContact}>+</button>
          <button id="add-contact" on:click={removeContact}>-</button>
        </div>
      </div>
      <div class="flex">
        <div class="lables">Email:</div>
        <div>
          <input type="text" bind:value={data.email} placeholder="Email">
        </div>
      </div>
    </div>
  </div>
  <div class="flex underline">
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
    flex-basis: 26%;
  }
  .underline{
    border-bottom: 1px solid blue;
  }
  .flex div{
    margin-bottom: 2px;
  }
</style>