<script>
    import {createEventDispatcher, onMount} from "svelte";

    export let selection = 2;
    export let dispatch = createEventDispatcher();

    let presets = [];

    onMount(async () => {
        let response = await fetch("/api/presets");
        let data = await response.json()
        presets = data.presets;
    })

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