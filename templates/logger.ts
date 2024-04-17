import winston, { format, transports, Logger } from 'winston';
import axios from 'axios';
import Transport from 'winston-transport';

interface LogData {
  level: string;
  message: string;
  path: string;
}

interface AuthTransportOptions extends Transport.TransportStreamOptions {
  url: string;
  auth: string;
}

class AuthTransport extends Transport {
  url: string;
  auth: string;

  constructor(opts: AuthTransportOptions) {
    super(opts);
    this.url = opts.url;
    this.auth = opts.auth;
  }

  log(info: LogData, callback: () => void): void {
    axios.post(this.url, info, {
      headers: {
        Authorization: this.auth
      }
    }).then(() => {
      callback();
    })
      .catch((error) => {
        console.error(error);
        callback();
      });
  }
}

const formatconfig = format.combine(
  format.timestamp({
    format: 'YYYY-MM-DD HH:mm:ss',
  }),
  format.simple(),
  format.json(),
  format.prettyPrint(),
  format.errors({ stack: true })
);

const createLog = (level: string): Logger =>
  winston.createLogger({
    transports: [
      new transports.Console({
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
      new transports.Console({
        format: formatconfig,
      }),
    ],
  });

const errorLogger = createLog('error');
const infoLogger = createLog('info');
const protectLogger = createLog('warn');

const logger = {
  info: (log: string, path: string) => {
    const logData: LogData = {
      level: 'info',
      message: log,
      path: path
    };
    infoLogger.info(logData);
  },
  error: (log: string, path: string) => {
    const logData: LogData = {
      level: 'error',
      message: log,
      path: path
    };
    errorLogger.error(logData);
  },
  warn: (log: string, path: string) => {
    const logData: LogData = {
      level: 'warn',
      message: log,
      path: path
    };
    protectLogger.warn(logData);
  },
};

export default logger;