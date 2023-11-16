<script>
    import {Confirm} from "svelte-confirm";

    export let things = [];

    export let sortFn = (a, b)=>{a-b};
    export let titleFn = (a)=>{ return JSON.stringify(a)};
    export let subTitleFn = (a)=>{ return JSON.stringify(a)};

    let sortedThings = [];

    $: sortedThings = sorted(things);

    let sorted = (things)=>{
        return things.sort(sortFn)
    }
</script>

<ol>
    {#each sortedThings as thing}
        <li>
            <span>
                { titleFn(thing) }
                {#if subTitleFn(thing) }<small>{subTitleFn(thing)}</small>{/if}
            </span>
            <a href={thing.url} class="transparent"><span class="material-symbols-outlined">tune</span></a>
        </li>
    {/each}
</ol>


<style>
    ol {
        list-style: none;
        padding: 0;
    }

    li {
        background-color: var(--glass-pane-1);
        padding: 15px;
        margin-bottom: 15px;
        border-radius: 15px;
    }

    li a {
        float: right;
    }
</style>