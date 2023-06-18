import _ from "lodash"
import { UseFetchOptions } from "nuxt/app"

export default class HttpWrapper {
    HttpGet(url: string, options?: UseFetchOptions<any>) {
        if (_.isNil(options)) {
            options = <UseFetchOptions<any>>{
                method: "GET",

            }
        } else {
            options.method = "GET"
        }

        const result = async (params: any) => {
            await useFetch(url, options)
        }

        return result
    }

    HttpPost(url: string, options?: UseFetchOptions<any>) {
        if (_.isNil(options)) {
            options = <UseFetchOptions<any>>{
                method: "POST"
            }
        } else {
            options.method = "POST"
        }

        const result = async (params: any) => {
            await useFetch(url, options)
        }

        return result
    }
}