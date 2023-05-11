<template>
    <Button label="Test console.log('asdf')" @click="buttonClick()">
    </Button>
</template>
<script>
export default {
    methods: {
        buttonClick() {
            console.log('asdf');
        }
    }
}
</script>