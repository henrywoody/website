<template>
    <a :href="href" target="_blank" rel="noopener noreferrer">
        <slot></slot>
    </a>
</template>

<script>
export default {
    name: "ExternalLink",
    props: {
        href: String,
    }
}
</script>