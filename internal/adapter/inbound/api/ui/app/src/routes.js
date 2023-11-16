import Home from "./pages/Home.svelte";
import Pattern from "./pages/Pattern.svelte";

import Bits from './pages/Bits.svelte';
import Bit from "./pages/Bit.svelte";
import Donations from './pages/Donations.svelte';
import Gifts from './pages/gifts.svelte';
import Gift from "./pages/gift.svelte";
import Resubs from "./pages/resubs.svelte";
import Resub from "./pages/resub.svelte";
import Donation from "./pages/Donation.svelte";


export default [
    {
        title: "Home",
        path: "/",
        component: Home
    },
    {
        title: "Pattern",
        path: "/pattern",
        component: Pattern
    },
    {
        title: "Resubs",
        path: "/resubs",
        component: Resubs
    },
    {
        title: "Resub alert",
        path: "/resubs/:id",
        component: Resub
    },
    {
        title: "Gift subs",
        path: "/gifts",
        component: Gifts
    },
    {
        title: "Gift sub alert",
        path: "/gifts/:id",
        component: Gift
    },
    {
        title: "Bits",
        path: "/bits",
        component: Bits
    },
    {
        title: "Bit Alert",
        path: "/bits/:id",
        component: Bit
    },
    {
        title: "Donations",
        path: "/donations",
        component: Donations
    },
    {
      title: "Donation",
      path: "/donations/:id",
      component: Donation
    }
];