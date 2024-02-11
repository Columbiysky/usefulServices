<template>
    <div class="container">
        <div class="panel">
            <Button label="test button"></Button>
            <PanelMenu :model="items">
            </PanelMenu>
        </div>
    </div>
</template>
<script setup lang="ts">
    import { ref } from "vue";
    const items = ref([
        {
            label: 'Check currencies courses',

        },
    ]);
</script>
<style>
.container {
    display: block;
    height: max-content;
    background-color: rgb(16, 45, 78);
}

p {
    color: aliceblue;
}
</style>