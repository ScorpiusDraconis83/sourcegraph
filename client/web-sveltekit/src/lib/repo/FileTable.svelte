<script lang="ts">
    import { mdiFileDocumentOutline, mdiFolderOutline } from '@mdi/js'

    import Icon from '$lib/Icon.svelte'
    import type { TreeEntryWithCommitInfo } from './FileTable.gql'
    import { replaceRevisionInURL } from '$lib/web'
    import Timestamp from '$lib/Timestamp.svelte'
    import type { TreeEntryFields } from './api/tree'

    export let entries: TreeEntryFields[]
    export let commitInfo: TreeEntryWithCommitInfo[]
    export let revision: string

    $: commitInfoByPath = new Map(commitInfo.map(entry => [entry.canonicalURL, entry]))
    $: hasCommitInfo = commitInfo.length > 0
</script>

<table class:hasCommitInfo>
    <thead>
        <tr>
            <th class="file-name">Name</th>
            {#if hasCommitInfo}
                <th class="commit-message">Last commit message</th>
                <th class="commit-date">Last commit date</th>
            {/if}
        </tr>
    </thead>
    <tbody>
        {#each entries as entry}
            <tr>
                <td class="file-name">
                    <Icon svgPath={entry.isDirectory ? mdiFolderOutline : mdiFileDocumentOutline} inline />
                    <a href={replaceRevisionInURL(entry.canonicalURL, revision)}>{entry.name}</a>
                </td>
                {#if hasCommitInfo}
                    {@const commit = commitInfoByPath.get(entry.canonicalURL)?.history.nodes[0]?.commit}
                    {#if commit}
                        <td class="commit-message">
                            <a href={commit.canonicalURL}>{commit.subject}</a>
                        </td>
                        <td class="commit-date">
                            <a href={commit.canonicalURL}>
                                <Timestamp date={commit.author.date} strict />
                            </a>
                        </td>
                    {:else}
                        <td />
                        <td />
                    {/if}
                {/if}
            </tr>
        {/each}
    </tbody>
</table>

<style lang="scss">
    table {
        width: 100%;
        table-layout: fixed;
    }

    th {
        background-color: var(--body-bg);
        border-bottom: 1px solid var(--border-color);
        padding: 0.25rem 0.5rem;
        font-weight: normal;
    }

    td {
        background-color: var(--color-bg-1);
        border-bottom: 1px solid var(--border-color);
        padding: 0.25rem 0.5rem;
    }

    th,
    td {
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        max-width: 100%;
        box-sizing: content-box;
    }

    .commit-message,
    .commit-date {
        a {
            color: var(--text-muted);
        }
    }

    .hasCommitInfo {
        .file-name {
            width: 30%;
        }
    }

    .commit-date {
        width: 120px;
        text-align: right;
    }
</style>
