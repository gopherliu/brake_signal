<script setup>
import { ref, reactive } from 'vue'
import { reqFetchBindCar, reqPutOnChain } from '../api'

const carVin = ref('')
const carAddress = reactive({
  vin: '',
  public_key: ''
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
    }
  }

  const payload = { time_stamp_nano: Date.now() * 1000, vin: carAddress.vin }
  const onChainRes = await reqPutOnChain(payload)
  console.log(onChainRes)
}
</script>

<template>
  <main class="container">
    <input type="text" v-model="carVin" placeholder="input car vin" />
    <button class="bind-car" @click="bindCarAddressHandle">bind</button>
    <template v-if="carAddress.vin">
      <p>vin: {{ carAddress.vin }}</p>
      <p>address: {{ carAddress.public_key }}</p>
      <p>chain status:{{ onChainStatus }}</p>
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
