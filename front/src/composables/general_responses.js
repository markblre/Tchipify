import { useNotification } from "@kyvg/vue3-notification";
import { useToast } from 'vue-toast-notification';

export function useGeneralResponses() {
    const toast = useToast();

    function manageError(error) {
        console.log(error)
        let message = ""
        if (error.response.data.message) {
            console.log("MESSAGE")
            console.log(error.response.data.message)
            message = error.response.data.message
        } else {
            if (error.response) {
                switch (error.response.status) {
                    case 401: {
                        message = "Unauthorized"
                        break;
                    }
                    case 403: {
                        message = "Forbidden"
                        break;
                    }
                    case 404: {
                        message = "Not found"
                        break;
                    }
                    default: {
                        message = "Something went wrong"
                        break;
                    }
                }
            }
        }
        toast.error(message)
    }

    function manageResponseNotification() {
        let message = "Action succeeded"
        toast.success(message)
    }

    return {manageError, manageResponseNotification}
}