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
    <DocsListRow title="Learn Chai">
        <DocPanel title="Getting Started Guide" href="/docs/guide">
            An introductory walkthrough of the Chai programming language.
            If you are new to Chai, start here.
        </DocPanel>
        <DocPanel title="Chai Koans" href="/docs/koans">
            A set of introductory exercises designed to get you familiar with 
            programming in Chai through application and example.
        </DocPanel>
        <DocPanel title="Style Guide" href="/docs/style-guide">
            A document explaining how to write clean, idiomatic Chai.
            This is a must read for any Chai programmer, new or old.
        </DocPanel>
    </DocsListRow>
</div>