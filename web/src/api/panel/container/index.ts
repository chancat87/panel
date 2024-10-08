import { request } from '@/utils'
import type { AxiosResponse } from 'axios'

export default {
  // 获取容器列表
  containerList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get('/container/list', { params: { page, limit } }),
  // 添加容器
  containerCreate: (config: any): Promise<AxiosResponse<any>> =>
    request.post('/container/create', config),
  // 删除容器
  containerRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/remove`, { id }),
  // 启动容器
  containerStart: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/start`, { id }),
  // 停止容器
  containerStop: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/stop`, { id }),
  // 重启容器
  containerRestart: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/restart`, { id }),
  // 暂停容器
  containerPause: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/pause`, { id }),
  // 恢复容器
  containerUnpause: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/unpause`, { id }),
  // 获取容器详情
  containerDetail: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/detail`, { params: { id } }),
  // 杀死容器
  containerKill: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/kill`, { id }),
  // 重命名容器
  containerRename: (id: string, name: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/rename`, { id, name }),
  // 获取容器状态
  containerStats: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/stats`, { params: { id } }),
  // 检查容器是否存在
  containerExist: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/exist`, { params: { id } }),
  // 获取容器日志
  containerLogs: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/logs`, { params: { id } }),
  // 清理容器
  containerPrune: (): Promise<AxiosResponse<any>> => request.post(`/container/prune`),
  // 获取网络列表
  networkList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get(`/container/network/list`, { params: { page, limit } }),
  // 创建网络
  networkCreate: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/network/create`, config),
  // 删除网络
  networkRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/network/remove`, { id }),
  // 检查网络是否存在
  networkExist: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/network/exist`, { params: { id } }),
  // 获取网络详情
  networkInspect: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/network/inspect`, { params: { id } }),
  // 连接容器到网络
  networkConnect: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/network/connect`, config),
  // 断开容器到网络的连接
  networkDisconnect: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/network/disconnect`, config),
  // 清理网络
  networkPrune: (): Promise<AxiosResponse<any>> => request.post(`/container/network/prune`),
  // 获取镜像列表
  imageList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get(`/container/image/list`, { params: { page, limit } }),
  // 检查镜像是否存在
  imageExist: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/image/exist`, { params: { id } }),
  // 拉取镜像
  imagePull: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/image/pull`, config),
  // 删除镜像
  imageRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/image/remove`, { id }),
  // 获取镜像详情
  imageInspect: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/image/inspect`, { params: { id } }),
  // 清理镜像
  imagePrune: (): Promise<AxiosResponse<any>> => request.post(`/container/image/prune`),
  // 获取卷列表
  volumeList: (page: number, limit: number): Promise<AxiosResponse<any>> =>
    request.get(`/container/volume/list`, { params: { page, limit } }),
  // 创建卷
  volumeCreate: (config: any): Promise<AxiosResponse<any>> =>
    request.post(`/container/volume/create`, config),
  // 检查卷是否存在
  volumeExist: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/volume/exist`, { params: { id } }),
  // 删除卷
  volumeRemove: (id: string): Promise<AxiosResponse<any>> =>
    request.post(`/container/volume/remove`, { id }),
  // 获取卷详情
  volumeInspect: (id: string): Promise<AxiosResponse<any>> =>
    request.get(`/container/volume/inspect`, { params: { id } }),
  // 清理卷
  volumePrune: (): Promise<AxiosResponse<any>> => request.post(`/container/volume/prune`)
}
