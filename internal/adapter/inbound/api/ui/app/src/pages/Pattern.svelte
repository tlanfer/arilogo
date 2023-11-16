<script>
    import {onMount} from "svelte";

    let presets = []
    let deviceAddr = "";

    let loadPresets = async ()=>{
        let response = await fetch("/api/presets");
        let data = await response.json()
        presets = data.presets;
    }

    let loadDeviceAddr = async ()=>{
        let response = await fetch("/api/device/addr");
        let data = await response.json()
        deviceAddr = "http://"+data.address;
    }

    onMount(async () => {
       await Promise.all([
           loadDeviceAddr(),
           loadPresets(),
       ])
    })
</script>

<div>
    <p>
        These are your available presets.
    </p>
    <ul id="pattern">
        {#each presets as p}
            <li>{p.name}</li>
        {/each}
    </ul>
    {#if deviceAddr !== ""}
    <p>
        <a href={ deviceAddr } target="_blank">Click here to create more presets</a>
    </p>
    {/if}
</div>

<style>
    #pattern {
        list-style:none;
        padding: 0;
        display: grid;
        grid-template-columns: repeat(2, minmax(0, 1fr));
    ;
        gap: 15px;
    }

    #pattern li {
        height: 50px;
        background-color: var(--glass-pane-1);
        padding: 15px;
        border-radius: 15px;

        display: flex;
        justify-content: center;
        align-items: center;
    }
</style>