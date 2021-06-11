<script>
    import feather from 'feather-icons'

    import {onMount} from 'svelte'

    export let iconName
    export let size
    export let color

    onMount(() => {
        feather.replace({
            // initialize these as empty properties so CSS can override them at will
            stroke: "",
            fill: ""
        })
    })
</script>

<span class="feather-icon">
    <i class="icon-{color}-hover" data-feather="{iconName}" width="{size}" height="{size}" stroke-width="0.7" fill="none"></i>
</span>