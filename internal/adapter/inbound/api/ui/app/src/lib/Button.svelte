<script>
    export let link = ()=>{};
    export let color = "var(--glass-pane-1)";
</script>

<button on:click={link} style="--color: {color}"><slot></slot></button>

<style>
    button {
        padding: 15px;
        background-color: var(--color);
        border-radius: 15px;
        border: none;
        font-size: medium;
    }
</style>