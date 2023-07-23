<script lang="ts">
    import links from '$lib/Constants/Navigation.json'
	import Typed from './Typed.svelte';

    const profileSubtitle = [
        "Développer Web Full Stack",
        "Entrepreneur",
        "Pationné de technologie"
    ]

    
</script>

<header class="mobile-nav">
    <div class="profile">
        <div>
            <p>Yoann Fort</p>
            <Typed strings={profileSubtitle}>
                <p class="typing"></p>
            </Typed>
        </div>
        <div>
            <div>
                panier
            </div>
            <div>
                dark
            </div>
            <div>
                button
            </div>
        </div>
         
    </div>

    <div class="top-menu">
        <ul>
            {#each links.links as component}
                <li style="width: calc(100% / {links.links.length});">
                    <a href={component.url}>
                        <i class={component.icon} />
                        <p>{component.name.toUpperCase()}</p>
                    </a>
                </li>
            {/each}
        </ul>
    </div>

</header>

<style>

    i {
        font-size: 18px;
    }
    
    .mobile-nav {
        background-color: white;
    }

    .profile {
        padding: 4px 18px;
        height: 42px;
        text-align: space;
        display: flex;
        flex-direction: row;
        flex-wrap:nowrap;
        justify-content: space-between;
    }

    .profile > div{
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
    }

    .profile > div:last-child{
        flex-direction: row;
    }

    .profile p {
        margin: 0;
        padding: 0;
    }

    .top-menu {
    }

    .top-menu > ul {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        list-style: none;
        border-top: 1px solid grey;
        border-bottom: 1px solid grey;
        padding-left: 0px;
        margin: 0;
    }

    .top-menu > ul > li:first-child{
        border-left: 1px solid grey;
    }

    .top-menu > ul > li {
        padding: 12px 0;
        border-right: 1px solid grey;
        margin: 10px 0;
    }

    .top-menu > ul > li > a {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        flex-wrap: wrap;
    }

    .top-menu > ul > li > a > p {
        font-size: 12px;
        margin: 4px;
    }

    .top-menu > ul > li > a {
        text-decoration: none;
    }
    .top-menu > ul > li > a:visited {
        text-decoration: none;
        
    }
</style>