<template>
  <div>
    <h1>about11111</h1>
    <ul>
      <li :key="index" v-for="(item, index) in list">{{ item }}</li>
    </ul>
  </div>
</template>

<script lang="ts">
import { computed } from "@vue/runtime-core";
import axios from "axios";

import { ref, defineComponent } from "vue";
export default defineComponent({
  name: "HelloWorld",
  props: {
    // msg: {
    //   type: String,
    //   required: true,
    // },
  },
  setup: () => {
    const count = ref(0);
    return { count };
  },
  data() {
    var list: string[];
    axios
      .get("/getfiles")
      .then((response) => {
        this.list = response.data.list;
      })
      .catch((error) => {
        console.error(error);
      });
    return { list };
  },
});
</script>

<style scoped>
a {
  color: #42b983;
}

label {
  margin: 0 0.5em;
  font-weight: bold;
}

code {
  background-color: #eee;
  padding: 2px 4px;
  border-radius: 4px;
  color: #304455;
}
</style>
