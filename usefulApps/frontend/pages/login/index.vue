<template>
    <div class="container">
        <div class="login-box">
            <label for="login">Login</label>
            <InputText id="login" v-model="login" />
            <label class="password-input" for="password">Password</label>

            <div class="password-input">
                <Password id="password" v-model="password" :feedback="false" />
            </div>

            <Button class="button-input" label="Login" @click="loginFunc()"></Button>
        </div>

        <div class="error-box">
            <span> {{ errorMessage }}</span>
        </div>
    </div>
</template>
<script lang="ts">
import type { IAccountLoginDto } from '~/data/account/IAccountLoginDto';
import type { ITokenValueDto } from '~/data/account/ITokenValueDto';
import type { IErrorMessage } from '~/data/general/IErrorMessage';
import HttpWrapper from '~/logic/webhooks/httpWrappers/httpWrapper';

export default {
    data() {
        return {
            login: '',
            password: '',
            errorMessage: '',
        }
    },
    methods: {
        loginFunc() {
            const _this = this;
            const wrapper = new HttpWrapper();
            wrapper.HttpPost('http://localhost:8081/login/', {
                onRequest({ request, options }) {
                    options.body = <IAccountLoginDto>{
                        account_login: _this.login,
                        account_password: _this.password
                    }
                },
                onResponse({ request, response, options }) {
                    if (response.status === 200) {
                        const res = response._data as ITokenValueDto;
                        console.log(res);
                    }
                },
                onRequestError({ request, options, error }) {
                    _this.errorMessage = error.message;
                },
                onResponseError({ request, response, options }) {
                    const err = response._data as IErrorMessage;
                    _this.errorMessage = err.message;
                }
            });
        }
    }
}
</script>

<style scoped>
.container {
    display: block;
    padding-top: 20%;
    padding-bottom: 11%;
    background-color: rgb(101, 110, 120);
}

.login-box {
    display: grid;
    grid-template-columns: 200px;
    grid-template-rows: 1fr 1fr 1fr;
    padding-left: 45%;
}

.password-input {
    padding-top: 20px;
}

.button-input {
    margin-top: 20px;
}

.error-box {
    color: red;
    padding-top: 50px;
    text-align: center;
}
</style>