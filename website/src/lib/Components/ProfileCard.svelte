<script lang="ts">
	import Card from "./Common/Card.svelte";
	import Typed from "./Common/Typed.svelte";
    import Links from '$lib/Constants/Navigation.json'
    import Profil from '$lib/Constants/Profil.json'
</script>

<Card id="profile">
    <div class="background">

    </div>
    <div class="profil">
        <div>
            <h3>Yoann Fort</h3>
            <Typed className="mobile-profil-typed" strings={Profil.typed}>
                <h4 class="typing"/>
            </Typed>
        </div>
        <div class="socials">
            <ul>
                {#each Links.socials as item }
                    <li>
                        <a target="_blank" href={item.url}>
                            <i class={item.icon}/>
                        </a>
                    </li>
                {/each}
            </ul>
        </div>
    </div>
    <div class="footer">
        <div><a href="#">Mon CV</a></div>
        <div><a href="#contact">Contact</a></div>
    </div>
</Card>

<style>
    .background {
        height: 50vh;
        min-height: 380px;
        width: 100%;
        background-image: url('/profile.jpeg');
        background-position: center center;
        background-size: cover;
        overflow: hidden;
        border:0;
    }

    .background::before{
        display: block;
        content: '';
        position: relative;
        left: -25%;
        bottom: -77%;
        width: 100%;
        height: 238px;
        background: #fff;
        transform: rotate(12deg);
        transform: rotate(12deg);
        z-index: 3;
    }

    .background::after{
        display: block;
        content: '';
        position: relative;
        left: auto;
        right: -25%;
        width: 100%;
        height: 238px;
        z-index: 3;
        bottom: -16%;
        background: #fff;
        transform: rotate(-8deg);
    }

    .profil {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-direction: column;
        width: 100%;
        margin-top: -80px;
        z-index: 4;
    }
    
    h4, h3 {
        margin: 0;
        margin-block-end: 0;
        margin-block-start: 0;
    }

     h4{
        height: 22px;
    }

    .profil ul {
        list-style: none;
        display:flex;
        justify-content: center;
        align-items: center;
        margin: 0;
        padding-inline-start: 0px;
    }
    
    .profil ul li {
        padding: 8px;
    }

    .footer {
        width: 100%;
        height: 70px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-direction: row;
    }
    .footer > div {
        width: 50%;
        display: flex;
        height: 100%;
        justify-content: center;
        align-items: center;
        border-top: 1px solid rgba(50, 50, 50, 0.2);
    }
</style>