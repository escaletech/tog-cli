const fs = require('fs')

const filePath = require('os').homedir() + '/.config/tog.json'

class Config {
  constructor ({ host, namespace, authToken } = {}) {
    /** @type {string} */
    this.host = host

    /** @type {string} */
    this.namespace = namespace

    /** @type {string} */
    this.authToken = authToken
  }

  /**
   * @returns {boolean}
   */
  isValid () {
    return this.host && this.authToken
  }
}

/**
 * @returns {Config}
 */
function readConfig () {
  try {
    const buffer = fs.readFileSync(filePath)
    return new Config(JSON.parse(buffer.toString()))
  } catch (err) {
    return new Config()
  }
}

/**
 * @param {Config} config
 */
function saveConfig (config) {
  fs.writeFileSync(filePath, Buffer.from(JSON.stringify(config, '', '  ')))
}

module.exports = {
  Config,
  readConfig,
  saveConfig
}
