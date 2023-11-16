<script>
    import PresetSelect from "./PresetSelect.svelte";
    import Button from "./Button.svelte";
    import {onMount} from "svelte";
    import page from "page";

    export let type = "Amount";
    export let uri;
    export let back;

   let reaction = {}

    let timer;
    const debounce = v => {
        clearTimeout(timer);
        timer = setTimeout(save, 750);
    }

    let load = async ()=>{
        let response = await fetch(uri)
        reaction = await response.json()
    }

    let save = async ()=>{
        await fetch( uri, {
            method: "POST",
            body: JSON.stringify(reaction),
        })
    }

    let go = async ()=>{
        await save();
        page(back);
    }

    let del = async ()=>{
        await fetch( uri, { method: "DELETE" })
        page(back);
    }

    let test = async ()=>{
        await fetch( "/api/test", {
            method: "POST",
            body: JSON.stringify(reaction),
        })
    }
    onMount(load);
</script>

<slot></slot>
<p></p>
<div id="menu">
    <label for="amount">{type}</label>
    <input type="number"
            bind:value={reaction.amount}
            on:input={debounce}>

    <label for="preset">Preset</label>
    <PresetSelect bind:selection={reaction.presetId} on:change={save}></PresetSelect>

    <label for="duration">Duration (millis)</label>
    <input type="number"
           min="500"
           max="100000"
           step="100"
           bind:value={reaction.duration}
           on:input={debounce}>

</div>
<div id="actions">
    <Button link={go}>
        <span class="material-symbols-outlined">undo</span> Back
    </Button>

    <Button link={test}>
        <span class="material-symbols-outlined">backlight_low</span> Test
    </Button>

    <Button link={del} color="#411">
        <span class="material-symbols-outlined">delete</span> Delete
    </Button>
</div>

<style>
    #menu {
        display: grid;
        grid-template-columns: 40% auto;
        grid-gap: 10px;
        margin-bottom: 2em;
    }
    #menu label {
        padding: 15px 0;
        font-size: larger;
    }

    #actions {
        display: flex;
        justify-content: space-between;
    }

    span.material-symbols-outlined {
        vertical-align: text-bottom;
    }
    input {
        text-align: right;
    }
</style>