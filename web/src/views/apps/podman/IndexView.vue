<script setup lang="ts">
defineOptions({
  name: 'apps-podman-index'
})

import Editor from '@guolao/vue-monaco-editor'
import { NButton, NPopconfirm } from 'naive-ui'

import podman from '@/api/apps/podman'
import systemctl from '@/api/panel/systemctl'

const currentTab = ref('status')
const status = ref(false)
const isEnabled = ref(false)
const registryConfig = ref('')
const storageConfig = ref('')

const statusStr = computed(() => {
  return status.value ? '正常运行中' : '已停止运行'
})

const getStatus = async () => {
  status.value = await systemctl.status('podman')
}

const getIsEnabled = async () => {
  isEnabled.value = await systemctl.isEnabled('podman')
}

const getConfig = async () => {
  registryConfig.value = await podman.registryConfig()
  storageConfig.value = await podman.storageConfig()
}

const handleSaveRegistryConfig = () => {
  useRequest(podman.saveRegistryConfig(registryConfig.value)).onSuccess(() => {
    window.$message.success('保存成功')
  })
}

const handleSaveStorageConfig = () => {
  useRequest(podman.saveStorageConfig(storageConfig.value)).onSuccess(() => {
    window.$message.success('保存成功')
  })
}

const handleStart = async () => {
  await systemctl.start('podman')
  window.$message.success('启动成功')
  await getStatus()
}

const handleStop = async () => {
  await systemctl.stop('podman')
  window.$message.success('停止成功')
  await getStatus()
}

const handleRestart = async () => {
  await systemctl.restart('podman')
  window.$message.success('重启成功')
  await getStatus()
}

const handleIsEnabled = async () => {
  if (isEnabled.value) {
    await systemctl.enable('podman')
    window.$message.success('开启自启动成功')
  } else {
    await systemctl.disable('podman')
    window.$message.success('禁用自启动成功')
  }
  await getIsEnabled()
}

onMounted(() => {
  getStatus()
  getIsEnabled()
  getConfig()
})
</script>

<template>
  <common-page show-footer>
    <template #action>
      <n-button
        v-if="currentTab == 'registryConfig'"
        class="ml-16"
        type="primary"
        @click="handleSaveRegistryConfig"
      >
        <TheIcon :size="18" icon="material-symbols:save-outline" />
        保存
      </n-button>
      <n-button
        v-if="currentTab == 'storageConfig'"
        class="ml-16"
        type="primary"
        @click="handleSaveStorageConfig"
      >
        <TheIcon :size="18" icon="material-symbols:save-outline" />
        保存
      </n-button>
    </template>
    <n-tabs v-model:value="currentTab" type="line" animated>
      <n-tab-pane name="status" tab="运行状态">
        <n-flex vertical>
          <n-alert type="info">
            Podman 是一个无守护进程的容器管理工具，处于停止状态为正常现象且不会影响使用！
          </n-alert>
          <n-card title="运行状态">
            <template #header-extra>
              <n-switch v-model:value="isEnabled" @update:value="handleIsEnabled">
                <template #checked> 自启动开 </template>
                <template #unchecked> 自启动关 </template>
              </n-switch>
            </template>
            <n-space vertical>
              <n-alert :type="status ? 'success' : 'error'">
                {{ statusStr }}
              </n-alert>
              <n-space>
                <n-button type="success" @click="handleStart">
                  <TheIcon :size="24" icon="material-symbols:play-arrow-outline-rounded" />
                  启动
                </n-button>
                <n-popconfirm @positive-click="handleStop">
                  <template #trigger>
                    <n-button type="error">
                      <TheIcon :size="24" icon="material-symbols:stop-outline-rounded" />
                      停止
                    </n-button>
                  </template>
                  确定要停止 Podman 吗？
                </n-popconfirm>
                <n-button type="warning" @click="handleRestart">
                  <TheIcon :size="18" icon="material-symbols:replay-rounded" />
                  重启
                </n-button>
              </n-space>
            </n-space>
          </n-card>
        </n-flex>
      </n-tab-pane>
      <n-tab-pane name="registryConfig" tab="注册表配置">
        <n-space vertical>
          <n-alert type="warning">
            此处修改的是 Podman 注册表配置文件（/etc/containers/registries.conf）
          </n-alert>
          <Editor
            v-model:value="registryConfig"
            language="ini"
            theme="vs-dark"
            height="60vh"
            mt-8
            :options="{
              automaticLayout: true,
              formatOnType: true,
              formatOnPaste: true
            }"
          />
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="storageConfig" tab="存储配置">
        <n-space vertical>
          <n-alert type="warning">
            此处修改的是 Podman 存储配置文件（/etc/containers/storage.conf）
          </n-alert>
          <Editor
            v-model:value="storageConfig"
            language="ini"
            theme="vs-dark"
            height="60vh"
            mt-8
            :options="{
              automaticLayout: true,
              formatOnType: true,
              formatOnPaste: true
            }"
          />
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="run-log" tab="运行日志">
        <realtime-log service="podman" />
      </n-tab-pane>
    </n-tabs>
  </common-page>
</template>
