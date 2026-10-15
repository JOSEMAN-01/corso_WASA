<template>
  <div class="min-h-screen p-6 flex flex-col items-center">
    <h1 class="text-4xl font-bold text-indigo-700 mb-8">
      Pokedex (Vue SFC - Options API)
    </h1>
    <!--Notare router view-->
    <!-- displays the component or template that corresponds to the current URL (routerlink)-->
    <router-view />
    <!---->
  </div>
</template>

<script>
// Come fa Vue ad "esportare" il template sopra, quando passo soltanto questo oggetto inutile?
export default { name: 'App' }
</script>
