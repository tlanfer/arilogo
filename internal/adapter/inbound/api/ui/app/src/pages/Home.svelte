<script>
    import PresetSelect from "../lib/PresetSelect.svelte";
    import {onMount} from "svelte";
    import DeviceScanner from "../lib/DeviceScanner.svelte";
    import Button from "../lib/Button.svelte";

    let idlePresetId;
    let twitchChannel;
    let streamlabsToken;
    let deviceAddr;

    let trigger;

    let getIdlePreset = async () => {
            let response = await fetch("/api/idle");
            let data = await response.json()
            idlePresetId = data.idlePresetId;
    }
    let setIdlePreset = async () => {
            await fetch("/api/idle", {
                    method: "POST",
                    body: JSON.stringify({
                            idlePresetId: idlePresetId,
                    })
            })
    }


    let getTwitchChannel = async () => {
            let response = await fetch("/api/twitch");
            let data = await response.json()
            twitchChannel = data.channel;
    }
    let setTwitchChannel = async () => {
            await fetch("/api/twitch", {
                    method: "POST",
                    body: JSON.stringify({
                            channel: twitchChannel,
                    })
            })
    };

    let getStreamlabsToken = async () => {
            let response = await fetch("/api/streamlabs");
            let data = await response.json()
            streamlabsToken = data.token;
    }
    let setStreamlabsToken = async () => {
            await fetch("/api/streamlabs", {
                    method: "POST",
                    body: JSON.stringify({
                            token: streamlabsToken,
                    })
            })
    };


    let getDeviceAddr = async () => {
            let response = await fetch("/api/device/addr");
            let data = await response.json()
            deviceAddr = data.address;
    }
    let setDeviceAddr = async () => {
        await fetch("/api/device/addr", {
            method: "POST",
            body: JSON.stringify({
                    address: deviceAddr,
            })
        })
        trigger = new Date();
    };


    let load = async () => {
            await Promise.all([
                    getIdlePreset(),
                    getTwitchChannel(),
                    getStreamlabsToken(),
                    getDeviceAddr(),
            ]);
    }

    let scanning = false;
    let enableScanner = ()=>{
            scanning = true;
    }

    let useDevice = (d)=>{
            deviceAddr = d;
            scanning = false;
            setDeviceAddr();
    }

    onMount(load);
</script>

<div id="menu">
        <label for="device">Device</label>
        {#if scanning}
                <DeviceScanner use={useDevice}></DeviceScanner>
        {:else}
                <div id="deviceInput">
                        <input type="text" bind:value={deviceAddr} on:change={setDeviceAddr}>
                        <Button link={enableScanner}><span class="material-symbols-outlined">wifi_find</span></Button>
                </div>
        {/if}


        <label for="idlePreset">Idle preset</label>
        <PresetSelect bind:trigger={trigger} bind:selection={idlePresetId} on:change={setIdlePreset}></PresetSelect>

        <label for="twitchChannel">Twitch channel</label>
        <input type="text" bind:value={twitchChannel} on:change={setTwitchChannel}>

        <label for="streamlabs">Streamlabs Token</label>
        <input type="password" bind:value={streamlabsToken} on:change={setStreamlabsToken}>
        <div class="description"><a href="https://streamlabs.com/dashboard#/settings/api-settings" target="_blank">Click here</a> and log in, then find API Tokens ➡ Your Socket API Token</div>

</div>

<style>
    #menu {
        display: grid;
        grid-template-columns: 40% auto;
        grid-gap: 10px;
    }
    #menu .description {
        font-size: smaller;
        grid-column: 1 / 3 ;
    }
    #menu label {
        padding: 15px 0;
        font-size: larger;
    }
    #deviceInput {
            display: flex;
            gap: 15px;
    }
    #deviceInput input {
            flex: 1 0 auto;
    }
</style>
