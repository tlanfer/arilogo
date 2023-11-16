<script>
    import {onMount} from "svelte";
    import OrderedList from "../lib/OrderedList.svelte";
    import Button from "../lib/Button.svelte";
    import page from "page";

    let alerts = []

    let load = async ()=>{
        let response = await fetch("/api/resubs")
        let data = await response.json()

        let all = [];
        for (const [key, value] of Object.entries(data.reactions)) {
            all.push({
                months: value.amount,
                url: "/resubs/"+key,
                presetId: value.presetId
            })
        }
        alerts = all;
    }

    let add = async ()=>{
        let response = await fetch("/api/resubs", {method:"POST"});
        let location = response.headers.get("Location");
        page("/resubs/"+location);
    }

    onMount(load);
</script>

<div id="actions">
    This is where we configure resub alerts.
    <Button link={add}>
        <span class="material-symbols-outlined">add</span> Add
    </Button>
</div>

<OrderedList
        things={alerts}
        sortFn={(a,b)=>{ return a.months-b.months}}
        titleFn={ a=>a.months+" months"}
        subTitleFn={a=>a.pattern}
></OrderedList>