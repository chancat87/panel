<script setup lang="ts">
import ws from '@/api/ws'
import type { LogInst } from 'naive-ui'

const props = defineProps({
  path: {
    type: String,
    required: false
  },
  service: {
    type: String,
    required: false
  }
})

const log = ref('')
const logRef = ref<LogInst | null>(null)
let logWs: WebSocket | null = null

const init = async () => {
  let cmd = ''
  if (props.path) {
    cmd = `tail -n 100 -f '${props.path}'`
  } else if (props.service) {
    cmd = `journalctl -u '${props.service}' -f`
  } else {
    window.$message.error('path 或 service 不能为空')
    return
  }
  ws.exec(cmd)
    .then((ws: WebSocket) => {
      logWs = ws
      ws.onmessage = (event) => {
        log.value += event.data + '\n'
        const lines = log.value.split('\n')
        if (lines.length > 500) {
          log.value = lines.slice(lines.length - 500).join('\n')
        }
      }
    })
    .catch(() => {
      window.$message.error('获取日志流失败')
    })
}

const close = () => {
  if (logWs) {
    logWs.close()
  }
  log.value = ''
}

watch(
  () => props.path,
  () => {
    close()
    init()
  }
)

watchEffect(() => {
  if (log.value) {
    nextTick(() => {
      logRef.value?.scrollTo({ position: 'bottom', silent: true })
    })
  }
})

onMounted(() => {
  init()
})

onUnmounted(() => {
  close()
})

defineExpose({
  init
})
</script>

<template>
  <n-log ref="logRef" :log="log" trim :rows="40" />
</template>

<style scoped lang="scss"></style>
