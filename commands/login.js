const opn = require('opn')
const uuid = require('uuid/v4')
const EventSource = require('eventsource')
const signale = require('signale')

const { readConfig, saveConfig } = require('../services/config')

module.exports = function login (args) {
  const config = readConfig()
  const host = config.host || args.host
  if (!host) {
    signale.error('No suitable host')
    return
  }

  const baseUrl = host.startsWith('http://') || host.startsWith('https://')
    ? host
    : `https://${host}`

  const token = uuid()
  const target = `${baseUrl}/auth/login?rd=/auth/cli-return&cli_token=${token}`

  opn(target)
    .then(() => signale.start(`Opening ${target}`))
    .then(() => signale.pending(`Waiting for authentication to complete`))
    .then(() => {
      const es = new EventSource(`${baseUrl}/auth/cli-notify/${token}`)
      es.onerror = e => signale.warn(e)
      es.onmessage = ({ data }) => {
        const { authToken } = JSON.parse(data)
        saveConfig({ host, authToken })
        signale.success('Login complete')
        es.close()
      }
    })
}
