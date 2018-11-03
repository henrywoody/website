<template>
    <pre><code :class="language">{{ code }}</code></pre>
</template>

<script>
import hljs from 'highlight.js/lib/highlight';
import python from 'highlight.js/lib/languages/python';

hljs.registerLanguage('python', python);

export default {
    name: "CodeBlock",
    props: {
        code: String,
        language: String,
    },
    mounted() {
        hljs.highlightBlock(this.$el);
    }
}
</script>

<style type="text/css">
/* https://highlightjs.org */

.hljs {
    transition: var(--dark-mode-transition);
}

/*LIGHT*/
:not(.dark) .hljs {
    display: block;
    background-color: var(--background-color);
    padding: 0.5em;
    color: #333333;
    overflow-x: auto;
}

:not(.dark) .hljs-comment,
:not(.dark) .hljs-meta {
    color: #969896;
}

:not(.dark) .hljs-string,
:not(.dark) .hljs-variable,
:not(.dark) .hljs-template-variable,
:not(.dark) .hljs-strong,
:not(.dark) .hljs-emphasis,
:not(.dark) .hljs-quote {
    color: #df5000;
}

:not(.dark) .hljs-keyword,
:not(.dark) .hljs-selector-tag,
:not(.dark) .hljs-type {
    color: #a71d5d;
}

:not(.dark) .hljs-literal,
:not(.dark) .hljs-symbol,
:not(.dark) .hljs-bullet,
:not(.dark) .hljs-attribute {
    color: #0086b3;
}

:not(.dark) .hljs-section,
:not(.dark) .hljs-name {
    color: #63a35c;
}

:not(.dark) .hljs-tag {
    color: #333333;
}

:not(.dark) .hljs-title,
:not(.dark) .hljs-attr,
:not(.dark) .hljs-selector-id,
:not(.dark) .hljs-selector-class,
:not(.dark) .hljs-selector-attr,
:not(.dark) .hljs-selector-pseudo {
    color: #795da3;
}

:not(.dark) .hljs-addition {
    color: #55a532;
    background-color: #eaffea;
}

:not(.dark) .hljs-deletion {
    color: #bd2c00;
    background-color: #ffecec;
}

:not(.dark) .hljs-link {
    text-decoration: underline;
}

/*DARK*/
.dark .hljs {
    display: block;
    overflow-x: auto;
    padding: 0.5em;
    background-color: var(--background-color);
}

.dark .hljs,
.dark .hljs-tag,
.dark .hljs-subst {
    color: #f8f8f2;
}

.dark .hljs-strong,
.dark .hljs-emphasis {
    color: #a8a8a2;
}

.dark .hljs-bullet,
.dark .hljs-quote,
.dark .hljs-number,
.dark .hljs-regexp,
.dark .hljs-literal,
.dark .hljs-link {
    color: #ae81ff;
}

.dark .hljs-code,
.dark .hljs-title,
.dark .hljs-section,
.dark .hljs-selector-class {
    color: #a6e22e;
}

.dark .hljs-strong {
    font-weight: bold;
}

.dark .hljs-emphasis {
    font-style: italic;
}

.dark .hljs-keyword,
.dark .hljs-selector-tag,
.dark .hljs-name,
.dark .hljs-attr {
    color: #f92672;
}

.dark .hljs-symbol,
.dark .hljs-attribute {
    color: #66d9ef;
}

.dark .hljs-params,
.dark .hljs-class .hljs-title {
    color: #f8f8f2;
}

.dark .hljs-string,
.dark .hljs-type,
.dark .hljs-built_in,
.dark .hljs-builtin-name,
.dark .hljs-selector-id,
.dark .hljs-selector-attr,
.dark .hljs-selector-pseudo,
.dark .hljs-addition,
.dark .hljs-variable,
.dark .hljs-template-variable {
    color: #e6db74;
}

.dark .hljs-comment,
.dark .hljs-deletion,
.dark .hljs-meta {
    color: #75715e;
}
</style>