const { inspect } = require('util')
const signale = require('signale')

const { readConfig, saveConfig } = require('../services/config')

const configKeys = ['host']

const format = value => inspect(value, false, 1, true)

module.exports = function config (args) {
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
  const config = readConfig()

  console.log(configKeys
    .filter(key => config[key] !== undefined)
    .map(key => `${key}=${format(config[key])}`)
    .join('\n'))
}

function read (args) {
  const key = args._[1]

  if (!configKeys.includes(key)) {
    signale.error(`Unknown configuration key '${key}'`)
    return 1
  }

  console.log(format(readConfig()[key]))
}

function write (args) {
  const [, key, value] = args._

  if (!configKeys.includes(key)) {
    signale.error(`Unknown configuration key '${key}'`)
    return 1
  }

  const config = readConfig()
  config[key] = value
  saveConfig(config)

  console.log(`${key}=${format(value)}`)
}
