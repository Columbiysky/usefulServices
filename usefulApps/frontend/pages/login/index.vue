<template>
    <div class="container">
        <div class="flex flex-column gap-2">
            <small for="login">Login</small>
            <InputText id="login" v-model="login" />
            <small for="password">Password</small>
            <Password v-model="password" :feedback="false" />

            <Button label="Login" @click="buttonClick()"></Button>
        </div>
    </div>
</template>
<script lang="ts">
export default {
    data() {
        return {
            login: '',
            password: '',
            error: '',
        }
    },
    methods: {
        buttonClick() {
            console.log(this.login)
            console.log(this.password)
        },
        onUpload() {
            // const toast = useToast();
            // toast.add({ severity: 'info', summary: 'Success', detail: 'File Uploaded', life: 3000 });
        }
    }
}
</script>

<style scoped>
.container {
    display: block;
    background-color: rgb(213, 222, 231);
}
</style>