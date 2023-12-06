import axios from 'axios'

const baseUrl = '' //`${process.env.VUE_APP_BASE_URL}` // eslint-disable-line
const fetch = axios.create({
  baseURL: baseUrl,
  timeout: 5000
  // headers: {}
})

export const reqFetchBindCar = (vin) => fetch.get(`/api/v1/pairs/${vin}`)
export const reqPutOnChain = (data) => fetch.post('/api/v1/onChain', data)
export const reqFetchOnChainInfo = (vin) => fetch.get(`/api/v1/getOnChainInfo/${vin}`)
