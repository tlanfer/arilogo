<script>
  import router from "page";
  import routes from './routes.js'
  import Menu from "./lib/Menu.svelte";

  let page;
  let title;
  let params;

  routes.forEach(route => {
    router(
      route.path,
      (ctx, next) => {
        params = ctx.params;
        next();
      },
      () => {
        page = route.component;
        title = route.title;
      }
    )
  })

  router.start();
</script>

  <header>
    <img src="/logo.png" alt="logo">
  </header>

  <nav>
    <Menu />
  </nav>

  <main>
    <h1>{title}</h1>
    <svelte:component this={page} params={params} />
  </main>

<style>


  header {
    grid-area: header;
    display: flex;
    align-items: center;

    font-family: 'Gloria Hallelujah', cursive;
  }

  header {
    grid-area: logo;
  }

  header img {
    width: 100%;
  }

  nav {
    grid-area: menu;
  }

  main {
    grid-area: main;
    width: 100%;

    background-color: var(--glass-pane-1);
    border-radius: 20px;
    box-sizing: border-box;
    padding: 20px 40px;
  }

  main h1 {
    font-family: 'Gloria Hallelujah', cursive;
    margin-top: 0;
  }
</style>
