<script>

    import ReactionEditor from "../lib/ReactionEditor.svelte";

    export let params = {};
    let id = params.id;
</script>

<ReactionEditor
        type="Donation"
        uri="/api/donations/{id}"
        back="/donations">
    This is where you configure donations. No currency conversion is done (for now). All values are in us dollar cents.
</ReactionEditor>