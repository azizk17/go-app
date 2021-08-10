<template>
  <div class="text-3xl font-bold text-red-500">Home</div>
  <div class=" text-gray-500">
    {{ b }}
  </div>
</template>
<script lang="ts">
import { defineComponent, computed, ref, onMounted } from "vue";
import axios from "axios";
export default defineComponent({
  name: "name",
  props: {},
  components: {},
  setup(props) {
      const b = ref()
      onMounted(async () => {
        b.value = await axios.get("api")
      });
    return {
      b
    };
  },
});
</script>
