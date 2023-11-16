<script>

    import ReactionEditor from "../lib/ReactionEditor.svelte";

    export let params = {};
    let id = params.id;
</script>

<ReactionEditor
        type="Months"
        uri="/api/resubs/{id}"
        back="/resubs">
</ReactionEditor>