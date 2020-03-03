const signale = require('signale')

const { client, config } = require('./util')

module.exports = function del (args) {
  const { namespace } = config(args)
  if (!namespace) {
    return signale.error('missing namespace')
  }

  const name = args._[1]
  if (!name) {
    return signale.error('missing name')
  }

  return client().delete(`/flags/${namespace}/${name}`)
    .then(res => signale.success(`flag ${namespace}/${name} deleted`))
    .catch(err => err.response && err.response.status === 404
      ? signale.error('flag not found')
      : signale.error(err))
}
