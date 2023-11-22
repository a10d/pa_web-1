const errorHandler = {

    handle(error: any) {

        console.error('Error: ', error);
    }
}

export const useErrorHandler = () => errorHandler;
