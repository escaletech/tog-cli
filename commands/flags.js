const { inspect } = require('util')
const signale = require('signale')
const axios = require('axios').default

const { readConfig } = require('../services/config')

const format = value => inspect(value, false, 1, true)

const client = () => {
  const { host, authToken } = readConfig()
  return axios.create({
    baseURL: host,
    headers: { 'Authorization': `Bearer ${authToken}` }
  })
}

module.exports = function flags (args) {
  switch (args._.length) {
    case 1:
      return list(args)

    case 2:
      return read(args)

    default:
      return write(args)
  }
}

function list (args) {
  const { namespace } = args
  return client().get(`/flags/${namespace}`)
    .then(res => console.log(res.data
      .map(({ name, state }) => `${name}=${format(state)}`)
      .join('\n')))
    .catch(err => signale.error(err))
}

function read (args) {
  const { namespace } = args
  const [, name] = args._

  return client().get(`/flags/${namespace}/${name}`)
    .then(res => console.log(`${res.data.name}=${format(res.data.state)}`))
    .catch(err => signale.error(err.response.data.message))
}

function write (args) {
  const { namespace } = args
  const [, name, state] = args._

  return client().put(`/flags/${namespace}/${name}`, { state: state === 'on' })
    .then(res => console.log(`${res.data.name}=${format(res.data.state)}`))
    .catch(err => signale.error(err.response.data.message))
}
