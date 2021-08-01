<template>
  <Disclosure v-slot="{ open }">
    <DisclosureButton class="py-2">
      <span>Is team pricing available?</span>

      <ChevronRightIcon :class="open ? 'transform rotate-90' : ''" />
    </DisclosureButton>
    <DisclosurePanel> </DisclosurePanel>
  </Disclosure>
</template>

<script >
import { Disclosure, DisclosureButton, DisclosurePanel } from "@headlessui/vue";
import { ChevronRightIcon } from "@heroicons/vue/solid";

export default {
  components: {
    Disclosure,
    DisclosureButton,
    DisclosurePanel,
    ChevronRightIcon,
  },
};
</script>
