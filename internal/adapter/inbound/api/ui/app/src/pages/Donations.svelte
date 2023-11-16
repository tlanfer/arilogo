<script>
    import {onMount} from "svelte";
    import OrderedList from "../lib/OrderedList.svelte";
    import Button from "../lib/Button.svelte";
    import page from "page";

    let alerts = []

    let load = async ()=>{
        let response = await fetch("/api/donations")
        let data = await response.json()

        let all = [];
        for (const [key, value] of Object.entries(data.reactions)) {
            all.push({
                bits: value.amount,
                url: "/donations/"+key,
                presetId: value.presetId
            })
        }
        alerts = all;
    }

    let add = async ()=>{
        let response = await fetch("/api/donations", {method:"POST"});
        let location = response.headers.get("Location");
        page("/donations/"+location);
    }

    onMount(load);
</script>


<div id="actions">
    This is where we configure bit alerts.
    <Button link={add}>
        <span class="material-symbols-outlined">add</span> Add
    </Button>
</div>

<OrderedList
        things={alerts}
        sortFn={(a,b)=>a.bits-b.bits}
        titleFn={a=>a.bits + " cents" }
        subTitleFn={a=>""}
></OrderedList>

