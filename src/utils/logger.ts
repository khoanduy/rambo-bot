import { Logger, createLogger, format, transports } from 'winston';

export const logger: Logger = createLogger({
  level: 'debug',
  format: format.combine(
    format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
    format.colorize(),
    format.errors({ stack: true }),
    format.printf(l => `${l.timestamp} ${l.level}: ${l.message}` + (l.splat !== undefined ? `${l.splat}` : ' ')),
  ),
  transports: [
    // - Write all logs with importance level of `error` or less to `error.log`
    // - Write all logs with importance level of `info` or less to `combined.log`
    new transports.Console(),
    new transports.File({ filename: 'error.log', level: 'error' }),
    new transports.File({ filename: 'combined.log' }),
  ],
});
