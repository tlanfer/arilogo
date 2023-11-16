<script>
    import {onMount} from "svelte";
    import Button from "./Button.svelte";

    export let use = (device)=>{};
    let devices = []
    let scanning = false;

    let load = async ()=>{
        scanning = true
        let response = await fetch("/api/device/scan");
        let reader = response.body.pipeThrough(new TextDecoderStream()).getReader();

        while(true){
            const {value, done} = await reader.read();
            if(done) break;

            devices = [...devices, value];
        }
        scanning = false;
    }

    onMount(load)
</script>

<ul>
    {#if scanning}
        Scanning...
    {/if}
    {#each devices as d}
        <li><Button link={()=>use(d)}><span class="material-symbols-outlined">highlight</span> Click to use {d}</Button></li>
    {/each}

</ul>

<style>
    ul {
        list-style: none;
        padding: 0;
    }
    li {
        width: 100%;
    }
    span.material-symbols-outlined {
        vertical-align: text-bottom;
    }
</style>