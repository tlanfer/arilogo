<script>

    import ReactionEditor from "../lib/ReactionEditor.svelte";

    export let params = {};
    let id = params.id;
</script>

<ReactionEditor
        type="Bits"
        uri="/api/bits/{id}"
        back="/bits">
</ReactionEditor>