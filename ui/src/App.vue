<template>
  <div>
    <!-- <img alt="Vue logo" src="./assets/logo.png" /> -->
    <el-button type="primary" icon="el-icon-edit"></el-button>
    <el-button type="primary" icon="el-icon-share"></el-button>
    <el-button type="primary" icon="el-icon-delete"></el-button>
    <el-button type="primary" icon="el-icon-search">搜索</el-button>
    <el-button type="primary"
      >上传<i class="el-icon-upload el-icon--right"></i
    ></el-button>
    <p>
      <router-link to="/">Go to Home</router-link>
      <router-link to="/about">Go to About</router-link>
    </p>
    <el-container>
      <el-aside width="200px">
        <el-menu>
          <el-menu-item>
            <router-link to="/">Go to Home</router-link>
          </el-menu-item>
          <el-menu-item>
            <router-link to="/about">Go to About</router-link>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main><router-view></router-view></el-main>
    </el-container>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import HelloWorld from "./components/HelloWorld.vue";

export default defineComponent({
  name: "App",
  components: {
    HelloWorld,
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
