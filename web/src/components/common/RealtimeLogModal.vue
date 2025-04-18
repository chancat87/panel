<script setup lang="ts">
import ws from '@/api/ws'
import type { LogInst } from 'naive-ui'
import { useGettext } from 'vue3-gettext'

const { $gettext } = useGettext()
const show = defineModel<boolean>('show', { type: Boolean, required: true })
const props = defineProps({
  path: {
    type: String,
    required: true
  }
})

const log = ref('')
const logRef = ref<LogInst | null>(null)
let logWs: WebSocket | null = null

const init = async () => {
  const cmd = `tail -n 100 -f '${props.path}'`
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
      window.$message.error($gettext('Failed to get log stream'))
    })
}

const handleClose = () => {
  if (logWs) {
    logWs.close()
  }
  log.value = ''
}

watch([() => props.path, () => show.value], () => {
  handleClose()
  if (show.value) {
    init()
  }
})
watchEffect(() => {
  if (log.value) {
    nextTick(() => {
      logRef.value?.scrollTo({ position: 'bottom', silent: true })
    })
  }
})

defineExpose({
  init
})
</script>

<template>
  <n-modal
    v-model:show="show"
    preset="card"
    :title="$gettext('Logs')"
    style="width: 80vw"
    size="huge"
    :bordered="false"
    :segmented="false"
    @close="handleClose"
    @mask-click="handleClose"
  >
    <n-log ref="logRef" :log="log" trim :rows="40" />
  </n-modal>
</template>

<style scoped lang="scss"></style>
