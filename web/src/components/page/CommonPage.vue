<script lang="ts" setup>
import { translateTitle } from '@/locales/menu'

interface Props {
  showFooter?: boolean
  showHeader?: boolean
  title?: string
}

withDefaults(defineProps<Props>(), {
  showFooter: false,
  showHeader: true,
  title: undefined
})

const route = useRoute()
</script>

<template>
  <app-page :show-footer="showFooter">
    <header v-if="showHeader" mb-15 min-h-45 flex items-center justify-between px-15>
      <slot v-if="$slots.header" name="header" />
      <template v-else>
        <div flex items-center>
          <slot v-if="$slots['title-prefix']" name="title-prefix" />
          <div mr-12 h-16 w-4 rounded-l-2 bg-primary></div>
          <h2 font-normal>
            {{ title ? title : route.meta.title ? translateTitle(route.meta.title) : '' }}
          </h2>
          <slot name="title-suffix" />
        </div>
        <slot name="action" />
      </template>
    </header>

    <n-card flex-1>
      <slot />
    </n-card>
  </app-page>
</template>
