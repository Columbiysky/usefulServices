<template>
    <Button label="Test console.log('asdf')" @click="buttonClick()">
    </Button>
</template>
<script lang="ts">
import { UseFetchOptions } from 'nuxt/app';
export default {
    methods: {
        buttonClick() {
            // const wrapper = new HttpWrapper();  потом напишу рабочий враппер, этот пока откисает
            // useFetch('http://localhost:8080/hello/') пример простого get без всего
            useFetch('http://localhost:8080/hello/', <UseFetchOptions<any>>{
                method: 'POST',
                body:
                {
                    "id": 1,
                    "name": "FROM CLIENT"
                }
            })
        }
    }
}
</script>