<script>

    import ReactionEditor from "../lib/ReactionEditor.svelte";

    export let params = {};
    let id = params.id;
</script>

<ReactionEditor
        type="Gift subs"
        uri="/api/gifts/{id}"
        back="/gifts">
</ReactionEditor>