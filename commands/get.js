const yaml = require('yaml')
const signale = require('signale')
const { highlight } = require('cli-highlight')

const { client, config } = require('./util')

module.exports = function get (args) {
  const { namespace } = config(args)
  if (!namespace) {
    return signale.error('missing namespace')
  }

  const name = args._[1]

  return client().get(`/flags/${namespace}/${name}`)
    .then(res => console.log(highlight(yaml.stringify(res.data), { language: 'yaml' })))
    .catch(err =>
      err.response && err.response.status === 404
        ? signale.error('flag not found')
        : signale.error(err)
    )
}
