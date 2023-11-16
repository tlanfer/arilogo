<script>
    import {createEventDispatcher, onMount} from "svelte";

    export let selection = 2;
    export let dispatch = createEventDispatcher();

    export let trigger = "";
    let presets = [];

    $: onTriggerChange(trigger)

    let onTriggerChange = async ()=>{
        await load();
    }

    export let load = async () => {
        let response = await fetch("/api/presets");
        let data = await response.json()
        presets = data.presets;
    }


    onMount(load);

    let change = ()=>{
        dispatch("change", {});
    }
</script>

<select bind:value={selection} on:change={change}>
    {#each presets as p}
    <option value={p.id}>{p.name}</option>
    {/each}
</select>

<style>
    select {
        width: 100%;
        text-align: right;
    }

    option {
        color: var(--main-font-color);
        background-color: var(--main-bg-color);
    }
</style>