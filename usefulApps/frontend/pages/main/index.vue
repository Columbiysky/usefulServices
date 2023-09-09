<template>
    <Button label="Test console.log('asdf')" @click="buttonClick()">
    </Button>
    <Button label="Youtube downloader"></Button>
</template>
<script lang="ts">
import HttpWrapper from '~/logic/webhooks/httpWrappers/httpWrapper';
export default {
    methods: {
        buttonClick() {
            const wrapper = new HttpWrapper();

            // GET
            wrapper.HttpGet('http://localhost:8080/helloFromPyServ/').then(
                (result) => {
                    console.log(result); //итоговое значение, которое хотим получить лежит в result.data.value
                },
                (error) => {
                    console.log("Error: ", error);
                }
            );

            // POST
            // wrapper.HttpPost('http://localhost:8080/hello/', {
            //     body: <CustomObject>{
            //         "id": 1,
            //         "name": "FROM CLIENT"
            //     },
            // }).then(
            //     (result) => {
            //         const value = result.data.value;

            //         // здесь вместо CustomObject можно воткуть что угодно,
            //         // из value нужно будет правильно вписать названия свойств, чтобы он их нашел и вставил
            //         const res = new CustomObject(value.id, value.name);

            //         console.log(res);
            //     },
            //     (error) => {
            //         console.log("Error: ", error);
            //     }
            // );
        }
    }
}
</script>