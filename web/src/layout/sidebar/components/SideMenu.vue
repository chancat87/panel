<script lang="ts" setup>
import { translateTitle } from '@/locales/menu'
import { usePermissionStore, useTabStore, useThemeStore } from '@/store'
import { isUrl, renderIcon } from '@/utils'

import type { MenuInst, MenuOption } from 'naive-ui'
import type { VNodeChild } from 'vue'
import type { Meta, RouteType } from '~/types/router'

const router = useRouter()
const currentRoute = useRoute()
const permissionStore = usePermissionStore()
const themeStore = useThemeStore()
const tabStore = useTabStore()

const menu = ref<MenuInst>()
watch(currentRoute, async () => {
  await nextTick()
  menu.value?.showOption()
})

const menuOptions = computed(() => {
  return permissionStore.menus.map((item) => getMenuItem(item))
})

function resolvePath(basePath: string, path: string) {
  if (isUrl(path)) return path
  return `/${[basePath, path]
    .filter((path) => !!path && path !== '/')
    .map((path) => path.replace(/(^\/)|(\/$)/g, ''))
    .join('/')}`
}

type MenuItem = MenuOption & {
  label: string
  key: string
  path: string
  children?: Array<MenuItem>
}

function getMenuItem(route: RouteType, basePath = ''): MenuItem {
  let menuItem: MenuItem = {
    label: route.meta?.title ? translateTitle(route.meta.title) : route.name,
    key: route.name,
    path: resolvePath(basePath, route.path),
    icon: getIcon(route.meta)
  }

  const visibleChildren = route.children
    ? route.children.filter(
        (item: RouteType) =>
          item.name && !item.isHidden && !permissionStore.hiddenRoutes.includes(item.name)
      )
    : []

  if (!visibleChildren.length) return menuItem

  if (visibleChildren.length === 1) {
    // 单个子路由处理
    const singleRoute = visibleChildren[0]
    menuItem = {
      label: singleRoute.meta?.title ? translateTitle(singleRoute.meta.title) : singleRoute.name,
      key: singleRoute.name,
      path: resolvePath(menuItem.path, singleRoute.path),
      icon: getIcon(singleRoute.meta)
    }
    const visibleItems = singleRoute.children
      ? singleRoute.children.filter((item: RouteType) => item.name && !item.isHidden)
      : []

    if (visibleItems.length === 1) menuItem = getMenuItem(visibleItems[0], menuItem.path)
    else if (visibleItems.length > 1)
      menuItem.children = visibleItems.map((item) => getMenuItem(item, menuItem.path))
  } else {
    menuItem.children = visibleChildren.map((item) => getMenuItem(item, menuItem.path))
  }

  return menuItem
}

function getIcon(meta?: Meta): (() => VNodeChild) | undefined {
  if (meta?.icon) return renderIcon(meta.icon, { size: 14, class: `${meta.icon} text-14` })
  return undefined
}

function handleMenuSelect(key: string, item: MenuOption) {
  const menuItem = item as MenuItem & MenuOption
  if (isUrl(menuItem.path)) {
    window.open(menuItem.path)
    return
  }
  if (menuItem.path === currentRoute.path && !currentRoute.meta?.keepAlive)
    tabStore.reloadTab(currentRoute.path)
  else router.push(menuItem.path)

  // 手机端自动收起菜单
  themeStore.isMobile && themeStore.setCollapsed(true)
}
</script>

<template>
  <n-menu
    ref="menu"
    :collapsed-icon-size="22"
    :collapsed-width="64"
    :indent="18"
    :options="menuOptions"
    :value="currentRoute.name as string"
    accordion
    class="side-menu"
    @update:value="handleMenuSelect"
  />
</template>

<style lang="scss">
.side-menu {
  .n-menu-item-content__icon {
    border: 1px solid rgb(229, 231, 235);
    border-radius: 4px;
  }

  .n-menu-item-content--child-active,
  .n-menu-item-content--selected {
    .n-menu-item-content__icon {
      border-color: var(--primary-color);
      background-color: var(--primary-color);

      i {
        color: #fff;
      }
    }
  }
}
</style>
