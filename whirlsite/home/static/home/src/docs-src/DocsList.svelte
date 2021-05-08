<script>
    import DocPanel from './DocPanel.svelte';
    import DocsListRow from './DocsListRow.svelte'
</script>

<style>
    .docs-list {
        margin-top: 2rem;

        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
    }
</style>

<div class="docs-list">
    <DocsListRow title="Learn Whirlwind">
        <DocPanel title="Getting Started Guide" href="/docs/guide">
            An introductory walkthrough of the Whirlwind programming language.
            If you are new to Whirlwind, start here.
        </DocPanel>
        <DocPanel title="Whirlwind Koans" href="/docs/koans">
            A set of introductory exercises designed to get you familiar with 
            programming in Whirlwind through application and example.
        </DocPanel>
        <DocPanel title="Style Guide" href="/docs/style-guide">
            A document explaining how to write clean, idiomatic Whirlwind.
            This is a must read for any Whirlwind programmer, new or old.
        </DocPanel>
    </DocsListRow>
</div>