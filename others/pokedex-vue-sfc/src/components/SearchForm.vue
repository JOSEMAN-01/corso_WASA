<template>
  <div class="flex space-x-2 mb-4">
    <!-- Notare binding bidirezionale di localId -->
    <!-- the form input elements update the Vue data instance, and a change in the Vue instance data updates the inputs -->
     <!-- v-model.number è una v-model che in più converte l'input in un numero-->
    <input type="number" v-model.number="localId" class="border p-2 rounded" />
    <!---->
    <button @click="emitSearch" class="bg-indigo-600 text-white px-4 py-2 rounded">
      Cerca
    </button>
  </div>
</template>

<script>
export default {
  name: "SearchForm",
  props: { initialId: { type: Number, default: 1 } },
  data() {
    return { localId: this.initialId }
  },
  methods: {
    emitSearch() {
      //Notare l'uso di $emit, che è un metodo built-in di vue per ativare eventi personalizzati
      //E' come se definessimo una proprietà che deve chiamare una funzione, e gli paserà come argomento localId
      this.$emit("search", this.localId)
    }
  }
}
</script>
