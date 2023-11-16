<script>
    import routes from "../routes.js";

    let here = location.pathname;
    const menuItems = {
        "Globals": [
            {
                title: "Home",
                url: "/",
                icon: "home",
            },
            {
                title: "Pattern",
                url: "/pattern",
                icon: "pattern",
            },
        ],
        "Events": [
            {
                title: "Re-Subs",
                url: "/resubs",
                icon: "event_repeat",
            },
            {
                title: "Gift subs",
                url: "/gifts",
                icon: "partner_exchange",
            },
            {
                title: "Bits",
                url: "/bits",
                icon: "send_money",
            },
            {
                title: "Donos",
                url: "/donations",
                icon: "euro",
                active: true,
            },
        ]
    }
</script>
{#key location.pathname}
{#each Object.entries(menuItems) as [key, value]}
    <h1>{key}</h1>
    {#each value as row}
        <a href={row.url} class:here={row.url===here} on:click={()=>{here=row.url}}><span class="material-symbols-outlined">{row.icon}</span> {row.title}</a>
    {/each}
{/each}
    {/key}

<style>
    a {
        display: block;
        padding: 15px;
        border-radius: 15px;

        text-transform: uppercase;
        font-weight: 600;
        margin-bottom: 5px;
    }

    a.here {
        background-color:  var(--glass-pane-2);
    }

    a:hover {
        background-color: var(--glass-pane-1);
    }

    h1 {
        font-family: 'Gloria Hallelujah', cursive;
    }

    .material-symbols-outlined {
        vertical-align: text-bottom;
        margin-right: 10px;
    }
</style>