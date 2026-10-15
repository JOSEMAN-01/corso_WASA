<script>
  import Stepper from './Stepper.vue'
  export default{//Devo rendere Stepper visibile all'interno di template
    components:{
      Stepper
    },
    data(){
      return {
        titles: ['sleepytime','gorilla','museum']
      }
    }
  }

</script>

<template>
  <h1> stepper tester </h1>
  <!--<Stepper title="sleepytime"></Stepper>-->
  <Stepper v-for="t in titles" :title="t"></Stepper>
  <!--Provare a modificare-->
</template>