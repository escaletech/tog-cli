#!/usr/bin/env node
const yargs = require('yargs')

const argv = yargs
  .command('login', 'authenticate with the server',
    yargs =>
      yargs.describe('host', 'tog server host').alias('h', 'host')
    , require('./commands/login'))
  .command('config', 'list configuration',
    yargs => {
      yargs.usage('tog config [key [value]]')
      yargs.positional('key', { description: 'configuration key to get/set' })
      yargs.positional('value', { implies: 'key', description: 'value to set' })
    }, require('./commands/config'))
  .command('flags', 'read or update flags',
    yargs => {
      yargs.usage('tog flags [name [state]]')
      yargs.positional('name', { description: 'flag name' })
      yargs.positional('state', { description: 'flag state', choices: ['on', 'off'], implies: 'name' })
      yargs.option('namespace', { alias: 'n', required: true })
    }, require('./commands/flags'))
  .argv

if (!argv) {
  console.log('nothing to see here')
}
