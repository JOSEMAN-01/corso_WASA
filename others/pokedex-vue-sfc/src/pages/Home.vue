<template>
  <div class="w-full max-w-4xl">
    <!-- Non capisco il binding di questo :initialId -->
    <!-- Pare che lo si fa per rendere 1 un valore intero, invece che una stringa, che il componente SearchForm non accetterebbe-->
    <SearchForm @search="onSearch" :initialId="1" />
    <!---->
    <div v-if="isLoading" class="mt-4 p-4 bg-yellow-100 border-l-4 border-yellow-500 rounded">
      Caricamento dati Pokémon...
    </div>

    <div v-else-if="error" class="mt-4 p-4 bg-red-600 text-white rounded">
      ERRORE: {{ error }}
    </div>

    <div v-else-if="pokemon" class="mt-6 bg-white p-6 rounded shadow">
      <PokemonSummary :pokemon="pokemon" />
      <!-- La seguente area viene dopo quella già definita nel template del component PokemonSummary-->
      <h3 class="text-xl font-semibold mt-6">Statistiche di Base</h3>
      <div class="grid grid-cols-2 gap-2">
        <!-- Come funziona questa for esattamente?-->
        <StatDetails
          v-for="stat in pokemon.stats"
          :key="stat.stat.name"
          :stat="stat"
        />
        <!-- -->
      </div>
    </div>
  </div>
</template>

<script>
//Quindi nel <template> di sopra, assumiamo che questi import sono stati già fatti
import SearchForm from '../components/SearchForm.vue'
import PokemonSummary from '../components/PokemonSummary.vue'
import StatDetails from '../components/StatDetails.vue'

export default {
  name: 'Home',
  components: { SearchForm, PokemonSummary, StatDetails },
  data() {
    return {
      pokemon: null,
      isLoading: false,
      error: null
    }
  },
  methods: {
    //Notare che id lo prenderemo da
    async onSearch(id) {
      this.isLoading = true
      this.error = null
      this.pokemon = null
      try {
        const r = await fetch(`https://pokeapi.co/api/v2/pokemon/${id}`)
        if (!r.ok) throw new Error("Pokémon non trovato!")
        this.pokemon = await r.json()
      } catch (e) {
        this.error = e.message
      } finally {
        this.isLoading = false
      }
    }
  }
}
</script>
