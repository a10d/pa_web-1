const errorHandler = {
    handle(error: any) {
        // Error reporting to analytics endpoint...
        // e.g. Sentry.captureException(error);
        // something like that.
        // Or just log it to console I guess.
        console.error('Error: ', error);
    }
}

export const useErrorHandler = () => errorHandler;
