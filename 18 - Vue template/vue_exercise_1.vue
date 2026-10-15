<script>
export default{ 
  data(){
    return {
      count:0,

    }
  },
  methods: {
    decr(){
      this.count=this.count-1;

    },
    incr(){
      this.count=this.count+1;
    }
  },
  computed:{
    counterStyle(){//pensare al getter e setter JS
      let col= this.count<0?"color:red;":"color:black;";
      let siz="font-size: "+Math.max(0.5, Math.abs(this.count))+'em';
      return col+siz
    }
  }

}
</script>

<template>
  <button @click="decr()">-</button>
  <span :style=counterStyle>{{count}}</span>
  <button @click="incr()">+</button>  

</template>

<style scoped>
button {color:red;}
span{color: blue;}
</style>