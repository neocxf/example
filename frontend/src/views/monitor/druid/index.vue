<template>
  <i-frame :src="url" />
</template>
<script>
import iFrame from "@/components/iFrame/index";
export default {
  name: "Druid",
  components: { iFrame },
  data() {
    return {
      url: import.meta.env.APP_BASE_API + "/druid/login.html"
    };
  },
};
</script>
