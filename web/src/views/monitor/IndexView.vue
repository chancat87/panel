<script setup lang="ts">
defineOptions({
  name: 'monitor-index'
})

import { LineChart } from 'echarts/charts'
import {
  DataZoomComponent,
  GridComponent,
  LegendComponent,
  TitleComponent,
  TooltipComponent
} from 'echarts/components'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { NButton } from 'naive-ui'
import VChart from 'vue-echarts'

import monitor from '@/api/panel/monitor'

use([
  CanvasRenderer,
  LineChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent
])

const start = ref(Math.floor(new Date(new Date().setHours(0, 0, 0, 0)).getTime()))
const end = ref(Math.floor(Date.now()))

useRequest(monitor.setting()).onSuccess(({ data }) => {
  monitorSwitch.value = data.enabled
  saveDay.value = data.days
})

const { loading, data } = useWatcher(monitor.list(start.value, end.value), [start, end], {
  initialData: {
    times: [],
    load: {},
    cpu: {},
    mem: {},
    swap: {},
    net: {}
  },
  debounce: [500],
  immediate: true
})

const monitorSwitch = ref(false)
const saveDay = ref(30)

const load = ref<any>({
  title: {
    text: '负载',
    textAlign: 'left',
    textStyle: {
      fontSize: 20
    }
  },
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    align: 'left',
    data: ['1分钟', '5分钟', '15分钟']
  },
  xAxis: [{ type: 'category', boundaryGap: false, data: data.value.times }],
  yAxis: [
    {
      type: 'value'
    }
  ],
  dataZoom: {
    show: true,
    realtime: true,
    start: 0,
    end: 100
  },
  series: [
    {
      name: '1分钟',
      type: 'line',
      smooth: true,
      data: data.value.load.load1,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    },
    {
      name: '5分钟',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.load.load5,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    },
    {
      name: '15分钟',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.load.load15,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    }
  ]
})

const cpu = ref<any>({
  title: {
    text: 'CPU',
    textAlign: 'left',
    textStyle: {
      fontSize: 20
    }
  },
  tooltip: {
    trigger: 'axis'
  },
  xAxis: [{ type: 'category', boundaryGap: false, data: data.value.times }],
  yAxis: [
    {
      name: '单位 %',
      min: 0,
      max: 100,
      type: 'value',
      axisLabel: {
        formatter: '{value} %'
      }
    }
  ],
  dataZoom: {
    show: true,
    realtime: true,
    start: 0,
    end: 100
  },
  series: [
    {
      name: '使用率',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.cpu.percent,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    }
  ]
})

const mem = ref<any>({
  title: {
    text: '内存',
    textAlign: 'left',
    textStyle: {
      fontSize: 20
    }
  },
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    align: 'left',
    data: ['内存', 'Swap']
  },
  xAxis: [{ type: 'category', boundaryGap: false, data: data.value.times }],
  yAxis: [
    {
      name: '单位 MB',
      min: 0,
      max: data.value.mem.total,
      type: 'value',
      axisLabel: {
        formatter: '{value} M'
      }
    }
  ],
  dataZoom: {
    show: true,
    realtime: true,
    start: 0,
    end: 100
  },
  series: [
    {
      name: '内存',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.mem.used,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    },
    {
      name: 'Swap',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.swap.used,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    }
  ]
})

const net = ref<any>({
  title: {
    text: '网络',
    textAlign: 'left',
    textStyle: {
      fontSize: 20
    }
  },
  tooltip: {
    trigger: 'axis'
  },
  legend: {
    align: 'left',
    data: ['总计出', '总计入', '每秒出', '每秒入']
  },
  xAxis: [{ type: 'category', boundaryGap: false, data: data.value.times }],
  yAxis: [
    {
      name: '单位 MB',
      type: 'value',
      axisLabel: {
        formatter: '{value} MB'
      }
    }
  ],
  dataZoom: {
    show: true,
    realtime: true,
    start: 0,
    end: 100
  },
  series: [
    {
      name: '总计出',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.net.sent,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    },
    {
      name: '总计入',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.net.recv,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    },
    {
      name: '每秒出',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.net.tx,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    },
    {
      name: '每秒入',
      type: 'line',
      smooth: true,
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: 'rgba(0, 0, 0, 0.5)'
        }
      },
      data: data.value.net.rx,
      markPoint: {
        data: [
          { type: 'max', name: '最大值' },
          { type: 'min', name: '最小值' }
        ]
      },
      markLine: {
        data: [{ type: 'average', name: '平均值' }]
      }
    }
  ]
})

const handleUpdate = async () => {
  useRequest(monitor.updateSetting(monitorSwitch.value, saveDay.value)).onSuccess(() => {
    window.$message.success('操作成功')
  })
}

const handleClear = async () => {
  useRequest(monitor.clear()).onSuccess(() => {
    window.$message.success('操作成功')
  })
}

// 监听 data 的变化
watch(data, () => {
  load.value.xAxis[0].data = data.value.times
  load.value.series[0].data = data.value.load.load1
  load.value.series[1].data = data.value.load.load5
  load.value.series[2].data = data.value.load.load15
  cpu.value.xAxis[0].data = data.value.times
  cpu.value.series[0].data = data.value.cpu.percent
  mem.value.xAxis[0].data = data.value.times
  mem.value.yAxis[0].max = data.value.mem.total
  mem.value.series[0].data = data.value.mem.used
  mem.value.series[1].data = data.value.swap.used
  net.value.xAxis[0].data = data.value.times
  net.value.series[0].data = data.value.net.sent
  net.value.series[1].data = data.value.net.recv
  net.value.series[2].data = data.value.net.tx
  net.value.series[3].data = data.value.net.rx
})
</script>

<template>
  <common-page show-footer>
    <template #action>
      <n-popconfirm @positive-click="handleClear">
        <template #trigger>
          <n-button type="error">
            <TheIcon :size="18" icon="material-symbols:delete-outline" />
            清除监控记录
          </n-button>
        </template>
        确定要清空吗？
      </n-popconfirm>
    </template>
    <n-card :segmented="true" flex items-center>
      <n-form
        inline
        label-placement="left"
        label-width="auto"
        require-mark-placement="right-hanging"
      >
        <n-flex items-center>
          <n-form-item label="开启监控">
            <n-switch v-model:value="monitorSwitch" @update-value="handleUpdate" />
          </n-form-item>
          <n-form-item label="保存天数">
            <n-input-number v-model:value="saveDay">
              <template #suffix> 天 </template>
            </n-input-number>
          </n-form-item>
          <n-form-item>
            <n-button type="primary" @click="handleUpdate">确定</n-button>
          </n-form-item>
          <n-form-item label="时间选择">
            <n-date-picker v-model:value="start" type="datetime" />
            -
            <n-date-picker v-model:value="end" type="datetime" />
          </n-form-item>
        </n-flex>
      </n-form>
    </n-card>
    <n-grid
      v-if="!loading"
      cols="1 s:1 m:1 l:2 xl:2 2xl:2"
      item-responsive
      responsive="screen"
      pt-20
    >
      <n-gi m-10>
        <n-card :bordered="false" style="height: 40vh">
          <v-chart class="chart" :option="load" autoresize />
        </n-card>
      </n-gi>
      <n-gi m-10>
        <n-card :bordered="false" style="height: 40vh">
          <v-chart class="chart" :option="cpu" autoresize />
        </n-card>
      </n-gi>
      <n-gi m-10>
        <n-card :bordered="false" style="height: 40vh">
          <v-chart class="chart" :option="mem" autoresize />
        </n-card>
      </n-gi>
      <n-gi m-10>
        <n-card :bordered="false" style="height: 40vh">
          <v-chart class="chart" :option="net" autoresize />
        </n-card>
      </n-gi>
    </n-grid>
    <n-skeleton v-else text :repeat="40" />
  </common-page>
</template>

<style scoped lang="scss"></style>
