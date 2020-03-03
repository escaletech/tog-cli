#!/usr/bin/env node
const yargs = require('yargs')

const argv = yargs
  .command('login', 'authenticate with the server',
    yargs =>
      yargs.describe('host', 'tog server host').alias('h', 'host')
    , require('./commands/login'))
  .command(['config', 'conf'], 'list or update configuration',
    yargs => {
      yargs.usage('tog config [key [value]]')
      yargs.positional('key', { description: 'configuration key to get/set' })
      yargs.positional('value', { implies: 'key', description: 'value to set' })
    }, require('./commands/config'))
  .command('list', 'list flags',
    yargs => {
      yargs.usage('tog list')
      yargs.option('namespace', { alias: 'n' })
    }, require('./commands/list'))
  .command('get', 'get a flag',
    yargs => {
      yargs.usage('tog get <name>')
      yargs.option('namespace', { alias: 'n' })
    }, require('./commands/get'))
  .command('set', 'create or update a flag',
    yargs => {
      yargs.usage('tog set <name> [options]')
      yargs.option('off')
      yargs.option('on')
      yargs.option('rollout', { alias: 'r' })
      yargs.option('description', { alias: 'd' })
      yargs.option('namespace', { alias: 'n' })
    }, require('./commands/set'))
  .command('delete', 'delete a flag',
    yargs => {
      yargs.usage('tog delete <name>')
      yargs.option('namespace', { alias: 'n' })
    }, require('./commands/delete'))
  .demandCommand(1, 'Choose a command')
  .argv

if (!argv) {
  console.log('nothing to see here')
}
