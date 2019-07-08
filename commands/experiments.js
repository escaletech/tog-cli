const signale = require('signale')
const axios = require('axios').default
const R = require('ramda')

const { readConfig } = require('../services/config')

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

    default:
      return read(args)
  }
}

function list (args) {
  const { namespace } = args
  return client().get(`/experiments/${namespace}`)
    .then(res => console.log(res.data))
    .catch(err => signale.error(err))
}

function read (args) {
  const { namespace } = args
  const [, name] = args._

  const modifiers = R.pick(['on', 'off', 'del', 'weight'], args)
  if (!R.isEmpty(modifiers)) {
    return client().get(`/experiments/${namespace}/${name}`)
      .then(({ data }) => data)
      .catch(err => err.response.status === 404 ? {} : Promise.reject(err))
      .then(exp => enrich(exp, modifiers))
      .then(exp => client().put(`/experiments/${namespace}/${name}`, exp))
      .then(({ data }) => signale.success('experiment saved') || console.log(data))
      .catch(err => signale.error(err.response.data))
  } else {
    return client().get(`/experiments/${namespace}/${name}`)
      .then(({ data }) => console.log(data))
      .catch(err => signale.error(err.response.data.message))
  }
}

function enrich (exp, { on, off, del, weight }) {
  const clone = R.clone(exp)

  if (on) {
    const onKeys = Array.isArray(on) ? on : [on]
    clone.flags = { ...clone.flags, ...R.fromPairs(R.map(k => [k, true], onKeys)) }
  }

  if (off) {
    const offKeys = Array.isArray(off) ? off : [off]
    clone.flags = { ...clone.flags, ...R.fromPairs(R.map(k => [k, false], offKeys)) }
  }

  if (del) {
    const delKeys = Array.isArray(del) ? del : [del]
    clone.flags = R.pipe(
      R.toPairs,
      R.filter(([k]) => !delKeys.includes(k)),
      R.fromPairs
    )(clone.flags)
  }

  if (weight !== undefined) {
    clone.weight = Number(weight)
  }

  return clone
}
