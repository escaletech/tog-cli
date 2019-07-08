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
  .command(['flag', 'flags'], 'read or update flags',
    yargs => {
      yargs.usage('tog flags [name [state]]')
      yargs.positional('name', { description: 'flag name' })
      yargs.positional('state', { description: 'flag state', choices: ['on', 'off'], implies: 'name' })
      yargs.option('namespace', { alias: 'n' })
    }, require('./commands/flags'))
  .command(['exp', 'experiment', 'experiments'], 'read or update experiments',
    yargs => {
      yargs.usage('tog exp [name]')
      yargs.positional('name', { description: 'experiment name' })
      yargs.option('namespace', { alias: 'n' })
      yargs.option('on', { description: 'Enable flags for the experiment' })
      yargs.option('off', { description: 'Disable flags for the experiment' })
      yargs.option('del', { description: 'Delete flags for the experiment' })
      yargs.option('weight', { alias: 'w', description: 'Sets the weight for the experiment' })
    }, require('./commands/experiments.js'))
  .demandCommand(1, 'Choose a command')
  .argv

if (!argv) {
  console.log('nothing to see here')
}
