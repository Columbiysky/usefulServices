<template>
    <div class="container">
        <div class="login-box">
            <label for="login">Login</label>
            <InputText id="login" v-model="login" />
            <label class="password-input" for="password">Password</label>

            <div class="password-input">
                <Password id="password" v-model="password" :feedback="false" />
            </div>

            <!-- <div class="button-input"> -->
            <Button class="button-input" label="Login" @click="buttonClick()"></Button>
            <!-- </div> -->
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
        }
    }
}
</script>

<style scoped>
.container {
    display: block;
    padding-top: 20%;
    padding-bottom: 15%;
    background-color: rgb(236, 241, 246);
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
</style>