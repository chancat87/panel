<script setup lang="ts">
defineOptions({
  name: 'apps-fail2ban-index'
})

import { NButton, NDataTable, NInput, NPopconfirm, NSwitch } from 'naive-ui'

import fail2ban from '@/api/apps/fail2ban'
import app from '@/api/panel/app'
import systemctl from '@/api/panel/systemctl'
import website from '@/api/panel/website'
import { renderIcon } from '@/utils'

const currentTab = ref('status')
const status = ref(false)
const isEnabled = ref(false)
const white = ref('')

const addJailModal = ref(false)
const addJailModel = ref({
  name: 'ssh',
  type: 'website',
  maxretry: 30,
  findtime: 300,
  bantime: 600,
  website_name: '',
  website_mode: 'cc',
  website_path: '/'
})

const jailModal = ref(false)
const jailCurrentlyBan = ref(0)
const jailTotalBan = ref(0)
const jailBanedList = ref<any[]>([])

const statusType = computed(() => {
  return status.value ? 'success' : 'error'
})
const statusStr = computed(() => {
  return status.value ? '正常运行中' : '已停止运行'
})

const jailsColumns: any = [
  {
    title: '名称',
    key: 'name',
    minWidth: 250,
    ellipsis: { tooltip: true }
  },
  {
    title: '状态',
    key: 'enabled',
    minWidth: 60,
    align: 'center',
    render(row: any) {
      return h(NSwitch, {
        size: 'small',
        rubberBand: false,
        value: row.enabled,
        disabled: true
      })
    }
  },
  { title: '最大尝试', key: 'max_retry', minWidth: 150, ellipsis: { tooltip: true } },
  { title: '封禁时间', key: 'ban_time', minWidth: 150, ellipsis: { tooltip: true } },
  { title: '周期', key: 'find_time', minWidth: 150, ellipsis: { tooltip: true } },
  {
    title: '操作',
    key: 'actions',
    width: 280,
    align: 'center',
    hideInExcel: true,
    render(row: any) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'warning',
            secondary: true,
            onClick: async () => {
              await getJailInfo(row.name)
              jailModal.value = true
            }
          },
          {
            default: () => '查看',
            icon: renderIcon('material-symbols:visibility', { size: 14 })
          }
        ),
        h(
          NPopconfirm,
          {
            onPositiveClick: () => handleDeleteJail(row.name)
          },
          {
            default: () => {
              return '确定删除规则' + row.name + '吗？'
            },
            trigger: () => {
              return h(
                NButton,
                {
                  size: 'small',
                  type: 'error',
                  style: 'margin-left: 15px'
                },
                {
                  default: () => '删除',
                  icon: renderIcon('material-symbols:delete-outline', { size: 14 })
                }
              )
            }
          }
        )
      ]
    }
  }
]

const banedIPColumns: any = [
  {
    title: 'IP',
    key: 'ip',
    minWidth: 200,
    resizable: true,
    ellipsis: { tooltip: true }
  },
  {
    title: '操作',
    key: 'actions',
    width: 100,
    align: 'center',
    hideInExcel: true,
    render(row: any) {
      return [
        h(
          NPopconfirm,
          {
            onPositiveClick: () => handleUnBan(row.name, row.ip)
          },
          {
            default: () => {
              return '确定解封' + row.ip + '吗？'
            },
            trigger: () => {
              return h(
                NButton,
                {
                  size: 'small',
                  type: 'error'
                },
                {
                  default: () => '解封',
                  icon: renderIcon('material-symbols:delete-outline', { size: 14 })
                }
              )
            }
          }
        )
      ]
    }
  }
]

const websites = ref<any[]>([])

const getWhiteList = async () => {
  white.value = await fail2ban.whitelist()
}

const handleSaveWhiteList = () => {
  useRequest(fail2ban.setWhitelist(white.value)).onSuccess(() => {
    window.$message.success('保存成功')
  })
}

const getWebsiteList = async (page: number, limit: number) => {
  const data = await website.list(page, limit)
  for (const item of data.items) {
    websites.value.push({
      label: item.name,
      value: item.name
    })
  }
  addJailModel.value.website_name = websites.value[0]?.value
}

const { loading, data, page, total, pageSize, pageCount, refresh } = usePagination(
  (page, pageSize) => fail2ban.jails(page, pageSize),
  {
    initialData: { total: 0, list: [] },
    initialPageSize: 20,
    total: (res: any) => res.total,
    data: (res: any) => res.items
  }
)

const getStatus = async () => {
  status.value = await systemctl.status('fail2ban')
}

const getIsEnabled = async () => {
  isEnabled.value = await systemctl.isEnabled('fail2ban')
}

const handleStart = async () => {
  await systemctl.start('fail2ban')
  window.$message.success('启动成功')
  await getStatus()
}

const handleIsEnabled = async () => {
  if (isEnabled.value) {
    await systemctl.enable('fail2ban')
    window.$message.success('开启自启动成功')
  } else {
    await systemctl.disable('fail2ban')
    window.$message.success('禁用自启动成功')
  }
  await getIsEnabled()
}

const handleStop = async () => {
  await systemctl.stop('fail2ban')
  window.$message.success('停止成功')
  await getStatus()
}

const handleRestart = async () => {
  await systemctl.restart('fail2ban')
  window.$message.success('重启成功')
  await getStatus()
}

const handleReload = async () => {
  await systemctl.reload('fail2ban')
  window.$message.success('重载成功')
  await getStatus()
}

const handleAddJail = () => {
  useRequest(fail2ban.add(addJailModel.value)).onSuccess(() => {
    refresh()
    window.$message.success('添加成功')
    addJailModal.value = false
  })
}

const handleDeleteJail = (name: string) => {
  useRequest(fail2ban.delete(name)).onSuccess(() => {
    refresh()
    window.$message.success('删除成功')
  })
}

const getJailInfo = async (name: string) => {
  const data = await fail2ban.jail(name)
  jailCurrentlyBan.value = data.currently_ban
  jailTotalBan.value = data.total_ban
  jailBanedList.value = data.baned_list
}

const handleUnBan = (name: string, ip: string) => {
  useRequest(fail2ban.unban(name, ip)).onSuccess(() => {
    window.$message.success('解封成功')
    getJailInfo(name)
  })
}

onMounted(() => {
  refresh()
  getStatus()
  getIsEnabled()
  getWhiteList()
  useRequest(app.isInstalled('nginx')).onSuccess(({ data }) => {
    if (data.installed) {
      getWebsiteList(1, 10000)
    }
  })
})
</script>

<template>
  <common-page show-footer>
    <template #action>
      <n-button
        v-if="currentTab == 'status'"
        class="ml-16"
        type="primary"
        @click="handleSaveWhiteList"
      >
        <TheIcon :size="18" icon="material-symbols:save-outline" />
        保存白名单
      </n-button>
      <n-button
        v-if="currentTab == 'jails'"
        class="ml-16"
        type="primary"
        @click="addJailModal = true"
      >
        <TheIcon :size="18" icon="material-symbols:add" />
        添加规则
      </n-button>
    </template>
    <n-tabs v-model:value="currentTab" type="line" animated>
      <n-tab-pane name="status" tab="运行状态">
        <n-space vertical>
          <n-card title="运行状态">
            <template #header-extra>
              <n-switch v-model:value="isEnabled" @update:value="handleIsEnabled">
                <template #checked> 自启动开 </template>
                <template #unchecked> 自启动关 </template>
              </n-switch>
            </template>
            <n-space vertical>
              <n-alert :type="statusType">
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
                  停止 Fail2ban 会导致所有规则失效，确定停止吗？
                </n-popconfirm>
                <n-button type="warning" @click="handleRestart">
                  <TheIcon :size="18" icon="material-symbols:replay-rounded" />
                  重启
                </n-button>
                <n-button type="primary" @click="handleReload">
                  <TheIcon :size="20" icon="material-symbols:refresh-rounded" />
                  重载
                </n-button>
              </n-space>
            </n-space>
          </n-card>
          <n-card title="IP 白名单">
            <n-input
              v-model:value="white"
              type="textarea"
              autosize
              placeholder="IP 白名单，以英文逗号,分隔"
            />
          </n-card>
        </n-space>
      </n-tab-pane>
      <n-tab-pane name="jails" tab="规则管理">
        <n-card title="规则列表" :segmented="true">
          <n-data-table
            striped
            remote
            :scroll-x="1000"
            :loading="loading"
            :columns="jailsColumns"
            :data="data"
            :row-key="(row: any) => row.name"
            v-model:page="page"
            v-model:pageSize="pageSize"
            :pagination="{
              page: page,
              pageCount: pageCount,
              pageSize: pageSize,
              itemCount: total,
              showQuickJumper: true,
              showSizePicker: true,
              pageSizes: [20, 50, 100, 200]
            }"
          />
        </n-card>
      </n-tab-pane>
      <n-tab-pane name="run-log" tab="运行日志">
        <realtime-log service="fail2ban" />
      </n-tab-pane>
    </n-tabs>
  </common-page>
  <n-modal v-model:show="addJailModal" title="添加规则">
    <n-card closable @close="() => (addJailModal = false)" title="添加规则" style="width: 60vw">
      <n-space vertical>
        <n-alert type="info">
          在设置周期内(秒)有超过最大重试(次)的IP访问，将禁止该IP禁止时间(秒)
        </n-alert>
        <n-alert type="warning">
          防护端口自动获取，如果修改了规则项对应的端口，请删除重新添加，否则防护可能不会生效
        </n-alert>

        <n-form :model="addJailModel">
          <n-form-item label="类型">
            <n-select
              v-model:value="addJailModel.type"
              :options="[
                { label: '网站', value: 'website' },
                { label: '服务', value: 'service' }
              ]"
            >
            </n-select>
          </n-form-item>
          <n-form-item v-if="addJailModel.type === 'website'" label="选择网站">
            <n-select
              v-model:value="addJailModel.website_name"
              :options="websites"
              placeholder="选择网站"
            />
          </n-form-item>
          <n-form-item v-if="addJailModel.type === 'website'" label="保护模式">
            <n-select
              v-model:value="addJailModel.website_mode"
              :options="[
                { label: 'CC', value: 'cc' },
                { label: '路径', value: 'path' }
              ]"
            >
            </n-select>
          </n-form-item>
          <n-form-item
            v-if="addJailModel.type === 'website' && addJailModel.website_mode === 'path'"
            label="保护路径"
          >
            <n-input v-model:value="addJailModel.website_path" placeholder="保护路径" />
          </n-form-item>
          <n-form-item v-if="addJailModel.type === 'service'" label="服务">
            <n-select
              v-model:value="addJailModel.name"
              :options="[
                { label: 'SSH', value: 'ssh' },
                { label: 'MySQL', value: 'mysql' },
                { label: 'Pure-Ftpd', value: 'pure-ftpd' }
              ]"
            >
            </n-select>
          </n-form-item>
          <n-form-item path="maxretry" label="最大尝试">
            <n-input-number v-model:value="addJailModel.maxretry" @keydown.enter.prevent :min="1" />
          </n-form-item>
          <n-form-item path="findtime" label="周期">
            <n-input-number v-model:value="addJailModel.findtime" @keydown.enter.prevent :min="1" />
          </n-form-item>
          <n-form-item path="bantime" label="禁止时间">
            <n-input-number v-model:value="addJailModel.bantime" @keydown.enter.prevent :min="1" />
          </n-form-item>
        </n-form>
        <n-button type="info" block @click="handleAddJail">提交</n-button>
      </n-space>
    </n-card>
  </n-modal>
  <n-modal v-model:show="jailModal" title="查看规则">
    <n-card closable @close="() => (jailModal = false)" title="查看规则" style="width: 60vw">
      <n-space vertical>
        <n-card title="规则信息" :segmented="true">
          <n-space vertical>
            <n-space>
              <n-text>当前封禁</n-text>
              <n-text>{{ jailCurrentlyBan }}</n-text>
            </n-space>
            <n-space>
              <n-text>总封禁</n-text>
              <n-text>{{ jailTotalBan }}</n-text>
            </n-space>
          </n-space>
        </n-card>
        <n-card title="封禁列表" :segmented="true">
          <n-data-table
            striped
            remote
            :scroll-x="300"
            :loading="false"
            :columns="banedIPColumns"
            :data="jailBanedList"
            :row-key="(row: any) => row.ip"
            :pagination="false"
          />
        </n-card>
      </n-space>
    </n-card>
  </n-modal>
</template>
