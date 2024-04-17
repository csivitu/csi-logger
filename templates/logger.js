import winston from 'winston';
import axios from 'axios';
import Transport from 'winston-transport'

class AuthTransport extends Transport {
  constructor(opts) {
    super(opts);
    this.url = opts.url
    this.auth = opts.auth
  }

  log(info, callback) {
    axios.post(this.url, info, {
      headers: {
        Authorization: this.auth
      }
    }).then(() => {
      callback();
    })
      .catch((error) => {
        console.error(error);
        callback(error);
      });
  }
}

const formatconfig = winston.format.combine(
  winston.format.timestamp({
    format: 'YYYY-MM-DD HH:mm:ss',
  }),
  winston.format.simple(),
  winston.format.json(),
  winston.format.prettyPrint(),
  winston.format.errors({ stack: true })
);

const createLog = (level) =>
  winston.createLogger({
    transports: [
      new winston.transports.Console({
        level,
        format: formatconfig,
      }),
      new AuthTransport({
        url: "http://localhost:3000/log",
        format: formatconfig,
        auth: "something cool"
      }),
    ],
    exceptionHandlers: [
      new winston.transports.Console({
        format: formatconfig,
      }),
    ],
  });

const errorLogger = createLog('error');
const infoLogger = createLog('info');
const protectLogger = createLog('warn');

const logger = {
  info: (log, path) => {
    const logData = {
      level: 'info',
      message: log,
      path: path
    };
    infoLogger.info(logData);
  },
  error: (log, path) => {
    const logData = {
      level: 'error',
      message: log,
      path: path
    };
    errorLogger.error(logData);
  },
  warn: (log, path) => {
    const logData = {
      level: 'warn',
      message: log,
      path: path
    };
    protectLogger.warn(logData);
  },
};

export default logger;
