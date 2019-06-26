#!/usr/bin/env node
const yargs = require('yargs')

const argv = yargs
  .command('login', 'authenticate with the server',
    yargs =>
      yargs.describe('host', 'tog server host').alias('h', 'host')
    ,
    require('./commands/login'))
  .argv

if (!argv) {
  console.log('nothing to see here')
}
