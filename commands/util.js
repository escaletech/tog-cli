const axios = require('axios').default

const { readConfig } = require('../services/config')

module.exports = {
  config (args) {
    const config = readConfig()
    const namespace = args.namespace || config.namespace
    return { ...config, namespace }
  },
  client () {
    const { host, authToken } = readConfig()
    return axios.create({
      baseURL: host,
      headers: { Authorization: `Bearer ${authToken}` }
    })
  }
}
