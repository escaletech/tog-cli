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
  .argv

if (!argv) {
  console.log('nothing to see here')
}
