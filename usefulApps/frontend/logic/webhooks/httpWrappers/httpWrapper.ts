import _ from "lodash";
import type { UseFetchOptions } from "nuxt/app";

export default class HttpWrapper {
    HttpGet(url: string, options?: UseFetchOptions<any>) {
        if (_.isNil(options)) {
            options = <UseFetchOptions<any>>{
                method: "GET",
            };
        } else {
            options.method = "GET";
        }

        return useFetch(url, options);
    }

    HttpPost(url: string, options?: UseFetchOptions<any>) {
        if (_.isNil(options)) {
            options = <UseFetchOptions<any>>{
                method: "POST",
            };
        } else {
            options.method = "POST";
        }

        return useFetch(url, options);
    }
}