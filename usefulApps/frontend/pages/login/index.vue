<template>
    <div class="container">
        <div class="login-box">
            <div class="label-and-control">
                <label for="login">Login</label>
                <InputText id="login" v-model="login" />
            </div>

            <div class="label-and-control">
                <label for="password">Password</label>
                <Password id="password" v-model="password" :feedback="false" />
            </div>

            <div class="label-and-control">
                <Button class="button-input" label="Login" @click="loginFunc()"></Button>
                <div class="error-box">
                    <span> {{ errorMessage }}</span>
                </div>
            </div>
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
                        localStorage.setItem('token', res.token_value);
                        navigateTo({ path: "/main" })
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
label {
    color: white;
}

.container {
    justify-content: center;
    display: flex;
    background-color: rgb(16, 45, 78);
    height: 100%;
    width: 100%;
}

.login-box {
    display: grid;
    grid-template-columns: 200px;
    grid-template-rows: 1fr 1fr 1fr 1fr;
    grid-row-gap: 20px;
    padding-top: 15%;
    padding-bottom: 20%;
    height: auto;
}

.label-and-control {
    display: grid;
    grid-template-rows: auto auto;
    grid-row-gap: 5px;
}

.button-input {
    width: -webkit-fill-available;
}

.error-box {
    color: red;
    padding-top: 10px;
    text-align: center;
}
</style>