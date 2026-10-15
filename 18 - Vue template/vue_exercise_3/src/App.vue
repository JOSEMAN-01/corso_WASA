<script>
  import Stepper from './Stepper.vue'
export default{//Devo rendere Stepper visibile all'interno di template
  components:{
    Stepper
  },
  data(){
    return {
      titles: ['sleepytime','gorilla','museum'],
      nextTitle: ""
    }
  
  },
  methods:{
    stepperAdder(){
      if(this.nextTitle!==""){
        this.titles.push(this.nextTitle);
        this.nextTitle="";
      }
    }
  }
}

</script>

<template>
  <h1> stepper tester </h1>
  <!--<Stepper title="sleepytime"></Stepper>-->
  <Stepper v-for="t in titles" :title="t"></Stepper>
  <!--Provare a modificare-->
  <input type="text" v-model="nextTitle"></input>
  <button @click="stepperAdder()">add stepper</button>
</template>