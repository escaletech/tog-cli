const yaml = require('yaml')
const signale = require('signale')
const { highlight } = require('cli-highlight')

const { client: getClient, config } = require('./util')

module.exports = function set (args) {
  const { namespace } = config(args)
  if (!namespace) {
    return signale.error('missing namespace')
  }

  const name = args._[1]

  const client = getClient()

  return client.get(`/flags/${namespace}/${name}`)
    .then(res => res.data)
    .catch(err => err.response && err.response.status === 404
      ? { namespace, name }
      : Promise.reject(err))
    .then(flag => {
      const { description, rollout, on, off } = args
      if (!flag.rollout && !(rollout || on || off)) {
        return signale.error('to create a flag, provide at least one rollout option')
      }

      if (!(description || rollout || on || off)) {
        return signale.error('please provide at least one option to update the flag')
      }

      if (description) {
        flag.description = description
      }

      if (on) {
        flag.rollout = [{ value: true }]
      }

      if (off) {
        flag.rollout = [{ value: false }]
      }

      if (rollout) {
        flag.rollout = yaml.parse(rollout)
      }

      return client.put(`/flags/${namespace}/${name}`, {
        description: flag.description,
        rollout: flag.rollout
      })
    })
    .then(({ data }) => {
      signale.success('flag saved')
      console.log(highlight(yaml.stringify(data), { language: 'yaml' }))
    })
    .catch(err => {
      signale.error(err.response.data)
    })
}
