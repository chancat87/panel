<script setup lang="ts">
import { NButton, NDataTable, NInput, NPopconfirm, NSpace, NTag } from 'naive-ui'

import cert from '@/api/panel/cert'

const props = defineProps({
  dnsProviders: {
    type: Array<any>,
    required: true
  }
})

const { dnsProviders } = toRefs(props)

const updateDNSModel = ref<any>({
  data: {
    ak: '',
    sk: ''
  },
  type: 'aliyun',
  name: ''
})
const updateDNSModal = ref(false)
const updateDNS = ref<any>()

const columns: any = [
  {
    title: '备注名称',
    key: 'name',
    minWidth: 200,
    resizable: true,
    ellipsis: { tooltip: true }
  },
  {
    title: '类型',
    key: 'type',
    width: 150,
    resizable: true,
    ellipsis: { tooltip: true },
    render(row: any) {
      return h(
        NTag,
        {
          type: 'info',
          bordered: false
        },
        {
          default: () => {
            const provider = dnsProviders.value.find((provider: any) => provider.value === row.type)
            if (provider) {
              return provider.label
            } else {
              return '未知'
            }
          }
        }
      )
    }
  },
  {
    title: '操作',
    key: 'actions',
    width: 200,
    align: 'center',
    hideInExcel: true,
    render(row: any) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            onClick: () => {
              updateDNS.value = row.id
              updateDNSModel.value.data.ak = row.dns_param.ak
              updateDNSModel.value.data.sk = row.dns_param.sk
              updateDNSModel.value.type = row.type
              updateDNSModel.value.name = row.name
              updateDNSModal.value = true
            }
          },
          {
            default: () => '修改'
          }
        ),
        h(
          NPopconfirm,
          {
            onPositiveClick: async () => {
              useRequest(cert.dnsDelete(row.id)).onSuccess(() => {
                refresh()
                window.$message.success('删除成功')
              })
            }
          },
          {
            default: () => {
              return '确定删除 DNS 吗？'
            },
            trigger: () => {
              return h(
                NButton,
                {
                  size: 'small',
                  type: 'error',
                  style: 'margin-left: 15px;'
                },
                {
                  default: () => '删除'
                }
              )
            }
          }
        )
      ]
    }
  }
]

const { loading, data, page, total, pageSize, pageCount, refresh } = usePagination(
  (page, pageSize) => cert.dns(page, pageSize),
  {
    initialData: { total: 0, list: [] },
    initialPageSize: 20,
    total: (res: any) => res.total,
    data: (res: any) => res.items
  }
)

const handleUpdateDNS = () => {
  useRequest(cert.dnsUpdate(updateDNS.value, updateDNSModel.value)).onSuccess(() => {
    refresh()
    updateDNSModal.value = false
    updateDNSModel.value.data.ak = ''
    updateDNSModel.value.data.sk = ''
    updateDNSModel.value.name = ''
    window.$message.success('更新成功')
  })
}

onMounted(() => {
  refresh()
  window.$bus.on('cert:refresh-dns', () => {
    refresh()
  })
})

onUnmounted(() => {
  window.$bus.off('cert:refresh-dns')
})
</script>

<template>
  <n-space vertical size="large">
    <n-data-table
      striped
      remote
      :scroll-x="1000"
      :loading="loading"
      :columns="columns"
      :data="data"
      :row-key="(row: any) => row.id"
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
  </n-space>
  <n-modal
    v-model:show="updateDNSModal"
    preset="card"
    title="修改 DNS"
    style="width: 60vw"
    size="huge"
    :bordered="false"
    :segmented="false"
  >
    <n-space vertical>
      <n-form :model="updateDNSModel">
        <n-form-item path="name" label="备注名称">
          <n-input v-model:value="updateDNSModel.name" type="text" placeholder="输入备注名称" />
        </n-form-item>
        <n-form-item path="type" label="DNS">
          <n-select
            v-model:value="updateDNSModel.type"
            placeholder="选择 DNS"
            clearable
            :options="dnsProviders"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'aliyun'" path="ak" label="Access Key">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入阿里云 Access Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'aliyun'" path="sk" label="Secret Key">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入阿里云 Secret Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'tencent'" path="ak" label="SecretId">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入腾讯云 SecretId"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'tencent'" path="sk" label="SecretKey">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入腾讯云 SecretKey"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'huawei'" path="ak" label="AccessKeyId">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入华为云 AccessKeyId"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'huawei'" path="sk" label="SecretAccessKey">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入华为云 SecretAccessKey"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'westcn'" path="sk" label="Username">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入西部数码 Username"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'westcn'" path="ak" label="API Password">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入西部数码 API Password"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'cloudflare'" path="ak" label="API Key">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Cloudflare API Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'godaddy'" path="ak" label="Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 GoDaddy Token"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'gcore'" path="ak" label="API Key">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 G-Core API Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'porkbun'" path="ak" label="API Key">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Porkbun API Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'porkbun'" path="sk" label="Secret Key">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入 Porkbun Secret Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'namecheap'" path="sk" label="API Username">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入 Namecheap API Username"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'namecheap'" path="ak" label="API Key">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Namecheap API Key"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'namesilo'" path="ak" label="API Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 NameSilo API Token"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'namecom'" path="sk" label="Username">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入 Name.com Username"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'namecom'" path="ak" label="Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Name.com Token"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'cloudns'" path="ak" label="Auth ID">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 ClouDNS Auth ID（使用Sub Auth ID请添加sub-前缀）"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'cloudns'" path="sk" label="Auth Password">
          <n-input
            v-model:value="updateDNSModel.data.sk"
            type="text"
            placeholder="输入 ClouDNS Auth Password"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'duckdns'" path="ak" label="Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Duck DNS Token"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'hetzner'" path="ak" label="Auth API Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Hetzner Auth API Token"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'linode'" path="ak" label="Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Linode Token"
          />
        </n-form-item>
        <n-form-item v-if="updateDNSModel.type == 'vercel'" path="ak" label="Token">
          <n-input
            v-model:value="updateDNSModel.data.ak"
            type="text"
            placeholder="输入 Vercel Token"
          />
        </n-form-item>
      </n-form>
      <n-button type="info" block @click="handleUpdateDNS">提交</n-button>
    </n-space>
  </n-modal>
</template>

<style scoped lang="scss"></style>
