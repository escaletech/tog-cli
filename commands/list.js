const signale = require('signale')
const Table = require('cli-table')
const yaml = require('yaml')
const { highlight } = require('cli-highlight')

const { client, config } = require('./util')

module.exports = function list (args) {
  const { namespace } = config(args)
  if (!namespace) {
    return signale.error('missing namespace')
  }

  return client().get(`/flags/${namespace}`)
    .then(res => {
      const table = new Table({
        head: ['name', 'description', 'rollout'],
        truncate: '...'
      })
      table.push(...res.data.map(({ name, description, rollout }) =>
        [name, description || '-', highlight(yaml.stringify(rollout).trim(), { language: 'yaml' })]))

      console.log('namespace: ' + namespace)
      console.log(table.toString())
      console.log('')
    })
    .catch(err => signale.error(err))
}
