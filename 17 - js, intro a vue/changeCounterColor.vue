<!--Single-File Component (SFC)-->

<script>
export default{ 
  data(){
    return {
      count:0,
      counterColor: "black"
    }
  },
  methods: {
    decr(){
      this.count=this.count-1;
      if(this.count<=0){
        this.counterColor="red";
      }
    },
    incr(){
      this.count=this.count+1;
    }
  }
}
</script>

<template>
  <button @click="decr()">-</button>
  <span :style="{color: counterColor}">{{count}}</span>
  <button @click="incr()">+</button>  

</template>

<style scoped>
button {color:red;}
span{color: blue;}
</style>