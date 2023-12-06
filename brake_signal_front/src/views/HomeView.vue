<script setup>
import { ref, reactive } from 'vue'
import { reqFetchBindCar, reqPutOnChain, reqFetchOnChainInfo } from '../api'

const carVin = ref('')
const carAddress = reactive({
  vin: '',
  public_key: ''
})
const onChainInfo = reactive({
  signal_hash: '',
  signal_info: '',
  status: ''
})
const onChainedInfo = reactive({
  last_on_chain_block: '-',
  last_on_chain_hash: '-',
  last_on_chain_info: '-',
  create_at: '-',
  update_at: '-'
})
const onChainStatus = ref('processing')
const bindCarAddressHandle = async () => {
  if (!carVin.value.trim()) {
    alert('plase input car vin')
  }
  console.log(carVin.value)
  const res = await reqFetchBindCar(carVin.value)
  // console.log(res)
  if (res.status === 200) {
    if (!res.data.code) {
      // reset car vin
      carVin.value = ''
      const { public_key, vin } = res.data.data
      carAddress.public_key = public_key
      carAddress.vin = vin
    } else {
      console.log(res.data.err_msg)
    }
  }

  const payload = { time_stamp_nano: String(Date.now() * 1000), vin: carAddress.vin }
  const onChainRes = await reqPutOnChain(payload)
  // console.log(onChainRes)
  if (onChainRes.status === 200) {
    if (!onChainRes.data.code) {
      const { signal_hash, signal_info, status } = onChainRes.data.data
      onChainInfo.signal_hash = signal_hash
      onChainInfo.signal_info = signal_info
      onChainInfo.status = status
    } else {
      console.log(onChainRes.data.err_msg)
    }
  }

  setInterval(() => {
    getOnChainInfo(carAddress.vin)
  }, 5000)
}

const getOnChainInfo = async (vin) => {
  const res = await reqFetchOnChainInfo(vin)
  console.log(res)
  if (res.status === 200) {
    if (!res.data.code) {
      const { last_on_chain_block, last_on_chain_hash, last_on_chain_info } = res.data.data
      onChainedInfo.last_on_chain_block = last_on_chain_block
      onChainedInfo.last_on_chain_hash = last_on_chain_hash
      onChainedInfo.last_on_chain_info = last_on_chain_info
    } else {
      console.log(res.data.err_msg)
    }
  }
}
</script>

<template>
  <main class="container">
    <input type="text" v-model="carVin" placeholder="input car vin" />
    <button class="bind-car" @click="bindCarAddressHandle">bind</button>
    <template v-if="carAddress.vin">
      <p>vin: {{ carAddress.vin }}</p>
      <p>address: {{ carAddress.public_key }}</p>
      <p>signal hash: {{ onChainInfo.signal_hash }}</p>
      <p>status: {{ onChainStatus }}</p>
      <p>last on chain block: {{ onChainedInfo.last_on_chain_block }}</p>
      <p>last on chain hash: {{ onChainedInfo.last_on_chain_hash }}</p>
      <p>last on chain info: {{ onChainedInfo.last_on_chain_info }}</p>
    </template>
  </main>
</template>

<style scoped>
.container .bind-car {
  cursor: pointer;
  margin-left: 5px;
  margin-bottom: 5px;
}
</style>
